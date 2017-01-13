package validate

import "github.com/astaxie/beego/validation"

// 验证参数结构体
type BarCodeValidation struct {
}

type BarCodeValidationParams struct {
	Name string `form:"name"` //食物名称
	Code string `form:"code"` //食物条码
}

// 验证接收“添加饮食数据”参数
func (v *BarCodeValidation) CheckBarCodeParams(params BarCodeValidationParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Name, "Name").Message("食物名称不能为空")
	valid.Required(params.Code, "Code").Message("条码不能为空")

	if valid.HasErrors() {
		// 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		ok = true
	}
	return ok, errMsg
}
