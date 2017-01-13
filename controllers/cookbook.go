package controllers

import (
	"fmt"
	"mysite/helper"
	"mysite/helper/apicode"
	"mysite/models"
	"mysite/utils"
	"mysite/validate"
)

type CookBookController struct {
	BaseController
}

// 初始化
func (this *CookBookController) Prepare() {

}

// 获取用户食谱
func (this *CookBookController) GetUserCookBook() {
	params := validate.CookBookParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 食物检测模型
	Validatiton := validate.DailyValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckGetUserCookBookParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	cookModel := models.CookBookModel{}
	// 设置参数
	datas, err := cookModel.GetCookBookByUid(params.Uid, params.Fooddate)
	if len(datas) == 0 {
		datas, err = cookModel.GetCookBookByUid(1, params.Fooddate)
	}
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), datas)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

// 设置用户食谱
func (this *CookBookController) SetUserCookBook() {
	params := validate.CookBookParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 食物检测模型
	Validatiton := validate.DailyValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckSetUserCookBookParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	cookModel := models.CookBookModel{}
	// 设置参数
	cook := models.CookBook{}
	cook.Uid = params.Uid
	cook.Source = params.Source
	cook.Indexid = params.Indexid
	cook.Diettime = params.Diettime
	cook.Ispurchase = params.Ispurchase
	cook.Goodorbad = params.Goodorbad
	datas, err := cookModel.SetCookBookByUid(cook)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), datas)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

// 设置用户食谱
func (this *CookBookController) RecommendUserCookBook() {
	ch_hk := models.CookBookModel{}
	list, err := ch_hk.SetRecommendByUid(0, helper.GetTimestamp())
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), list)
	} else {
		fmt.Println(err)
	}
}

// 获取用户食谱
func (this *CookBookController) GetRecommendUserCookBook() {
	uid, err1 := this.GetInt64("uid")
	date := this.GetString("date")
	if err1 != nil || len(date) < 0 {
		this.RenderApiJsonSlice(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam), err1.Error())
	}

	ch_hk := models.CookBookModel{}
	list, err := ch_hk.GetCookBookByUid(uid, date)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), list)
	} else {
		fmt.Println(err)
	}
}
