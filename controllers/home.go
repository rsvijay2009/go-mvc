package controllers

import (
	"github.com/astaxie/beego"
)

// HomeController define the base controller
type HomeController struct {
	beego.Controller
}

// Home used to get the login view
func (c *HomeController) Home() {
	loggedUser := c.GetSession("username")

	c.Data["Title"] = "Welcome to Beego MVC project"
	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Nav"] = "navbar.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"

	if loggedUser != nil {
		c.Data["UserName"] = loggedUser
		c.TplName = "home.tpl"
		c.Render()
	}

	c.Redirect(beego.URLFor("UserController.Login"), 302)
}
