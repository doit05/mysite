package models

import (
	"github.com/astaxie/beego/orm"
	// "fmt"
	// "strconv"
)

// '运动信息';
type SportsData struct {
	Id          int64   //`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Uid         int64   //`uid` int(11) NOT NULL COMMENT ' 用户ID',
	Name        string  //`name` varchar(200) NOT NULL COMMENT '运动名称',
	Type        int64   // `type` tinyint(4) DEFAULT NULL COMMENT 'sport操作来源类型 1-系统， 2-自定义',
	Img_url     string  // `img_url` varchar(255) DEFAULT NULL COMMENT '图片地址',
	Switch      int64   // `switch` tinyint(4) DEFAULT '0' COMMENT '运动记录是否删除 0，标示删除',
	Value       float64 // `value` double DEFAULT NULL COMMENT '数值',
	Energy      float64 // `energy` double DEFAULT NULL COMMENT '消耗能量',
	Unit        string  // `unit` varchar(200) DEFAULT NULL COMMENT '单位',
	Date        int64   // `date` int(11) DEFAULT NULL COMMENT '日期',
	Source      int64   // `source` tinyint(4) DEFAULT NULL COMMENT '数据读取来源 1 －手动 2-读仪器',
	Create_time int64   //`create_time` int(11) DEFAULT NULL COMMENT '创建时间',
	Update_time int64   //`update_time` int(11) DEFAULT NULL COMMENT '更新时间',
}

// 数据表名称
func (this *SportsData) TableName() string {
	return "sports_data"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(SportsData))
}

type SportsDataModel struct { //业务模型
}


//删除用户运动数据
func (this *SportsDataModel) DelSportsData(id int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("sports_data").Filter("id", id).Delete()
	return
}


//更新一条数据,如果不存在，则创建新的运动
func (this *SportsDataModel) UpdateSportsData(sport SportsData) (Id int64, err error) {
	o := orm.NewOrm()
	var created bool
	if created, Id, err = o.ReadOrCreate(&sport, "Name", "Uid", "Date"); err == nil {
		if !created {
			Id, err = o.Update(&sport, "Value", "Update_time", "Energy")
		}
	}
	return
}

//获取用户运动数据
func (this *SportsDataModel) GetSportsData(uid, start, end int64, name string) (datas []SportsData, err error) {
	o := orm.NewOrm()
	if len(name) > 0{
		select_str := "select * from sports_data where uid = ? and name = '?' and date > ? and date < ?"
		_, err = o.Raw(select_str, uid, name, start, end).QueryRows(&datas)
		return
	}
	select_str := "select * from sports_data where uid = ? and date > ? and date < ?"
	_, err = o.Raw(select_str, uid, start, end).QueryRows(&datas)
	return
}
