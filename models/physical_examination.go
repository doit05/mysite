package models

import (
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"mysite/helper"
	"mysite/utils"
	//"fmt"
	"regexp"
	"strings"
	//"strconv"
)

// '体检数据数据';
type ExamResult struct {
	Id       int64  //  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id '
	Uid      int64  //  `uid` int(11) NOT NULL default 0 COMMENT 'uid '
	Fileld1  string //fileld1 varchar(255) NOT NULL COMMENT ' 机构名称',
	Fileld2  string //fileld2 varchar(255) NOT NULL COMMENT ' 病人档案ID',
	Fileld3  string //fileld3 varchar(255) NOT NULL COMMENT ' 检测时间',
	Fileld4  string //fileld4 varchar(255) NOT NULL COMMENT ' 样本来源',
	Fileld31 string //fileld31 varchar(255) NOT NULL COMMENT '取奶样本',

	Fileld5  string //fileld5 varchar(255) NOT NULL COMMENT ' 婴儿信息',
	Fileld6  string //fileld6 varchar(255) NOT NULL COMMENT ' 母亲信息',
	Fileld7  string //fileld7 varchar(255) NOT NULL COMMENT ' 姓名',
	Fileld32 string //fileld32 varchar(255) NOT NULL COMMENT '姓名2',
	Fileld8  string //fileld8 varchar(255) NOT NULL COMMENT ' 性别',
	Fileld9  string //fileld9 int NOT NULL COMMENT ' 年龄(岁)',
	Fileld10 string //fileld10 string NOT NULL COMMENT ' 身高(cm)',
	Fileld11 string //fileld11 string NOT NULL COMMENT ' 体重(KG)',
	Fileld33 string //fileld33 string NOT NULL COMMENT ' 身高1(cm)',
	Fileld34 string //fileld34 string NOT NULL COMMENT ' 体重2(KG)',
	Fileld12 string //fileld12 string NOT NULL COMMENT '头围(cm)',
	//
	Fileld13 string //fileld13 varchar(255) NOT NULL COMMENT '是否早产',
	Fileld14 string //fileld14 varchar(255) NOT NULL COMMENT 'BMI',
	Fileld15 string //fileld15 varchar(255) NOT NULL COMMENT '分娩方式',
	Fileld16 string //fileld16 varchar(255) NOT NULL COMMENT ' 情绪',
	Fileld35 string //fileld35 varchar(255) NOT NULL COMMENT '不良习惯',
	//
	Fileld17 string //fileld17 varchar(255) NOT NULL COMMENT ' Apgar评分',
	Fileld18 string //fileld18 varchar(255) NOT NULL COMMENT ' 胎次',
	Fileld19 string //fileld19 varchar(255) NOT NULL COMMENT ' 住院号',
	Fileld20 string //fileld20 varchar(255) NOT NULL COMMENT ' 出生日期',
	//
	Fileld21 string //fileld21 varchar(255) NOT NULL COMMENT ' 门诊号',
	Fileld22 string //fileld22 varchar(255) NOT NULL COMMENT ' 床号',
	Fileld36 string //fileld36 varchar(255) NOT NULL COMMENT ' 开奶时间',
	Fileld23 string //fileld23 varchar(255) NOT NULL COMMENT ' 送检科室',
	Fileld24 string //fileld24 varchar(255) NOT NULL COMMENT ' 送检医生',
	//
	Fileld25 string //fileld25 varchar(255) NOT NULL COMMENT '手机号',
	//
	Fat     string //FAT string NOT NULL COMMENT '脂肪',
	Snf     string //SNF string NOT NULL COMMENT '脱脂干物质',
	Density string //Density string NOT NULL COMMENT '密度',
	//
	Protein  string //Protein string NOT NULL COMMENT '蛋白质',
	Lactose  string //Lactose string NOT NULL COMMENT '乳糖',
	Minerals string //Minerals string NOT NULL COMMENT '矿物质',
	Freezing string //Freezing string NOT NULL COMMENT '冰点',
	//
	Energy       string //Energy string NOT NULL COMMENT '能量',
	Watercontent string //WaterContent string NOT NULL COMMENT '含水量',
	Carbohydrate string //Carbohydrate string NOT NULL COMMENT '碳水化合物',
	Grayscale    string //grayscale string NOT NULL COMMENT '灰度',

	Fattips     string //FATtips string NOT NULL COMMENT '脂肪提示',
	Snftips     string //SNFtips string NOT NULL COMMENT '脱脂干物质提示',
	Densitytips string //Densitytips string NOT NULL COMMENT '密度提示',
	//
	Proteintips  string //Proteintips string NOT NULL COMMENT '蛋白质提示',
	Lactosetips  string //Lactosetips string NOT NULL COMMENT '乳糖提示',
	Mineralstips string //Mineralstips string NOT NULL COMMENT '矿物质提示',
	Freezingtips string //Freezingtips string NOT NULL COMMENT '冰点提示',
	//
	Energytips       string //Energytips string NOT NULL COMMENT '能量提示',
	Watercontenttips string //WaterContenttips string NOT NULL COMMENT '含水量提示',
	Carbohydratetips string //Carbohydratetips string NOT NULL COMMENT '碳水化合物提示',
	Grayscaletips    string //grayscaletips string NOT NULL COMMENT '灰度提示',
	Result           string //result text NOT NULL COMMENT '诊断结果',
	//
	Fileld26 string //fileld26 varchar(255) NOT NULL COMMENT '手报告者',
	Fileld27 string //fileld27 varchar(255) NOT NULL COMMENT '手管理员审核者',
	Datetime string //datetime varchar(255) NOT NULL COMMENT '日期',
	Affirms  string //affirms varchar(255) NOT NULL COMMENT '申明',
	//
	Create_time int64 //create_time int(11) DEFAULT NULL COMMENT '创建时间',
	Update_time int64 //update_time int(11) DEFAULT NULL COMMENT '更新时间',
}

