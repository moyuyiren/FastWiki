package MySql

import "FastWiki/Model"

// SelectLv1Menu 查询一级菜单数据
func SelectLv1Menu(p *[]Model.Community) (err error) {
	if err = dbEngine.Find(p); err != nil {
		return err
	}
	return
}

// 根据结果查询二级菜单数据
func SelectLv2Menu(p *[]Model.Seccommunity, s int64) (err error) {
	if err = dbEngine.Where("community_id = ?", s).Find(p); err != nil {
		return err
	}
	//
	return
}
