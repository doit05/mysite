package models

import "github.com/astaxie/beego/orm"

// '中国食物营养信息';
type ChDietGroup struct {
	Groupid          string `orm:"pk;column(Groupid);"`
	Chinesegroupname string // `ChineseGroupName` varchar(255) DEFAULT NULL,
	Groupname        string // `GroupName` varchar(255) DEFAULT NULL,
}

// 数据表名称
func (this *ChDietGroup) TableName() string {
	return "chinesefoodgroup"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(ChDietGroup))
}

type ChDietGroupModel struct { //业务模型
}

//获取中国饮食
func (this *ChDietGroupModel) GetChDietGroup() (groups []ChDietGroup, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("chinesefoodgroup").Limit(100).All(&groups)
	return
}
