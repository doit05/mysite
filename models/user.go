package models

import (
	"mysite/helper"
	"mysite/validate"
	// "time"
	// "fmt"

	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

// '用户信息';
type User struct {
	Id       int64  //`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id ',
	Name     string `valid:"Required;MaxSize(30)"`            //`name` varchar(32) COMMENT '用户名',
	Password string `valid:"Required;MinSize(7);MaxSize(30)"` //`password` varchar(32) COMMENT '用户密码',
	Email    string `valid:"Email;MaxSize(30)"`               //`email` varchar(32) DEFAULT NULL COMMENT '用户邮箱',
	Mobile   string `valid:"Mobile"`                          //`mobile` varchar(20) DEFAULT NULL COMMENT '用户手机号',

	Truename    string  //user_truename  varchar(40)  DEFAULT NULL COMMENT '用户姓名'
	Personimage string  //`personImage` blob COMMENT '用户图像',
	Idcard      string  //`IDCard` varchar(20) DEFAULT NULL COMMENT '身份证号码',
	Sourcetype  int64   //'0-register by user,1-register by QQ'  DEFAULT '0' COMMENT "注册来源"
	Sex         int64   // '1-man 2-woman'  DEFAULT NULL COMMENT "性别"
	Birthdate   string  //`birthDate` date DEFAULT NULL COMMENT '出生日期',
	Age         float64 //`age` int(11) DEFAULT NULL COMMENT '年龄',
	Height      int64   //`height` int(11) DEFAULT NULL COMMENT "身高"
	Weight      int64   //`weight` int(11) DEFAULT NULL COMMENT '体重kg',
	Memo        string  //`memo` text COMMENT '备注，用户可再此输入和自己相关的一些信息',
	Usertype    int64   //`userType` int(11) DEFAULT NULL COMMENT '0-用户 1-营养师 用户类型，不同类型的用户注册时必须输入的信息不同',
	Certno      string  //`certNo` varchar(30) DEFAULT NULL COMMENT '资格证编号，仅对用户类型为1的用户有效',
	Certdate    string  //`certDate` date DEFAULT NULL COMMENT '取得营养师资格的时间',

	Create_time       int64  //`create_time` int(11) DEFAULT NULL COMMENT '创建时间',
	Update_time       int64  //`update_time` int(11) DEFAULT NULL COMMENT '更新时间',
	Pal               int64  // `PAL` int(11) DEFAULT NULL,
	Gluteninallery    int64  //   `gluteninAllery` int(11) DEFAULT NULL,
	Flowerallery      int64  //   `flowerAllery` int(11) DEFAULT NULL,
	Peanutallery      int64  //   `peanutAllery` int(11) DEFAULT NULL,
	Eggallery         int64  //   `eggAllery` int(11) DEFAULT NULL,
	Treenutallery     int64  //   `treeNutAllery` int(11) DEFAULT NULL,
	Fishallery        int64  //   `fishAllery` int(11) DEFAULT NULL,
	Eathearthealth    int64  //   `eatHeartHealth` int(11) DEFAULT NULL,
	Eatlowcholesterol int64  //   `eatLowCholesterol` int(11) DEFAULT NULL,
	Eatlowbloodpres   int64  //   `eatLowBloodPres` int(11) DEFAULT NULL,
	Eatforpregnant    int64  //   `eatForPregnant` int(11) DEFAULT NULL,
	Eatformombaby     int64  //   `eatForMomBaby` int(11) DEFAULT NULL,
	Avoidgmo          int64  //   `AvoidGMO` int(11) DEFAULT NULL,
	Avoidpork         int64  //   `avoidPork` int(11) DEFAULT NULL,
	Avoidmeatandfish  int64  //   `avoidMeatAndFish` int(11) DEFAULT NULL,
	Eatvegan          int64  //   `eatVegan` int(11) DEFAULT NULL,
	Paleodiet         int64  //   `paleoDiet` int(11) DEFAULT NULL,
	Specialfood       string //   `specialFood` varchar(255) DEFAULT NULL,
	Autogeneratestps  int64  //   `autoGenerateStps` int(11) DEFAULT NULL,
	Autodate          string //   `autoDate` time DEFAULT NULL,
	Pregnantlevel     int64  //
	Monitor_id        int64
}

// 数据表名称
func (this *User) TableName() string {
	return "user_basic"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(User))
}

