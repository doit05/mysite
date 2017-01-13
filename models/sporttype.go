package models

import "github.com/astaxie/beego/orm"

// '运动类型信息';
type SportsType struct {
	Id          int64   //`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Uid         int64   //`uid` int(11) NOT NULL AUTO_INCREMENT COMMENT 'uid ',
	Name        string  // `name` text CHARACTER SET utf8 NOT NULL,
	Type        int64   // `type` decimal(10,0) DEFAULT NULL,
	Energy      float64 // `enery` decimal(10,0) DEFAULT NULL,
	Unit        string  // `unit` int(11) DEFAULT NULL,
	Imgurl      string  // 'imgurl' varchar(255) default ""
	Switch      int64   // `switch` tinyint DEFAULT NULL,
	Create_time int64
	Update_time int64
}

// 数据表名称
func (this *SportsType) TableName() string {
	return "sporttype"
}

// 注册定义的model
func init() {
	orm.RegisterModelWithPrefix("", new(SportsType))
}

//获取所有的运动类型
func (this *SportsDataModel) GetSportTypes(uid int64) (types []SportsType, err error) {
	o := orm.NewOrm()
	sql := "select *from sporttype where uid=? order by convert(name using gbk) collate gbk_chinese_ci asc limit 200"
	_, err = o.Raw(sql, uid).QueryRows(&types)
	return
}

//添加所有的运动类型
func (this *SportsDataModel) AddSportType(sporttype SportsType) (id int64, err error) {
	o := orm.NewOrm()
	exist := o.QueryTable("sporttype").Filter("uid", sporttype.Uid).Filter("name", sporttype.Name).Exist()
	if !exist {
		sporttype.Switch = 1
		id, err = o.Insert(&sporttype)
	}
	return
}

//删除所有的运动类型
func (this *SportsDataModel) DelSportType(id int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sporttype").Filter("id", id).Delete()
	return
}
