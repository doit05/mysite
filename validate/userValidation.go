package validate

import (
	"github.com/astaxie/beego/validation"
	"time"
)

//用户注册参数
type UserRegisterParams struct {
	Name     string `form:"name"`
	Password string `form:"password"`
	Email    string `form:"email"`
	Mobile   string `form:"mobile"`
}

//用户登录参数
type UserLoginParams struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

//用户Id参数
type UserIdParams struct {
	Uid int64 `form:"uid"`
}

//用户账户信息参数
type UserInfoParams struct {
	Id          int64   `form:"uid"`
	Name        string  `form:"name"`
	Password    string  `form:"password"`
	Email       string  `form:"email"`
	Mobile      string  `form:"mobile"`
	Truename    string  `form:"truename"`
	Personimage string  `form:"personimage"`
	Idcard      string  `form:"idcard"`
	Sourcetype  int64   `form:"sourcetype"`
	Sex         int64   `form:"sex"`
	Birthdate   string  `form:"birthdate"`
	Age         float64 `form:"age"`
	Height      int64   `form:"height"`
	Weight      int64   `form:"weight"`
	Memo        string  `form:"memo"`
	Usertype    int64   `form:"usertype"`
	Certno      string  `form:"certno"`
	Certdate    string  `form:"certdate"`

	Pal               int64  `form:"pal"`
	Gluteninallery    int64  `form:"gluteninallery"`
	Flowerallery      int64  `form:"flowerallery"`
	Peanutallery      int64  `form:"peanutallery"`
	Eggallery         int64  `form:"eggallery"`
	Treenutallery     int64  `form:"treenutallery"`
	Fishallery        int64  `form:"fishallery"`
	Eathearthealth    int64  `form:"eathearthealth"`
	Eatlowcholesterol int64  `form:"eatlowcholesterol"`
	Eatlowbloodpres   int64  `form:"eatlowbloodpres"`
	Eatforpregnant    int64  `form:"eatforpregnant"`
	Eatformombaby     int64  `form:"eatformombaby"`
	Avoidgmo          int64  `form:"avoidgmo"`
	Avoidpork         int64  `form:"avoidpork"`
	Avoidmeatandfish  int64  `form:"avoidmeatandfish"`
	Eatvegan          int64  `form:"eatvegan"`
	Paleodiet         int64  `form:"paleodiet"`
	Specialfood       string `form:"specialfood"`
	Autogeneratestps  int64  `form:"autogeneratestps"`
	Autodate          string `form:"Autodate"`
	Pregnantlevel     int64  `form:"pregnantlevel"`
}

// 验证参数结构体
type UserValidationParams struct {
}

// 验证接收“用户登录”参数
func (v *UserValidationParams) CheckLoginParams(params UserLoginParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Name, "name").Message("用户名不能为空")

	valid.Required(params.Password, "password").Message("用户密码不能为空不能为空")

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		ok = true
	}
	return ok, errMsg
}

// 验证接收“用户Id”参数
func (v *UserValidationParams) CheckUserIdParams(params UserIdParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Uid, "Uid").Message("Uid 不能为空")

	if valid.HasErrors() { // 验证不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		ok = true
	}

	if params.Uid < 1 {
		errMsg += "[Uid必须为大于1的整数]" // 将错误信息拼接起来
		ok = false
	}

	return ok, errMsg

}

//验证接收用户账户信息参数
func (v *UserValidationParams) CheckSetInfoParams(params UserInfoParams) (ok bool, errMsg string) {
	valid := validation.Validation{}

	valid.Required(params.Id, "Id").Message("uid不能为空")
	var err1 error
	var err2 error
	if len(params.Certdate) > 0{
		_, err1 = time.Parse("2006-01-02", params.Certdate)
	}
	if len(params.Birthdate) > 0{
		_, err2 = time.Parse("2006-01-02", params.Birthdate)
	}
	if valid.HasErrors() { //验证 不通过
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	} else {
		if err1 != nil {
			errMsg += "[ certdate 营养师证日期格式错误 ]"
		}else if err2 != nil {
			errMsg += "[ certdate 出生日期格式错误 ]"
		}else{
			ok = true
		}
	}
	return ok, errMsg

}
