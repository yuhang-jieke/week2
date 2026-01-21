package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/week2/wei/api-away/handler/server"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/create", server.Create)
	r.POST("/upload", server.Upload)
	return r
}
