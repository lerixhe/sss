package handler

import (
	"context"
	"encoding/json"
	"image"
	"image/png"
	"net/http"
	GETIMAGECD "sss/GetImageCd/proto/GetImageCd"
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
		"errno":  utils.RECODE_SESSIONERR,
		"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
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
