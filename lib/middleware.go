package lib

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"log"
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
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	//a, _ := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source.

	// 会自动在数据库中创建表（casbin_rule），在表中添加策略即可
	/**  表字段
	+-------+---------------------+------+-----+---------+----------------+
	| Field | Type                | Null | Key | Default | Extra          |
	+-------+---------------------+------+-----+---------+----------------+
	| id    | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
	| ptype | varchar(100)        | YES  | MUL | NULL    |                |
	| v0    | varchar(100)        | YES  |     | NULL    |                |
	| v1    | varchar(100)        | YES  |     | NULL    |                |
	| v2    | varchar(100)        | YES  |     | NULL    |                |
	| v3    | varchar(100)        | YES  |     | NULL    |                |
	| v4    | varchar(100)        | YES  |     | NULL    |                |
	| v5    | varchar(100)        | YES  |     | NULL    |                |
	+-------+---------------------+------+-----+---------+----------------+
	*/

	/** 添加策略
	insert into casbin_rule (ptype,v0,v1,v2) values('p', 'member', '/depts', 'GET'),('p', 'admin', '/depts', 'POST');
	insert into casbin_rule (ptype,v0,v1) values('g','admin', 'member'),('g','sjt','admin');
	+----+-------+--------+--------+------+------+------+------+
	| id | ptype | v0     | v1     | v2   | v3   | v4   | v5   |
	+----+-------+--------+--------+------+------+------+------+
	|  1 | g     | admin  | member | NULL | NULL | NULL | NULL |
	|  2 | g     | sjt    | admin  | NULL | NULL | NULL | NULL |
	|  4 | p     | admin  | /depts | POST | NULL | NULL | NULL |
	|  3 | p     | member | /depts | GET  | NULL | NULL | NULL |
	+----+-------+--------+--------+------+------+------+------+
	 */

	a, err := gormadapter.NewAdapterByDB(DB) // Your driver and data source.

	if err !=nil {
		log.Fatal(err)
	}

	e, err := casbin.NewEnforcer("resources/model.conf", a)
	if err !=nil {
		log.Fatal(err)
	}

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	e.LoadPolicy()
	// Save the policy back to DB.
	//e.SavePolicy()
	return func(ctx *gin.Context) {
		role := ctx.Request.Header.Get("token")
		log.Println("role:", role)
		access, err:= e.Enforce(role, ctx.Request.RequestURI, ctx.Request.Method)
		if err !=nil || !access {
			log.Println("permission denied")
			ctx.AbortWithStatusJSON(403, "permission denied")
		} else {
			ctx.Next()
		}

	}
	
}
