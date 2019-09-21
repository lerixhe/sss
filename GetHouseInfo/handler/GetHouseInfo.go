package handler

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	_ "github.com/astaxie/beego/cache/redis"

	_ "github.com/gomodule/redigo/redis"

	GETHOUSEINFO "sss/GetHouseInfo/proto/GetHouseInfo"
	"sss/ihomeWeb/models"
	"sss/ihomeWeb/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
)

type GetHouseInfo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetHouseInfo) CallGetHouseInfo(ctx context.Context, req *GETHOUSEINFO.Request, rsp *GETHOUSEINFO.Response) error {
	beego.Info("获取房源详细信息 GetHouseInfo /api/v1.0/houses/:id")
	// 初始化回复
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 获取请求参数
	sessionID := req.GetSessionID()
	houseID, _ := strconv.Atoi(req.GetHouseID())
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
		beego.Info("redis连接失败", err)
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
	userID, err := redis.Int(reply, nil)
	if err != nil {
		beego.Info("缓存数据类型错误", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	rsp.UserID = strconv.Itoa(userID)
	// 尝试在redis中直接获取房屋信息key=house_info_{houseID}
	reply = bm.Get("house_info_" + req.GetHouseID())
	if reply != nil {
		// 查询到，则直接返回并结束任务
		beego.Info("缓存查询到数据！")
		houseInfoBytes, err := redis.Bytes(reply, nil)
		if err != nil {
			beego.Info("缓存数据类型错误", err)
			rsp.Error = utils.RECODE_DATAERR
			rsp.ErrMsg = utils.RecodeText(rsp.Error)
			return err
		}
		rsp.HouseInfoBytes = houseInfoBytes
		return nil
	}
	beego.Info("缓存未查询到数据，开始查询数据库！")
	// 数据库查询房屋数据
	// 根据id查询house与user表，得到该id下的house列表
	// 注意需要关联查询，才能完整查到对应的area信息。
	house := models.House{Id: houseID}
	o := orm.NewOrm()
	err = o.Read(&house)
	// 使用LoadReated进行关联查询
	_, err = o.LoadRelated(&house, "Area")
	_, err = o.LoadRelated(&house, "User")
	_, err = o.LoadRelated(&house, "Images")
	_, err = o.LoadRelated(&house, "Facilities")
	if err != nil {
		beego.Info("房屋数据库查询错误", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	beego.Info("房屋信息数据库读出完毕", house)
	// 关联查询完毕，house中存储的信息已经相对完整，开始序列化
	houseInfoBytes, err := json.Marshal(house)
	if err != nil {
		beego.Info("房屋信息数据序列化错误", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 先不发送，存储到redis缓存，方便下次下旬
	err = bm.Put("house_info_"+req.GetHouseID(), houseInfoBytes, time.Second*3600)
	if err != nil {
		beego.Info("redis存储房屋缓存失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	beego.Info("房屋信息成功存储到缓存")
	// 发送数据
	rsp.HouseInfoBytes = houseInfoBytes
	return nil
}
