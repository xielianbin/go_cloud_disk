package controller

import (
	"clouddisk/lib"
	"clouddisk/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 全部文件页面
func Files(c *gin.Context) {
	fmt.Println("访问files,调用控制层的Files")
	openId, _ := c.Get("openId")
	fId := c.DefaultQuery("fId", "0")
	//获取用户信息
	user := model.GetUserInfo(openId)

	//获取当前目录所有文件
	files := model.GetUserFile(fId, user.FileStoreId)
	//获取当前目录所有文件夹
	fileFolder := model.GetFileFolder(fId, user.FileStoreId)

	//获取父级的文件夹信息
	parentFolder := model.GetParentFolder(fId)

	//获取当前目录所有父级
	currentAllParent := model.GetCurrentAllParent(parentFolder, make([]model.FileFolder, 0))

	//获取当前目录信息
	currentFolder := model.GetCurrentFolder(fId)

	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)
	fmt.Println("文件信息获取成功，返回网页信息")
	//	使用Gin框架渲染HTML模板的典型示例‌1。c.HTML()方法接收三个参数：
	//	http.StatusOK表示HTTP 200成功状态码‌2
	//	"files.html"指定要渲染的模板文件
	//	gin.H{}是Gin提供的快捷方式，用于创建map类型模板数据‌1
	//该方法会将这些数据注入到files.html模板中，最终生成完整的HTML响应返回给客户端‌1。这种模式在Gin框架中常用于构建动态网页应用，通过模板引擎实现前后端数据交互‌13。

	c.HTML(http.StatusOK, "files.html", gin.H{
		"currAll":          "active",
		"user":             user,
		"fId":              currentFolder.Id,
		"fName":            currentFolder.FileFolderName,
		"files":            files,
		"fileFolder":       fileFolder,
		"parentFolder":     parentFolder,
		"currentAllParent": currentAllParent,
		"fileDetailUse":    fileDetailUse,
	})
}

// 处理新建文件夹
func AddFolder(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := model.GetUserInfo(openId)

	folderName := c.PostForm("fileFolderName")
	parentId := c.DefaultPostForm("parentFolderId", "0")

	//新建文件夹数据
	model.CreateFolder(folderName, parentId, user.FileStoreId)

	//获取父文件夹信息
	parent := model.GetParentFolder(parentId)

	c.Redirect(http.StatusMovedPermanently, "/cloud/files?fId="+parentId+"&fName="+parent.FileFolderName)
}

func DownloadFile(c *gin.Context) {
	fId := c.Query("fId")

	file := model.GetFileInfo(fId)
	if file.FileHash == "" {
		return
	}

	//从oss获取文件
	fileData := lib.DownloadLocal(file.FileHash, file.Postfix)
	//下载次数+1
	model.DownloadNumAdd(fId)

	c.Header("Content-disposition", "attachment;filename=\""+file.FileName+file.Postfix+"\"")
	c.Data(http.StatusOK, "application/octect-stream", fileData)
}

// 删除文件
func DeleteFile(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := model.GetUserInfo(openId)

	fId := c.DefaultQuery("fId", "")
	folderId := c.Query("folder")
	if fId == "" {
		return
	}

	//删除数据库文件数据
	model.DeleteUserFile(fId, folderId, user.FileStoreId)

	c.Redirect(http.StatusMovedPermanently, "/cloud/files?fid="+folderId)
}

// 删除文件夹
func DeleteFileFolder(c *gin.Context) {
	fId := c.DefaultQuery("fId", "")
	if fId == "" {
		return
	}
	//获取要删除的文件夹信息 取到父级目录重定向
	folderInfo := model.GetCurrentFolder(fId)

	//删除文件夹并删除文件夹中的文件信息
	model.DeleteFileFolder(fId)

	c.Redirect(http.StatusMovedPermanently, "/cloud/files?fId="+strconv.Itoa(folderInfo.ParentFolderId))
}

// 修改文件夹名
func UpdateFileFolder(c *gin.Context) {
	fileFolderName := c.PostForm("fileFolderName")
	fileFolderId := c.PostForm("fileFolderId")

	fileFolder := model.GetCurrentFolder(fileFolderId)

	model.UpdateFolderName(fileFolderId, fileFolderName)

	c.Redirect(http.StatusMovedPermanently, "/cloud/files?fId="+strconv.Itoa(fileFolder.ParentFolderId))
}