type UserModel struct {
	//业务模型
}

//exect SQL
func (this *UserModel) ExecSql(sql string) (rets []orm.ParamsList, err error) {

	o := orm.NewOrm()
	query := sql
	_, err = o.Raw(query).ValuesList(&rets)
	return
}

//插入一条数据
func (this *UserModel) InsertUser(user User) (Id int64, err error) {
	user.Create_time = helper.GetTimestamp()
	user.Update_time = user.Create_time
	o := orm.NewOrm()
	str := "INSERT INTO user_basic (name, password, mobile, email, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := o.Raw(str, user.Name, user.Password, user.Mobile, user.Email, user.Create_time, user.Update_time).Exec()
	if err == nil {
		Id, err = res.LastInsertId()
	}

	// Id, err = o.Insert(&user)
	return
}

//用户登录时，通过用户名和密码验证用户名和密码是否匹配
func (this *UserModel) CheckUserPass(u User) (id int64, err error) {
	user, err1 := this.GetUserByName(u.Name)
	if err1 == nil && user.Password == helper.Md5(u.Password) {
		id = user.Id
	}
	return
}

//用户名是否存在
//通过用户名获取用户
func (this *UserModel) ExistUserName(name string) (ok bool, err error) {
	ok = true
	_, err = this.GetUserByName(name)
	if err == orm.ErrNoRows {
		ok = false
	}
	return
}

//通过用户名获取用户
func (this *UserModel) GetUserByName(name string) (u User, err error) {
	o := orm.NewOrm()
	query := "select * from user_basic where name = '" + name + "'"
	err = o.Raw(query).QueryRow(&u)
	return
}

//用户注册时，检测参数
func (v *UserModel) CheckRegisterParams(params validate.UserRegisterParams) (ok bool, errMsg string) {
	valid := validation.Validation{}
	u := User{Name: params.Name, Password: params.Password, Email: params.Email, Mobile: params.Mobile}
	b, err := valid.Valid(&u)
	if err != nil {
		errMsg = err.Error()
	}
	if !b {
		for _, err := range valid.Errors {
			errMsg += "[" + err.Message + "]" // 将错误信息拼接起来
		}
	}
	if b && err == nil {
		ok = true
	}
	return
}

