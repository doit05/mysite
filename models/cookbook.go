package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"mysite/helper"
	"mysite/utils"
	//"strconv"
	//"time"
)

// '用户食谱数据';
type CookBook struct {
	Id          int64  //  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Uid         int64  // `uid` int(11) NOT NULL COMMENT ' 用户ID',
	Foodid      string // `foodid` varchar(50) NOT NULL COMMENT '食物id',
	Diettime    int64  // `diettime` tinyint(4) DEFAULT NULL COMMENT '早中晚',
	Source      string // `source` varchar(255) NOT NULL COMMENT 'b表名',
	Week        int64  // `week` tinyint(4) NOT NULL COMMENT '周一至周日',
	Fooddate    string // `fooddate` date DEFAULT NULL COMMENT '日期',
	Foodname    string // `foodname` varchar(255) DEFAULT '“”' COMMENT '食物名称',
	Indexid     int64  //食物id
	Rule        string //规则名称
	Create_time int64  // `create_time` int(11) DEFAULT NULL COMMENT '创建时间',
	Update_time int64  // `update_time` int(11) DEFAULT NULL COMMENT '更新时间',
	Goodorbad   int64  // `update_time` int(11) DEFAULT NULL COMMENT '好评或差评',
	Ispurchase  int64  // `update_time` int(11) DEFAULT NULL COMMENT '是否购买',
}

//rules kind的限制
type RulesKind struct {
	Category []string
	DayKinds int64
	WeeKinds int64
	Name     string
}

// 数据表名称
func (this *CookBook) TableName() string {
	return "cookbook"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(CookBook))
}

type CookBookModel struct {
	//业务模型
}

//获取用户饮食数据
func (this *CookBookModel) GetCookBookByUid(uid int64, fooddate string) (datas []CookBook, err error) {
	o := orm.NewOrm()
	var nums int64 = 0
	nums, err = o.QueryTable("cookbook").Filter("uid", uid).Filter("fooddate", fooddate).OrderBy("diettime").All(&datas)
	if err == nil && nums == 0 {
		go this.SetRecommendByUid(uid, helper.StrToTimestamp(fooddate))
	}

	return
}

// //根据食物id获取用户饮食数据
// func (this *CookBookModel) GetCookBookById(uid int64, fooddate string) (datas []CookBook, err error) {
// 	o := orm.NewOrm()
// 	sql := "SELECT id, user_name FROM cookbook as a, " +
// 	" WHERE id = ?"
// 	_, err := o.Raw("", 1).QueryRows(&users)
// 	return
// }

//设置用户饮食数据
func (this *CookBookModel) SetCookBookByUid(cook CookBook) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&cook)
	return
}

//recommend food rules
const (
	count int64  = 10          //查询倍数
	kind1 string = "谷薯芋、杂豆、主食" // 3, 5
	kind2 string = "坚果、大豆及制品"  // 2, 5
	kind3 string = "蛋类、肉类及制品"  // 3, 5
	kind4 string = "奶类及制品"     //2, 5
	kind5 string = "蔬果和菌藻"     // 4 10

)

//设置用户推荐食物
func (th *CookBookModel) SetRecommendByUid(uid, date int64) (diets []Chfoodnutri_hk, err error) {
	rules_kind := make([]RulesKind, 0)
	if date == 0 {
		date = helper.GetTimestamp()
	}
	datetime := ""
	umodel := UserModel{}
	dris, err := umodel.GetUserDris(uid)
	if err != nil {
		return
	}
	total_energy := dris.Nutrvaluerni
	var curr_energy float64 = 0
	var i int64 = 0
	for i = 0; i < 7; i++ {
		curr_energy = 0
		datetime = helper.Date("Y-m-d", date+i*86400)
		fmt.Println(i, " --- ", datetime, " int64 : ", date+i*86400)

		kind_1 := RulesKind{th.getSubGroups(kind1), 3, 5, kind1}
		kind_2 := RulesKind{th.getSubGroups(kind2), 2, 5, kind2}
		kind_3 := RulesKind{th.getSubGroups(kind3), 3, 5, kind3}
		kind_4 := RulesKind{th.getSubGroups(kind4), 2, 5, kind4}
		kind_5 := RulesKind{th.getSubGroups(kind5), 2, 10, kind5}
		rules_kind = append(rules_kind, kind_1)
		rules_kind = append(rules_kind, kind_2)
		rules_kind = append(rules_kind, kind_3)
		rules_kind = append(rules_kind, kind_4)
		rules_kind = append(rules_kind, kind_5)
		md := new(ChDietnutriModel)
		for _, item := range rules_kind {
			datas, err1 := md.GetFoodBySubGroup(i*20, item.DayKinds*count, item.Category)
			if err1 == nil {
				for _, food := range datas {
					if th.serachdiets(food, diets) == false && !th.MaxEnergy(food, curr_energy, total_energy) {
						diets = append(diets, food)
						curr_energy += food.Energy
					}
				}

			} else {
				err = err1
				return
			}
		}
		if len(diets) > 0 {
			name := "kind1 - kind5"
			err2 := th.putInto(datetime, name, uid, diets)
			if err2 != nil {
				utils.Log.Error("插入cookbook : %v", err2)
			}

		}
	}

	return
}

func (th *CookBookModel) MaxEnergy(food Chfoodnutri_hk, curr, total float64) (ret bool) {
	ret = curr+food.Energy > total
	return

}

func (th *CookBookModel) serachdiets(food Chfoodnutri_hk, diets []Chfoodnutri_hk) (ret bool) {
	if len(diets) < 1 {
		ret = false
	}
	for _, item := range diets {
		if food.Id == item.Id {
			return true
		}
		if food.Subgroupid == item.Subgroupid {
			return true
		}
	}
	return false

}

func (th *CookBookModel) putInto(date, name string, uid int64, datas []Chfoodnutri_hk) (err error) {

	for index, one := range datas {
		cook := CookBook{}
		cook.Rule = name
		cook.Create_time = helper.GetTimestamp()
		cook.Update_time = cook.Create_time
		cook.Uid = uid
		cook.Source = "chinesefoodnutrifromhk"
		cook.Foodname = one.Foodchinesename
		cook.Indexid = one.Id
		cook.Diettime = int64(index%3 + 1)
		//tmp := index / 3
		cook.Fooddate = date
		_, err = th.SetCookBookByUid(cook)
		if err != nil {
			utils.Log.Error("recommend coookbook %v", err)
		}
	}
	return
}

func (th *CookBookModel) getGroups(kind string) (ret []string) {
	if len(kind) > 0 {
		switch kind {
		case kind1:
			ret = append(ret, "01")
		case kind2:
			ret = append(ret, "02", "05")
		case kind3:
			ret = append(ret, "09", "06")
		case kind4:
			ret = append(ret, "10")
		case kind5:
			ret = append(ret, "03")
		}

	}
	return
}

func (th *CookBookModel) getSubGroups(kind string) (ret []string) {
	groups := th.getGroups(kind)
	if len(groups) > 0 {
		tmp := make([]string, 0)
		for _, groupId := range groups {
			o := orm.NewOrm()
			query := fmt.Sprintf("SELECT SubGroupID FROM yeat.chineseFoodSubGroup where GroupID = %s", groupId)
			num, err := o.Raw(query).QueryRows(&tmp)
			if num > 0 && err == nil {
				for _, item := range tmp {
					if helper.InSlice(item, ret) == false {
						ret = append(ret, item)
					}
				}

			}
		}
	}
	return
}
