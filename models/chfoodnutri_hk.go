package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// '中国香港食物营养信息';
type Chfoodnutri_hk struct {
	Id              int64   //`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Groupid         string  //  `GroupID` varchar(30) NOT NULL,
	Subgroupid      string  //  `SubGroupID` varchar(30) NOT NULL,
	Foodid          string  //`FoodID` varchar(30) NOT NULL,
	Foodchinesename string  // `FoodChineseName` varchar(255) DEFAULT NULL,
	Foodname        string  //   `FoodName` decimal(10,0) DEFAULT NULL,
	Foodalias       string  //   `FoodAlias` varchar(255) DEFAULT NULL,
	Datasource      string  //   `DataSource` varchar(255) DEFAULT NULL,
	Foodableweight  string  //   `FoodableWeight` varchar(30) DEFAULT NULL,
	Energy          float64 //   `Energy` decimal(10,0) DEFAULT NULL,
	Water           float64 // `Water` decimal(10,0) DEFAULT NULL,
	Protein         float64 //   `Protein` decimal(10,0) DEFAULT NULL,
	Carbohydrate    float64 //   `Carbohydrate` decimal(10,0) DEFAULT NULL,
	Fat             float64 //   `Fat` decimal(10,0) DEFAULT NULL,
	Fiber           float64 //  `Fiber` decimal(10,0) DEFAULT NULL,
	Sugar           float64 //  `Sugar` decimal(10,0) DEFAULT NULL,
	Vitaminc        float64 //`VitaminC` decimal(10,0) DEFAULT NULL,
	Vitamina        float64 //   `VitaminA` decimal(10,0) DEFAULT NULL,
	Vitamine        float64 //   `VitaminE` decimal(10,0) DEFAULT NULL,
	Vitaminb1       float64 //   `VitaminB1` decimal(10,0) DEFAULT NULL,
	Vitaminb2       float64 //   `VitaminB2` decimal(10,0) DEFAULT NULL,
	Cholesterol     float64 //  `Cholesterol` decimal(10,0) DEFAULT NULL,
	Magnesiummg     float64 //  `MagnesiumMg` decimal(10,0) DEFAULT NULL,
	Calciumca       float64 //  `CalciumCa` decimal(10,0) DEFAULT NULL,
	Ironfe          float64 //  `IronFe` decimal(10,0) DEFAULT NULL,
	Zinczn          float64 //   `ZincZn` decimal(10,0) DEFAULT NULL,
	Coppercu        float64 //   `CopperCu` decimal(10,0) DEFAULT NULL,
	Manganesemn     float64 //   `ManganeseMn` decimal(10,0) DEFAULT NULL,
	Kaliumk         float64 // `KaliumK` decimal(10,0) DEFAULT NULL,
	Phosphorp       float64 // `PhosphorP` decimal(10,0) DEFAULT NULL,
	Sodiumna        float64 // `SodiumNa` decimal(10,0) DEFAULT NULL,
	Saturatedfat    float64 // `SaturatedFat` decimal(10,0) DEFAULT NULL,
	Unsaturatedfat  float64 // `UnSaturatedFat` decimal(10,0) DEFAULT NULL,
	updatedflag     float64 // `updatedFlag` int(11) DEFAULT '0',
	CreateTime      int64   //`create_time` int(11) DEFAULT NULL COMMENT '创建时间',
	UpdateTime      int64   //`update_time` int(11) DEFAULT NULL COMMENT '更新时间',
}

// 数据表名称
func (this *Chfoodnutri_hk) TableName() string {
	return "chinesefoodnutrifromhk"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(Chfoodnutri_hk))
}

//更新一条数据,如果不存在，则创建新的目标
func (this *ChDietnutriModel) UpdateChDietnutriHKData(diet Chfoodnutri_hk) (Id int64, err error) {
	o := orm.NewOrm()
	var created bool
	if created, Id, err = o.ReadOrCreate(&diet, "Groupid", "Dietid"); err == nil {
		if !created {
			Id, err = o.Update(&diet, "Dietname", "Update_time", "Energy")
		}
	}
	return
}

//获取中国饮食
func (this *ChDietnutriModel) GetChDietnutriHK() (diets []Chfoodnutri_hk, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("chinesefoodnutrifromhk").Limit(100).All(&diets)
	return
}

//获取一组中国饮食数据
func (this *ChDietnutriModel) GetChDietsByGroupIdHK(groupid string) (diets []Chfoodnutri_hk, err error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("chinesefoodnutrifromhk").Filter("Groupid", groupid).Limit(100).All(&diets)
	fmt.Println(num)
	return
}

//获取食物类别
func (this *ChDietnutriModel) GetChDietsGroupsHK() (list orm.ParamsList, err error) {
	o := orm.NewOrm()
	query := "select distinct Groupid from chinesefoodnutrifromhk"
	_, err = o.Raw(query).ValuesFlat(&list)
	return
}

//搜索香港饮食数据
func (this *ChDietnutriModel) SearchChDietnutriByNameHK(groupid int64, name string) (diets []Chfoodnutri_hk, err error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("chinesefoodnutrifromhk").Filter("Groupid", groupid).Filter("Foodchinesename__icontains", name).Limit(100).All(&diets)
	fmt.Println(num)
	return
}

//根据获取类别获取食物
func (this *ChDietnutriModel) GetFoodBySubGroup(start,num int64, subgroups []string) (datas []Chfoodnutri_hk, err error) {
				o := orm.NewOrm()
				length := len(subgroups)
				count := int64(num)/int64(length)

				for _, item := range subgroups{
								tmp := make([]Chfoodnutri_hk,0)
								nums, err1 := o.QueryTable("chinesefoodnutrifromhk").Limit(count, start).Filter("SubGroupID", item).All(&tmp)
								if nums > 0 && err == nil {
												for _,food := range tmp {
																datas = append(datas, food)
												}
								}else {
												err = err1
								}
				}

				return
}
