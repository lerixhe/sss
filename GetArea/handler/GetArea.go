package handler

import (
	"context"
	"encoding/json"
	GETAREA "sss/GetArea/proto/GetArea"
	"sss/ihomeWeb/models"
	"sss/ihomeWeb/utils"
	"time"

	_ "github.com/astaxie/beego/cache/redis"

	"github.com/astaxie/beego/cache"

	// _ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

type GetArea struct{}

// // Call is a single request handler called via client.Call or the generated client code
// func (e *GetArea) Call(ctx context.Context, req *GetArea.Request, rsp *GetArea.Response) error {
// 	log.Log("Received GetArea.Call request")
// 	rsp.Msg = "Hello " + req.Name
// 	return nil
// }

// 服务所属方法的具体定义
func (e *GetArea) GetAreas(ctx context.Context, req *GETAREA.Request, rsp *GETAREA.Response) error {
	beego.Info("请求地区信息：GetArea api/v1.0/areas")
	// 初始化错误码，默认成功信息
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)
	/* 从缓存中获取数据，有数据则直接返回给前端
	没有数据，则去数据库查找数据
	将数据存到缓存中
	把数据发送给前端。
	*/

	// {"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}
	redisConf := map[string]string{
		"key":      utils.G_server_name,
		"conn":     utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum":    utils.G_redis_dbnum,
		"password": utils.G_redis_auth,
	}
	// 将map转换为json
	redisConfJSON, _ := json.Marshal(redisConf)

	// 链接redis
	bm, err := cache.NewCache("redis", string(redisConfJSON))
	if err != nil {
		beego.Info("缓存查询失败", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 缓存中获取数据，指定key为area_info
	areaValueJSON := bm.Get("area_info")
	if areaValueJSON != nil {
		// 获取到数据了,打包发送给前端
		beego.Info("缓存查询成功")
		// beego.Info("JSON:", string(areaValueJSON.([]byte)))
		// 查询到的数据仍是json，这里将其转回map,返回的是一组地址，需要使用切片接收
		areaMap := make([]map[string]interface{}, 0)
		json.Unmarshal(areaValueJSON.([]byte), &areaMap)
		beego.Info("缓存中的数据：")
		// 把map打包成rsp
		for index, area := range areaMap {
			beego.Info(index, area)
			tmp := GETAREA.Response_Areas{Aid: int32(area["aid"].(float64)), Aname: area["aname"].(string)}
			rsp.Data = append(rsp.Data, &tmp)
		}
		// 如果缓存有，则无需去数据库了
		return nil
	}
	beego.Info("缓存未找到，开始查询数据库")
	// 创建orm句柄
	o := orm.NewOrm()

	// 查什么？用什么接收
	areas := []models.Area{}
	num, err := o.QueryTable("area").All(&areas, "Id", "Name")
	if err != nil {
		beego.Info("数据库查询失败", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return err
	}
	if num == 0 {
		beego.Info("没有查询到数据")
		rsp.Error = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	// 将查询到的数据，存入缓存中
	areaJSON, err := json.Marshal(areas)
	err = bm.Put("area_info", areaJSON, time.Second*time.Duration(utils.G_redis_expire))
	if err != nil {
		beego.Info("数据库缓存失败", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
	}
	// 将查询到的数据，按照proto的格式发送给web
	for key, value := range areas {
		beego.Info(key, value)
		tmp := GETAREA.Response_Areas{Aid: int32(value.Id), Aname: value.Name}
		rsp.Data = append(rsp.Data, &tmp)
	}
	return nil
}
