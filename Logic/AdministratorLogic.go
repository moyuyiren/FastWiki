package Logic

import (
	"FastWiki/Dao/MySql"
	"FastWiki/Dao/Redis"
	"FastWiki/Model"
	"FastWiki/RabbitMQ"
	"errors"
	"go.uber.org/zap"
)

var ErrorPremIDMsg error

// UserRightsVerify 用户权限二次验证
func UserRightsVerify(p *Model.Userpermissions, user *Model.Userpermissions) (err error) {
	if err = MySql.UserPerm(p.Userid, user); err != nil {
		return err
	}
	if user.Userpermissions != p.Userpermissions {
		zap.L().Error("This Token Is fake")
		//通过rabbitMQ推送重大漏洞信息至信息收集系统
		return &PwdError{msg: "用户Token可能出现泄漏 ，请提醒用户修改密码"}
	}
	return err

}

// FlushRedisMessageForMain 刷新redis缓存信息
func FlushRedisMessageForMain(p *Model.Userpermissions) (err error) {
	//验证用户权限是否真实
	user := new(Model.Userpermissions)
	if err = UserRightsVerify(p, user); err != nil {
		if errr := RabbitMQ.SendMessage("The password of the super administrator has been leaked. Handle it in a timely manner"); errr != nil {
			return errors.New(err.Error() + errr.Error())
		}
		return err
	}

	/*开始刷新Redis流程*/
	//从mysql中读取一级菜单写入Redis
	Lv1Menu := make([]Model.Community, 0)
	if err = MySql.SelectLv1Menu(&Lv1Menu); err != nil {
		zap.L().Error("Select Lv1 Menu Error", zap.Error(err))
		return err
	}
	//刷新一级菜单
	if err = Redis.FlushMenuLv1(&Lv1Menu); err != nil {
		zap.L().Error("Flush Lv1 Menu Error", zap.Error(err))
		return err
	}
	//======================================================================================================================
	//从mysql中读取二级菜单写入Redis
	for i := 0; i < len(Lv1Menu); i++ {
		Lv2Menu := make([]Model.Seccommunity, 0)
		if err = MySql.SelectLv2Menu(&Lv2Menu, Lv1Menu[i].Community_id); err != nil {
			zap.L().Error("Select Lv2 Menu Error", zap.Error(err))
			return err
		}
		//刷新二级菜单
		if err = Redis.FlushMenuLv2(&Lv2Menu, Lv1Menu[i].Community_name); err != nil {
			zap.L().Error("Flush Lv2 Menu Error", zap.Error(err))
			return err
		}
	}

	return err
}

// GetAllUserPerm 获取所有用户权限信息
func GetAllUserPerm(pageNum, pageSize int64, p *[]Model.ReturnUserpermissions) (err error) {
	if err = MySql.GetAllUserPrem(pageNum, pageSize, p); err != nil {
		zap.L().Error("全量分页查询用户信息失败", zap.Error(err))
		return err
	}
	return
}

// GetOneUserPerm 获取所有用户权限信息
func GetOneUserPerm(str string, p *[]Model.ReturnUserpermissions) (err error) {
	if err = MySql.GetOneUserPrem(str, p); err != nil {
		zap.L().Error("查询指定用户信息失败", zap.Error(err))
		return err
	}
	return
}

// ExchangeUserPrem 修改用户权限信息
func ExchangeUserPrem(p *Model.Userpermissions) (err error) {
	if err = MySql.ChangeUserPrem(p); err != nil {
		zap.L().Error("修改用户权限失败", zap.Error(err))
		return
	}
	return
}
