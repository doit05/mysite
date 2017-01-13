package validate

import "github.com/astaxie/beego/validation"

// 验证参数结构体
type ReferenceValidationParams struct {
}

type ChnutrirefValidationParams struct {
	Agestart float64 `form:"agestart"` //年龄
	Ageend   float64 `form:"ageend"`   //年龄
	Sex      int64   `form:"sex"`      //性别
	PAL      int64   `form:"pal"`      //身体活动水平
}

// 验证接收“添加饮食数据”参数
func (v *ReferenceValidationParams) CheckChRefParams(params ChnutrirefValidationParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Agestart, "Agestart").Message("年龄范围起始值不能为空")
	valid.Required(params.Ageend, "Ageend").Message("年龄范围终止值不能为空")
	valid.Required(params.Sex, "Sex").Message("性别不能为空")
	valid.Required(params.PAL, "PAL").Message("身体活动水平不能为空")

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
