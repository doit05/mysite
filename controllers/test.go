package controllers

import (
//	"fmt"
	"github.com/weilaihui/fdfs_client"
	"mysite/helper"
	"mysite/helper/apicode"
	"mysite/utils"
	"os"
)

type TestController struct {
	BaseController
}

func (c *TestController) Get() {
	path := "/Users/doit/Downloads/test.jpg"
	res, fileMd5, err := c.testFastdfs(path)
	if err == nil {
		url := c.toUrl(res, fileMd5)
		c.RenderApiJson(apicode.Success, apicode.Msg(apicode.Success), url)
	} else {
		c.RenderApiJson(apicode.SystemError, apicode.Msg(apicode.SystemError), err)
	}
}

func (c *TestController) testFastdfs(path string) (uploadResponse *fdfs_client.UploadFileResponse, fileMd5 string, err error) {
	fdfsClient, err := fdfs_client.NewFdfsClient("/Users/doit/programeyard/Go/src/github.com/weilaihui/fdfs_client/client.conf")
	if err != nil {
		utils.Log.Error("New FdfsClient error %v", err)
		c.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), err)
	}

	file, err := os.Open(path) // For read access.
	if err != nil {
		utils.Log.Error("open err %v", err)
	}

	var fileSize int64 = 0
	if fileInfo, err := file.Stat(); err == nil {
		fileSize = fileInfo.Size()
	}
	fileBuffer := make([]byte, fileSize)
	_, err = file.Read(fileBuffer)
	if err != nil {
		utils.Log.Error("open err %v", err)
	}
	fileMd5 = helper.Md5Bytes(fileBuffer)

	uploadResponse, err = fdfsClient.UploadByBuffer(fileBuffer, "jpg")

	if err != nil {
		utils.Log.Error("open err %v", err)
	}

	//fdfsClient.DeleteFile(uploadResponse.RemoteFileId)
	return uploadResponse, fileMd5, err
}

func (c *TestController) toUrl(res *fdfs_client.UploadFileResponse, fileMd5 string) (ret string) {
	fileName := res.RemoteFileId
	basic_host := "http://esx.bigo.sg"
	ret = basic_host + "/live/" + fileName
	return ret

}

// func test() {
// 	fdfsClient, err := fdfs_client.NewFdfsClient("/Users/doit/programeyard/Go/src/github.com/weilaihui/fdfs_client/client.conf")
// 	if err != nil {
// 		fmt.Printf("New FdfsClient error %s", err.Error())
// 		return
// 	}

// 	uploadResponse, err := fdfsClient.UploadByFilename("main.go")
// 	if err != nil {
// 		fmt.Errorf("UploadByfilename error %s", err.Error())
// 	}
// 	fmt.Println(uploadResponse.GroupName)
// 	fmt.Println(uploadResponse.RemoteFileId)
// 	// fdfsClient.DeleteFile(uploadResponse.RemoteFileId)
// }
