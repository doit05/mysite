package models

import (
	"github.com/astaxie/beego/orm"
)

// '中国食物营养信息';
type ChDietnutri struct {
	Id              int64   //`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Groupid         string  //  `GroupID` varchar(30) NOT NULL,
	Dietid          string  //`DietID` varchar(30) NOT NULL,
	Dietchinesename string  // `DietChineseName` varchar(255) DEFAULT NULL,
	Dietname        string  //   `DietName` decimal(10,0) DEFAULT NULL,
	Diettype        string  //   `DietType` varchar(255) DEFAULT NULL,
	Dietweight      string  //   `DietWeight` varchar(30) DEFAULT NULL,
	Energy          float64 //   `Energy` decimal(10,0) DEFAULT NULL,
	Protein         float64 //   `Protein` decimal(10,0) DEFAULT NULL,
	Carbohydrate    float64 //   `Carbohydrate` decimal(10,0) DEFAULT NULL,
	Fat             float64 //   `Fat` decimal(10,0) DEFAULT NULL,
	Fiber           float64 //  `Fiber` decimal(10,0) DEFAULT NULL,
	Sugar           float64 //  `Sugar` decimal(10,0) DEFAULT NULL,
	Vitaminc        float64 //`VitaminC` decimal(10,0) DEFAULT NULL,
	Dietsteps       string  //  `DietSteps` varchar(8000) DEFAULT NULL,
	Meterial        string  // `Meterial` int(11) DEFAULT NULL,
	Imagesrc        string  //   `ImageSrc` text,
	Source          int64   //   `Source` int(11) DEFAULT '0',
	Vitamina        float64 //   `VitaminA` decimal(10,0) DEFAULT NULL,
	Vitamine        float64 //   `VitaminE` decimal(10,0) DEFAULT NULL,
	Vitaminb1       float64 //   `VitaminB1` decimal(10,0) DEFAULT NULL,
	Vitaminb2       float64 //   `VitaminB2` decimal(10,0) DEFAULT NULL,
	Vitaminb3       float64 //   `VitaminB3` decimal(10,0) DEFAULT NULL,
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
	Seleniumse      float64 // `SeleniumSe` decimal(10,0) DEFAULT NULL,
	CreateTime      int64   //`create_time` int(11) DEFAULT NULL COMMENT '创建时间',
	UpdateTime      int64   //`update_time` int(11) DEFAULT NULL COMMENT '更新时间',
	Barcode         string  // `BarCode` varchar(50) DEFAULT NULL
	Evaluation      string  // `Evaluation` varchar(255) DEFAULT NULL
}

// 数据表名称
func (this *ChDietnutri) TableName() string {
	return "ch_dietnutri"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(ChDietnutri))
}

type ChDietnutriModel struct {
	//业务模型
}

//更新一条数据,如果不存在，则创建新的目标
func (this *ChDietnutriModel) UpdateChDietnutriData(diet ChDietnutri) (Id int64, err error) {
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
func (this *ChDietnutriModel) GetChDietnutri(index, len int64) (diets []ChDietnutri, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("ch_dietnutri").Limit(len, index).All(&diets)
	return
}

//获取收藏食物详情
func (this *ChDietnutriModel) GetChDietnutriCollect(uid, start, end, collection int64) (diets []ChDietnutri, err error) {
	o := orm.NewOrm()
	if collection == 1 {
		query := "select foodid from food_data where uid=?"
		query += " and date>?"
		query += " and date<?"
		query += " and collection=?"
		var fooddata []string
		nums, err := o.Raw(query, uid, start, end, collection).QueryRows(&fooddata)
		if err == nil && nums > 0 {
			_, err = o.QueryTable("ch_dietnutri").Filter("id__in", fooddata).All(&diets)
		}
	} else {
		err.Error()
	}
	return
}

//获取常见食物
func (this *ChDietnutriModel) GetCommonFoods(uid int64) (diets []ChDietnutri, err error) {
	o := orm.NewOrm()
	sql := "call sp_common_food("
	sql += "?)"
	_, err = o.Raw(sql, uid).QueryRows(&diets)
	return
}

//获取中国食物详情
func (this *ChDietnutriModel) GetChDietnutriByName(name, nutriName string) (maps []orm.Params, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("ch_dietnutri").Filter("DietChineseName", name).Limit(1).Values(&maps)
	return
}

//根据食物名搜索中国食物
func (this *ChDietnutriModel) SearchChDietnutriByName(name string) (diets []ChDietnutri, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("ch_dietnutri").Filter("DietChineseName__icontains", name).Limit(100).All(&diets)
	return
}

//根据食物条码搜索中国食物
func (this *ChDietnutriModel) SearchChDietnutriByBarCode(barcode string) (diet ChDietnutri, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("ch_dietnutri").Filter("BarCode__icontains", barcode).Limit(1).All(&diet)
	return
}