//更新用户账户信息数据
func (this *UserModel) UpdateUserInfo(userinfo User) (row int64, err error) {
	o := orm.NewOrm()
	obj := orm.Params{
		"update_time": userinfo.Update_time,
	}
	if len(userinfo.Name) > 0 {
		obj["Name"] = userinfo.Name
	}
	if len(userinfo.Password) > 0 {
		obj["Password"] = userinfo.Password
	}
	if len(userinfo.Email) > 0 {
		obj["Email"] = userinfo.Email
	}
	if len(userinfo.Mobile) > 0 {
		obj["Mobile"] = userinfo.Mobile
	}
	if len(userinfo.Truename) > 0 {
		obj["Truename"] = userinfo.Truename
	}
	if len(userinfo.Personimage) > 0 {
		obj["Personimage"] = userinfo.Personimage
	}
	if len(userinfo.Idcard) > 0 {
		obj["Idcard"] = userinfo.Idcard
	}
	if (userinfo.Sex) > 0 {
		obj["Sex"] = userinfo.Sex
	}
	if len(userinfo.Birthdate) > 0 {
		obj["Birthdate"] = userinfo.Birthdate
	}
	if userinfo.Age > 0 {
		obj["Age"] = userinfo.Age
	}
	if userinfo.Height > 0 {
		obj["Height"] = userinfo.Height
	}
	if userinfo.Weight > 0 {
		obj["Weight"] = userinfo.Weight
	}
	if len(userinfo.Memo) > 0 {
		obj["Memo"] = userinfo.Memo
	}
	if userinfo.Usertype > 0 {
		obj["Usertype"] = userinfo.Usertype
	}
	if len(userinfo.Certno) > 0 {
		obj["Certno"] = userinfo.Certno
	}
	if len(userinfo.Certdate) > 0 {
		obj["Certdate"] = userinfo.Certdate
	}
	if userinfo.Pal > 0 {
		obj["Pal"] = userinfo.Pal
	}
	if (userinfo.Gluteninallery) > 0 {
		obj["Gluteninallery"] = userinfo.Gluteninallery
	}
	if userinfo.Flowerallery > 0 {
		obj["Flowerallery"] = userinfo.Flowerallery
	}
	if userinfo.Peanutallery > 0 {
		obj["Peanutallery"] = userinfo.Peanutallery
	}
	if userinfo.Eggallery > 0 {
		obj["Eggallery"] = userinfo.Eggallery
	}
	if (userinfo.Treenutallery) > 0 {
		obj["Treenutallery"] = userinfo.Treenutallery
	}
	if userinfo.Fishallery > 0 {
		obj["Fishallery"] = userinfo.Fishallery
	}
	if userinfo.Eathearthealth > 0 {
		obj["Eathearthealth"] = userinfo.Eathearthealth
	}
	if userinfo.Eatlowcholesterol > 0 {
		obj["Eatlowcholesterol"] = userinfo.Eatlowcholesterol
	}
	if (userinfo.Eatlowbloodpres) > 0 {
		obj["Eatlowbloodpres"] = userinfo.Eatlowbloodpres
	}
	if userinfo.Eatforpregnant > 0 {
		obj["Eatforpregnant"] = userinfo.Eatforpregnant
	}
	if userinfo.Eatformombaby > 0 {
		obj["Eatformombaby"] = userinfo.Eatformombaby
	}
	if userinfo.Avoidgmo > 0 {
		obj["Avoidgmo"] = userinfo.Avoidgmo
	}
	if userinfo.Avoidpork > 0 {
		obj["Avoidpork"] = userinfo.Avoidpork
	}
	if userinfo.Avoidmeatandfish > 0 {
		obj["Avoidmeatandfish"] = userinfo.Avoidmeatandfish
	}
	if userinfo.Eatvegan > 0 {
		obj["Eatvegan"] = userinfo.Eatvegan
	}
	if userinfo.Paleodiet > 0 {
		obj["Paleodiet"] = userinfo.Paleodiet
	}
	if len(userinfo.Specialfood) > 0 {
		obj["Specialfood"] = userinfo.Specialfood
	}
	if userinfo.Autogeneratestps > 0 {
		obj["Autogeneratestps"] = userinfo.Autogeneratestps
	}
	// if userinfo.Autodate > 0{
	// 	obj["Autodate"] = userinfo.Autodate
	// }
	if userinfo.Pregnantlevel > 0 {
		obj["Pregnantlevel"] = userinfo.Pregnantlevel
	}
	_, err = o.QueryTable("user_basic").Filter("id", userinfo.Id).Update(obj)
	return
}

//通过用户Id获取用户账户信息
func (this *UserModel) GetUserInfoById(id int64) (user User, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("user_basic").Filter("id", id).One(&user)
	return
}

//通过用户Id获取用户账户信息
func (this *UserModel) GetUserDris(uid int64) (data ChDris, err error) {
	user, err1 := this.GetUserInfoById(uid)
	if err1 != nil {
		err = err1
		return
	}
	agestart, ageend := helper.GetAgeScore(user.Age)
	dris_model := FoodDataModel{}
	data, err = dris_model.GetChdriByAge("能量", user.Pal, user.Sex, agestart, ageend)
	fmt.Println(data, "err : ", err)
	return
}
