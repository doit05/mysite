package validate

import (
	"time"

	"github.com/astaxie/beego/validation"
)

//用户运动数据
type SportsAddParams struct {
	Uid         int64   `form:"uid"`
	Name        string  `form:"name"`
	Type        int64   `form:"type"`
	Img_url     string  `form:"imgurl"`
	Value       float64 `form:"value"`
	Energy      float64 `form:"energy"`
	Unit        string  `form:"unit"`
	Date        int64   `form:"date"`
	Source      int64   `form:"source"`
	Create_time int64   `form:"createtime"`
	Update_time int64   `form:"updatetime"`
	Start       int64   `form:"start"`
	End         int64   `form:"end"`
}

//用户运动类型
type SportTypeParams struct {
	Id          int64   `form:"id"`
	Uid         int64   `form:"uid"`
	Name        string  `form:"name"`
	Type        int64   `form:"type"`
	Energy      float64 `form:"energy"`
	Unit        string  `form:"unit"`
	Imgurl      string  `form:"imgurl"`
	Create_time int64   `form:"createtime"`
	Update_time int64   `form:"updatetime"`
}

// '首页图表数据';用户获取首页数据参数
type IndexDataParams struct {
	Id            int64   `form:"id"`
	Uid           int64   `form:"uid"`
	Monitor_id    int64   `form:"monitorid"`
	Max_energy    float64 `form:"maxenergy"`
	Normal_sum    float64 `form:"normalsum"`
	Normal_energy float64 `form:"normalenergy"`
	Total_sum     float64 `form:"totalsum"`
	Total_energy  float64 `form:"totalenergy"`
	Datetime      int64   `form:"datetime"`
	Create_time   int64
	Update_time   int64
}

// '用户食谱数据'参数;
type CookBookParams struct {
	Id          int64  `form:"id"`
	Uid         int64  `form:"uid"`
	Foodid      int64  `form:"foodid"`
	Diettime    int64  `form:"diettime"`
	Source      string `form:"source"`
	Week        int64  `form:"week"`
	Fooddate    string `form:"fooddate"`
	Foodname    string `form:"foodname"`
	Indexid     int64  `form:"indexid"`
	Create_time int64
	Update_time int64
	Ispurchase  int64 `form:"ispurchase"`
	Goodorbad   int64 `form:"goodorbad"`
}

// 验证参数结构体
type DailyValidationParams struct {
}

// 验证添加用户食谱参数
func (v *DailyValidationParams) CheckSetUserCookBookParams(params CookBookParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Uid, "Uid").Message("用户id不能为空")
	valid.Required(params.Source, "Source").Message("食物来源表不能为空")
	valid.Required(params.Indexid, "Indexid").Message("食物id不能为空")
	valid.Required(params.Diettime, "Diettime").Message("早中晚")

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

// 验证获取用户食谱参数
func (v *DailyValidationParams) CheckGetUserCookBookParams(params CookBookParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Uid, "Uid").Message("用户id不能为空")
	valid.Required(params.Fooddate, "Fooddate").Message("监控id不能为空")
	_, err1 := time.Parse("2006-01-02", params.Fooddate)
	if valid.HasErrors() {
		// 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}

	} else {
		if err1 != nil {
			errMsg += "[ certdate 日期格式错误 ]"
		} else {
			ok = true
		}

	}
	return ok, errMsg
}

// 验证设置首页监控id
func (v *DailyValidationParams) CheckSetIndexMonitorParams(params IndexDataParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Uid, "Uid").Message("用户id不能为空")
	valid.Required(params.Monitor_id, "Monitor_id").Message("监控id不能为空")

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

// 验证接收“获取首页数据”参数
func (v *DailyValidationParams) CheckIndexDataParams(params IndexDataParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Uid, "Uid").Message("用户id不能为空")
	valid.Required(params.Datetime, "Datetime").Message("日期不能为空")

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

// 验证接收“用户上传运动数据”参数
func (v *DailyValidationParams) CheckAddSportTypeParams(params SportTypeParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Uid, "Uid").Message("用户id不能为空")
	valid.Required(params.Name, "Name").Message("运动名称不能为空")
	valid.Required(params.Energy, "Energy").Message("运动消耗不能为空")
	valid.Required(params.Unit, "Unit").Message("单位不能为空")
	valid.Required(params.Imgurl, "Imgurl").Message("图片地址不能为空")

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

// 验证接收“用户上传运动数据”参数
func (v *DailyValidationParams) CheckAddSportsParams(params SportsAddParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Uid, "Uid").Message("用户id不能为空")
	valid.Required(params.Name, "Name").Message("运动名称不能为空")
	valid.Required(params.Img_url, "Img_url").Message("图片地址不能为空")
	valid.Required(params.Value, "Value").Message("运动值不能为空")
	valid.Required(params.Energy, "Energy").Message("运动总消耗不能为空")
	valid.Required(params.Unit, "Unit").Message("运动消耗不能为空")
	valid.Required(params.Date, "Date").Message("运动日期不能为空")
	valid.Required(params.Source, "Source").Message("运动数据来源不能为空")

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

// 验证读取“用户运动数据”参数
func (v *DailyValidationParams) CheckGetSportsParams(params SportsAddParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Uid, "Uid").Message("用户id不能为空")
	// valid.Required(params.Name, "Name").Message("运动名称不能为空")
	valid.Required(params.Start, "Start").Message("起始时间不能为空")
	valid.Required(params.End, "End").Message("结束时间不能为空")

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
