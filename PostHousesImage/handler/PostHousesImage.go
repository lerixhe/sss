package handler

import (
	"context"
	"encoding/json"
	"path"
	"reflect"
	"sss/ihomeWeb/models"
	"strconv"

	"github.com/astaxie/beego/orm"

	POSTHOUSESIMAGE "sss/PostHousesImage/proto/PostHousesImage"
	"sss/ihomeWeb/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/gomodule/redigo/redis"
)

type PostHousesImage struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *PostHousesImage) CallPostHousesImage(ctx context.Context, req *POSTHOUSESIMAGE.Request, rsp *POSTHOUSESIMAGE.Response) error {
	beego.Info("上传房源照片请求 PostAvatar/api/v1.0/houses/:id/images")
	// 初始化回复
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 获取请求参数
	fileExt := path.Ext(req.GetFileName())[1:]
	fileSize := req.GetFileSize()
	houseID, _ := strconv.Atoi(req.GetHouseID())
	image := req.GetImage()
	sessionID := req.GetSessionID()
	// 取得用户id，主要为验证上传者身份是否为房屋拥有者
	// 读取redis链接配置
	redisConf := map[string]string{
		"key":      utils.G_server_name,
		"conn":     utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum":    utils.G_redis_dbnum,
		"password": utils.G_redis_auth,
	}
	// 将map转换为json
	redisConfJSON, _ := json.Marshal(redisConf)
	// 开始链接redis
	bm, err := cache.NewCache("redis", string(redisConfJSON))
	if err != nil {
		beego.Info("缓存查询失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	reply := bm.Get(sessionID + "user_id")
	if reply == nil {
		beego.Info("缓存查询结果为空")
		rsp.Error = utils.RECODE_NODATA
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	beego.Info(reply, reflect.TypeOf(reply))
	id, err := redis.Int(reply, nil)
	if err != nil {
		beego.Info("缓存数据类型错误", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	beego.Info("用户id,session", id, sessionID)

	// 文件完整性校验
	if len(image) != int(fileSize) {
		beego.Info("传输数据丢失")
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	// 校验完毕，开始上传
	filed, _ := utils.UploadByBuffer(image, fileExt)
	beego.Info("图片上传返回结果：", filed)
	// 根据houseID取得house信息
	house := models.House{Id: houseID}
	o := orm.NewOrm()
	err = o.Read(&house)
	if err != nil {
		beego.Info("未找到房屋信息")
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	beego.Info("读取到房屋信息", house)
	if house.User.Id != id {
		beego.Info("目标房屋不属于当前登录身份")
		rsp.Error = utils.RECODE_ROLEERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	// 查看是否存在主图,没有则设置为主图
	if house.Index_image_url == "" {
		house.Index_image_url = filed
		_, err = o.Update(&house, "Index_image_url")
		if err != nil {
			beego.Info("设置主图失败")
			rsp.Error = utils.RECODE_DBERR
			rsp.ErrMsg = utils.RecodeText(rsp.Error)
			return err
		}
	}
	// 把图片加入，图片表中
	houseImage := models.HouseImage{
		House: &house,
		Url:   filed,
	}
	_, err = o.Insert(&houseImage)
	if err != nil {
		beego.Info("插入到图片表失败")
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 对house表更新
	house.Images = append(house.Images, &houseImage)
	_, err = o.Update(&house)
	if err != nil {
		beego.Info("更新house表失败")
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	rsp.Filed = filed
	return nil
}
