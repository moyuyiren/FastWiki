package Model

// 主分类
type Community struct {
	Community_id   int64  `xorm:"community_id"`
	Community_name string `xorm:"community_name"`
	Introduction   string `xorm:"introduction"`
}

// 次分类
type Seccommunity struct {
	Community_id       int64  `xorm:"community_id"`       //大分类
	Sec_community_id   int64  `xorm:"sec_community_id"`   //分类
	Sec_community_info string `xorm:"sec_community_info"` //条目名
}
