package controller

import (
	"clouddisk/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	fmt.Println("访问index，调用Index")
	openId, _ := c.Get("openId")
	//获取用户信息
	user := model.GetUserInfo(openId)
	//获取用户仓库信息
	userFileStore := model.GetUserFileStore(user.Id)
	fmt.Println("获取用户仓库信息，仓库信息为：", userFileStore)
	//获取用户文件数量
	fileCount := model.GetUserFileCount(user.FileStoreId)
	fmt.Println("获取用户文件数量，文件数量为：", userFileStore)
	//获取用户文件夹数量
	fileFolderCount := model.GetUserFileFolderCount(user.FileStoreId)
	fmt.Println("获取用户文件夹数量，文件夹数量为：", fileFolderCount)
	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":            user,
		"currIndex":       "active",
		"userFileStore":   userFileStore,
		"fileCount":       fileCount,
		"fileFolderCount": fileFolderCount,
		"fileDetailUse":   fileDetailUse,
	})
}
