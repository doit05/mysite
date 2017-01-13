package controllers

import (
	"mysite/helper/apicode"
	"mysite/models"
	"mysite/utils"
	"mysite/validate"
)

type ChdrisController struct {
	BaseController
}

// 初始化
func (this *ChdrisController) Prepare() {

}

//获取中国营养食物参考
func (this *ChdrisController) GetChNutriRefs() {
	params := validate.ChnutrirefValidationParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 检测模型
	Validatiton := validate.ReferenceValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckChRefParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	chNutriRefModel := models.ChDrisModel{}
	//设置参数
	refs, err := chNutriRefModel.GetChDietNutriRef(params.Agestart, params.Ageend, params.Sex, params.PAL)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), refs)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}
