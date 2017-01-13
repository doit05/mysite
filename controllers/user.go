package controllers

import (
	"encoding/json"
	"mysite/helper"
	"mysite/helper/apicode"
	"mysite/models"
	"mysite/utils"
	"mysite/validate"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	// "fmt"
)

type UserController struct {
	BaseController
}

var (
	globalSessions *session.Manager
)

// 初始化
func (this *UserController) Prepare() {
	cc := `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "admin", "cookieLifeTime": 3600, "providerConfig": "www.doit05.top:6379,100,redis123"}`
	cf2 := new(session.ManagerConfig)
	cf2.EnableSetCookie = true
	err := json.Unmarshal([]byte(cc), cf2)
	if err == nil {
		globalSessions, err = session.NewManager("redis", cf2)
		if err != nil {
			utils.Log.Error("redis err : %v", err)
		}
		go globalSessions.GC()
	} else {

	}
}

// 执行SQL
func (this *UserController) ExecSql() {
	sql := this.GetString("sql")
	if sql == "" {
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}
	// 用户模型
	um := models.UserModel{}
	rets, err := um.ExecSql(sql)
	if err != nil {
		utils.Log.Error("exect SQL error : %s , err: %v", sql, err) // 记录log
		this.RenderApiJson("506", "执行出错", err)
		return
	}
	this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), rets)

}

func (this *UserController) Login() {

	params := validate.UserLoginParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定Upload参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 用户登录验证模型
	Validatiton := validate.UserValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckLoginParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	w := this.Ctx.ResponseWriter
	r := this.Ctx.Request
	us, err := globalSessions.SessionStart(w, r)
	defer us.SessionRelease(w)

	if err != nil {
		this.RenderApiJson("507", err.Error(), http.StatusInternalServerError)
		return
	}

	sessionid := this.Ctx.GetCookie("gosessionid")
	utils.Log.Error("cookeis : %s \n", this.Ctx.Request.Cookies())
	if len(sessionid) > 0 {
		uid := us.Get(sessionid)
		utils.Log.Error("302 sessionid : %s uname : %s \n", sessionid, us.Get(uid))

		if us.Get(uid) == params.Name {
			w.WriteHeader(http.StatusFound)
			this.RenderApiJsonSlice("302", apicode.Msg(apicode.Success), uid)
			return
		}
	}

	// 用户模型
	um := models.UserModel{}

	//用户名不存在
	exist, err := um.ExistUserName(params.Name)
	if exist == false {
		utils.Log.Error("查询出错 : %v", err) // 记录log
		this.RenderApiJsonSlice(apicode.NameUnFound, apicode.Msg(apicode.NameUnFound), err)
		return
	}

	//检测密码是否匹配
	user := models.User{Name: params.Name, Password: params.Password}
	if ok, err := um.CheckUserPass(user); ok > 0 && err == nil {
		uid := strconv.FormatInt(ok, 10)
		us.Set(uid, params.Name)
		us.Set(sessionid, uid)

		// this.Ctx.SetCookie("sessionid", uid, 86400, "/")

		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), ok)
	} else { //检测不通过
		this.RenderApiJsonEmpty(apicode.PasswordUnmached, apicode.Msg(apicode.PasswordUnmached))

	}
	return
}

// 初始化
func (this *UserController) Register() {

	params := validate.UserRegisterParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定Upload参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	// 用户模型
	um := models.UserModel{}
	// 验证参数
	if ok, errMsg := um.CheckRegisterParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %+v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	//用户名重复,返回空
	if u, err1 := um.GetUserByName(params.Name); err1 == nil && u.Id > 0 {
		this.RenderApiJsonEmpty(apicode.UserNameExist, apicode.Msg(apicode.UserNameExist))
		return
	}

	user := models.User{}
	user.Name = params.Name
	user.Password = helper.Md5(params.Password)
	user.Mobile = params.Mobile
	user.Email = params.Email
	user.Create_time = helper.GetTimestamp()
	user.Update_time = user.Create_time
	// 设置参数

	id, err := um.InsertUser(user)
	if err == nil && id > 0 {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), id)
	} else {
		utils.Log.Error("插入数据库错误 ： %v", err)
		this.RenderApiJson(apicode.InsertUserFailed, apicode.Msg(apicode.InsertUserFailed), err)
	}
	return
}

