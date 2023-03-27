package Middlewares

import (
	"FastWiki/Controller"
	"FastWiki/Controller/reciprocal"
	"FastWiki/Token"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleWare 认证中间件
func JWTAuthMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeInvalidToken, "请求头缺少Auth Token")
			c.Abort()
			return
		}

		//按空格分割

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeInvalidToken, "Token格式不对")
			c.Abort()
			return
		}
		//获取token后使用pasreToken进行解析
		mc, err := Token.ParseToken(parts[1])
		if err != nil {
			reciprocal.ResponseError(c, reciprocal.CodeNeedLogin)
			c.Abort()
			return
		}
		c.Set(Controller.CtxUserIDKey, mc.Userid)
		c.Set(Controller.CtxUserPermissions, mc.Userpermissions)
		//将数据写入请求上下文
		c.Next()

	}
}
