package handler

import (
	"context"
	"encoding/json"
	POSTSESSION "sss/PostSession/proto/PostSession"
	"sss/ihomeWeb/models"
	"sss/ihomeWeb/utils"
	"time"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/gomodule/redigo/redis"

	"github.com/astaxie/beego"
)

type PostSession struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *PostSession) CallPostSession(ctx context.Context, req *POSTSESSION.Request, rsp *POSTSESSION.Response) error {
	beego.Info("用户登录 PostSession api/v1.0/session")
	// 初始化回复
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 获取请求参数
	mobile := req.GetMobile()
	password := req.GetPassword()
	// 构造数据结构
	user := models.User{Mobile: mobile}
	// 数据库中查询
	o := orm.NewOrm()
	// err := o.Read(&user, "Mobile")
	// if err != nil {
	// 	beego.Info("数据库未找到用户", err)
	// 	rsp.Error = utils.RECODE_NODATA
	// 	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 	return err
	// }
	err := o.QueryTable("User").Filter("Mobile", mobile).One(&user)
	if err != nil {
		beego.Info("数据库未找到用户，该用户可能没有注册", err)
		rsp.Error = utils.RECODE_NODATA
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 找到了用户记录,需要验证密码
	if utils.GetMd5String(password) != user.Password_hash {
		beego.Info("密码错误")
		rsp.Error = utils.RECODE_PWDERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	// 验证通过，则用户状态写入缓存
	// 1.链接redis
	// 2.得到sessionID
	// 3.构造数据后存储到redis
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
		beego.Info("缓存链接失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	sessionID := utils.GetMd5String(mobile + password)
	rsp.SessionID = sessionID
	// 用户信息存入缓存，格式key：sessionID+用户字段 vaule：字段值
	// 可能没有设置用户名，则设置为手机号
	if user.Name == "" {
		user.Name = user.Mobile
	}
	bm.Put(sessionID+"user_id", user.Id, time.Second*3600)
	bm.Put(sessionID+"user_name", user.Name, time.Second*3600)
	bm.Put(sessionID+"user_mobile", user.Mobile, time.Second*3600)

	return nil
}
