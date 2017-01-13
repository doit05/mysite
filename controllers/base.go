package controllers

import (
	"mysite/helper"
	// "mysite/models"
	// "mysite/utils"
	"reflect"

	"github.com/astaxie/beego"
)

// 基础控制器
type BaseController struct {
	beego.Controller
	isNewVersion bool // 迁移到golang的版本标识
}

// 设置新旧版本标识
func (this *BaseController) SetIsNewVersion() {
	if len(this.Ctx.Input.Header("dcc-ver")) > 0 { // 新版的dcc_ver值为2， 旧版的没有值
		this.isNewVersion = true
	}
}

// 获取 from POST 和 PUT request body 和 query string
func (this *BaseController) getInputParams() map[string]string {
	// 包含get请求和post请求的所以参数
	params := make(map[string]string)

	for key, val := range this.Input() {
		// 注意：如果参数为 val的值是为数组 ids[]=888&ids[]=999 会被忽略
		if len(val) == 1 {
			params[key] = val[0]
		}
	}

	return params
}

// // 校验api_sign是否正确
// func (this *BaseController) CheckSign() (ok bool, errMsg string) {
// 	params := this.getInputParams()
//
// 	// 获取api_sign字段的值
// 	apiSign, signExist := params["api_sign"]
//
// 	// api_sign字段不存在 或为空
// 	if signExist == false || apiSign == "" {
// 		errMsg = "缺少api_sign参数"
// 		return
// 	}
//
// 	appId, idExist := params["app_id"]
//
// 	// app_id字段不存在 或为空
// 	if idExist == false || appId == "" {
// 		errMsg = "缺少app_id参数"
// 		return
// 	}
//
// 	// 删除api_sign字段
// 	delete(params, "api_sign")
//
// 	// 获取app_key
// 	smosModel := models.SmosModel{}
// 	appKey, err := smosModel.GetAppkey(appId)
//
// 	if err != nil { // 获取app_key出错了
// 		utils.Log.Error("获取app_key出错: appId: " + appId + ", err:" + err.Error()) // 记录log
// 		errMsg = "获取app_id对应的app_key失败"
// 		return
// 	}
//
// 	if len(appKey) == 0 { // api_key不存在
// 		utils.Log.Error("获取app_key不存在: appId: " + appId) // 记录log
// 		errMsg = "未找到app_id对应的app_key"
// 		return
// 	}
//
// 	// 验证api_sign
// 	if helper.MakeSign(params, appKey) == apiSign {
// 		ok = true
// 		return
// 	} else {
// 		errMsg = "签名不正确"
// 		return
// 	}
// }

// 渲染api json格式数据
func (this *BaseController) RenderApiJsonExtro(status string, msg string, extro, data interface{}) {
	res := helper.ApiResExtro{}
	res.Status = status
	res.Msg = msg
	res.Data = data
	res.Extro = extro
	this.Data["json"] = &res
	this.ServeJSON()
}

// 渲染api json格式数据
func (this *BaseController) RenderApiJsonExtroData(status string, msg string, extro, extrodata, data interface{}) {
	res := helper.ApiResExtro{}
	res.Status = status
	res.Msg = msg
	res.Data = data
	res.Extro = extro
	res.DataExtro = extrodata
	this.Data["json"] = &res
	this.ServeJSON()
}

// 渲染api json格式数据
func (this *BaseController) RenderApiJson(status string, msg string, data interface{}) {
	res := helper.InitApiRes()
	res.Status = status
	res.Msg = msg
	res.Data = data

	this.Data["json"] = &res
	this.ServeJSON()
}

// 渲染旧版本的api json格式数据
func (this *BaseController) RenderOldApiJson(success bool, msg interface{}, data interface{}) {
	res := new(helper.OldApiRes)

	res.Success = success
	res.ErrMsg = msg
	res.Data = data

	this.Data["json"] = &res
	this.ServeJSON()
}

// 返回空， ［］
func (this *BaseController) RenderApiJsonEmpty(status string, msg string) {
	this.RenderApiJson(status, msg, make([]interface{}, 0))
	return
}

// 渲染api jsonArray格式数据
func (this *BaseController) RenderApiJsonSlice(status string, msg string, data interface{}) {
	dataType := reflect.ValueOf(data).Kind()
	if dataType == reflect.Slice || dataType == reflect.Array {
		this.RenderApiJson(status, msg, data)
	} else {
		retData := make([]interface{}, 0)
		retData = append(retData, data)
		this.RenderApiJson(status, msg, retData)
	}
}
