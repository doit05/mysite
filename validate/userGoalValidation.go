package validate

import (
	"github.com/astaxie/beego/validation"
)

//用户目标参数
type UserGoalParams struct {
	Uid             int64   `form:"uid"`
	Name            string  `form:"name"`
	Type            int64   `form:"type"`
	Img_url         string  `form:"imgurl"`
	Switch          int64   `form:"switch"`
	Start           float64 `form:"start"`
	End             float64 `form:"end"`
	Weekly_change   float64 `form:"weeklychange"`
	Increase        int     `form:"increase"`
	Date            int64   `form:"date"`
	Daily_entry     float64 `form:"dailyentry"`
	Daily_foodpoint float64 `form:"dailyfoodpoint"`
}

// 验证接收“系统添加目标”参数
func (v *UserValidationParams) CheckAddSystemGoalParams(params UserGoalParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Name, "Name").Message("名称不能为空")

	valid.Required(params.Img_url, "Img_url").Message("图片地址不能为空")

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		ok = true
	}
	return ok, errMsg
}

// 验证接收“获取系统目标”参数
func (v *UserValidationParams) CheckGetSystemGoalParams(params UserGoalParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		ok = true
	}
	return ok, errMsg
}

// 验证接收“用户添加目标”参数
func (v *UserValidationParams) CheckAddGoalParams(params UserGoalParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Name, "Name").Message("目标名不能为空")

	valid.Required(params.Uid, "Uid").Message("用户ID不能为空")

	valid.Required(params.Date, "Date").Message("目标截止日期不能为空")

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		ok = true
	}
	return ok, errMsg
}

// 验证接收“用户添加监控”参数
func (v *UserValidationParams) CheckAddMonitorParams(params UserGoalParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Name, "Name").Message("监控名不能为空")

	valid.Required(params.Uid, "Uid").Message("用户ID不能为空")

	valid.Required(params.Img_url, "Img_url").Message("图片地址不能为空")

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		ok = true
	}
	return ok, errMsg
}

// 验证接收“用户删除监控”参数
func (v *UserValidationParams) CheckDeleteMonitorParams(params UserGoalParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Name, "Name").Message("监控名不能为空")

	valid.Required(params.Uid, "Uid").Message("用户ID不能为空")

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		ok = true
	}
	return ok, errMsg
}
