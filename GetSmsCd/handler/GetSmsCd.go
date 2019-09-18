package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sss/ihomeWeb/models"
	"sss/ihomeWeb/utils"
	"strconv"
	"time"

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
	// 初始化返回信息为成功，对于用户来讲，即使出现某些错误信息，也不该全盘展现，尽量展现无感知的默认成功结果。
	// 除非需要，才显示对应的错误信息
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
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
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	reply := bm.Get(uuid)
	if reply == nil {
		beego.Info("缓存查询结果为空")
		rsp.Error = utils.RECODE_NODATA
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	str, _ := redis.String(reply, nil)
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
		beego.Info("校验成功，开始发送验证码")
	}
	/* 数据校验完毕，开始短信验证流程 */
	// 1.生成验证码
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	SMScode := r.Intn(9999) + 1001 //1001~10999
	// 2.发送到用户手机
	// 	用户登录名称 sms@1909310674080489.onaliyun.com
	// AccessKey ID LTAI4Fmft2zgUcC8FQNHeA8m
	// AccessKeySecret fnu82dTJXYW94yTYLXM8eUn4xlflbi
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4Fmft2zgUcC8FQNHeA8m", "fnu82dTJXYW94yTYLXM8eUn4xlflbi")
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = mobile
	request.SignName = "代码研习社"
	request.TemplateCode = "SMS_174270094"
	request.TemplateParam = "{\"code\":\"" + strconv.Itoa(SMScode) + "\"}"
	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	if response.IsSuccess() {
		beego.Info("验证码发送成功")
		// 缓存验证码
		err := bm.Put(mobile, SMScode, time.Second*300)
		if err != nil {
			beego.Info("缓存验证码失败")
			return err
		}
	} else {
		// 短信发送不成功，返回错误信息
		rsp.Error = utils.RECODE_SMSERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
	}
	return nil
}
