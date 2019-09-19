package handler

import (
	"context"
	"encoding/json"
	POSTREG "sss/PostReg/proto/PostReg"
	"sss/ihomeWeb/models"
	"sss/ihomeWeb/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"

	_ "github.com/astaxie/beego/cache/redis"
	"github.com/gomodule/redigo/redis"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

type PostReg struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *PostReg) CallPostReg(ctx context.Context, req *POSTREG.Request, rsp *POSTREG.Response) error {
	beego.Info("发送注册表单 PostReg /api/v1.0/users")
	// 初始化返回
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 获取请求中的参数
	mobile := req.GetMobile()
	password := req.GetPassword()
	smsCode := req.GetSmsCode()
	// get缓存
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
	// 取出验证码信息
	reply := bm.Get(mobile)
	// 	测试账号绿色通道.1~100，短信验证码固定为123
	if i, _ := strconv.Atoi(mobile); i <= 100 && i >= 1 {
		reply = "123"
	}
	if reply == nil {
		beego.Info("缓存查询结果为空:验证码1")
		rsp.Error = utils.RECODE_NODATA
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	str, _ := redis.String(reply, nil)
	if str != smsCode {
		beego.Info("短信验证码错误", str, "?", smsCode)
		rsp.Error = utils.RECODE_SMSCDERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	//短信通过验证，开始注册用户
	// 1.构建用户数据结构
	user := models.User{
		Name:          mobile,
		Password_hash: utils.GetMd5String(password),
		Mobile:        mobile,
	}
	// 2.链接并操作数据库
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		beego.Info("数据库插入用户失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	beego.Info("用户注册成功，用户ID:", id)
	// 给用户生成SessionID并返回
	sessionID := utils.GetMd5String(mobile + password)
	rsp.SessionID = sessionID
	// 用户信息存入缓存，格式key：sessionID+用户字段 vaule：字段值
	bm.Put(sessionID+"user_id", id, time.Second*3600)
	bm.Put(sessionID+"user_name", mobile, time.Second*3600)
	bm.Put(sessionID+"user_mobile", mobile, time.Second*3600)
	return nil
}
