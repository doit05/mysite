package models

import (
	"mysite/helper"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// '首页图表数据';
type IndexData struct {
	Id            int64   //  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Uid           int64   // `uid` int(11) NOT NULL COMMENT ' 用户ID',
	Monitor_id    int64   //   `monitor_id` int(11) NOT NULL COMMENT '监控ID',
	Max_energy    float64 //   `max_energy` double NOT NULL COMMENT '最大能量值',
	Normal_sum    float64 //   `normal_sum` double NOT NULL COMMENT '标准消耗能量',
	Normal_energy float64 //   `normal_energy` double NOT NULL COMMENT '标准摄入能量',
	Total_sum     float64 //   `total_sum` double NOT NULL COMMENT '用户总消耗',
	Total_energy  float64 //   `total_energy` double NOT NULL COMMENT '用户总摄入',
	Datetime      int64   //   `date` int(11) NOT NULL COMMENT '日期',
	Create_time   int64   //   `create_time` int(11) NOT NULL COMMENT '创建时间',
	Update_time   int64   // `update_time` int(11) DEFAULT NULL COMMENT '更新时间',
}

// 数据表名称
func (this *IndexData) TableName() string {
	return "tbl_index"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(IndexData))
}

type IndexDataModel struct { //业务模型
}

//获取用户首页数据
func (this *IndexDataModel) GetIndexData(uid, date int64) (datas []IndexData, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("tbl_index").Filter("uid", uid).Filter("date", date).All(&datas)
	return
}

//获取用户首页数据详情
func (this *IndexDataModel) GetIndexDataDetail(uid, date, pal, sex int64, agestart, ageend float64) (data IndexData, err error) {
	o := orm.NewOrm()
	var list orm.ParamsList
	start := date
	end := date + 86400
	sql := "select NutrValueRNI from ch_dris where agestart = ? and ageend = ? and pal = ? and sex = ? and NutrName= '能量'"
	_, err = o.Raw(sql, agestart, ageend, pal, sex).ValuesFlat(&list)
	sum, ok := list[0].(string)
	if ok && len(sum) > 0 {
		data.Normal_energy, _ = strconv.ParseFloat(sum, 64)
	}
	sql_sum := "SELECT sum(a.energy) FROM food_data as a where a.uid = ? and a.date > ? and a.date < ? "
	_, err = o.Raw(sql_sum, uid, start, end).ValuesFlat(&list)
	sum, ok = list[0].(string)
	if ok {
		data.Total_energy, _ = strconv.ParseFloat(sum, 64)
	}
	sql_sum = "SELECT sum(a.energy) FROM sports_data as a where a.uid = ? and a.date > ? and a.date < ? "
	_, err = o.Raw(sql_sum, uid, start, end).ValuesFlat(&list)
	sum, ok = list[0].(string)
	if ok {
		data.Total_sum, _ = strconv.ParseFloat(sum, 64)
	}
	data.Monitor_id = 23
	data.Max_energy = 4000
	data.Uid = uid

	// query := "SELECT sum(a.energy) as total_energy, c.NutrValueRNI as normal_energy, sum(d.energy) as total_sum, 4000 as max_energy FROM food_data as a, user_monitor as b, ch_dris as c, sports_data as d "
	// query += " where a.uid = " + strconv.FormatInt(uid, 10)
	// query +=  " and a.date > " + strconv.FormatInt(start, 10) + " and a.date < " +  strconv.FormatInt(end, 10)
	// query += " and d.uid = " + strconv.FormatInt(uid, 10)
	// query +=  " and d.date > " + strconv.FormatInt(start, 10) + " and d.date < " +  strconv.FormatInt(end, 10)
	// query += " and b.uid = " + strconv.FormatInt(uid, 10)
	// // query +=  " and d.date > " + strconv.FormatInt(start, 10) + " and d.date < " +  strconv.FormatInt(start, 10)
	// query += " and c.id = " + strconv.FormatInt(36, 10)
	// err = o.Raw(query).QueryRow(&data)
	// fmt.Println(query, "\n", data)
	return
}

//添加用户饮食数据
func (this *IndexDataModel) AddIndexData(data IndexData) (Id int64, err error) {
	o := orm.NewOrm()
	Id, err = o.Insert(&data)
	return
}

//删除用户饮食数据
func (this *IndexDataModel) DelIndexData(id int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("tbl_index").Filter("id", id).Delete()
	return
}
func (this *IndexDataModel) SetIndexMonitorId(uid, monitorid int64) (num int64, err error) {
	o := orm.NewOrm()
	obj := orm.Params{
		"update_time": helper.GetTimestamp(),
		"Monitor_id":  monitorid,
	}
	num, err = o.QueryTable("user_basic").Filter("id", uid).Update(obj)
	return
}

//更改用户首页数据
func (this *IndexDataModel) UpdateIndexData(data IndexData) (num int64, err error) {
	// o := orm.NewOrm()
	// obj := orm.Params{
	// 	"update_time": food.Update_time,
	// }
	// if len(food.Img_url) > 0{
	// 	obj["img_url"] = food.Img_url
	// }
	// if food.Value > 0{
	// 	obj["value"] = food.Value
	// }
	// if food.Energy > 0{
	// 	obj["energy"] = food.Energy
	// }
	// if len(food.Unit) > 0{
	// 	obj["unit"] = food.Unit
	// }
	// if food.Date > 0{
	// 	obj["date"] = food.Date
	// }
	// if food.Diettime > 0{
	// 	obj["diettime"] = food.Diettime
	// }
	// if food.Collection > 0{
	// 	obj["collection"] = food.Collection
	// }
	// num, err = o.QueryTable("food_data").Filter("id", food.Id).Update(obj)
	return
}

// 设置首页监控数据id
func (this *IndexDataModel) GetIndexMonitor(uid int64) (obj User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("user_basic").Filter("id", uid).One(&obj, "Monitor_id")
	return
}
