package apicode

const (
	Success          string = "0000" // 成功
	SystemError      string = "-1"   // 系统错误
	MissParam        string = "-2"   // 参数缺失
	InvalidParam     string = "-3"   // 参数错误
	ApiNotFound      string = "-4"   // 接口未找到
	SignError        string = "-5"   // 签名错误
	RequestExpire    string = "-6"   // 请求已失效
	QueryError       string = "-7"   //查询出错
	DeleteError      string = "-8"   //查询出错
	InsertDataFailed string = "-9"   //插入数据失败
	SaveFileFailed   string = "-10"  //保存文件失败
	//
	DockerClientError        string = "-11" //docker client err
	DockerContainerListError string = "-12" //docker ContainerList err

	// 错误代码从4位开始
	MsgInsertFailed      string = "0001" // 插入消息失败
	MsgProcessing        string = "0002" // 同一消息正在处理中
	MsgUploaded          string = "0003" // 同一消息已经处理完成
	UserNameExist        string = "0004" //同一操作记录正在处理
	GetGoalsFailed       string = "0005" //获取用户目标失败
	PasswordUnmached     string = "0006" //用户名和密码不匹配
	NameUnFound          string = "0007" //用户名不存在
	GetSystemGoalsFailed string = "0008" //获取系统目标失败

	InsertUserFailed    string = "0010" // 插入用户失败
	ParseMetadataFailed string = "0011" // 解析Metadata数据失败
	InsertGoalFailed    string = "0100" // 插入用户目标失败
	ReceivedMsg         string = "0101" //消息已经上传

	InsertSportDataFailed string = "0110" //插入运动数据失败
	SetUserInfoFailed     string = "0111" //更新账户信息失败
	GetUserInfoFailed     string = "1000" //获取账户信息失败
	DeleteGoalFailed      string = "1001" //删除用户目标失败
	GetSexFailed          string = "2000" //用户性别没有设置
	GetAgeFailed          string = "2001" //用户年龄没有设置
	GetPalFailed          string = "2002" //"用户活动水平没有设置",
	GetWeightFailed       string = "2003" //"婴儿体重没有设置",
	GetMonitorFailed      string = "2004" //"首页默认监控没有设置",
	GetSysGoalFailed      string = "2005" //"系统中不存在这样的目标或监控",
	GetUserGoalFailed     string = "2006" //"用户没有添加这个的目标或监控",
	GetUserFoodFailed     string = "2007" //"获取用户食物失败",
	//

)

// 获取错误代码对应的提示信息
func Msg(code string) string {
	msgMap := map[string]string{
		Success:       "success",
		SystemError:   "系统错误",
		MissParam:     "参数缺失",
		InvalidParam:  "参数错误",
		ApiNotFound:   "接口未找到",
		SignError:     "签名错误",
		RequestExpire: "请求已失效",
		QueryError:    "查询出错",
		DeleteError:   "删除出错",

		MsgInsertFailed:  "上传消息失败",
		MsgProcessing:    "同一消息正在处理中",
		MsgUploaded:      "同一消息已经处理完成",
		ReceivedMsg:      "消息已经上传",
		PasswordUnmached: "用户名和密码不匹配",
		NameUnFound:      "用户名不存在",

		UserNameExist: "用户名已经存在",

		InsertUserFailed:    "插入用户失败",
		ParseMetadataFailed: "解析Metadata数据失败",

		InsertGoalFailed:      "插入用户目标失败",
		GetGoalsFailed:        "获取用户目标失败",
		GetSystemGoalsFailed:  "获取系统目标失败",
		InsertSportDataFailed: "插入运动数据失败",
		SetUserInfoFailed:     "更新账户信息失败",
		GetUserInfoFailed:     "获取账户信息失败",
		DeleteGoalFailed:      "删除用户目标失败",
		InsertDataFailed:      "插入数据失败",
		GetSexFailed:          "用户性别没有设置",
		GetAgeFailed:          "用户年龄没有设置",
		GetPalFailed:          "用户活动水平没有设置",
		GetWeightFailed:       "婴儿体重没有设置",
		GetMonitorFailed:      "首页默认监控没有设置",
		GetSysGoalFailed:      "系统中不存在这样的目标或监控",
		GetUserGoalFailed:     "用户没有添加这个的目标或监控",
		GetUserFoodFailed:     "获取用户食物失败",
		SaveFileFailed:        "保存文件失败",

		DockerClientError:        "docker client err",
		DockerContainerListError: "docker ContainerList err",
	}

	msg, exist := msgMap[code]

	if exist == false { // 不存在默认返回空字符串
		msg = ""
	}

	return msg
}