// 数据表名称
func (this *ExamResult) TableName() string {
	return "tbl_exammination"
}

func init() {
	// 注册定义的model
	orm.RegisterModelWithPrefix("", new(ExamResult))
}

//根据用户名称获取用户体检报告
func (this *IndexDataModel) GetResultsByName(name string, date int64) (datas []ExamResult, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("tbl_exammination").Filter("Fileld7__icontains", name).All(&datas)
	return
}

//添加用户体检报告
func (this *IndexDataModel) AddResultsData(result ExamResult) (Id int64, err error) {
	o := orm.NewOrm()
	Id, err = o.Insert(&result)
	return
}

//解析用户体检报告
func (this *IndexDataModel) Getexam(filepath string) (id int64, err error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		utils.Log.Error("%v \n", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	j := 0 //有效行数索引
	obj := new(ExamResult)
	if len(lines) > 0 {
		obj.Create_time = helper.GetTimestamp()
		obj.Update_time = obj.Create_time
		for _, line := range lines {
			if len(line) > 0 {
				j++
				if strings.Contains(line, "\u0000") {
					re := regexp.MustCompile("\u0000")
					line = re.ReplaceAllString(line, "")
				}

				line = strings.TrimSpace(line)
				switch {
				case j == 1:
					obj.Fileld1 = strings.TrimSpace(line)
				case strings.Contains(line, "病人档案ID"):
					obj.Fileld2, obj.Fileld3 = this.extract2(line)
				case strings.Contains(line, "样本来源"):
					obj.Fileld4, obj.Fileld31 = this.extract3(line)
				case strings.Contains(line, "婴儿信息"):
					obj.Fileld5, obj.Fileld6 = this.extract4(line)
				case strings.Contains(line, "年龄"):
					obj.Fileld7, obj.Fileld8, obj.Fileld9, obj.Fileld32 = this.extract5(line)
				case strings.Contains(line, "身高"):
					obj.Fileld10, obj.Fileld11, obj.Fileld33, obj.Fileld34 = this.extract6(line)
				case strings.Contains(line, "头围"):
					obj.Fileld12, obj.Fileld13, obj.Fileld14, obj.Fileld15 = this.extract7(line)
				case strings.Contains(line, "情绪"):
					obj.Fileld16, obj.Fileld35 = this.extract8(line)
				case strings.Contains(line, "Apgar"):
					obj.Fileld17, obj.Fileld18, obj.Fileld19 = this.extract9(line)
				case strings.Contains(line, "出生日期"):
					obj.Fileld20, obj.Fileld21, obj.Fileld22 = this.extract10(line)
				case strings.Contains(line, "开奶"):
					obj.Fileld36 = this.extract11(line)
				case strings.Contains(line, "送检科室"):
					obj.Fileld23, obj.Fileld24, obj.Fileld25 = this.extract12(line)
				case strings.Contains(line, "FAT"):
					obj.Fat, obj.Fattips = this.extract15(line)
				case strings.Contains(line, "SNF"):
					obj.Snf, obj.Snftips = this.extract16(line)
				case strings.Contains(line, "Density"):
					obj.Density, obj.Densitytips = this.extract17(line)
				case strings.Contains(line, "Protein"):
					obj.Protein, obj.Proteintips = this.extract25(line)
				case strings.Contains(line, "Lactose"):
					obj.Lactose, obj.Lactosetips = this.extract18(line)
				case strings.Contains(line, "Minerals"):
					obj.Minerals, obj.Mineralstips = this.extract19(line)
				case strings.Contains(line, "Freezing"):
					obj.Freezing, obj.Freezingtips = this.extract20(line)
				case strings.Contains(line, "Energy"):
					obj.Energy, obj.Energytips = this.extract21(line)
				case strings.Contains(line, "water"):
					obj.Watercontent, obj.Watercontenttips = this.extract22(line)
				case strings.Contains(line, "Carbohydrate"):
					obj.Carbohydrate, obj.Carbohydratetips = this.extract23(line)
				case strings.Contains(line, "grayscale"):
					obj.Grayscale, obj.Grayscaletips = this.extract24(line)
				case strings.Contains(line, "报告者"):
					obj.Fileld26, obj.Fileld27, obj.Datetime = this.extract40(line)
				case strings.Contains(line, "当前结果"):
				case strings.Contains(line, "项目"):
				case strings.Contains(line, "诊断结果"):
				case strings.Contains(line, "此报告仅对"):
					obj.Affirms = strings.TrimSpace(line)
				default:
					obj.Result += line + "\n"
				}
			}

		}
	}
	re := regexp.MustCompile("\n\n")
	obj.Result = re.ReplaceAllString(obj.Result, "")
	if obj.Create_time > 0 {
		id, err = this.AddResultsData(*obj)
	}

	return

}

func (this *IndexDataModel) extract2(line string) (pid, datetime string) {
	tmp := strings.Split(line, " ")
	if len(tmp) > 0 {
		//fmt.Printf("%v  %d \n", tmp, len(tmp))
		i := 0
		date := make([]string, 2)
		for _, str := range tmp {
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					pid = str
				case 3:
					date_arr := strings.Split(str, ":")
					date = append(date, date_arr[1])

				case 4:

					date = append(date, str)
					datetime = strings.Join(date, " ")
				}

			}

		}
	}
	return

}

