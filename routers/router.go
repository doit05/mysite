package routers

import (
	"mysite/controllers"

	"github.com/astaxie/beego"
)

func init() {
//测试
beego.Router("/system/test", &controllers.FastdfsController{}, "get:Get")

//docker
beego.Router("/docker/client/ps", &controllers.DockerClientController{}, "get:DockerPs")


	beego.Router("/sql", &controllers.UserController{}, "get:ExecSql")

	//条码插入
	beego.Router("/system/addbarcode", &controllers.BarCodeController{}, "post:AddBarCode")

	//系统路由
	beego.Router("/system/news", &controllers.ADController{})
	beego.Router("/upload", &controllers.MainController{}, "post:Upload")

	//系统用户
	beego.Router("/system/register", &controllers.UserController{}, "post:Register")
	beego.Router("/system/login", &controllers.UserController{}, "post:Login")

	//系统目标
	beego.Router("/system/setgoal", &controllers.UserGoalController{}, "post:SetSystemGoal")
	beego.Router("/system/getgoals", &controllers.UserGoalController{}, "get:GetSystemGoals")
	beego.Router("/system/deletegoal", &controllers.UserGoalController{}, "post:SetSystemGoal")

	//系统监控
	beego.Router("/system/setmonitor", &controllers.UserGoalController{}, "post:SetSystemGoal")
	beego.Router("/system/getmonitors", &controllers.UserGoalController{}, "get:GetSystemGoals")
	beego.Router("/system/deletemonitor", &controllers.UserGoalController{}, "post:SetSystemGoal")

	/**用户基础数据路由*/

	//3.用户账户
	beego.Router("basic/setuserinfo", &controllers.UserController{}, "post:SetUserInfo")
	beego.Router("basic/getuserinfo", &controllers.UserController{}, "get:GetUserInfo")

	//首页数据
	beego.Router("basic/getindexdata", &controllers.IndexDataController{}, "get:GetIndexData")
	beego.Router("basic/setindexmonitor", &controllers.IndexDataController{}, "get:SetIndexData")

	//获取体检数据
	beego.Router("basic/getexam", &controllers.IndexDataController{}, "get:GetExamData")

	//我的食谱
	beego.Router("basic/getcookbook", &controllers.CookBookController{}, "get:GetUserCookBook")
	beego.Router("basic/setcookbook", &controllers.CookBookController{}, "post:SetUserCookBook")
	//获取推荐食谱
	beego.Router("basic/getrecommendcookbook", &controllers.CookBookController{}, "post:GetRecommendUserCookBook")
	beego.Router("basic/recommendcookbook", &controllers.CookBookController{}, "get:RecommendUserCookBook")

	// 1. 用户目标
	beego.Router("/basic/choosegoal", &controllers.UserGoalController{}, "post:ChooseGoal")
	beego.Router("/basic/deletegoal", &controllers.UserGoalController{}, "post:ChooseGoal")
	beego.Router("/basic/getusergoals", &controllers.UserGoalController{}, "get:GetUserGoals")

	//2.用户监控

	beego.Router("/basic/choosemonitor", &controllers.UserGoalController{}, "post:ChooseGoal")
	beego.Router("/basic/deletemonitor", &controllers.UserGoalController{}, "post:DeleteGoal")
	beego.Router("/basic/getusermonitors", &controllers.UserGoalController{}, "get:GetUserGoals")

	//3.用户监控数据
	beego.Router("/basic/upmonitordata", &controllers.UserMonitorController{}, "post:UpData")
	beego.Router("/basic/downmonitordata", &controllers.UserMonitorController{}, "get:DownData")
	beego.Router("/basic/getusermonitorlast", &controllers.UserMonitorController{}, "get:GetUserMonitorLast")

	/**用户日常运动数据路由*/
	beego.Router("/daily/addsports", &controllers.SportsController{}, "post:AddSports")
	beego.Router("/daily/delsports", &controllers.SportsController{}, "get:DeleteSports")
	beego.Router("/daily/getsports", &controllers.SportsController{}, "get:GetSports")
	beego.Router("/daily/getsporttype", &controllers.SportsController{}, "get:GetSportTypes")
	beego.Router("/system/getsporttype", &controllers.SportsController{}, "get:GetSportTypes")
	beego.Router("/daily/addsporttype", &controllers.SportsController{}, "post:AddSportType")
	beego.Router("/daily/deletesporttype", &controllers.SportsController{}, "get:DelSportType")

	/**用户日常饮食数据路由*/
	beego.Router("/daily/getchdietnutri", &controllers.FoodController{}, "get:GetChDietnutris")
	beego.Router("/daily/dietsbygroup", &controllers.FoodController{}, "get:GetChDietsByGroupHK")
	beego.Router("/daily/getchdietgroup", &controllers.FoodController{}, "get:GetChDietgroups")
	beego.Router("/daily/getfooddetail", &controllers.FoodController{}, "get:GetFooddetail")
	beego.Router("/daily/searchfood", &controllers.FoodController{}, "get:SearchChDietsByName")
	beego.Router("/daily/searchfoodbybarcode", &controllers.FoodController{}, "get:SearchChDietsByBarCode")

	beego.Router("/daily/getfoodsbyuid", &controllers.FoodController{}, "get:GetFoodsByUid")
	beego.Router("/daily/addfoodsdatabyuid", &controllers.FoodController{}, "post:AddFoodsDataByUid")
	beego.Router("/daily/updatefoodsdatabyid", &controllers.FoodController{}, "post:UpdateFoodsDataById")
	beego.Router("/daily/deletefoodbyid", &controllers.FoodController{}, "get:DelFoodsData")
	beego.Router("/daily/getfoodcollection", &controllers.FoodController{}, "get:GetCollectFoods")
	beego.Router("/daily/getcommonfood", &controllers.FoodController{}, "get:GetCommonFoods")
	beego.Router("/daily/concelcollectfood", &controllers.FoodController{}, "get:ConcelCollectFoods")
	//中国居民膳食成分参考路由
	beego.Router("/system/getchnutriref", &controllers.ChdrisController{}, "post:GetChNutriRefs")

}
