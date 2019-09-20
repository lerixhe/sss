package handler

import (
	"context"
	"encoding/json"
	"path"
	"reflect"
	POSTAVATAR "sss/PostAvatar/proto/PostAvatar"
	"sss/ihomeWeb/models"
	"sss/ihomeWeb/utils"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"

	"github.com/astaxie/beego"
)

type PostAvatar struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *PostAvatar) CallPostAvatar(ctx context.Context, req *POSTAVATAR.Request, rsp *POSTAVATAR.Response) error {
	beego.Info("上传用户头像 PostAvatar /api/v1.0/user/avatar")
	// 初始化rsp
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 获取req参数，并校验
	filebuf := req.GetAvatar()
	fileName := req.GetFileExt()
	fileSize := req.GetFileSize()
	sessionID := req.GetSessionID()
	// 文件完整性校验
	if len(filebuf) != int(fileSize) {
		beego.Info("传输数据丢失")
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	// 获取文件拓展名,不要点
	ext := path.Ext(fileName)[1:]
	// 校验完毕，开始上传
	filed, err := utils.UploadByBuffer(filebuf, ext)
	beego.Info("图片上传返回结果：", filed)

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
	/// 通过用户id更新用户头像字段数据
	user := models.User{
		Id:         id,
		Avatar_url: filed,
	}
	o := orm.NewOrm()
	_, err = o.Update(&user, "Avatar_url")
	if err != nil {
		beego.Info("数据库更新错误", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	beego.Info("头像url已存储")
	// 返回前端数据
	rsp.AvatarUrl = filed
	return nil
}