//更新用户账户信息
func (this *UserController) SetUserInfo() {
	params := validate.UserInfoParams{}
	//绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定Upload参数出错, err: %v", err) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}
	//用户账户信息模型
	Validatiton := validate.UserValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckSetInfoParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		return
	}

	userInfoModel := models.UserModel{}
	//设置参数
	userinfo := models.User{}
	userinfo.Id = params.Id
	userinfo.Name = params.Name
	if len(params.Password) > 0 {
		userinfo.Password = helper.Md5(params.Password)
	}
	userinfo.Email = params.Email
	userinfo.Mobile = params.Mobile
	userinfo.Truename = params.Truename
	userinfo.Personimage = params.Personimage
	userinfo.Idcard = params.Idcard
	userinfo.Sourcetype = params.Sourcetype
	userinfo.Sex = params.Sex
	userinfo.Birthdate = params.Birthdate
	userinfo.Age = params.Age
	userinfo.Height = params.Height
	userinfo.Weight = params.Weight
	userinfo.Memo = params.Memo
	userinfo.Usertype = params.Usertype
	userinfo.Certno = params.Certno
	userinfo.Certdate = params.Certdate
	userinfo.Pal = params.Pal
	userinfo.Gluteninallery = params.Gluteninallery
	userinfo.Flowerallery = params.Flowerallery
	userinfo.Peanutallery = params.Peanutallery
	userinfo.Eggallery = params.Eggallery
	userinfo.Treenutallery = params.Treenutallery
	userinfo.Fishallery = params.Fishallery
	userinfo.Eathearthealth = params.Eathearthealth
	userinfo.Eatlowcholesterol = params.Eatlowcholesterol
	userinfo.Eatlowbloodpres = params.Eatlowbloodpres
	userinfo.Eatforpregnant = params.Eatforpregnant
	userinfo.Eatformombaby = params.Eatformombaby
	userinfo.Avoidgmo = params.Avoidgmo
	userinfo.Avoidpork = params.Avoidpork
	userinfo.Avoidmeatandfish = params.Avoidmeatandfish
	userinfo.Eatvegan = params.Eatvegan
	userinfo.Paleodiet = params.Paleodiet
	userinfo.Specialfood = params.Specialfood
	userinfo.Autogeneratestps = params.Autogeneratestps
	userinfo.Autodate = params.Autodate
	userinfo.Pregnantlevel = params.Pregnantlevel
	userinfo.Update_time = helper.GetTimestamp()

	row, erro := userInfoModel.UpdateUserInfo(userinfo)
	if erro == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), row)
	} else {
		utils.Log.Error("更新数据库错误 ： %v", erro)
		this.RenderApiJson(apicode.SetUserInfoFailed, apicode.Msg(apicode.SetUserInfoFailed), erro)
	}
	return
}

//获取用户账户信息
func (this *UserController) GetUserInfo() {
	params := validate.UserIdParams{}
	// 绑定参数
	if err := this.ParseForm(&params); err != nil {
		utils.Log.Error("绑定Uid参数出错, err: %v", err) // 记录log
		this.RenderApiJsonSlice(apicode.MissParam, apicode.Msg(apicode.MissParam), err)
		return
	}
	// 用户Id检测模型
	Validatiton := validate.UserValidationParams{}

	// 验证参数
	if ok, errMsg := Validatiton.CheckUserIdParams(params); !ok {
		utils.Log.Error("验证参数失败, errMsg: %s, params: %v", errMsg, params) // 记录log
		if this.isNewVersion {
			this.RenderApiJsonEmpty(apicode.InvalidParam, apicode.Msg(apicode.InvalidParam))
		} else {
			this.RenderOldApiJson(false, "参数错误", nil)
		}
		return
	}

	userModel := models.UserModel{}
	id := params.Uid

	userInfo, err := userModel.GetUserInfoById(id)
	if err == nil {
		this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), userInfo)
	} else {
		utils.Log.Error("查询数据库错误： ％v", err)
		this.RenderApiJson(apicode.GetUserInfoFailed, apicode.Msg(apicode.GetUserInfoFailed), err)
	}

}
