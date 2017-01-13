package models

import (
	"github.com/astaxie/beego/orm"
	//"fmt"
)

// '中国食物营养信息';
type ChDris struct {
	Id             int64   //`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Agestart       float64 //`ageStart` double NOT NULL COMMENT ' 两个字段组合成0-6个月、7-12个月、1-3岁、4-6岁、7-10岁、11-13岁、14-18岁、19-49岁、50-64岁、65-80岁和80岁以上',
	Ageend         float64 //`ageEnd` double NOT NULL COMMENT '0.6、1、3、6、10、13、18、49、64、79以及80以上',
	Sex            int64   //`sex` tinyint(4) NOT NULL COMMENT '性别 0-男 1－女',
	Nutrno         string  //`NutrNo` varchar(255) DEFAULT NULL COMMENT '营养素编号',
	Nutrname       string  //`NutrName` varchar(255) DEFAULT NULL COMMENT '营养素名称',
	Nutrlist       string  //`NutrList` varchar(255) DEFAULT '0' COMMENT '营养素单位',
	Nutrvaluetype  int64   //`NutrValueType` tinyint(4) DEFAULT NULL COMMENT '营养素数值计量类型。当NutrValueType为0时，NutrValue字段有效；当NutrValueType为1时，NutrValueStart和NutrValueEnd有效。',
	Nutrvaluerni   float64 //`NutrValueRNI` double DEFAULT NULL COMMENT '营养素值参考摄入量(RNI)',
	Nutrvalueul    float64 //`NutrValueUL` double DEFAULT NULL COMMENT '营养素值最高摄入量(UL)',
	Nutrvaluestart float64 //`NutrValueStart` double DEFAULT NULL COMMENT '营养素值范围1',
	Nutrvalueend   float64 //`NutrValueEnd` double DEFAULT NULL COMMENT '营养素值范围2',
	Pal            int64   //`PAL` tinyint(4) DEFAULT NULL COMMENT '身体活动水平 0-不适用 1-轻活动水平(1.5) 2-中活动水平(1.75) 3-重活动水平(2)',
	Create_time    int64   //`create_time` int(11) DEFAULT NULL COMMENT '创建时间',
	Update_time    int64   //`update_time` int(11) DEFAULT NULL COMMENT '更新时间',
}

// 数据表名称
func (this *ChDris) TableName() string {
	return "ch_dris"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(ChDris))
}

type ChDrisModel struct {
	//业务模型
}

func (this *FoodDataModel) GetChdriByAge(name string, pal, sex int64, agestart, ageend float64) (data ChDris, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("ch_dris").Filter("Nutrname", name).Filter("Sex", sex).Filter("Pal", pal).Filter("Agestart", agestart).Filter("Ageend", ageend).One(&data)
	return
}

//根据年龄、性别、身体活动水平获取中国膳食参考
func (this *ChDrisModel) GetChDietNutriRef(agestart, ageend float64, sex, pal int64) (refs []ChDris, err error) {
	o := orm.NewOrm()
	agestart = agestart - 1
	_, err = o.QueryTable("ch_dris").Filter("Agestart", agestart).Filter("Ageend", ageend).Filter("Pal__in", 0, pal).Filter("Sex__in", 0, sex).All(&refs)
	return
}
