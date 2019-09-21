package handler

import (
	"context"
	"encoding/json"
	"reflect"
	"sss/ihomeWeb/models"
	"time"

	"github.com/astaxie/beego/orm"

	GETINDEX "sss/GetIndex/proto/GetIndex"
	"sss/ihomeWeb/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

type GetIndex struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetIndex) CallGetIndex(ctx context.Context, req *GETINDEX.Request, rsp *GETINDEX.Response) error {
	beego.Info("获取首页轮播图 GetIndex api/v1.0/house/index")
	// 初始化回复
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	//先去缓存读index数据，如果有直接返回
	// 读取redis链接配置
	redisConf := map[string]string{
		"key":      utils.G_server_name,
		"conn":     utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum":    utils.G_redis_dbnum,
		"password": utils.G_redis_auth,
	}
	// 将map转换为json
	redisConfJSON, _ := json.Marshal(redisConf)
	// 开始链接redis
	bm, err := cache.NewCache("redis", string(redisConfJSON))
	if err != nil {
		beego.Info("缓存连接失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	reply := bm.Get("home_page_data")
	if reply != nil {
		beego.Info("查询到缓存", reflect.TypeOf(reply))
		rsp.IndexBytes, _ = redis.Bytes(reply, nil)
		return nil
	}
	beego.Info("未从缓存中查询到数据，开始数据库查询")
	houses := []models.House{}
	o := orm.NewOrm()
	_, err = o.QueryTable("House").Limit(models.HOME_PAGE_MAX_HOUSES).All(&houses)
	if err != nil {
		beego.Info("数据库查询失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	for i, _ := range houses {
		o.LoadRelated(&houses[i], "Area")
		o.LoadRelated(&houses[i], "User")
		o.LoadRelated(&houses[i], "Images")
		o.LoadRelated(&houses[i], "Facilities")
	}

	beego.Info("数据库中查询到的数据", houses)
	// 将数据打包成前端所需要的格式
	// housesmap := map[string]interface{}{}
	// housesmap["houses"] = houses
	bytes, err := json.Marshal(houses)
	if err != nil {
		beego.Info("json编码失败", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	rsp.IndexBytes = bytes
	// 存入缓存todo

	err = bm.Put("home_page_data", bytes, time.Second*3600)
	if err != nil {
		beego.Info("缓存index失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	beego.Info("index缓存成功", err)

	return nil
}
