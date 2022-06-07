package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 单个文件上传
func Upload(c *gin.Context) {
	// 文件上传
	//file, err := c.FormFile("file")
	//if err != nil {
	//	c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	//	return
	//}
	//
	//basePath := "./data/"
	//filename := basePath + filepath.Base(file.Filename)
	//if err := c.SaveUploadedFile(file, filename); err != nil {
	//	c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	//	return
	//}
	//c.String(http.StatusOK,fmt.Sprintf("文件 %s 上传成功 ", file.Filename))

	// 单个文件上传
	// file与文件上传的标志是一致的
	file, err := c.FormFile("file")
	log.Println(file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(file.Filename)
	dst := fmt.Sprintf("/goPj/data/%s", file.Filename)
	// 上传文件到指定的目录
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}

// 多个文件上传
func Uploads(c *gin.Context) {
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for index, file := range files {
		log.Println(file.Filename)
		dst := fmt.Sprintf("/goPj/data//%s_%d", file.Filename, index)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d files uploaded!", len(files)),
	})
}
