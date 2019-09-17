package handler

import (
	"context"
	GETAREA "sss/GetArea/proto/GetArea"
	"sss/ihomeWeb/models"
	"sss/ihomeWeb/utils"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

type GetArea struct{}

// // Call is a single request handler called via client.Call or the generated client code
// func (e *GetArea) Call(ctx context.Context, req *GetArea.Request, rsp *GetArea.Response) error {
// 	log.Log("Received GetArea.Call request")
// 	rsp.Msg = "Hello " + req.Name
// 	return nil
// }

// 服务所属方法的具体定义
func (e *GetArea) GetAreas(ctx context.Context, req *GETAREA.Request, rsp *GETAREA.Response) error {
	beego.Info("请求地区信息：GetArea api.v1.0/areas")
	// 初始化错误码，默认成功信息
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)
	/* 从缓存中获取数据，有数据则直接返回给前端
	没有数据，则去数据库查找数据
	将数据存到缓存中
	把数据发送给前端。
	*/

	// 创建orm句柄
	o := orm.NewOrm()

	// 查什么？用什么接收
	areas := []models.Area{}
	num, err := o.QueryTable("area").All(&areas)
	if err != nil {
		beego.Info("数据库查询失败", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return err
	}
	if num == 0 {
		beego.Info("没有查询到数据")
		rsp.Error = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	// 将查询到的数据，按照proto的格式发送给web
	for key, value := range areas {
		beego.Info(key, value)
		tmp := GETAREA.Response_Areas{Aid: int32(value.Id), Aname: value.Name}
		rsp.Data = append(rsp.Data, &tmp)
	}
	return nil
}
