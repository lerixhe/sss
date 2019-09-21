package handler

import (
	"context"
	"encoding/json"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	DELETESESSION "sss/DeleteSession/proto/DeleteSession"
	GETAREA "sss/GetArea/proto/GetArea"
	GETHOUSEINFO "sss/GetHouseInfo/proto/GetHouseInfo"
	GETIMAGECD "sss/GetImageCd/proto/GetImageCd"
	GETINDEX "sss/GetIndex/proto/GetIndex"
	GETSESSION "sss/GetSession/proto/GetSession"
	GETSMSCD "sss/GetSmsCd/proto/GetSmsCd"
	GETUSERHOUSES "sss/GetUserHouses/proto/GetUserHouses"
	GETUSERINFO "sss/GetUserInfo/proto/GetUserInfo"
	POSTAVATAR "sss/PostAvatar/proto/PostAvatar"
	POSTHOUSES "sss/PostHouses/proto/PostHouses"
	POSTHOUSESIMAGE "sss/PostHousesImage/proto/PostHousesImage"
	POSTREG "sss/PostReg/proto/PostReg"
	POSTSESSION "sss/PostSession/proto/PostSession"
	POSTUSERAUTH "sss/PostUserAuth/proto/PostUserAuth"
	PUTUSERINFO "sss/PutUserInfo/proto/PutUserInfo"
	"sss/ihomeWeb/utils"
	"strconv"

	"github.com/afocus/captcha"

	"github.com/astaxie/beego"

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
		beego.Info("RPC错误")
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
		beego.Info("RPC错误")
		http.Error(w, err.Error(), 500)
		return
	}
	// 由于前端所需接口有个json,这里构造一下结构
	data := make(map[string]string)
	data["name"] = rsp.GetName()
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.ErrMsg,
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
	service := grpc.NewService()
	service.Init()
	getIndexService := GETINDEX.NewGetIndexService("go.micro.srv.GetIndex", service.Client())
	rsp, err := getIndexService.CallGetIndex(context.TODO(), &GETINDEX.Request{})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误", err)
		http.Error(w, err.Error(), 500)
		return
	}
	// 这里直接接收并返回，让服务端提前把格式处理好
	houses := []models.House{}
	bytes := rsp.GetIndexBytes()
	err = json.Unmarshal(bytes, &houses)
	if err != nil {
		beego.Info("json解码失败", err)
		http.Error(w, err.Error(), 500)
		return
	}
	beego.Info("获取到index数据", houses)
	// 构造前端可接受的json格式
	housesJSON := []interface{}{}
	for _, house := range houses {
		h := map[string]interface{}{
			"house_id":    house.Id,
			"title":       house.Title,
			"price":       house.Price,
			"area_name":   house.Area.Name,
			"img_url":     utils.AddDomain2Url(house.Index_image_url),
			"room_count":  house.Room_count,
			"order_count": house.Order_count,
			"address":     house.Address,
			"user_avatar": utils.AddDomain2Url(house.User.Avatar_url),
			"ctime":       house.Ctime.Format("2006-01-02 15:04:05"),
		}
		housesJSON = append(housesJSON, h)
	}
	housesDATA := make(map[string]interface{})
	housesDATA["houses"] = housesJSON
	response := map[string]interface{}{
		"errno":  rsp.GetError(),
		"errmsg": rsp.GetErrMsg(),
		"data":   housesJSON,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		beego.Info("encode失败：", err)
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

// 登录

func PostSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("用户登录 PostSession api/v1.0/session")
	// 获取客户端提交的表单
	data := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		beego.Info("表单解析失败", err)
		http.Error(w, err.Error(), 500)
		return
	}
	beego.Info("用户提交信息", data["mobile"], data["password"])
	// 取得数据并校验
	if data["mobile"] == "" || data["password"] == "" {
		beego.Info("表单数据不完整")
		http.Error(w, err.Error(), 500)
		return
	}

	// 调用微服务，回复是否通过登录验证
	service := grpc.NewService()
	service.Init()
	sessionService := POSTSESSION.NewPostSessionService("go.micro.srv.PostSession", service.Client())
	rsp, err := sessionService.CallPostSession(context.TODO(), &POSTSESSION.Request{
		Mobile:   data["mobile"],
		Password: data["password"],
	})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误", err)
		http.Error(w, err.Error(), 500)
		return
	}

	// 若通过验证，拿到服务回复中的sessioniD
	sessionID := rsp.GetSessionID()
	beego.Info("验证通过,开始查看本地cookies")
	// 读取cookie
	cookie, err := r.Cookie("userlogin")
	// 如果没读取到或者出错，则设置cookie
	if err != nil || cookie.Value == "" {
		beego.Info("本地cookies不存在，开始设置cookie")
		newcookie := http.Cookie{
			Name:   "userlogin",
			Value:  sessionID,
			Path:   "/",
			MaxAge: 600,
		}
		http.SetCookie(w, &newcookie)
		beego.Info("cookies设置成功")

	}
	response := map[string]interface{}{
		"errno":  rsp.GetError(),
		"errmsg": rsp.GetErrMsg(),
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

// 退出登录
func DeleteSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("用户退出登录 DeleteSession api/v1.0/session")
	// 从cookies中获取sessionID
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// 说明用户本没有登录，返回对应信息即可
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
	// 调用微服务，将服务器端session删除，并返回删除结果
	service := grpc.NewService()
	service.Init()
	sessionService := DELETESESSION.NewDeleteSessionService("go.micro.srv.DeleteSession", service.Client())
	rsp, err := sessionService.CallDeleteSession(context.TODO(), &DELETESESSION.Request{
		SessionID: cookie.Value,
	})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误")
		http.Error(w, err.Error(), 500)
		return
	}

	// 设置cookie失效,可以不做，但最好做
	newCookie := http.Cookie{
		Name:   "userlogin",
		Path:   "/",
		MaxAge: -1,
		Value:  "",
	}
	http.SetCookie(w, &newCookie)
	//返回给前端的数据
	response := map[string]interface{}{
		"errno":  rsp.GetError(),
		"errmsg": rsp.GetErrMsg(),
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

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
		beego.Info("RPC错误")
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
		beego.Info("RPC错误")
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
func PostReg(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("发送注册表单 PostReg /api/v1.0/users")
	// 获取web发来的表单(json)使用map来接收
	requestInfo := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&requestInfo)
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误")
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
		beego.Info("RPC错误")
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

// 获取用户信息
func GetUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取用户信息 GetUserInfo /api/v1.0/user")
	// 从cookies中获取sessionID
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// 说明用户本没有登录，返回对应信息即可
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			beego.Info("json转码错误")
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	// 调用微服务
	service := grpc.NewService()
	service.Init()
	getUserInfoService := GETUSERINFO.NewGetUserInfoService("go.micro.srv.GetUserInfo", service.Client())
	rsp, err := getUserInfoService.CallGetUserInfo(context.TODO(), &GETUSERINFO.Request{
		SessionID: cookie.Value,
	})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误")
		http.Error(w, err.Error(), 500)
		return
	}
	// 构造前端接受的data结构，接收rsp中的参数
	data := make(map[string]interface{})
	data["user_id"] = rsp.GetUserID()
	data["name"] = rsp.GetName()
	data["mobile"] = rsp.GetMobile()
	data["real_name"] = rsp.GetRealName()
	data["id_card"] = rsp.GetIDCard()
	data["avatar_url"] = utils.AddDomain2Url(rsp.GetAvatarUrl())

	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

