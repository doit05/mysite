package validate

import (
	"github.com/astaxie/beego/validation"
)

//用户监控参数
type UserMonitorParams struct {
	Uid         int64   `form:"uid"`
	Monitor_id  int64   `form:"monitorid"`
	Create_time int64   `form:"createtime"`
	Value       float64 `form:"value"`
	Value1      float64 `form:"value1"`
	Comment     string `form:"comment"`
	Value_time  int64   `form:"valuetime"`
	Start       int64   `form:"start"`
	End         int64   `form:"end"`
}

// 验证接收“用户添加监控数据”参数
func (v *UserValidationParams) CheckMonitorDataUpParams(params UserMonitorParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Monitor_id, "Monitor_id").Message("监控id不能为空")

	valid.Required(params.Uid, "Uid").Message("用户ID不能为空")

	valid.Required(params.Value, "Value").Message("监控值不能为空")

	valid.Required(params.Value_time, "Value_time").Message("监控值的时间不能为空")

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		ok = true
	}
	return ok, errMsg
}

// 验证下载“用户监控数据”参数
func (v *UserValidationParams) CheckMonitorDataDownParams(params UserMonitorParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Monitor_id, "Monitor_id").Message("监控id不能为空")
	valid.Required(params.Uid, "Uid").Message("用户ID不能为空")
	valid.Required(params.Start, "Start").Message("起始时间不能为空")
	valid.Required(params.End, "End").Message("结束时间不能为空")

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		ok = true
	}
	return ok, errMsg
}
