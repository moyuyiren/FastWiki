package RouterS

import (
	"FastWiki/Controller"
	"FastWiki/Logger"
	"FastWiki/Middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	//Gin Work Mod
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(Logger.GinLogger(), Logger.GinRecovery(true))
	/*无需登录*/
	//首页内容=====================================================================================

	/*需要登录*/
	User := r.Group("/api/Gcot")
	//注册And登录And修改密码
	User.POST("/signup", Controller.SignUpHandler)
	User.POST("/index", Controller.LoginUser)
	User.POST("/ExchangePassword", Controller.ExchangePassword)

	//需要登录
	User.Use(Middlewares.JWTAuthMiddleWare())
	{
		//管理员操作
		//启动后由管理员手动刷新Redis缓存
		User.PUT("/ZeroToOne", Controller.ControlZeroToOne)
		User.GET("/GetAllUserPremMessage", Controller.GetAllUserMessage) //管理员查询所有用户权限
		User.GET("/GetAllUserPremMessage", Controller.GetOneUserMessage) //管理员查询单个用户权限
		User.POST("/SetUserLevel", Controller.SetUserPremLevel)          //管理员设置用户权限

		/*=============================================================================================================================================*/
		//用户创建词条
		User.POST("/post", Controller.CreatePostHandler)
		/*=============================================================================================================================================*/

	}

	//测试
	r.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "Link start",
		})
	})
	return r

}