// 上传用户头像
func PostAvatar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("上传用户头像 PostAvatar /api/v1.0/user/avatar")
	// 从r中接收图片流
	file, header, err := r.FormFile("avatar")
	if err != nil {
		// 直接给前端返回错误
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//输出日志
	beego.Info("文件大小：", header.Size)
	beego.Info("文件名称：", header.Filename)
	// 使用切片接收文件流
	filebuf := make([]byte, header.Size)
	_, err = file.Read(filebuf)
	if err != nil {
		// 直接给前端返回错误
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// 从cookies中获取sessionID
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// 说明用户本没有登录，返回对应信息即可
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
	// 开始调用微服务
	service := grpc.NewService()
	service.Init()
	postAvatarService := POSTAVATAR.NewPostAvatarService("go.micro.srv.PostAvatar", service.Client())
	rsp, err := postAvatarService.CallPostAvatar(context.TODO(), &POSTAVATAR.Request{
		SessionID: cookie.Value,
		Avatar:    filebuf,
		FileExt:   header.Filename,
		FileSize:  header.Size,
	})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误")
		http.Error(w, err.Error(), 500)
		return
	}
	// 构造前端接受的data结构，接收rsp中的参数
	data := make(map[string]interface{})
	data["avatar_url"] = utils.AddDomain2Url(rsp.GetAvatarUrl())

	// 给前端返回数据
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

// 更新用户名
func PutUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("更新用户名 PutUserInfo /api/v1.0/user/name")
	// 获取客户端提交的json表单
	data := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		beego.Info("表单解析失败", err)
		http.Error(w, err.Error(), 500)
		return
	}
	beego.Info("用户提交新用户名", data["name"])
	// 获取用户cookie中的sessionID
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// 说明用户本没有登录，返回对应信息即可
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

	// 开始调用微服务
	service := grpc.NewService()
	service.Init()
	putUserInfoService := PUTUSERINFO.NewPutUserInfoService("go.micro.srv.PutUserInfo", service.Client())
	rsp, err := putUserInfoService.CallPutUserInfo(context.TODO(), &PUTUSERINFO.Request{
		SessionID: cookie.Value,
		NewName:   data["name"],
	})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误")
		http.Error(w, err.Error(), 500)
		return
	}
	// 构造前端接受的data结构，接收rsp中的参数
	databack := make(map[string]interface{})
	databack["name"] = rsp.GetNewName()

	// 给前端返回数据
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.ErrMsg,
		"data":   databack,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

