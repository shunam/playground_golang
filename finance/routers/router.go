package routers

import (
	"finance/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/users", &controllers.UsersController{}, "get:GetAll")
}
