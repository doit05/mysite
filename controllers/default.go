package controllers

import (
	"github.com/astaxie/beego"
	"mysite/helper/apicode"
	"mysite/models"
	"mysite/utils"
	//"fmt"
)

type MainController struct {
	BaseController
}

// 初始化˙
func init() {
}

func (c *MainController) Index() {
	c.Data["Message"] = "www.doit05.cn"
	c.TplName = "default.tpl"
	c.Render()
}

func (this *MainController) Upload() {
	filetype, err := this.GetInt("filetype")

	f, h, err := this.GetFile("myfile")
	if err != nil {
		utils.Log.Error("获取文件出错 ： %v", err)
		this.RenderApiJson(apicode.SaveFileFailed, apicode.Msg(apicode.SaveFileFailed), err)
	}
	path := beego.AppConfig.String("file_dir")
	if filetype == 5 {
		path = beego.AppConfig.String("exam_dir")
	}
	path += h.Filename
	f.Close()
	err = this.SaveToFile("myfile", path)
	if err != nil {
		utils.Log.Error("保存文件出错 ： %v", err)
		this.RenderApiJson(apicode.SaveFileFailed, apicode.Msg(apicode.SaveFileFailed), err)
	}
	url := "{'url': 'http://www.doit05.top/" + h.Filename + "'}"
	if filetype == 5 {
		indexModel := models.IndexDataModel{}
		_, err1 := indexModel.Getexam(path)
		if err1 != nil {
			utils.Log.Error("分析文件失败"+path+" ： %v", err)
		}
	}
	this.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), url)
}
