package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"sss/ihomeWeb/utils"

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
