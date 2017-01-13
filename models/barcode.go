package models

import (
	"github.com/astaxie/beego/orm"
)

type BarCode struct {
	Id   int64  //  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Code string //`Code` varchar(50) NOT NULL COMMENT '条码',
	Name string //`Name` varchar(50) DEFAULT EMPTY STRING COMMENT '条码',
}

// 数据表名称
func (this *BarCode) TableName() string {
	return "barcode"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(BarCode))
}

type BarCodeModel struct {
	//业务模型
}

func (this *BarCodeModel) AddBarCode(barcode BarCode) (num, Id int64, err error) {
	o := orm.NewOrm()
	isExist := o.QueryTable("barcode").Filter("Code", barcode.Code).Filter("Name", barcode.Name).Exist()
	if !isExist {
		codeIsExist := o.QueryTable("barcode").Filter("Code", barcode.Code).Exist()
		if codeIsExist {
			var barcodeTemp BarCode
			o.QueryTable("barcode").Filter("Code", barcode.Code).One(&barcodeTemp)
			barcode.Id = barcodeTemp.Id
			num, err = o.Update(&barcode, "Name")
		} else {
			Id, err = o.Insert(&barcode)
		}
	}
	return
}
