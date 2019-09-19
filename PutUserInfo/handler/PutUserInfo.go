package handler

import (
	"context"
	"encoding/json"
	"reflect"
	"sss/ihomeWeb/models"
	"sss/ihomeWeb/utils"

	PUTUSERINFO "sss/PutUserInfo/proto/PutUserInfo"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

type PutUserInfo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *PutUserInfo) CallPutUserInfo(ctx context.Context, req *PUTUSERINFO.Request, rsp *PUTUSERINFO.Response) error {
	beego.Info("更新用户名 PutUserInfo /api/v1.0/user/name")
	// 初始化回复
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 获取请求参数
	sessionID := req.GetSessionID()
	newName := req.GetNewName()

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
	// 通过用户id查询用户所有数据
	user := models.User{Id: id, Name: newName}
	o := orm.NewOrm()
	o.Update(&user, "Name")
	beego.Info(user)
	// 更新完毕，返回新用户名
	rsp.NewName = newName
	beego.Info("用户名更新完毕：", newName)
	return nil
}
