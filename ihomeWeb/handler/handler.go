package handler

import (
	"context"
	"encoding/json"
	"image"
	"image/png"
	"net/http"
	"reflect"
	"regexp"
	GETIMAGECD "sss/GetImageCd/proto/GetImageCd"
	GETSESSION "sss/GetSession/proto/GetSession"
	GETSMSCD "sss/GetSmsCd/proto/GetSmsCd"
	POSTREG "sss/PostReg/proto/PostReg"
	"sss/ihomeWeb/utils"

	"github.com/afocus/captcha"

	"github.com/astaxie/beego"

	GETAREA "sss/GetArea/proto/GetArea"

	"sss/ihomeWeb/models"

	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro/service/grpc"
)

// func IhomeWebCall(w http.ResponseWriter, r *http.Request) {
// 	// decode the incoming request as json
// 	var request map[string]interface{}
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// }

// 调用远程方法的函数:获取地址
func GetArea(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取地址列表 GetArea api/v1.0/areas")
	service := grpc.NewService()
	service.Init()
	areaService := GETAREA.NewGetAreaService("go.micro.srv.GetArea", service.Client())
	res, err := areaService.GetAreas(context.TODO(), &GETAREA.Request{})
	// 若发生错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// 接收数据
	/* 1.准备一个切片  2.读取回复中的data部分*/
	areaList := []models.Area{}
	for _, value := range res.Data {
		temp := models.Area{Id: int(value.Aid), Name: value.Aname}
		areaList = append(areaList, temp)
	}
	response := map[string]interface{}{
		"errno":  res.Error,
		"errmsg": res.Errmsg,
		"data":   areaList,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

// 调用远程方法的函数：获取session
func GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取登录状态 GetISession api/v1.0/session")
	// 取出cookies
	cookie, err := r.Cookie("userlogin")
	if err != nil {
		// 用户未登录，直接返回
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	// 调用微服务
	service := grpc.NewService()
	service.Init()
	sessionService := GETSESSION.NewGetSessionService("go.micro.srv.GetSession", service.Client())
	rsp, err := sessionService.CallGetSession(context.TODO(), &GETSESSION.Request{
		SessionID: cookie.Value,
	})

	// 若发生错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// 由于前端所需接口有个json,这里构造一下结构
	data := make(map[string]string)
	data["name"] = rsp.GetName()
	response := map[string]interface{}{
		"errno":  rsp.GetError(),
		"errmsg": rsp.GetErrMsg(),
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

// 调用远程方法的函数:获取首页轮播图
func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取首页轮播图 GetIndex api/v1.0/house/index")

	// service := grpc.NewService()
	// service.Init()
	// areaService := GETAREA.NewGetAreaService("go.micro.srv.GetArea", service.Client())
	// res, err := areaService.GetAreas(context.TODO(), &GETAREA.Request{})
	// // 若发生错误
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	response := map[string]interface{}{
		"errno":  utils.RECODE_DBERR,
		"errmsg": utils.RecodeText(utils.RECODE_DBERR),
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

// 调获取验证码图片
func GetImageCode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取验证码图片 GetImageCode api/v1.0/imagecode/:uuid")
	// 获取uuid
	uuid := ps.ByName("uuid")
	service := grpc.NewService()
	service.Init()
	imageCdService := GETIMAGECD.NewGetImageCdService("go.micro.srv.GetImageCd", service.Client())
	rsp, err := imageCdService.CallGetImageCd(context.TODO(), &GETIMAGECD.Request{
		Uuid: uuid,
	})
	// 若发生错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// 若成功，则收到rsp为一堆图片零件，这里用图片对象接收
	// 注意取得指针内容后赋值
	var img image.RGBA
	img.Pix = rsp.Pix
	img.Stride = int(rsp.Stride)
	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)
	var captchaImg captcha.Image
	captchaImg.RGBA = &img
	w.Header().Set("Content-Type", "application/png")
	// 将图片发送给浏览器
	png.Encode(w, captchaImg)

	// response := map[string]interface{}{
	// 	"errno":  utils.RECODE_DBERR,
	// 	"errmsg": utils.RecodeText(utils.RECODE_DBERR),
	// }
	// w.Header().Set("Content-Type", "application/json")
	// if err := json.NewEncoder(w).Encode(response); err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	// return
}

// 调获取短信验证码
func GetSmsCode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取短信验证码 GetSmsCode /api/v1.0/smscode/:mobile")
	// 获取手机号,由于是直接使用：接收的，故可直接通过Name获取
	mobile := ps.ByName("mobile")
	// 获取url中的参数部分，？后面的都不属于url的部分，而是其携带的参数
	// api/v1.0/smscode/22222?text=2222&id=a992d1e5-fc77-4963-a0fe-d52693469c5c
	beego.Info("URL参数：", r.URL.Query())
	// 获取输入的图片验证码、uuid
	text := r.URL.Query()["text"][0]
	uuid := r.URL.Query()["id"][0]
	// 初步数据校验:手机号格式校验
	//创建正则句柄
	myreg := regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)
	//进行正则匹配
	bo := myreg.MatchString(mobile)
	// 手机号验证不通过，则直接返回相应从错误
	if !bo {
		response := map[string]interface{}{
			"errno":  utils.RECODE_MOBILEERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	service := grpc.NewService()
	service.Init()
	smsCdSercice := GETSMSCD.NewGetSmsCdService("go.micro.srv.GetSmsCd", service.Client())
	rsp, err := smsCdSercice.CallGetSmsCd(context.TODO(), &GETSMSCD.Request{
		Mobile: mobile,
		Uuid:   uuid,
		Text:   text,
	})

	// 若发生错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.ErrMsg,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

// 调用发送注册表单函数
func PostReg(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("发送注册表单 PostReg /api/v1.0/users")
	// 获取web发来的表单(json)使用map来接收
	requestInfo := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&requestInfo)
	// 若发生错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// 遍历获取到的参数
	for key, value := range requestInfo {
		beego.Info(key + ":" + value + ":" + reflect.TypeOf(value).String())
	}
	// 数据基本校验：验空,失败则直接返回错误信息，并结束。
	if requestInfo["mobile"] == "" || requestInfo["password"] == "" || requestInfo["sms_code"] == "" {
		response := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	// 校验完数据，开始GRPC调用微服务
	service := grpc.NewService()
	service.Init()
	postRegService := POSTREG.NewPostRegService("go.micro.srv.PostReg", service.Client())
	rsp, err := postRegService.CallPostReg(context.TODO(), &POSTREG.Request{
		Mobile:   requestInfo["mobile"],
		Password: requestInfo["password"],
		SmsCode:  requestInfo["sms_code"],
	})
	// 若发生错误
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.ErrMsg,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}
