package helper

import (
				"crypto/md5"
				"encoding/hex"
				"fmt"
				"log"
				"mysite/helper/apicode"
				"os"
				"path/filepath"
				"reflect"
				"regexp"
				"sort"
				"strings"
				"time"
)

func GetAgeScore(age float64) (agestart, ageend float64) {
				switch {
				case age < 0.6:
								agestart = 0
								ageend = 0.5
				case age < 1:
								agestart = 0.6
								ageend = 1
				case age > 10 && age < 14:
								agestart = 11
								ageend = 13
				case age > 13 && age < 18:
								agestart = 14
								ageend = 17
				case age > 17 && age < 50:
								agestart = 18
								ageend = 49
				case age > 49 && age < 65:
								agestart = 50
								ageend = 64
				case age > 64 && age < 80:
								agestart = 65
								ageend = 80
				case age > 79:
								agestart = 80
								ageend = 200
				default:
								ageend = age
								agestart = age
				}
				return
}

func GetNum(str string) (num string) {
				re := regexp.MustCompile(`\d+`)
				return re.FindString(str)
}

func GetType(url string) (goaltype int64) {
				goaltype = 1
				re := regexp.MustCompile("monitor")
				if re.MatchString(url) {
								goaltype = 2
				}
				return goaltype
}

func CollectType(url string) (goaltype int64) {
				goaltype = 0
				re := regexp.MustCompile("collection")
				if re.MatchString(url) {
								goaltype = 1
				}
				return goaltype
}

//判断是否系统的，还是用户请求
func GetSystemType(url string) (gettype int64) {
				gettype = 1
				re := regexp.MustCompile("system")
				if re.MatchString(url) {
								gettype = 2
				}
				return
}

//
func ChangeSwitch(url string) (value int64) {
				value = 1
				re := regexp.MustCompile("delete")
				if re.MatchString(url) {
								value = 0
				}
				return
}

// md5加密
func Md5(s string) string {
				h := md5.New()
				h.Write([]byte(s))
				return hex.EncodeToString(h.Sum(nil))
}

// md5加密
func Md5Bytes(s []byte) string {
				h := md5.New()
				h.Write(s)
				return hex.EncodeToString(h.Sum(nil))
}

// 判断某个字符串是否在slice中, 类似php的in_array()函数
func InSlice(value string, slice []string) bool {
				for _, v := range slice {
								if v == value {
												return true
								}
				}
				return false
}

// 判断某个对象是否在slice中, 类似php的in_array()函数
func InSliceObj(value interface{}, slice []interface{}) bool {
				for _, v := range slice {
								if v == value {
												return true
								}
				}
				return false
}

// slice数组相加
func ContractSlice(list1, list2 []interface{}) (list []interface{}) {

				return
}