// 获取用户实名认证状态
func GetUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取用户实名认证状态 GetUserAuth /api/v1.0/user/auth")
	// 从cookies中获取sessionID
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// 说明用户本没有登录，返回对应信息即可
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			beego.Info("json转码错误")
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	// 调用微服务
	service := grpc.NewService()
	service.Init()
	getUserInfoService := GETUSERINFO.NewGetUserInfoService("go.micro.srv.GetUserInfo", service.Client())
	rsp, err := getUserInfoService.CallGetUserInfo(context.TODO(), &GETUSERINFO.Request{
		SessionID: cookie.Value,
	})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误")
		http.Error(w, err.Error(), 500)
		return
	}
	// 构造前端接受的data结构，接收rsp中的参数
	data := make(map[string]interface{})
	data["user_id"] = rsp.GetUserID()
	data["name"] = rsp.GetName()
	data["mobile"] = rsp.GetMobile()
	data["real_name"] = rsp.GetRealName()
	data["id_card"] = rsp.GetIDCard()
	data["avatar_url"] = utils.AddDomain2Url(rsp.GetAvatarUrl())

	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

// 发送进行实名认证请求
func PostUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("发送进行实名认证请求 PostUserAuth /api/v1.0/user/auth")
	// r中获取参数
	data := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		beego.Info("json解码错误")
		http.Error(w, err.Error(), 500)
		return
	}
	for k, v := range data {
		beego.Info(k, v)
	}
	// 从cookies中获取sessionID
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// 说明用户本没有登录，返回对应信息即可
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			beego.Info("json转码错误")
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// 数据校验
	// 是否为空
	if data["real_name"] == "" || data["id_card"] == "" {
		// 说明用户本没有登录，返回对应信息即可
		response := map[string]interface{}{
			"errno":  utils.RECODE_NODATA,
			"errmsg": utils.RecodeText(utils.RECODE_NODATA),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			beego.Info("json转码错误")
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	// 调用微服务
	service := grpc.NewService()
	service.Init()
	postUserAuthService := POSTUSERAUTH.NewPostUserAuthService("go.micro.srv.PostUserAuth", service.Client())
	rsp, err := postUserAuthService.CallPostUserAuth(context.TODO(), &POSTUSERAUTH.Request{
		SessionID: cookie.Value,
		RealName:  data["real_name"],
		IDCard:    data["id_card"],
	})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误")
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

// 获取用户房源
func GetUserHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取用户房源 GetUserHousers /api/v1.0/user/houses")
	// 从cookies中获取sessionID
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// 说明用户本没有登录，返回对应信息即可
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			beego.Info("json转码错误", err)
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	// 调用微服务
	service := grpc.NewService()
	service.Init()
	getUserHouses := GETUSERHOUSES.NewGetUserHousesService("go.micro.srv.GetUserHouses", service.Client())
	rsp, err := getUserHouses.CallGetUserHouses(context.TODO(), &GETUSERHOUSES.Request{
		SessionID: cookie.Value,
	})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误")
		http.Error(w, err.Error(), 500)
		return
	}
	//接收rsp中的二进制流,复杂数据类型传输，这里采用json的二进制流形式，
	data := rsp.GetMix()
	houseList := []models.House{}

	err = json.Unmarshal(data, &houseList)
	if err != nil {
		beego.Info("json转码错误2", err)
		http.Error(w, err.Error(), 500)
		return
	}
	for _, house := range houseList {
		beego.Info("请求到的房屋信息\n", house, house.Area)

	}
	// json接口文档要求格式为map包裹数组,这里得到的是个数组，
	// 需要构造一个map，存入其键houses中
	// 注意!!!不能直接回传[]models.House{}，此结构体中有指针，会丢失信息
	// 使用工具函数解决
	houseListJSON := []map[string]interface{}{}
	for _, house := range houseList {
		houseJSON := house.To_house_info()
		houseListJSON = append(houseListJSON, houseJSON)
	}
	houses := make(map[string]interface{})
	houses["houses"] = houseListJSON
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.ErrMsg,
		"data":   houses,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

// 发布房源请求
func PostHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("发布房源请求 PostHouses /api/v1.0/user/houses")
	// 从cookies中获取sessionID
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// 说明用户本没有登录，返回对应信息即可
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			beego.Info("json转码错误")
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	// 从r中获取post表单得到字节流
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		beego.Info("表单获取失败")
		http.Error(w, err.Error(), 500)
		return
	}
	// 调用微服务
	service := grpc.NewService()
	service.Init()
	postHouses := POSTHOUSES.NewPostHousesService("go.micro.srv.PostHouses", service.Client())
	rsp, err := postHouses.CallPostHouses(context.TODO(), &POSTHOUSES.Request{
		SessionID: cookie.Value,
		HouseInfo: body,
	})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误", err)
		http.Error(w, err.Error(), 500)
		return
	}
	data := map[string]string{}
	data["house_id"] = rsp.GetHousID()
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

