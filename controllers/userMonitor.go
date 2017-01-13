package controllers

import (
	"fmt"
	"mysite/helper"
	"mysite/helper/apicode"
	"mysite/models"
	"mysite/utils"
	"mysite/validate"
)

type UserMonitorController struct {
	BaseController
}

// 初始化
func (this *UserMonitorController) Prepare() {

}

// 用户上传监控数据upmonitordata
func (this *UserMonitorController) UpData() {

	params := validate.UserMonitorParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		fmt.Printf("绑定参数出错, errMsg: %s, params: %v", err, params)
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 用户监控检测模型
	Validatiton := validate.UserValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckMonitorDataUpParams(params); !ok {
		fmt.Printf("验证参数失败, errMsg: %s, params: %v", errMsg, params)
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		if this.isNewVersion {
			this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		} else {
			this.RenderOldApiJson(false, "参数错误", nil)
		}
		return
	}

	monitorModel := models.UserMonitorModel{}
	// 设置参数
	monitor := models.UserMonitor{}
	monitor.Uid = params.Uid
	monitor.Monitor_id = params.Monitor_id
	monitor.Value_time = params.Value_time
	monitor.Value = params.Value
	monitor.Value1 = params.Value1
	monitor.Comment = params.Comment
	monitor.Create_time = helper.GetTimestamp()

	id, err := monitorModel.UpdateUserMonitor(monitor)
	if err == nil {
		_, err = monitorModel.SetUserMonitorLast(params.Monitor_id, id)
		if err != nil {
			utils.Log.Error("修改用户监控lastid失败错误 ： %v", err)
		}
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), id)
	} else {
		utils.Log.Error("插入数据库错误 ： %v", err)
		this.RenderApiJson(apicode.InsertDataFailed, apicode.Msg(apicode.InsertDataFailed), err)
	}
	return
}

// 获取用户监控数据
func (this *UserMonitorController) DownData() {
	params := validate.UserMonitorParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonSlice(apicode.MissParam, apicode.Msg(apicode.MissParam), err)
		return
	}

	// 用户Id检测模型
	Validatiton := validate.UserValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckMonitorDataDownParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	monitorModel := models.UserMonitorModel{}
	//判断是否是微量元素
	goalModel := models.UserGoalModel{}
	goal, err := goalModel.GetUserGoalById(params.Monitor_id)
	if err != nil {
		utils.Log.Error("查询数据库错误： ％v", err)
		this.RenderApiJson(apicode.GetUserGoalFailed, apicode.Msg(apicode.GetUserGoalFailed), err)
	}
	enName := helper.GetEnName(goal.Name)
	if len(enName) > 0 {
		monitors := []models.UserMonitor{}
		start := params.Start
		end := params.End
		for ; start < end; start += 86400 {
			monitor, err := monitorModel.GetUserNutriInfo(params.Uid, start, start+86400, enName)
			monitor.Uid = params.Uid
			monitor.Monitor_id = params.Monitor_id
			if err != nil {
				utils.Log.Error("查询数据库错误： ％v", err)
				continue
				// this.RenderApiJson(apicode.GetUserFoodFailed, apicode.Msg(apicode.GetUserFoodFailed), err)
			}
			monitor.Value_time = start
			monitors = append(monitors, monitor)

		}
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), monitors)
	}
	monitors, err := monitorModel.GetUserMonitorsById(params.Uid, params.Monitor_id, params.Start, params.End)

	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), monitors)
	} else {
		utils.Log.Error("查询数据库错误： ％v", err)
		this.RenderApiJson(apicode.GetGoalsFailed, apicode.Msg(apicode.GetGoalsFailed), err)
	}
}

// 获取用户监控最后一条数据
func (this *UserMonitorController) GetUserMonitorLast() {
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
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	monitorModel := models.UserMonitorModel{}

	monitors, err := monitorModel.GetUserMonitorLast(params.Uid)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), monitors)
	} else {
		utils.Log.Error("查询数据库错误： ％v", err)
		this.RenderApiJson(apicode.GetGoalsFailed, apicode.Msg(apicode.GetGoalsFailed), err)
	}
}
