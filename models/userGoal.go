package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

// '用户信息';
type UserGoal struct {
	Id              int64   //`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Uid             int64   //`uid` int(11) NOT NULL COMMENT ' 用户ID',
	Name            string  //`name` varchar(200) DEFAULT NULL COMMENT '目标名称',
	Type            int64   // type` tinyint default null comment '目标类型 1-目标， 2-监控',
	Img_url         string  // `img_url` varchar(255) default null comment '图片地址',
	Switch          int64   // `Switch` tinyint(4) DEFAULT '0' COMMENT '目标是否满足',
	Unit            string  //Unit string comment '单位'
	Start           float64 // `start` double DEFAULT NULL COMMENT '开始值',
	End             float64 // `end` double DEFAULT NULL COMMENT '目标值',
	Weekly_change   float64 // `weeklyChange` double DEFAULT NULL COMMENT '每周变化',
	Increase        int     // `increase` tinyint DEFAULT '0' COMMENT '0-表示增加 1-表示减少',
	Date            int64   // `Date` date DEFAULT NULL COMMENT '截止日期',
	Daily_entry     float64 // `dailyEntry` double DEFAULT NULL COMMENT '推荐的能量摄入数量',
	Daily_foodpoint float64 // `dailyFoodpoint` double DEFAULT NULL COMMENT '推荐的食物份数摄入量',
	Create_time     int64   //`create_time` int(11) DEFAULT NULL COMMENT '创建时间',
	Update_time     int64   //`update_time` int(11) DEFAULT NULL COMMENT '更新时间',
	Max_value       float64
	Min_value       float64
}

// 单位对应表

var Unit = map[string]string{
	"体重":    "kg",
	"血糖":    "mol",
	"血压":    "mmHg",
	"心率":    "bpm",
	"钙":     "mg",
	"体温":    "℃",
	"围度":    "cm",
	"血氧":    "%",
	"铜":     "mg",
	"铁":     "mg",
	"镁":     "mg",
	"锰":     "mg",
	"钾":     "mg",
	"锌":     "ug",
	"磷":     "mg",
	"维生素A":  "ug",
	"维生素B1": "ug",
	"维生素B2": "ug",
	"维生素C":  "ug",
	"维生素E":  "ug",
}

// 数据表名称
func (this *UserGoal) TableName() string {
	return "user_goal"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(UserGoal))
}

type UserGoalModel struct { //业务模型
}

//更新一条数据,如果不存在，则创建新的目标
func (this *UserGoalModel) UpdateSystemGoal(goal UserGoal) (Id int64, err error) {
	o := orm.NewOrm()
	var created bool
	if created, Id, err = o.ReadOrCreate(&goal, "Type", "Name"); err == nil {
		if !created {
			Id, err = o.Update(&goal, "Switch", "Update_time")
		}
	}
	return
}

//更新一条数据,如果不存在，则创建新的目标
func (this *UserGoalModel) UpdateUserGoal(goal UserGoal) (num int64, err error) {
	goal.Unit = Unit[goal.Name]
	o := orm.NewOrm()
	var created bool
	if created, num, err = o.ReadOrCreate(&goal, "Uid", "Name", "Type"); err == nil {
		if !created {
			r, err1 := o.Raw("UPDATE user_goal SET switch = ?, update_time = ?  WHERE uid = ? and name = ?", 1, goal.Update_time, goal.Uid, goal.Name).Exec()
			if err1 != nil {
				err = err1
				return
			}
			num, err = r.RowsAffected()
		}
	}
	return num, err
}

//删除一个目标
func (this *UserGoalModel) DeleteUserGoal(goal UserGoal) (num int64, err error) {

	o := orm.NewOrm()
	r, err := o.Raw("UPDATE user_goal SET switch = ?, update_time = ?  WHERE uid = ? and name = ?", goal.Switch, goal.Update_time, goal.Uid, goal.Name).Exec()
	num, err = r.RowsAffected()
	return
}

//插入一条数据
func (this *UserGoalModel) InsertUserGoal(goal UserGoal) (Id int64, err error) {
	o := orm.NewOrm()
	Id, err = o.Insert(&goal)
	return
}

//获取用户全部目标
func (this *UserGoalModel) GetUserGoalsById(uid, goaltype int64) (goals []UserGoal, err error) {
	o := orm.NewOrm()
	select_str := "select * from user_goal where uid = '" + strconv.FormatInt(uid, 10) + "'"
	select_str += " and type = '" + strconv.FormatInt(goaltype, 10) + "'"
	select_str += " and switch = 1"
	select_str += " order by id"

	num, err := o.Raw(select_str).QueryRows(&goals)
	if num == 0 && err != nil { //如果，数据库中未找到，则返回nil和err
		return nil, err
	}
	// fmt.Println(select_str)
	return
}

/**
*通过用户Id和目标名称获取
 */
func (this *UserGoalModel) GetUserGoalInfo(uid int64, Goalname string) (goal UserGoal, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("user_goal").Filter("uid", uid).Filter("name", Goalname).One(&goal)
	return
}

/**
*通过Id获取
 */
func (this *UserGoalModel) GetUserGoalById(id int64) (goal UserGoal, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("user_goal").Filter("id", id).One(&goal)
	return
}
