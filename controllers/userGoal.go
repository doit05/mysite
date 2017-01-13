package controllers

import (
	"mysite/helper"
	"mysite/helper/apicode"
	"mysite/models"
	"mysite/utils"
	"mysite/validate"
	// "fmt"
)

type UserGoalController struct {
	BaseController
}

// 初始化
func (this *UserGoalController) Prepare() {

}

// 设置系统目标
func (this *UserGoalController) SetSystemGoal() {

	params := validate.UserGoalParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 用户目标检测模型
	Validatiton := validate.UserValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckAddSystemGoalParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	goalModel := models.UserGoalModel{}
	// 设置参数
	goal := models.UserGoal{}
	goal.Uid = 1
	goal.Name = params.Name
	goal.Img_url = params.Img_url
	goal.Type = helper.GetType(this.Ctx.Request.URL.String())
	goal.Switch = helper.ChangeSwitch(this.Ctx.Request.URL.String())
	goal.Create_time = helper.GetTimestamp()
	goal.Update_time = goal.Create_time

	id, err := goalModel.UpdateSystemGoal(goal)
	if err == nil && id > 0 {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), id)
	} else {
		utils.Log.Error("更新数据库错误 ： %v", err)
		this.RenderApiJson(apicode.InsertGoalFailed, apicode.Msg(apicode.InsertGoalFailed), err)
	}
	return
}

// 获取系统全部目标
func (this *UserGoalController) GetSystemGoals() {
	params := validate.UserGoalParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定AddGoal参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 用户目标检测模型
	Validatiton := validate.UserValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckGetSystemGoalParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	goalModel := models.UserGoalModel{}
	var uid int64 = 1

	goaltype := helper.GetType(this.Ctx.Request.URL.String())

	goals, err := goalModel.GetUserGoalsById(uid, goaltype)

	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), goals)
	} else {
		utils.Log.Error("查询数据库错误： ％v", err)
		this.RenderApiJson(apicode.GetGoalsFailed, apicode.Msg(apicode.GetGoalsFailed), err)
	}
}

// 用户选择目标，初始化
func (this *UserGoalController) ChooseGoal() {

	params := validate.UserGoalParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定AddGoal参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 用户目标检测模型
	Validatiton := validate.UserValidationParams{}
	if helper.GetType(this.Ctx.Request.URL.String()) == 1 {
		// 验证参数
		if ok, errMsg := Validatiton.CheckAddGoalParams(params); !ok {
			utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
			this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
			return
		}
	} else {
		// 验证参数
		if ok, errMsg := Validatiton.CheckAddMonitorParams(params); !ok {
			utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
			this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
			return
		}
	}

	goalModel := models.UserGoalModel{}
	sysgoal, err1 := goalModel.GetUserGoalInfo(1, params.Name)
	if err1 != nil {
		utils.Log.Error("查询系统目标失败 ： ", err1)
		this.RenderApiJson(apicode.GetSysGoalFailed, apicode.Msg(apicode.GetSysGoalFailed), err1)
	}

	// 设置参数
	goal := models.UserGoal{}
	goal.Max_value = sysgoal.Max_value
	goal.Min_value = sysgoal.Min_value
	goal.Uid = params.Uid
	goal.Name = params.Name
	goal.Img_url = sysgoal.Img_url
	goal.Start = params.Start
	goal.End = params.End
	goal.Weekly_change = params.Weekly_change
	goal.Increase = params.Increase
	goal.Date = params.Date
	goal.Type = helper.GetType(this.Ctx.Request.URL.String())
	goal.Switch = helper.ChangeSwitch(this.Ctx.Request.URL.String())
	goal.Create_time = helper.GetTimestamp()
	goal.Update_time = goal.Create_time
	id, err := goalModel.UpdateUserGoal(goal)
	if err == nil && id > 0 {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), id)
	} else {
		utils.Log.Error("插入数据库错误 ： %v", err)
		this.RenderApiJson(apicode.InsertGoalFailed, apicode.Msg(apicode.InsertGoalFailed), err)
	}
	return
}

// 用户删除目标
func (this *UserGoalController) DeleteGoal() {

	params := validate.UserGoalParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 用户目标检测模型
	Validatiton := validate.UserValidationParams{}
	// 验证参数
	if helper.GetType(this.Ctx.Request.URL.String()) == 1 {

		if ok, errMsg := Validatiton.CheckAddGoalParams(params); !ok {
			utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
			this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
			return
		}
	} else {
		if ok, errMsg := Validatiton.CheckDeleteMonitorParams(params); !ok {
			utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
			this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
			return
		}
	}

	goalModel := models.UserGoalModel{}
	// 设置参数
	goal := models.UserGoal{}
	goal.Uid = params.Uid
	goal.Name = params.Name
	goal.Switch = 0
	goal.Update_time = helper.GetTimestamp()

	num, err := goalModel.DeleteUserGoal(goal)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), num)
	} else {
		utils.Log.Error("更新数据库错误 ： %v", err)
		this.RenderApiJson(apicode.DeleteGoalFailed, apicode.Msg(apicode.DeleteGoalFailed), err)
	}
	return
}

// 获取用户全部目标
func (this *UserGoalController) GetUserGoals() {
	params := validate.UserIdParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定Uid参数出错, err: %v", err) // 记录log
		this.RenderApiJsonSlice(apicode.MissParam, apicode.Msg(apicode.MissParam), err)
		return
	}
	// 用户Id检测模型
	Validatiton := validate.UserValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckUserIdParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		if this.isNewVersion {
			this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		} else {
			this.RenderOldApiJson(false, "参数错误", nil)
		}
		return
	}

	goalModel := models.UserGoalModel{}
	uid := params.Uid
	goalType := helper.GetType(this.Ctx.Request.URL.String())
	goals, err := goalModel.GetUserGoalsById(uid, goalType)

	if err == nil {
		indexModel := models.IndexDataModel{}
		obj, err1 := indexModel.GetIndexMonitor(uid)
		if err1 == nil {
			this.RenderApiJsonExtro(apicode.Success, apicode.Msg(apicode.Success), obj.Monitor_id, goals)
		}
		utils.Log.Error("查询首页监控id错误： ％v", err1)
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), goals)
	} else {
		utils.Log.Error("查询数据库错误： ％v", err)
		this.RenderApiJson(apicode.GetGoalsFailed, apicode.Msg(apicode.GetGoalsFailed), err)
	}
}
