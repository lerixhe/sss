package handler

import (
	"context"
	"encoding/json"
	DELETESESSION "sss/DeleteSession/proto/DeleteSession"
	"sss/ihomeWeb/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
)

type DeleteSession struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *DeleteSession) CallDeleteSession(ctx context.Context, req *DELETESESSION.Request, rsp *DELETESESSION.Response) error {
	beego.Info("用户退出登录 DeleteSession api/v1.0/session")
	// 初始化rsp
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 获取请求中的参数
	sessionID := req.GetSessionID()
	// 链接redis
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
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 删除缓存
	bm.Delete(sessionID + "user_id")
	bm.Delete(sessionID + "user_name")
	bm.Delete(sessionID + "user_mobile")

	return nil
}
