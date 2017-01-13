package controllers

import (
	"fmt"
	"mysite/helper"
	"mysite/helper/apicode"
	"mysite/models"
	"mysite/utils"
	"mysite/validate"
)

type FoodController struct {
	BaseController
}

// 初始化
func (this *FoodController) Prepare() {

}

// 获取中国饮食类别
func (this *FoodController) GetChDietgroups() {

	foodModel := models.ChDietGroupModel{}
	// 设置参数
	diets, err := foodModel.GetChDietGroup()
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), diets)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

// 搜索中国饮食
func (this *FoodController) SearchChDietsByName() {

	params := validate.FoodNameParams{}

	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 食物检测模型
	Validatiton := validate.FoodValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckSearchChDietsByNameParams(params); !ok {
		fmt.Printf("验证参数失败, errMsg: %s, params: %v", errMsg, params)      // 记录log
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	foodModel := models.ChDietnutriModel{}
	// 设置参数
	if params.Type == 1 {
		diets, err := foodModel.SearchChDietnutriByName(params.Name)
		if err == nil {
			this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), diets)
		} else {
			utils.Log.Error("查询数据库错误 ： %v", err)
			this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
		}
	} else {
		var goupid int64 = 26
		diets, err := foodModel.SearchChDietnutriByNameHK(goupid, params.Name)
		if err == nil {
			this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), diets)
		} else {
			utils.Log.Error("查询数据库错误 ： %v", err)
			this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
		}
	}

	return
}

// 获取中国饮食
func (this *FoodController) GetChDietnutris() {
	params := validate.IndexLenParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 食物检测模型
	Validatiton := validate.FoodValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckIndexLenParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	foodModel := models.ChDietnutriModel{}
	// 设置参数
	diets, err := foodModel.GetChDietnutri(params.Index, params.Length)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), diets)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

// 获取中国饮食
func (this *FoodController) GetChDietsByGroupHK() {

	params := validate.ChDietsGroupParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 食物检测模型
	Validatiton := validate.FoodValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckGetChDietsByGroupParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	foodModel := models.ChDietnutriModel{}
	// 设置参数
	diets, err := foodModel.GetChDietsByGroupIdHK(params.Groupid)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), diets)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

// 获取用户饮食数据
func (this *FoodController) GetFoodsByUid() {

	params := validate.FoodDataParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 食物检测模型
	Validatiton := validate.FoodValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckGetFoodDataParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}
	foodModel := models.FoodDataModel{}
	// 设置参数
	foods, err := foodModel.GetFoodsByUid(params.Uid, params.Start, params.End)
	if err == nil {
		var total float64 = 0
		for _, data := range foods {
			total += data.Energy
		}
		this.RenderApiJsonExtro(apicode.Success, apicode.Msg(apicode.Success), total, foods)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

//获取常见食物
func (this *FoodController) GetCommonFoods() {
	params := validate.CommonParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 食物检测模型
	Validatiton := validate.FoodValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckCommonParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	foodModel := models.ChDietnutriModel{}
	// 设置参数
	commondiets, err := foodModel.GetCommonFoods(params.Uid)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), commondiets)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

// 添加用户饮食数据
func (this *FoodController) AddFoodsDataByUid() {

	params := validate.FoodDataParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		fmt.Printf("验证参数失败, err: %s, params: %v", err, params)
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.MissParam, apicode.Msg(apicode.MissParam))
		return
	}

	// 食物检测模型
	Validatiton := validate.FoodValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.AddFoodDataParams(params); !ok {
		fmt.Printf("验证参数失败, errMsg: %s, params: %v", errMsg, params)
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	foodModel := models.FoodDataModel{}
	food := models.FoodData{}
	food.Uid = params.Uid
	food.Name = params.Name
	food.Foodid = params.Foodid
	food.Type = params.Type
	food.Img_url = params.Img_url
	food.Value = params.Value
	food.Energy = params.Energy
	food.Unit = params.Unit
	food.Date = params.Date
	food.Source = params.Source
	food.Location = params.Location
	food.Diettime = params.Diettime
	food.Switch = 1
	food.Create_time = helper.GetTimestamp()
	food.Update_time = food.Create_time
	food.Collection = params.Collection
	// 设置参数
	id, err := foodModel.AddFoodsData(food)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), id)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

