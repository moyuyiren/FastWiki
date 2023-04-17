package Controller

import (
	"FastWiki/Controller/reciprocal"
	"FastWiki/Logic"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler 创建文档
func CreatePostHandler(c *gin.Context) {
	userid, userPerm, err := getCurrentUserID(c)
	if err != nil {
		reciprocal.ResponseError(c, reciprocal.CodeNeedLogin)
		return
	} else if userPerm < 1 {
		reciprocal.ResponseError(c, reciprocal.CodeUserPermissionDenied)
	}
	fmt.Println(userid)

	//传入逻辑层
	if err := Logic.UserCreatePost(); err != nil {
		zap.L().Error("User Create Post Fail", zap.Error(err))
		reciprocal.ResponseError(c, reciprocal.CodeServerBusy)
	}
	reciprocal.ResponseSuccess(c, "新建文档成功")

}

// GetMessage 查询文档
func GetMessage(c *gin.Context) {
	substance := c.Query("substance")
	if substance == "" {
		reciprocal.ResponseError(c, reciprocal.CodeInvalidParams)
	}
	//model.
	//reciprocal.ResponseSuccess(c)

}
