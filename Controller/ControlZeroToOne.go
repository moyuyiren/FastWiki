package Controller

import (
	"FastWiki/Controller/reciprocal"
	"FastWiki/Logic"
	"FastWiki/Model"
	"github.com/gin-gonic/gin"
)

// ControlZeroToOne 刷新服务器Redis列表相关缓存
func ControlZeroToOne(c *gin.Context) {
	//获取用户token信息
	userid, userPerm, err := getCurrentUserID(c)
	if err != nil {
		reciprocal.ResponseError(c, reciprocal.CodeNeedLogin)
		return
	}
	if userPerm < 3 {
		reciprocal.ResponseError(c, reciprocal.CodeUserPermissionDenied)
		return
	}
	user := new(Model.Userpermissions)
	user.Userid = userid
	user.Userpermissions = userPerm

	//传入Logic进行判断
	if err := Logic.FlushRedisMessageForMain(user); err != nil {
		reciprocal.ResponseError(c, reciprocal.CodeServerBusy)
		return
	}
	reciprocal.ResponseSuccess(c, "Redis刷新完成")

}
