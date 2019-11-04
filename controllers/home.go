package controllers

import (
	"github.com/astaxie/beego"
	"github.com/rsvijay2009/go-mvc/models"
)

// HomeController define the base controller
type HomeController struct {
	beego.Controller
}

// Home used to get the login view
func (c *HomeController) Home() {
	loggedUserID := c.GetSession("uId")
	loggedUserName := c.GetSession("username")

	if loggedUserID != nil {
		user := models.FindUserByID(loggedUserID.(int))
		c.Layout = "layout.tpl"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Nav"] = "navbar.tpl"
		c.LayoutSections["Footer"] = "footer.tpl"
		c.Data["User"] = user
		c.Data["UserName"] = loggedUserName
		c.Data["UserID"] = user.Id
		c.Data["Title"] = "Welcome"
		c.TplName = "home.tpl"
		c.Render()
	}

	c.Redirect(beego.URLFor("UserController.Login"), 302)
}
