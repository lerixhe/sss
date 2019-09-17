package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"sss/ihomeWeb/models"
	"sss/ihomeWeb/utils"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"

	GETSMSCD "sss/GetSmsCd/proto/GetSmsCd"

	"github.com/astaxie/beego"
)

type GetSmsCd struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetSmsCd) CallGetSmsCd(ctx context.Context, req *GETSMSCD.Request, rsp *GETSMSCD.Response) error {
	beego.Info("获取短信验证码 GetSmsCode /api/v1.0/smscode/:mobile")
	// 获取web传来的参数
	mobile := req.GetMobile()
	text := req.GetText()
	uuid := req.GetUuid()
	// 开始校验
	// 1.图片验证码校验
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
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	reply := bm.Get(uuid)
	if reply == nil {
		beego.Info("缓存查询结果为空")
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	str, err := redis.String(reply, err)
	if str != text {
		beego.Info("图片验证码错误", str, "?", text)
		rsp.Error = utils.RECODE_IMAGECDERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	/* 2.手机号是否已注册 */
	o := orm.NewOrm()
	user := &models.User{Mobile: mobile}
	err = o.Read(user, "Mobile")
	if err == nil {
		// 说明找到了对应用户
		beego.Info("该用户已经注册")
		rsp.Error = utils.RECODE_ALREADREG
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	} else {
		beego.Info("该用户未注册开始注册")
	}
	/* 数据校验完毕，开始短信验证流程 */
	// 	用户登录名称 sms@1909310674080489.onaliyun.com
	// AccessKey ID LTAI4Fmft2zgUcC8FQNHeA8m
	// AccessKeySecret fnu82dTJXYW94yTYLXM8eUn4xlflbi
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4Fmft2zgUcC8FQNHeA8m", "fnu82dTJXYW94yTYLXM8eUn4xlflbi")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return nil
}
