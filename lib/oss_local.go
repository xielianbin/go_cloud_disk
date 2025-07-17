package lib

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

// 上传文件至本地
func Uploadlocal(filename, fileHash string) {
	// 获取文件后缀
	fmt.Println("调用本地上传")
	fileSuffix := path.Ext(filename)
	conf := LoadServerConfig()

	// 创建目标目录（如果不存在）
	destDir := filepath.Join(conf.UploadLocation, "files")
	if err := os.MkdirAll(destDir, 0755); err != nil {
		fmt.Println("创建目录Error:", err)
		return
	}

	// 构建源文件和目标文件路径
	srcPath := filepath.Join(conf.UploadLocation, filename)
	destPath := filepath.Join(destDir, fileHash+fileSuffix)

	// 复制文件到目标位置
	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("打开源文件Error:", err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		fmt.Println("创建目标文件Error:", err)
		return
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		fmt.Println("文件复制Error:", err)
		return
	}

	fmt.Println("文件上传成功:", destPath)
}

// 从本地下载文件
func DownloadLocal(fileHash, fileType string) []byte {
	conf := LoadServerConfig()
	filePath := filepath.Join(conf.UploadLocation, fileHash+fileType)

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return data
}
