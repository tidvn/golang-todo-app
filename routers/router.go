// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"todo-app/controllers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/todo",
			beego.NSRouter("/", &controllers.TodoController{}, "Post:Create"),
			beego.NSRouter("/", &controllers.TodoController{}, "Get:GetAll"),
			beego.NSRouter("/:id", &controllers.TodoController{}, "Post:Edit"),
			beego.NSRouter("/delete/:id", &controllers.TodoController{}, "Get:Delete"),
		),
	)
	beego.AddNamespace(ns)
}
