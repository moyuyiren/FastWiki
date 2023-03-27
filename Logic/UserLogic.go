package Logic

import (
	"FastWiki/Dao/MySql"
	"FastWiki/Model"
	"FastWiki/Token"
	"FastWiki/Utils"
	"fmt"
	"go.uber.org/zap"
)

var ErrorSnowIDMsg error

// UserSignUp 用户注册
func UserSignUp(p *Model.GetUserInfo) (err error) {
	SnowID := Utils.GenID()
	if SnowID == 0 {
		return ErrorSnowIDMsg
	}
	EncrtptPWD := Utils.EncryptPassword(p.Password)
	user := &Model.User{
		Userid:      SnowID,
		Author_id:   p.Author_id,
		Password:    EncrtptPWD,
		Email:       p.Email,
		Phonenumber: p.Phonenumber,
	}

	userpermissions := &Model.Userpermissions{
		Userid:          SnowID,
		Userpermissions: 0,
	}

	//传入Dao层
	if err := MySql.UserRegister(user); err != nil {
		return err
	}
	if err := MySql.UserRegisterPermissions(userpermissions); err != nil {
		return err
	}
	return err
}

// UserLogin 用户登录
func UserLogin(p *Model.UserLoginMessage) (token string, err error) {
	//根据用户手机号码查询用户密码，判断密码是否正确
	nowUser := new(Model.User)
	if err = MySql.UserInfo(p, nowUser); err != nil {
		return "", err
	}
	if nowUser.Password != Utils.EncryptPassword(p.Password) {
		zap.L().Error("the user password err" + string(nowUser.Userid))
		return "", err
	}
	//查询用户的权限等级，生成当前用户Token返回
	nowUserPerm := new(Model.Userpermissions)
	if err := MySql.UserPerm(nowUser.Userid, nowUserPerm); err != nil {
		zap.L().Error("the user permissions not Get" + string(nowUser.Userid))
		return "", err
	}

	return Token.GenToken(nowUserPerm.Userid, nowUserPerm.Userpermissions)
}

// ExchangePasswd 用户修改密码
func ExchangePasswd(p *Model.UserExchangePassword) (err error) {
	//根据用户注册手机号查询用户密码
	pwd, err := MySql.SelectPasswdForPhonenumber(p.Phonenumber)
	if err != nil {
		return err
	}
	oldPwd := Utils.EncryptPassword(p.Password)
	fmt.Println(pwd, oldPwd)
	//密码校验
	if pwd == oldPwd {
		fmt.Println("进来了")
		//密码正确更新密码，密码错误修改失败返回
		newPwd := Utils.EncryptPassword(p.NewPassword)
		if err = MySql.ChangeUserPasswd(p.Phonenumber, newPwd); err != nil {
			zap.L().Error("更新用户密码失败", zap.Error(err))
			return err
		}
	}
	return &PwdError{
		msg: "旧密码输入错误",
		Pwd: p.Password,
	}
}