// 上传房源照片请求
func PostHousesImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("上传房源照片请求 PostAvatar/api/v1.0/houses/:id/images")
	// 从url接收房屋id
	houseID := ps.ByName("id")
	beego.Info("照片所属房屋id", houseID)
	// 从r中接收图片流
	file, header, err := r.FormFile("house_image")
	if err != nil {
		// 直接给前端返回错误
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//输出日志
	beego.Info("文件大小：", header.Size)
	beego.Info("文件名称：", header.Filename)
	// 使用切片接收文件流
	filebuf := make([]byte, header.Size)
	_, err = file.Read(filebuf)
	if err != nil {
		// 直接给前端返回错误
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	// 从cookies中获取sessionID
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// 说明用户本没有登录，返回对应信息即可
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
	// 开始调用微服务
	service := grpc.NewService()
	service.Init()
	postHousesImage := POSTHOUSESIMAGE.NewPostHousesImageService("go.micro.srv.PostHousesImage", service.Client())
	rsp, err := postHousesImage.CallPostHousesImage(context.TODO(), &POSTHOUSESIMAGE.Request{
		SessionID: cookie.Value,
		HouseID:   houseID,
		Image:     filebuf,
		FileName:  header.Filename,
		FileSize:  header.Size,
	})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误")
		http.Error(w, err.Error(), 500)
		return
	}
	// 构造前端接受的data结构，接收rsp中的参数
	data := make(map[string]interface{})
	data["url"] = utils.AddDomain2Url(rsp.GetFiled())

	// 给前端返回数据
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}

// 获取房源详细信息
func GetHouseInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取房源详细信息 GetHouseInfo /api/v1.0/houses/:id")
	// 获取房屋id
	houseID := ps.ByName("id")
	// 从cookies中获取sessionID
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// 说明用户本没有登录，返回对应信息即可
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			beego.Info("json转码错误")
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	// 调用微服务
	service := grpc.NewService()
	service.Init()
	getHouseInfo := GETHOUSEINFO.NewGetHouseInfoService("go.micro.srv.GetHouseInfo", service.Client())
	rsp, err := getHouseInfo.CallGetHouseInfo(context.TODO(), &GETHOUSEINFO.Request{
		SessionID: cookie.Value,
		HouseID:   houseID,
	})
	// 若发生错误
	if err != nil {
		beego.Info("RPC错误")
		http.Error(w, err.Error(), 500)
		return
	}
	//接收rsp中的房屋详情信息流
	data := rsp.GetHouseInfoBytes()
	house := new(models.House)

	err = json.Unmarshal(data, &house)
	if err != nil {
		beego.Info("json转码错误")
		http.Error(w, err.Error(), 500)
		return
	}
	beego.Info("请求到的房屋信息\n", house.User, house.Area, house.Facilities)
	houses := make(map[string]interface{})
	houses["house"] = house.To_one_house_desc()
	userID, _ := strconv.Atoi(rsp.GetUserID())
	response := map[string]interface{}{
		"errno":   rsp.Error,
		"errmsg":  rsp.ErrMsg,
		"data":    houses,
		"user_id": userID,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	return
}
