package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"mysite/helper"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	configPrifix := helper.GetConfigPrifix() // 获取配置前缀
	// Yeat/数据库(Test)
	err := orm.RegisterDataBase("default", "mysql", beego.AppConfig.String(configPrifix+"Test"))
	if err != nil {
		panic("设置Test_db数据库配置失败, err: " + err.Error())
	}
}
