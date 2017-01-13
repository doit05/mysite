package controllers

import (
	"mysite/helper/apicode"
	"mysite/models"
	"mysite/utils"
	"mysite/validate"
)

type BarCodeController struct {
	BaseController
}

// 初始化
func (this *BarCodeController) Prepare() {

}

func (this *BarCodeController) AddBarCode() {
	params := validate.BarCodeValidationParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 检测模型
	Validatiton := validate.BarCodeValidation{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckBarCodeParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}
	BarCodeModel := models.BarCodeModel{}
	BarCode := models.BarCode{}
	BarCode.Code = params.Code
	BarCode.Name = params.Name
	//设置参数
	num, id, err := BarCodeModel.AddBarCode(BarCode)
	if id >= 0 && err == nil {
		if num > 0 {
			this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), num)
		} else {
			this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), id)
		}
	} else {
		utils.Log.Error("插入数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}