// 更新用户饮食数据
func (this *FoodController) UpdateFoodsDataById() {
	params := validate.FoodDataParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.MissParam, apicode.Msg(apicode.MissParam))
		return
	}
	// 食物检测模型
	Validatiton := validate.FoodValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckUpdateFoodDataParams(params); !ok {
		fmt.Printf("验证参数失败, errMsg: %s, params: %v", errMsg, params)
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	foodModel := models.FoodDataModel{}
	food := models.FoodData{}
	food.Id = params.Id
	food.Img_url = params.Img_url
	food.Value = params.Value
	food.Energy = params.Energy
	food.Unit = params.Unit
	food.Date = params.Date
	food.Diettime = params.Diettime
	food.Collection = params.Collection
	food.Update_time = helper.GetTimestamp()

	// 设置参数
	id, err := foodModel.UpdateFoodsData(food)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), id)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

// 删除用户食物数据
func (this *FoodController) DelFoodsData() {
	id, err := this.GetInt64("id")
	// 绑定参数
	if err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}
	// 验证参数
	if id < 1 {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", "id必须大于0", id) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	foodModel := models.FoodDataModel{}
	num, err := foodModel.DelFoodsData(id)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), num)
	} else {
		utils.Log.Error("删除数据库错误 ： %v", err)
		this.RenderApiJson(apicode.DeleteError, apicode.Msg(apicode.DeleteError), err)
	}
	return
}

// 删除用户食物数据
func (this *FoodController) GetFooddetail() {
	id, err := this.GetInt64("id")
	// 绑定参数
	if err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}
	// 验证参数
	if id < 1 {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %s", "id 必须大于1", id) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}
	foodModel := models.FoodDataModel{}
	ret, err := foodModel.GetFooddetail(id)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), ret)
	} else {
		utils.Log.Error("删除数据库错误 ： %v", err)
		this.RenderApiJson(apicode.DeleteError, apicode.Msg(apicode.DeleteError), err)
	}
	return
}

//获取用户收藏食物
func (this *FoodController) GetCollectFoods() {
	params := validate.FoodDataParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 食物检测模型
	Validatiton := validate.FoodValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckGetFoodDataParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}
	collection := helper.CollectType(this.Ctx.Request.URL.String())

	foodModel := models.ChDietnutriModel{}
	// 设置参数
	collectdiets, err := foodModel.GetChDietnutriCollect(params.Uid, params.Start, params.End, collection)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), collectdiets)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}

// 根据条码获取中国食物
func (this *FoodController) SearchChDietsByBarCode() {
	params := validate.BarCodeParams{}

	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 食物检测模型
	Validatiton := validate.FoodValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckBarCodeParams(params); !ok {
		fmt.Printf("验证参数失败, errMsg: %s, params: %v", errMsg, params)      // 记录log
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	foodModel := models.ChDietnutriModel{}
	// 设置参数
	if params.Type == 1 {
		diet, err := foodModel.SearchChDietnutriByBarCode(params.Barcode)
		if err == nil {
			if len(diet.Barcode) == 0 {
				this.RenderApiJsonEmpty(apicode.Success, apicode.Msg(apicode.Success))
			} else {
				this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), diet)
			}
		} else {
			utils.Log.Error("查询数据库错误 ： %v", err)
			this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
		}
	} else {
		//todo 预留查其它表
	}

	return
}

//取消收藏食物
func (this *FoodController) ConcelCollectFoods() {
	params := validate.CocelCollectParams{}

	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 食物检测模型
	Validatiton := validate.FoodValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckCocelCollectParams(params); !ok {
		fmt.Printf("验证参数失败, errMsg: %s, params: %v", errMsg, params)      // 记录log
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	foodModel := models.FoodDataModel{}
	// 设置参数
	err := foodModel.ConcelCollect(params.Uid, params.Foodid)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), nil)
	} else {
		utils.Log.Error("查询数据库错误 ： %v", err)
		this.RenderApiJson(apicode.QueryError, apicode.Msg(apicode.QueryError), err)
	}
	return
}
