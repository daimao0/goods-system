package utils

import (
	"github.com/labstack/gommon/log"
	"os"
)

// WriteFile 写文件
func WriteFile(content, filePath string) {

	// 创建或打开文件（如果路径不存在则创建）
	file, err := os.Create(filePath)
	if err != nil {
		log.Error("Error creating file:", err)
		return
	}
	defer file.Close() // 确保关闭文件，即使出现错误或者函数返回

	// 将字符串写入文件
	_, err = file.WriteString(content)
	if err != nil {
		log.Error("Error writing to file:", err)
		return
	}

}
