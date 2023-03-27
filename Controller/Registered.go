package Controller

import (
	"FastWiki/Controller/reciprocal"
	"FastWiki/Logic"
	"FastWiki/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 用户注册默认为一般用户
func SignUpHandler(c *gin.Context) {
	//获取用户传入参数
	p := new(Model.GetUserInfo)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("用户注册参数解析失败", zap.Error(err))
		//判断错误类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			reciprocal.ResponseError(c, reciprocal.CodeInvalidParams)
			return
		}
		reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeInvalidParams, reciprocal.RemoveTopStruct(errs.Translate(reciprocal.Trans)))
		return
	}

	////传入逻辑层处理
	if err := Logic.UserSignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		reciprocal.ResponseError(c, reciprocal.CodeServerBusy)
		return
	}
	reciprocal.ResponseSuccess(c, nil)
}

// LoginUser  用户登录
func LoginUser(c *gin.Context) {
	p := new(Model.UserLoginMessage)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("用户注册参数解析失败", zap.Error(err))
		//判断错误类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			reciprocal.ResponseError(c, reciprocal.CodeInvalidParams)
			return
		}
		reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeInvalidParams, reciprocal.RemoveTopStruct(errs.Translate(reciprocal.Trans)))
		return
	}

	//传入逻辑层进行登录逻辑判断及处理
	token, err := Logic.UserLogin(p)
	fmt.Println(token)
	if err != nil {
		zap.L().Error("logic.UserLogin failed", zap.Error(err))
		reciprocal.ResponseError(c, reciprocal.CodeInvalidPassword)
		return
	}

	reciprocal.ResponseSuccess(c, token)

}

// ExchangePassword 修改密码
func ExchangePassword(c *gin.Context) {
	p := new(Model.UserExchangePassword)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("用户注册参数解析失败", zap.Error(err))
		//判断错误类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			reciprocal.ResponseError(c, reciprocal.CodeInvalidParams)
			return
		}
		reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeInvalidParams, reciprocal.RemoveTopStruct(errs.Translate(reciprocal.Trans)))
		return
	}
	//传入逻辑层处理
	if err := Logic.ExchangePasswd(p); err != nil {
		reciprocal.ResponseError(c, reciprocal.CodeExchangeUserPasswd)
		return
	}
	reciprocal.ResponseSuccess(c, "密码修改完成")
}

// SetUserPremLevel  修改用户权限
func SetUserPremLevel(c *gin.Context) {
	//解析用户Token
	userID, userPerm, err := getCurrentUserID(c)
	if err != nil {
		reciprocal.ResponseError(c, reciprocal.CodeNeedLogin)
		return
	}
	if userPerm < 3 {
		reciprocal.ResponseError(c, reciprocal.CodeUserPermissionDenied)
		return
	}
	p := new(Model.Userpermissions)
	p.Userid = userID
	p.Userpermissions = userPerm
	user := new(Model.Userpermissions)
	//当前Root管理员用户权限验证
	if err = Logic.UserRightsVerify(p, user); err != nil {
		reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeUserPermissionDenied, "请尽快修改密码")
		return
	}
	//获取要修改的用户权限
	p = new(Model.Userpermissions)
	if err = c.ShouldBindJSON(p); err != nil {
		zap.L().Error("用户权限变更参数解析失败", zap.Error(err))
		//判断错误类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			reciprocal.ResponseError(c, reciprocal.CodeInvalidParams)
			return
		}
		reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeInvalidParams, reciprocal.RemoveTopStruct(errs.Translate(reciprocal.Trans)))
		return

	}
	if err = Logic.ExchangeUserPrem(p); err != nil {
		reciprocal.ResponseError(c, reciprocal.CodeServerBusy)
		return
	}
	reciprocal.ResponseSuccess(c, "用户权限修改成功")

}

// GetAllUserMessage  查询所有用户权限
func GetAllUserMessage(c *gin.Context) {
	//解析用户Token
	userID, userPerm, err := getCurrentUserID(c)
	if err != nil {
		reciprocal.ResponseError(c, reciprocal.CodeNeedLogin)
		return
	}
	if userPerm < 3 {
		reciprocal.ResponseError(c, reciprocal.CodeUserPermissionDenied)
		return
	}
	p := new(Model.Userpermissions)
	p.Userid = userID
	p.Userpermissions = userPerm
	user := new(Model.Userpermissions)
	//逻辑层操作
	//根据用户id查询用户权限
	if err = Logic.UserRightsVerify(p, user); err != nil {
		reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeUserPermissionDenied, "请尽快修改密码")
		return
	}
	//查询所有用户权限
	//获取分页信息
	pageNum, pageSize, pageErr := GetPageInfo(c)
	if pageErr != nil {
		zap.L().Error("Get Page Info Error", zap.Error(pageErr))
		reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeServerBusy, "系统忙，请稍后查询")
		return
	}
	AllUserPrem := new([]Model.ReturnUserpermissions)
	if err = Logic.GetAllUserPerm(pageNum, pageSize, AllUserPrem); err != nil {
		reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeServerBusy, "系统忙，请稍后查询")
		return
	}
	if len(*AllUserPrem) == 0 {
		reciprocal.ResponseError(c, reciprocal.CodePageError)
		return
	}
	reciprocal.ResponseSuccess(c, AllUserPrem)

}

// GetOneUserMessage 查询单个用户权限
func GetOneUserMessage(c *gin.Context) {
	//解析用户Token
	userID, userPerm, err := getCurrentUserID(c)
	if err != nil {
		reciprocal.ResponseError(c, reciprocal.CodeNeedLogin)
		return
	}
	if userPerm < 3 {
		reciprocal.ResponseError(c, reciprocal.CodeUserPermissionDenied)
		return
	}
	p := new(Model.Userpermissions)
	p.Userid = userID
	p.Userpermissions = userPerm
	user := new(Model.Userpermissions)
	//逻辑层操作
	//根据用户id查询用户权限
	if err = Logic.UserRightsVerify(p, user); err != nil {
		reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeUserPermissionDenied, "请尽快修改密码")
		return
	}
	//查询指定用户权限
	str := c.Query("Author_id")
	OneUserPrem := new([]Model.ReturnUserpermissions)
	if err = Logic.GetOneUserPerm(str, OneUserPrem); err != nil {
		reciprocal.ResponseErrorWithMsg(c, reciprocal.CodeServerBusy, "系统忙，请稍后查询")
		return
	}
	if len(*OneUserPrem) == 0 {
		reciprocal.ResponseError(c, reciprocal.CodePageError)
		return
	}
	reciprocal.ResponseSuccess(c, OneUserPrem)

}
