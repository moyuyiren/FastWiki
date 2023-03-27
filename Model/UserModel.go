package Model

/*===================================================前端数据结构=====================================================*/
// 用户注册前端数据结构
type GetUserInfo struct {
	Author_id   string `json:"author_Id"    binding:"required"`
	Password    string `json:"password"     binding:"required"`
	Repassword  string `json:"Repassword"   binding:"required,eqfield=Password"`
	Email       string `json:"email"        binding:"required"`
	Phonenumber string `json:"phonenumber"  binding:"required"`
}

// 用户登录前端数据结构
type UserLoginMessage struct {
	Phonenumber string `json:"phonenumber"  binding:"required"`
	Password    string `json:"password"     binding:"required"`
}

// 用户修改密码前端数据结构
type UserExchangePassword struct {
	Phonenumber string `json:"phonenumber"  binding:"required"`
	Password    string `json:"password"     binding:"required"`
	NewPassword string `json:"newPassword"     binding:"required"`
}

// 返回前端数据结构
type ReturnUserpermissions struct {
	Author_id       string `xorm:"author_id"`
	Userid          int64  `xorm:"userid"`
	Userpermissions int8   `xorm:"userpermissions"`
}

/*===================================================后端数据结构======================================================*/
// 用户注册数据结构
type User struct {
	Userid      int64  `xorm:"userid"`
	Author_id   string `xorm:"author_id"`
	Password    string `xorm:"password"`
	Email       string `xorm:"email"`
	Phonenumber string `xorm:"phonenumber"`
}

/*==================================================公用数据结构======================================================*/
// 用户权限数据结构
type Userpermissions struct {
	Userid          int64 `xorm:"userid" json:"userid"`
	Userpermissions int8  `xorm:"userpermissions" json:"userpermissions"`
}
