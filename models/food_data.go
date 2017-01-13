package models

import (
	"github.com/astaxie/beego/orm"
)

// '用户日常饮食数据';
type FoodData struct {
	Id          int64   //  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Uid         int64   // `uid` int(11) NOT NULL COMMENT ' 用户ID',
	Foodid      int64   //`foodid` int(11) NOT NULL COMMENT 'FOODID',
	Name        string  // `name` varchar(200) NOT NULL COMMENT '食物名称',
	Type        int64   // `type` tinyint(4) DEFAULT NULL COMMENT '早中晚',
	Img_url     string  // `img_url` text COMMENT '图片地址',
	Switch      int64   // `switch` tinyint(4) DEFAULT '0' COMMENT '食物是否删除 0，标示删除',
	Value       float64 // `value` double DEFAULT NULL COMMENT '数值',
	Energy      float64 // `energy` double DEFAULT NULL COMMENT '消耗能量',
	Unit        string  // `unit` varchar(200) DEFAULT NULL COMMENT '单位',
	Date        int64   // `date` int(11) DEFAULT NULL COMMENT '日期',
	Source      int64   // `source` tinyint(4) DEFAULT NULL COMMENT '数据读取来源 1 －手动 2-读图片',
	Location    string  // `location` text COMMENT '地点位置',
	Create_time int64   // `create_time` int(11) DEFAULT NULL COMMENT '创建时间',
	Update_time int64   // `update_time` int(11) DEFAULT NULL COMMENT '更新时间',
	Diettime    int64   //早中晚
	Collection  int64   // `Collection` tinyint DEFAULT NULL COMMENT '用户食物收藏',
	Tablesource int64   //`tablesource` tinyint DEFAULT NULL COMMENT '表来源'
}

// 数据表名称
func (this *FoodData) TableName() string {
	return "food_data"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(FoodData))
}

type FoodDataModel struct { //业务模型
}

//获取用户饮食数据
func (this *FoodDataModel) GetFoodsByUid(uid, start, end int64) (foods []FoodData, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("food_data").Filter("uid", uid).Filter("date__gt", start).Filter("date__lt", end).OrderBy("diettime").Limit(100).All(&foods)
	return
}

//添加用户饮食数据
func (this *FoodDataModel) AddFoodsData(food FoodData) (Id int64, err error) {
	o := orm.NewOrm()
	Id, err = o.Insert(&food)
	return
}

//删除用户饮食数据
func (this *FoodDataModel) DelFoodsData(id int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("food_data").Filter("id", id).Delete()
	return
}

//获取饮食数据详情
func (this *FoodDataModel) GetFooddetail(id int64) (maps []orm.Params, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("food_data").Filter("id", id).Values(&maps, "id", "DietChineseName", "DietID", "DietSteps")
	return
}

//更改用户饮食数据
func (this *FoodDataModel) UpdateFoodsData(food FoodData) (num int64, err error) {
	o := orm.NewOrm()
	obj := orm.Params{
		"update_time": food.Update_time,
	}
	if len(food.Img_url) > 0 {
		obj["img_url"] = food.Img_url
	}
	if food.Value > 0 {
		obj["value"] = food.Value
	}
	if food.Energy > 0 {
		obj["energy"] = food.Energy
	}
	if len(food.Unit) > 0 {
		obj["unit"] = food.Unit
	}
	if food.Date > 0 {
		obj["date"] = food.Date
	}
	if food.Diettime > 0 {
		obj["diettime"] = food.Diettime
	}
	if food.Collection > 0 {
		obj["collection"] = food.Collection
	}
	num, err = o.QueryTable("food_data").Filter("id", food.Id).Update(obj)
	return
}

//取消食物收藏
func (this *FoodDataModel) ConcelCollect(uid, foodid int64) (err error) {
	o := orm.NewOrm()
	sql := "update food_data set collection=0 where uid=?"
	sql += " and foodid =?"
	sql += " and collection =1"
	_, err = o.Raw(sql, uid, foodid).Exec()
	return
}
