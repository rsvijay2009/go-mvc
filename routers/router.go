package routers

import (
	"github.com/rsvijay2009/go-mvc/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// SignUp, SignIn & Register related routes
	beego.Router("/", &controllers.UserController{}, "get,post:Login")
	beego.Router("/register", &controllers.UserController{}, "get,post:Register")
	beego.Router("/logout/", &controllers.UserController{}, "get:Logout")

	// Member related routes
	beego.Router("/member/", &controllers.HomeController{}, "get:Home")

	// User notes related routes
	beego.Router("/notepad/", &controllers.NotePadController{}, "get,post:CreateNote")
	beego.Router("/notepad/:mode/:id([0-9]+", &controllers.NotePadController{}, "get:GetNotes;post:UpdateNotes")
	beego.Router("/notepad/delete/:id([0-9]+", &controllers.NotePadController{}, "get:DeleteNotes")
}