func (this *IndexDataModel) extract3(line string) (str1, str2 string) {
	tmp := strings.Split(line, " ")
	if len(tmp) > 0 {
		i := 0
		for _, str := range tmp {
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					for _, word := range str {
						if string(word) != "取" {
							str1 += string(word)
						} else {
							break
						}

					}
				case 3:
					str2 = str
				}

			}

		}
	}
	return

}

func (this *IndexDataModel) extract4(line string) (str1, str2 string) {
	tmp := strings.Split(line, ":")
	if len(tmp) > 0 {
		i := 0
		for _, str := range tmp {
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					str1 = strings.Split(str, "母")[0]
				case 3:
					str2 = str
				}

			}

		}
	}
	return

}

func (this *IndexDataModel) extract5(line string) (str1, str2, str4, str3 string) {
	tmp := strings.Split(line, ":")
	if len(tmp) > 0 {
		i := 0
		for _, str := range tmp {
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					for _, word := range str {
						if string(word) != "性" {
							str1 += string(word)
						} else {
							break
						}

					}
				case 3:
					match, _ := regexp.MatchString("男", str)
					if match {
						str2 = "男"
					} else {
						str2 = "女"
					}
				case 4:
					for _, word := range str {
						if string(word) != "年" {
							str3 += string(word)
						} else {
							break
						}

					}
				case 5:
					str4 = str
				}

			}

		}
	}
	return

}

func (this *IndexDataModel) extract6(line string) (str1, str2, str3, str4 string) {
	tmp := strings.Split(line, ":")
	if len(tmp) > 0 {
		i := 0
		for _, str := range tmp {
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					str1 = helper.GetNumFromStr(str)
				case 3:
					str2 = helper.GetNumFromStr(str)
				case 4:
					str3 = helper.GetNumFromStr(str)
				case 5:
					str4 = helper.GetNumFromStr(str)
				}
			}

		}
	}
	return
}

func (this *IndexDataModel) extract7(line string) (str1, str2, str3, str4 string) {
	tmp := strings.Split(line, ":")
	if len(tmp) > 0 {
		i := 0
		for _, str := range tmp {
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					str1 = helper.GetNumFromStr(str)
				case 3:
					match, _ := regexp.MatchString("否", str)
					if match {
						str2 = "否"
					} else {
						str2 = "是"
					}
				case 4:
					str3 = helper.GetNumFromStr(str)
				case 5:
					str4 = str

				}

			}

		}
	}
	return
}

