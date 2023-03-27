package Controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var ErrorUserNotLogin = errors.New("用户未登录")

const CtxUserIDKey = "userID"
const CtxUserPermissions = "Userpermissions"

// getCurrentUserID 获取当前登录的用户的id和权限
func getCurrentUserID(c *gin.Context) (userID int64, userPerm int8, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	uperm, ok1 := c.Get(CtxUserPermissions)
	userID, ok = uid.(int64)
	userPerm, ok = uperm.(int8)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	if !ok1 {
		err = ErrorUserNotLogin
		return
	}

	if !ok1 {
		err = ErrorUserNotLogin
		return
	}
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// GetPageInfo 获取当前分页信息
func GetPageInfo(c *gin.Context) (int64, int64, error) {
	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")

	var (
		pageNum  int64
		pageSize int64
		err      error
	)

	pageNum, err = strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		pageNum = 1
	}

	pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 10
	}
	return pageNum - 1, pageSize, err
}
