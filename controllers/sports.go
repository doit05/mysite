package controllers

import (
	"fmt"
	"mysite/helper"
	"mysite/helper/apicode"
	"mysite/models"
	"mysite/utils"
	"mysite/validate"
)

type SportsController struct {
	BaseController
}

// 初始化
func (this *SportsController) Prepare() {

}

// 添加系统运动
func (this *SportsController) AddSports() {
	params := validate.SportsAddParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 运动检测模型
	Validatiton := validate.DailyValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckAddSportsParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		fmt.Printf("验证参数失败, errMsg: %s, params: %v", errMsg, params)
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	sportsModel := models.SportsDataModel{}
	// 设置参数
	sport := models.SportsData{}
	sport.Uid = 1
	if params.Uid > 1 {
		sport.Uid = params.Uid
	}
	sport.Name = params.Name
	sport.Type = helper.GetSystemType(this.Ctx.Request.URL.String())
	sport.Img_url = params.Img_url
	sport.Switch = helper.ChangeSwitch(this.Ctx.Request.URL.String())
	sport.Value = params.Value
	sport.Energy = params.Energy
	sport.Unit = params.Unit
	sport.Date = params.Date
	sport.Source = params.Source
	sport.Create_time = helper.GetTimestamp()
	sport.Update_time = sport.Create_time

	id, err := sportsModel.UpdateSportsData(sport)
	if err == nil && id > 0 {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), id)
	} else {
		utils.Log.Error("更新数据库错误 ： %v", err)
		this.RenderApiJson(apicode.InsertSportDataFailed, apicode.Msg(apicode.InsertSportDataFailed), err)
	}
	return
}

// 删除系统运动
func (this *SportsController) DeleteSports() {
	id, err := this.GetInt64("id")
	// 绑定参数
	if err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}
	// 验证参数
	if id < 1 {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", "id必须大于0", id) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	sportModel := models.SportsDataModel{}
	num, err := sportModel.DelSportsData(id)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), num)
	} else {
		utils.Log.Error("删除数据库错误 ： %v", err)
		this.RenderApiJson(apicode.DeleteError, apicode.Msg(apicode.DeleteError), err)
	}
	return

}

// 获取系统运动
func (this *SportsController) GetSports() {
	params := validate.SportsAddParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 运动检测模型
	Validatiton := validate.DailyValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckGetSportsParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		fmt.Printf("验证参数失败, errMsg: %s, params: %v", errMsg, params)
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	sportsModel := models.SportsDataModel{}
	// 获取数据
	uid := params.Uid
	name := params.Name
	start := params.Start
	end := params.End

	datas, err := sportsModel.GetSportsData(uid, start, end, name)
	if err == nil {
		var total float64 = 0
		for _, data := range datas {
			total += data.Energy
		}
		this.RenderApiJsonExtro(apicode.Success, apicode.Msg(apicode.Success), total, datas)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

// 获取运动类型
func (this *SportsController) GetSportTypes() {
	var uid int64 = 0
	//如果是非系统请求，即用户请求
	if helper.GetSystemType(this.Ctx.Request.URL.String()) == 1 {
		params := validate.UserIdParams{}
		// 绑定参数
		if err := this.ParseForm(&params); err != nil {
			utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
			this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
			return
		}

		// 运动检测模型
		Validatiton := validate.UserValidationParams{}

		// 验证参数
		if ok, errMsg := Validatiton.CheckUserIdParams(params); !ok {
			utils.Log.Error("验证参数失败, errMsg: %s, params: %s", errMsg, params) // 记录log
			this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
			return
		}
		uid = params.Uid
	}

	sportModel := models.SportsDataModel{}
	// 获取数据
	types, err := sportModel.GetSportTypes(uid)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), types)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return

}

// 添加运动类型
func (this *SportsController) AddSportType() {
	params := validate.SportTypeParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		fmt.Printf("绑定参数出错, err: %v", err)
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 运动检测模型
	Validatiton := validate.DailyValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckAddSportTypeParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", errMsg, params) // 记录log
		fmt.Printf("验证参数失败, errMsg: %s, params: %s", errMsg, params)
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	sportModel := models.SportsDataModel{}
	sporttype := models.SportsType{}
	sporttype.Uid = params.Uid
	sporttype.Name = params.Name
	sporttype.Type = 0
	sporttype.Energy = params.Energy
	sporttype.Unit = params.Unit
	sporttype.Imgurl = params.Imgurl
	sporttype.Create_time = helper.GetTimestamp()
	sporttype.Update_time = sporttype.Create_time

	// 获取数据
	id, err := sportModel.AddSportType(sporttype)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), id)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return

}

// 添加运动类型
func (this *SportsController) DelSportType() {
	params := validate.SportTypeParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 验证参数
	if params.Id < 1 {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", "id 小于1", params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	sportModel := models.SportsDataModel{}

	// 获取数据
	num, err := sportModel.DelSportType(params.Id)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), num)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return

}
