package lib

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

const TokenKey = "Casbin-Token"

// token="user name"
func CheckLogin() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		userName := ctx.Request.Header.Get("token")
		if userName == "" {
			ctx.AbortWithStatusJSON(401, "missed token")
		} else {
			ctx.Set(TokenKey, userName)
			ctx.Next()
		}
	}
}

func RBAC() gin.HandlerFunc  {
	e:= casbin.NewEnforcer("resources/model.conf","resources/p.csv")
	return func(ctx *gin.Context) {
		role := ctx.Request.Header.Get("token")
		fmt.Println("role:", role)
		ok:= e.Enforce(role, ctx.Request.RequestURI, ctx.Request.Method)
		if !ok {
			fmt.Println("permission denied")
			ctx.AbortWithStatusJSON(403, "permission denied")
		} else {
			ctx.Next()
		}
	}
	
}
