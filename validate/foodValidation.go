package validate

import (
	"github.com/astaxie/beego/validation"
)

//中国饮食参数
type ChDietnutrisParams struct {
}

//饮食类别参数
type ChDietsGroupParams struct {
	Groupid          string `form:"groupid"`
	Chinesegroupname string `form:"chinesegroupname"`
	Groupname        string `form:"groupname"`
}

//饮食名称
type FoodNameParams struct {
	Name string `form:"name"`
	Type int64  `form:"type"`
}

// 食物数据参数
type FoodDataParams struct {
	Id         int64   `form:"id"`
	Uid        int64   `form:"uid"`
	Foodid     int64   `form:"foodid"`
	Name       string  `form:"name"`
	Type       int64   `form:"type"`
	Img_url    string  `form:"imgurl"`
	Value      float64 `form:"value"`
	Energy     float64 `form:"energy"`
	Unit       string  `form:"unit"`
	Date       int64   `form:"date"`
	Source     int64   `form:"source"`
	Location   string  `form:"location"`
	Diettime   int64   `form:"diettime"`
	Collection int64   `form:"collection"`
	Start      int64   `form:"start"`
	End        int64   `form:"end"`
}

// 开始值， length值 数据参数
type IndexLenParams struct {
	Index  int64 `form:"index"`
	Length int64 `form:"length"`
}

//常见食物接口请求参数
type CommonParams struct {
	Uid int64 `form:"uid"`
}

//根据条码获取中国食物接口请求参数
type BarCodeParams struct {
	Barcode string `form:"barcode"`
	Type    int64  `form:"type"`
}

//根据Foodid获取中国食物接口请求参数
type CocelCollectParams struct {
	Uid    int64 `form:"uid"`
	Foodid int64 `form:"foodid"`
}

// 验证参数结构体
type FoodValidationParams struct {
}

// 验证接收“添加饮食数据”参数
func (v *FoodValidationParams) CheckIndexLenParams(params IndexLenParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Index, "Index").Message("索引值不能为空")
	valid.Required(params.Length, "Length").Message("总长度不能为空")

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

// 验证接收“添加饮食数据”参数
func (v *FoodValidationParams) AddFoodDataParams(params FoodDataParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Foodid, "Foodid").Message("食物id不能为空")
	valid.Required(params.Uid, "Uid").Message("用户id不能为空")
	valid.Required(params.Name, "Name").Message("食物名称不能为空")
	valid.Required(params.Value, "Value").Message("食物量不能为空")
	valid.Required(params.Unit, "Unit").Message("食物单位不能为空")
	valid.Required(params.Source, "Source").Message("食物输入来源不能为空")
	valid.Required(params.Date, "Date").Message("食物输入日期不能为空")

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

// 验证接收更改“饮食数据”参数
func (v *FoodValidationParams) CheckUpdateFoodDataParams(params FoodDataParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Id, "Id").Message("id不能为空")

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

// 验证“获取用户饮食数据”参数
func (v *FoodValidationParams) CheckGetFoodDataParams(params FoodDataParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Uid, "Uid").Message("用户id不能为空")
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

// 验证接收“获取中国饮食”参数
func (v *FoodValidationParams) CheckGetChDietnutrisParams(params ChDietnutrisParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	// valid.Required(params.Uid, "Uid").Message("用户id不能为空")

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

// 验证接收“根据分组获取中国饮食”参数
func (v *FoodValidationParams) CheckGetChDietsByGroupParams(params ChDietsGroupParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Groupid, "Groupid").Message("分组id不能为空")
	valid.Numeric(params.Groupid, "Groupid").Message("分组id必须为数字")

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

// 验证接收“根据分组获取中国饮食”参数
func (v *FoodValidationParams) CheckSearchChDietsByNameParams(params FoodNameParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Name, "Name").Message("食物名称不能为空")

	valid.Required(params.Type, "Type").Message("搜索类型必须为数字")

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

// 验证接收“常见食物”参数
func (v *FoodValidationParams) CheckCommonParams(params CommonParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Uid, "Index").Message("用户id不能为空")

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

// 验证接收“通过条码查找中国食物”参数
func (v *FoodValidationParams) CheckBarCodeParams(params BarCodeParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Barcode, "Barcode").Message("条码值不能为空")
	valid.Required(params.Type, "Type").Message("类型不能为空")

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

// 验证接收“通过foodid取消食物收藏”参数
func (v *FoodValidationParams) CheckCocelCollectParams(params CocelCollectParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Uid, "Uid").Message("Uid不能为空")
	valid.Required(params.Foodid, "Foodid").Message("Foodid不能为空")

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
