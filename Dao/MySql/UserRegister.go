package MySql

import (
	"FastWiki/Model"
	"fmt"
	"go.uber.org/zap"
)

// UserRegister 用户注册（写入用户注册表）
func UserRegister(user *Model.User) error {
	affected, err := dbEngine.Insert(user)
	fmt.Println("查看用户数据写入", affected)
	return err
}

// UserRegisterPermissions 用户注册（写入用户权限表）
func UserRegisterPermissions(userpermissions *Model.Userpermissions) error {
	affected, err := dbEngine.Insert(userpermissions)
	fmt.Println("查看用户数据写入", affected)
	return err
}

// UserInfo 查询用户基本信息
func UserInfo(b *Model.UserLoginMessage, p *Model.User) error {
	has, err := dbEngine.Where("phonenumber = ?", b.Phonenumber).Get(p)
	if !has {
		zap.L().Error("The user not register ", zap.Error(err))
		return err
	}
	return nil
}

// UserPerm 查询用户权限
func UserPerm(b int64, p *Model.Userpermissions) error {
	has, err := dbEngine.Where("userid = ?", b).Get(p)
	if !has {
		zap.L().Error("The user permissions query lose", zap.Error(err))
		return err
	}
	return nil
}

// SelectPasswdForPhonenumber 根据手机号码查询用户密码
func SelectPasswdForPhonenumber(num string) (pwd string, err error) {
	fmt.Println(num)
	sql := "select password from tb_user where phonenumber = ?"
	pp, err := dbEngine.Query(sql, num)
	pwd = string(pp[0]["password"])
	return
}

// ChangeUserPasswd 修改用户密码
func ChangeUserPasswd(phonenum, newpasswd string) (err error) {
	sql := "update tb_user set password = ? where phonenumber = ?"
	res, err := dbEngine.Exec(sql, newpasswd, phonenum)
	fmt.Println(res)
	return err
}

// GetAllUserPrem 获取所有用户权限
func GetAllUserPrem(pageNum, pageSize int64, p *[]Model.ReturnUserpermissions) (err error) {
	fmt.Println(pageNum, pageSize)
	err = dbEngine.SQL("select b.* ,a.author_id from tb_User a join tb_Userpermissions b where a.userid =b.userid limit ?,?", pageNum, pageSize).Find(p)
	return
}

// GetOneUserPrem 查询指定用户权限
func GetOneUserPrem(str string, p *[]Model.ReturnUserpermissions) (err error) {
	err = dbEngine.SQL("select a.author_id,b.userid,b.userpermissions from tb_user a,tb_userpermissions b where a.author_id = ?", str).Find(p)
	return
}

// ChangeUserPrem  修改用户权限
func ChangeUserPrem(p *Model.Userpermissions) (err error) {
	_, err = dbEngine.Exec("update tb_userpermissions set userpermissions=? where userid = ?", p.Userpermissions, p.Userid)
	return
}
