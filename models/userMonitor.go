package models

import (
	"fmt"
	"mysite/helper"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// '用户监控数据信息';
type UserMonitor struct {
	Id          int64   //`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Uid         int64   //`uid` int(11) NOT NULL COMMENT ' 用户ID',
	Monitor_id  int64   //`monitor_id` int(11) DEFAULT NULL COMMENT '监控的id',
	Create_time int64   //`create_time` int(11) DEFAULT NULL COMMENT '监控创建时间',
	Update_time int64   //`Update_time` int(11) DEFAULT NULL COMMENT '监控更新时间',
	Value       float64 // `Value` double DEFAULT NULL COMMENT '每周变化',
	Value1      float64 // `Value` double DEFAULT NULL COMMENT '每周变化',
	Comment     string  // `Comment` varchar(255) DEFAULT NULL COMMENT '备注',
	Value_time  int64   // `Value_time` date DEFAULT NULL COMMENT '当前时间戳的输入的值',
}

// '用户最后监控信息';
type UserMonitorLast struct {
	Id         int64 //`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Uid        int64 //`uid` int(11) NOT NULL COMMENT ' 用户ID',
	Name       string
	Img_url    string
	Value      float64 // `Value` double DEFAULT NULL COMMENT '每周变化',
	Value1     float64 // `Value` double DEFAULT NULL COMMENT '每周变化',
	Value_time int64   // `Value_time` date DEFAULT NULL COMMENT '当前时间戳的输入的值',
	Comment    string  // `Comment` varchar(255) DEFAULT NULL COMMENT '备注',
	Unit       string  //Unit string comment '单位'
}

// 数据表名称
func (this *UserMonitor) TableName() string {
	return "user_monitor"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(UserMonitor))
}

type UserMonitorModel struct { //业务模型
}

//更新一条数据,如果不存在，则创建新的监控
func (this *UserMonitorModel) UpdateUserMonitor(monitor UserMonitor) (Id int64, err error) {
	o := orm.NewOrm()
	one := UserMonitor{}
	err = o.QueryTable("user_monitor").Filter("Uid", monitor.Uid).Filter("Monitor_id", monitor.Monitor_id).Filter("Value_time", monitor.Value_time).One(&one)

	if one.Id > 0 {
		obj := orm.Params{
			"update_time": helper.GetTimestamp(),
		}
		if monitor.Value > 0 {
			obj["Value"] = monitor.Value
		}
		if monitor.Value1 > 0 {
			obj["Value1"] = monitor.Value1
		}
		_, err = o.QueryTable("user_monitor").Filter("id", one.Id).Update(obj)
	} else {
		Id, err = this.InsertUserMonitor(monitor)
	}
	return
}

//插入一条数据
func (this *UserMonitorModel) InsertUserMonitor(monitor UserMonitor) (Id int64, err error) {
	o := orm.NewOrm()
	Id, err = o.Insert(&monitor)
	return
}

//获取用户全部监控
func (this *UserMonitorModel) GetUserMonitorsById(uid, monitorid, start, end int64) (monitors []UserMonitor, err error) {
	o := orm.NewOrm()
	select_str := "select * from user_monitor where uid = '" + strconv.FormatInt(uid, 10) + "'"
	select_str += " and Monitor_id = '" + strconv.FormatInt(monitorid, 10) + "'"
	select_str += " and value_time > '" + strconv.FormatInt(start, 10) + "'"
	select_str += " and value_time < '" + strconv.FormatInt(end, 10) + "'"
	select_str += " order by value_time"
	_, err = o.Raw(select_str).QueryRows(&monitors)
	return
}

func (this *UserMonitorModel) GetUserMonitorsByEn(uid, monitorid, start, end int64, enName string) (monitors []UserMonitor, err error) {
	o := orm.NewOrm()
	select_str := "select * from user_monitor where uid = '" + strconv.FormatInt(uid, 10) + "'"
	select_str += " and Monitor_id = '" + strconv.FormatInt(monitorid, 10) + "'"
	select_str += " and value_time > '" + strconv.FormatInt(start, 10) + "'"
	select_str += " and value_time < '" + strconv.FormatInt(end, 10) + "'"
	select_str += " order by value_time"
	_, err = o.Raw(select_str).QueryRows(&monitors)
	return
}

//根据id获取监控
func (this *UserMonitorModel) GetUserMonitor(id int64) (monitor UserMonitor, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("user_monitor").Filter("Id", id).One(&monitor)
	return
}

//根据id获取监控
func (this *UserMonitorModel) GetUserMonitorLast(uid int64) (monitors []UserMonitorLast, err error) {
	o := orm.NewOrm()
	query := "SELECT a.id, a.uid, a.name,a.img_url, b.value, b.value1, b.value_time, b.comment, a.unit FROM user_goal as a left join user_monitor as b on (a.lastid = b.id)"
	query += " where a.switch = 1 and a.type = 2 and a.uid = " + strconv.FormatInt(uid, 10)
	_, err = o.Raw(query).QueryRows(&monitors)
	fmt.Println(query)
	return
}

/**
*通过用户Id和监控名称获取
 */
func (this *UserMonitorModel) GetUserMonitorInfo(Uid int64, Monitorname string) (monitor UserMonitor, err error) {
	o := orm.NewOrm()
	query := "select * from user_monitor where monitor_name = '" + Monitorname + "' and uid = '" + strconv.FormatInt(monitor.Uid, 10) + "'"
	err = o.Raw(query).QueryRow(&monitor)
	return
}

//根据id设置监控最后一条数据
func (this *UserMonitorModel) SetUserMonitorLast(monitorid, dataid int64) (num int64, err error) {
	o := orm.NewOrm()
	r, err := o.Raw("UPDATE user_goal SET lastid = ? WHERE id = ?", dataid, monitorid).Exec()
	num, err = r.RowsAffected()
	return
}

//获取营养素数据
func (this *UserMonitorModel) GetUserNutriInfo(uid, start, end int64, nutriName string) (monitor UserMonitor, err error) {
	foodModel := FoodDataModel{}
	chModel := ChDietnutriModel{}
	foods, err1 := foodModel.GetFoodsByUid(uid, start, end)
	if err1 != nil {
		err = err1
		return
	}
	for _, data := range foods {
		maps, err2 := chModel.GetChDietnutriByName(data.Name, nutriName)
		if err2 != nil {
			err = err2
			return
		}
		if maps == nil {
			continue
		} else {
			value, _ := helper.GetFloat(maps[0][nutriName])
			monitor.Value += value
		}
	}
	return
}
