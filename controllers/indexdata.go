package controllers

import (
	"mysite/helper"
	"mysite/helper/apicode"
	"mysite/models"
	"mysite/utils"
	"mysite/validate"
)

type IndexDataController struct {
	BaseController
}

// 初始化
func (this *IndexDataController) Prepare() {

}

// 获取首页数据
func (this *IndexDataController) GetIndexData() {
	params := validate.IndexDataParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 检测模型
	Validatiton := validate.DailyValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckIndexDataParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	um := models.UserModel{}
	user, err := um.GetUserInfoById(params.Uid)
	if err != nil {

		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
		return
	}
	if user.Sex == 0 {
		// this.RenderApiJsonEmpty(apicode.GetSexFailed, apicode.Msg(apicode.GetSexFailed))
		// return
		user.Sex = 1
	}

	if user.Age == 0 {
		// this.RenderApiJsonEmpty(apicode.GetAgeFailed, apicode.Msg(apicode.GetAgeFailed))
		// return
		user.Age = 26
	}
	agestart, ageend := helper.GetAgeScore(user.Age)
	indexModel := models.IndexDataModel{}
	if user.Age < 2 {
		user.Pal = 0
	} else {
		if user.Pal == 0 {
			user.Pal = 1
		}
	}
	if agestart > 0.6 && user.Weight == 0 {
		// this.RenderApiJsonEmpty(apicode.GetPalFailed, apicode.Msg(apicode.GetPalFailed))
		// return
	}
	if agestart < 4 && user.Weight == 0 {
		// this.RenderApiJsonEmpty(apicode.GetWeightFailed, apicode.Msg(apicode.GetWeightFailed))
		// return
	}
	if user.Monitor_id == 0 {
		// this.RenderApiJsonEmpty(apicode.GetMonitorFailed, apicode.Msg(apicode.GetMonitorFailed))
		// return
	}

	// 设置参数
	data, err := indexModel.GetIndexDataDetail(params.Uid, params.Datetime, user.Pal, user.Sex, agestart, ageend)
	userGoalModel := models.UserGoalModel{}
	userMonitorModel := models.UserMonitorModel{}

	if err != nil {
		utils.Log.Error("查询首页数据错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}

	if agestart < 4 {
		data.Normal_energy *= float64(user.Weight)
	}

	extro, err1 := userGoalModel.GetUserGoalById(user.Monitor_id)
	if err1 != nil {
		utils.Log.Error("查询首页监控错误 ： %v", err)
		this.RenderApiJsonExtroData(apicode.Success, apicode.Msg(apicode.Success), nil, nil, data)
	}
	//判断是否是微量元素
	monitorModel := models.UserMonitorModel{}
	enName := helper.GetEnName(extro.Name)
	monitors := []models.UserMonitor{}
	if len(enName) > 0 {
		start := params.Datetime - 86400*6
		end := params.Datetime + 86400
		for ; start < end; start += 86400 {
			monitor, err := monitorModel.GetUserNutriInfo(params.Uid, start, start+86400, enName)
			monitor.Uid = params.Uid
			monitor.Monitor_id = extro.Id
			if err != nil {
				utils.Log.Error("查询微量元素数据错误： ％v", err)
				continue
			}
			monitor.Value_time = start
			monitors = append(monitors, monitor)
		}
	} else {
		monitors, err = userMonitorModel.GetUserMonitorsById(params.Uid, user.Monitor_id, params.Datetime-86400*6, params.Datetime+86400)
	}

	if err != nil {
		utils.Log.Error("查询首页监控数据错误 ： %v", err)
	}
	this.RenderApiJsonExtroData(apicode.Success, apicode.Msg(apicode.Success), extro, monitors, data)
	return
}

// 设置首页监控数据id
func (this *IndexDataController) SetIndexData() {
	params := validate.IndexDataParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 检测模型
	Validatiton := validate.DailyValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckSetIndexMonitorParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	indexModel := models.IndexDataModel{}
	// 设置参数
	data, err := indexModel.SetIndexMonitorId(params.Uid, params.Monitor_id)

	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), data)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

func (this *IndexDataController) GetExamData() {

	indexModel := models.IndexDataModel{}
	//indexModel.Getexam("/Users/doit/tmp/20160716143816001.txt")
	// 设置参数
	datas, err := indexModel.GetResultsByName("BBB", 0)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), datas)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}
