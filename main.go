package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juankis/importExcel/utiles"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		utiles.ImportExcel(c)
	})
	router.Run(":9900")
}
