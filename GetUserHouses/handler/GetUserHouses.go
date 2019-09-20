package handler

import (
	"context"
	"encoding/json"
	"reflect"
	"sss/ihomeWeb/models"

	"github.com/astaxie/beego/orm"

	GETUSERHOUSES "sss/GetUserHouses/proto/GetUserHouses"
	"sss/ihomeWeb/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

type GetUserHouses struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetUserHouses) CallGetUserHouses(ctx context.Context, req *GETUSERHOUSES.Request, rsp *GETUSERHOUSES.Response) error {
	beego.Info("获取用户房源 GetUserHousers /api/v1.0/user/houses")
	// 初始化回复
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 获取请求参数
	sessionID := req.GetSessionID()

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
		beego.Info("缓存查询失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 验证sessionID，并得到id
	reply := bm.Get(sessionID + "user_id")
	if reply == nil {
		beego.Info("缓存查询结果为空")
		rsp.Error = utils.RECODE_NODATA
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	beego.Info(reply, reflect.TypeOf(reply))
	id, err := redis.Int(reply, nil)
	if err != nil {
		beego.Info("缓存数据类型错误", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	houseList := []models.House{}
	// 根据id查询house与user表，得到该id下的house列表
	o := orm.NewOrm()
	_, err = o.QueryTable("House").Filter("User", id).All(&houseList)
	if err != nil {
		beego.Info("房屋数据查询失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	beego.Info("该用户下的房屋列表：")
	for i, v := range houseList {
		beego.Info(i, v)
	}
	// 转为json
	data, err := json.Marshal(houseList)
	rsp.Mix = data
	return nil
}
