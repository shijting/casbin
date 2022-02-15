package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/casbin/lib"
)

func main()  {
	r := gin.New()
	r.Use(lib.CheckLogin(), lib.RBAC())
	r.GET("/depts", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "部门列表"})
	})
	r.POST("/depts", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "创建部门"})
	})
	r.Run(":8081")
}
