package models

import (
	"github.com/astaxie/beego/orm"
	// "time"
)

// '用户食谱数据';
type Chfood_goods_nutri struct {
	Id             int64   // `id` int(11) NOT NULL AUTO_INCREMENT,
	Foodname       string  // `FoodName` text,
	Foodableweight string  // `FoodableWeight` text,
	Energy         string // `Energy` text,
	Protein        string // `Protein` text,
	Fat            string // `Fat` text,
	Fiber          string // `Fiber` text,
	Carbohydrate   string // `Carbohydrate` text,
	Vitamina       string // `VitaminA` text,
	Imgurl         string  // `imgurl` text,
	Category       string  // `category` text,
}

// 数据表名称
func (this *Chfood_goods_nutri) TableName() string {
	return "chinesefoodgoodsnutri"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(Chfood_goods_nutri))
}

//根据获取类别获取食物
func (this *CookBookModel) GetFoodByCategory(num int64, category string) (datas []Chfood_goods_nutri, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("chinesefoodgoodsnutri").Filter("category", category).Limit(num).All(&datas)
	return
}
