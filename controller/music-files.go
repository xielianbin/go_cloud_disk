package controller

import (
	"clouddisk/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MusicFiles(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := model.GetUserInfo(openId)

	fmt.Println("访问 music-files，调用 MusicFiles")
	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)
	//获取音频类型文件
	musicFiles := model.GetTypeFile(4, user.FileStoreId)

	c.HTML(http.StatusOK, "music-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"musicFiles":    musicFiles,
		"musicCount":    len(musicFiles),
		"currMusic":     "active",
		"currClass":     "active",
	})
}