// 写入log到文件
func Log2File(file string, content string) {

				// log 目录
				logDir := GetRootPath() + "/logs/" + time.Now().Format("20060102")

				// 创建目录
				err := os.MkdirAll(logDir, 0666)
				if err != nil {
								log.Fatalf("error can't mkdir logs: %v", err)
				}

				// log文件路径
				file = logDir + "/" + file + ".log"

				// 打开文件
				f, ferr := os.OpenFile(file, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
				if ferr != nil {
								log.Fatalf("error opening file: %v", ferr)
				}
				defer f.Close()

				// 记录错误到文件
				log.SetOutput(f)

				log.Println(" " + content)
}

// 获取入口文件的绝对路径
func GetRootPath() string {
				dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
				if err != nil {
								log.Fatal(err)
				}
				return dir
}

// 是否是生产环境
func IsProductionEnv() bool {
				if os.Getenv("PAY_HOST") == "prod" {
								// 生产环境
								return true
				}
				// 开发或测试环境
				return false
}

// 获取配置的前缀
// 例如开发环境返回 dev:: 生产环境返回 prod::
func GetConfigPrifix() string {
				str := "::"
				if IsProductionEnv() {
								str = "prod" + str
				} else {
								str = "dev" + str
				}

				return str
}

// 手机号码格式是否正确
func IsMobile(mobile string) bool {
				var validID = regexp.MustCompile(`^\d{11}$`)
				return validID.MatchString(mobile)
}

// 生成api_sign
func MakeSign(params map[string]string, screct_key string) string {

				if len(params) == 0 {
								return ""
				}

				// 获取所有键名，用于排序
				var keys []string
				for k := range params {
								keys = append(keys, k)
				}

				// 按键正向排序
				sort.Strings(keys)

				// 将map拼接成字符串
				var sign_data string
				for _, v := range keys {
								sign_data += v + "=" + params[v] + "&"
				}

				// 末尾添加密钥
				// 例如：action=get_order_list&client_time=1451830538&user_id=3332&4b111cc14a33b88e37e2e2934f493458
				sign_data += screct_key

				return Md5(sign_data)
}

type ApiRes struct {
				Status string      `json:"status"`
				Msg    string      `json:"msg"`
				Data   interface{} `json:"data"`
}

type ApiResExtro struct {
				Status    string      `json:"status"`
				Msg       string      `json:"msg"`
				Data      interface{} `json:"data"`
				Extro     interface{} `json:"extro"`
				DataExtro interface{} `json:"dataextro"`
}

type OldApiRes struct {
				Success bool
				ErrMsg  interface{}
				Data    interface{}
}

// 初始化api返回数据
func InitApiRes() *ApiRes {
				apiRes := new(ApiRes)
				apiRes.Status = apicode.Success
				apiRes.Msg = apicode.Msg(apicode.Success)
				apiRes.Data = make(map[string]interface{})

				return apiRes
}

// 获取错误代码对应的提示信息
func GetEnName(code string) string {
				Ch2EnMap := map[string]string{
								"蛋白质":   "Protein",
								"碳水化合物": "Carbohydrate",
								"脂肪":    "Fat",
								"纤维素":   "Fiber",
								"糖":     "Sugar",
								"维生素C":  "Vitaminc",
								"维生素A":  "Vitamina",
								"维生素E":  "Vitamine",
								"维生素B1": "Vitaminb1",
								"维生素B2": "Vitaminb2",
								"维生素B3": "Vitaminb3",
								"胆固醇":   "Cholesterol",
								"镁":     "Magnesiummg",
								"钙":     "Calciumca",
								"铁":     "Ironfe",
								"锌":     "Zinczn",
								"铜":     "Coppercu",
								"锰":     "Manganesemn",
								"钾":     "Kaliumk",
								"磷":     "Phosphorp",
								"钠":     "Sodiumna",
								"硒":     "Seleniumse",
				}
				msg, exist := Ch2EnMap[code]

				if exist == false {
								// 不存在默认返回空字符串
								msg = ""
				}
				return msg
}

func GetFloat(unk interface{}) (float64, error) {
				var floatType = reflect.TypeOf(float64(0))
				v := reflect.ValueOf(unk)
				v = reflect.Indirect(v)
				if !v.Type().ConvertibleTo(floatType) {
								return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
				}
				fv := v.Convert(floatType)
				return fv.Float(), nil
}

func GetNumFromStr(str string) string {
				re := regexp.MustCompile("[0-9]+")
				return re.FindString(str)
}

func GetFloatFromStr(str, tail string) string {
				re := regexp.MustCompile("([0-9]+.[0-9]+)" + tail)
				return re.FindString(str)
}

func GetDegressStr(str string) (ret string) {
				if strings.Contains(str, "偏高") {
								return "偏高"
				}
				if strings.Contains(str, "偏低") {
								return "偏低"
				}
				return
}

func checkFileExist(path string) (ret bool,err error) {
				_, err = os.Open(path) // For read access.
				ret = err == nil
				return

}

// fileName:文件名字(带全路径)
// content: 写入的内容
func appendToFile(fileName string, content string) error {
				// 以只写的模式，打开文件
				f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
				if err != nil {
								return err
				} else {
								// 查找文件末尾的偏移量
								n, _ := f.Seek(0, os.SEEK_END)
								// 从末尾的偏移量开始写入内容
								_, err = f.WriteAt([]byte(content), n)
				}
				defer f.Close()
				return err
}



//读取文件需要经常进行错误检查，这个帮助方法可以精简下面的错误检查过程。
func check(e error) {
				if e != nil {
								panic(e)
				}
}