func (this *IndexDataModel) extract8(line string) (str1, str2 string) {
	tmp := strings.Split(line, ":")
	if len(tmp) > 0 {
		i := 0
		for _, str := range tmp {
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					str1 = helper.GetNumFromStr(str)
				case 3:
					str2 = helper.GetNumFromStr(str)
				}
			}

		}
	}
	return
}

func (this *IndexDataModel) extract9(line string) (str1, str2, str3 string) {
	tmp := strings.Split(line, ":")
	if len(tmp) > 0 {
		i := 0
		for _, str := range tmp {
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					str1 = helper.GetNumFromStr(str)
				case 3:
					re := regexp.MustCompile("(.+)住院")
					str_arr := re.FindAllStringSubmatch(str, -1)
					if len(str_arr) > 0 {
						str2 = str_arr[0][1]
					}
				case 4:
					str3 = helper.GetNumFromStr(str)
				}
			}

		}
	}
	return
}

func (this *IndexDataModel) extract10(line string) (str1, str2, str3 string) {
	tmp := strings.Split(line, " ")
	if len(tmp) > 0 {
		i := 0
		for _, str := range tmp {
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					str1 += str
				case 3:
					str1 += " " + str
				case 4:
				case 5:
					str2 = helper.GetNumFromStr(str)
				case 6:
				case 7:
					str3 = helper.GetNumFromStr(str)
				}
			}

		}
	}
	return
}

func (this *IndexDataModel) extract11(line string) (str1 string) {
	tmp := strings.Split(line, " ")
	if len(tmp) > 0 {
		i := 0
		for _, str := range tmp {
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					str1 += str
				case 3:
					str1 += " " + str
				}
			}

		}
	}
	return
}

func (this *IndexDataModel) extract12(line string) (str1, str2, str3 string) {
	tmp := strings.Split(line, ":")
	if len(tmp) > 0 {
		i := 0
		for _, str := range tmp {
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					re := regexp.MustCompile("(.+)送检")
					str_arr := re.FindAllStringSubmatch(str, -1)
					if len(str_arr) > 0 {
						str1 = str_arr[0][1]
					}
				case 3:
					re := regexp.MustCompile("(.+)手机")
					str_arr := re.FindAllStringSubmatch(str, -1)
					if len(str_arr) > 0 {
						str2 = str_arr[0][1]
					}
				case 4:
					str3 += helper.GetNumFromStr(str)
				}
			}

		}
	}
	return
}

/**
报错
ret, err := strconv.Parsestring(str_arr, 10)
0 err :  strconv.Parsestring: parsing "89.910\x00": invalid syntax

*/
//func (this *IndexDataModel)extract15(line string) (ret string) {
//				str_arr := helper.GetstringFromStr(line)
//				str_arr = strings.TrimSpace(str_arr)
//				fmt.Println(len(str_arr))
//				if len(str_arr) > 0 {
//								ret, err := strconv.Parsestring(str_arr, 10)
//								fmt.Println(ret,"err : ", err)
//				}
//
//				return
//}

//FAT
func (this *IndexDataModel) extract15(line string) (ret, tips string) {
	str_arr := helper.GetFloatFromStr(line, ".")
	str_arr = strings.TrimSpace(str_arr)
	tips = helper.GetDegressStr(line)
	ret = str_arr
	return
}

//SNF
func (this *IndexDataModel) extract16(line string) (ret, tips string) {
	re := regexp.MustCompile("[0-9]+.[0-9]+")
	str_arr := re.FindString(line)
	tips = helper.GetDegressStr(line)
	ret = str_arr
	return
}

//Density
func (this *IndexDataModel) extract17(line string) (ret, tips string) {
	re := regexp.MustCompile("[0-9]+.[0-9]{3}")
	str_arr := re.FindString(line)
	tips = helper.GetDegressStr(line)
	ret = str_arr
	return
}

//Protein
func (this *IndexDataModel) extract25(line string) (ret, tips string) {
	re := regexp.MustCompile("[0-9]+.[0-9]{2}")
	str_arr := re.FindString(line)
	tips = helper.GetDegressStr(line)
	ret = str_arr
	return
}

//Lactose
func (this *IndexDataModel) extract18(line string) (ret, tips string) {
	re := regexp.MustCompile("[0-9]+.[0-9]{2}")
	str_arr := re.FindString(line)
	tips = helper.GetDegressStr(line)
	ret = str_arr
	return
}

