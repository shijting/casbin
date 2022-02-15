package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/casbin/lib"
)

func main()  {
	// /depts
	//sub:= "lisi" // 想要访问资源的用户。
	//obj:= "/depts" // 将被访问的资源。
	//act:= "POST" // 用户对资源执行的操作。
	//e:= casbin.NewEnforcer("resources/model.conf","resources/p.csv")
	//
	//ok:= e.Enforce(sub, obj, act)
	//if ok {
	//	log.Println("通过")
	//} else {
	//	log.Println("禁止访问")
	//}
	r := gin.New()
	r.Use(lib.CheckLogin(), lib.RBAC())
	r.GET("/depts", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "部门列表"})
	})
	r.POST("/depts", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"result": "创建部门"})
	})
	r.Run(":8080")
}