//Minerals
func (this *IndexDataModel) extract19(line string) (ret, tips string) {
	re := regexp.MustCompile("[0-9]+.[0-9]{2}")
	str_arr := re.FindString(line)
	tips = helper.GetDegressStr(line)
	ret = str_arr
	return
}

//Freezing
func (this *IndexDataModel) extract20(line string) (ret, tips string) {
	re := regexp.MustCompile("[0-9]+.[0-9]{3}")
	str_arr := re.FindString(line)
	tips = helper.GetDegressStr(line)
	ret = str_arr
	return
}

//Energy
func (this *IndexDataModel) extract21(line string) (ret, tips string) {
	re := regexp.MustCompile("[0-9]+.[0-9]{3}")
	str_arr := re.FindString(line)
	tips = helper.GetDegressStr(line)
	ret = str_arr
	return
}

//water content
func (this *IndexDataModel) extract22(line string) (ret, tips string) {
	re := regexp.MustCompile("[0-9]+.[0-9]{3}")
	str_arr := re.FindString(line)
	tips = helper.GetDegressStr(line)
	ret = str_arr
	return
}

//Carbohydrate
func (this *IndexDataModel) extract23(line string) (ret, tips string) {
	re := regexp.MustCompile("[0-9]+.[0-9]{2}")
	str_arr := re.FindString(line)
	tips = helper.GetDegressStr(line)
	ret = str_arr
	return
}

//grayscale
func (this *IndexDataModel) extract24(line string) (str, tips string) {
	re := regexp.MustCompile("度(.+)")
	str_arr := re.FindStringSubmatch(line)
	if len(str_arr) > 1 {
		str = str_arr[1]
	}
	tips = helper.GetDegressStr(line)
	return
}

//报告者: 管理员 审核者:管理员           日期:2016/7/16 14:38:16
func (this *IndexDataModel) extract40(line string) (str1, str2, str3 string) {
	tmp := strings.Split(line, " ")
	if len(tmp) > 0 {
		i := 0
		for _, str := range tmp {
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				i++
				switch i {
				case 1:
				case 2:
					arr2 := strings.Split(str, ":")
					if len(arr2) > 1 {
						str2 = arr2[1]
					}
					re := regexp.MustCompile("(.+)审核")
					str_arr := re.FindStringSubmatch(arr2[0])
					if len(str_arr) > 1 {
						str1 = str_arr[1]
					}

				case 3:
					str3 += strings.Split(str, ":")[1]

				case 4:
					str3 += " " + str
				}
			}

		}
	}
	return
}

//列对应的中文
func (this *IndexDataModel) MapNames() (map_names map[string]string) {
	map_names["fileld1"] = "机构名称"
	map_names["fileld2"] = "病人档案ID"
	map_names["fileld3"] = "检测时间"
	map_names["fileld4"] = "样本来源"
	map_names["fileld5"] = "婴儿信息"
	map_names["fileld6"] = "母亲信息"
	map_names["fileld7"] = "姓名"
	map_names["fileld8"] = "性别"
	map_names["fileld9"] = "年龄(岁)"
	map_names["fileld10"] = "身高(cm)"
	map_names["fileld11"] = "体重(KG)"
	map_names["fileld12"] = "头围(cm)"
	map_names["fileld13"] = "是否早产"
	map_names["fileld14"] = "BMI"
	map_names["fileld15"] = "分娩方式"

	map_names["fileld16"] = "情绪"
	map_names["fileld17"] = "Apgar评分"
	map_names["fileld18"] = "胎次"
	map_names["fileld19"] = "住院号"
	map_names["fileld20"] = "出生日期"
	map_names["fileld21"] = "门诊号"

	map_names["fileld22"] = "床号:"
	map_names["fileld23"] = "送检科室"
	map_names["fileld24"] = "送检医生"
	map_names["fileld25"] = "手机号"
	map_names["FAT"] = "FAT"
	map_names["SNF"] = "SNF"
	map_names["Density"] = "Density"

	map_names["Protein"] = "Protein"
	map_names["Lactose"] = "Lactose"
	map_names["Minerals"] = "Minerals"
	map_names["Freezing"] = "Freezing"
	map_names["Energy"] = "Energy"
	map_names["waterContent"] = "waterContent"

	map_names["Carbohydrate"] = "Carbohydrate:"
	map_names["grayscale"] = "grayscale"
	map_names["result"] = "诊断结果"

	map_names["fileld38"] = "报告者"
	map_names["fileld39"] = "管理员审核者"
	map_names["datetime"] = "日期"
	return
}
