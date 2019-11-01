package controllers

import (
	"html/template"

	"github.com/rsvijay2009/go-mvc/models"
	"github.com/rsvijay2009/go-mvc/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// UserController define the base controller
type UserController struct {
	beego.Controller
}

// Login to handle user login
func (c *UserController) Login() {
	if c.Ctx.Input.IsPost() {
		form := models.LoginForm{}
		c.ParseForm(&form)
		user := models.FindUser(form.Email)

		if utils.CheckPasswordHash(form.Password, user.Password) {
			c.SetSession("username", user.FirstName)
			c.SetSession("uId", user.Id)
			c.Redirect("/member", 302)
		} else {
			c.Data["ErrorMsg"] = "Invalid credentials. Please try again"
			c.Layout = "layout.tpl"
			c.LayoutSections = make(map[string]string)
			c.LayoutSections["Nav"] = "navbar.tpl"
			c.LayoutSections["Footer"] = "footer.tpl"
			c.Data["Title"] = "Login"
			c.TplName = "index.tpl"
			c.Render()
		}
	}
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Data["Title"] = "Login"
	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Nav"] = "navbar.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.Data["Title"] = "Login"
	c.TplName = "index.tpl"
	c.Render()
}

// Register to handle user registration
func (c *UserController) Register() {
	if c.Ctx.Input.IsPost() {
		form := models.User{}
		c.ParseForm(&form)
		firstName := c.GetString("first_name")
		lastName := c.GetString("last_name")
		email := c.GetString("email")
		password := c.GetString("password")
		confirmPassword := c.GetString("confirm_password")
		phone := c.GetString("phone")

		valid := validation.Validation{}
		valid.Required(firstName, "First Name")
		valid.Required(lastName, "Last Name")
		valid.Email(email, "Email")
		valid.Required(password, "Password")

		if password != confirmPassword {
			valid.SetError("Confirm Password", "not match with password")
		}
		valid.Required(phone, "Phone")

		c.Data["firstName"] = firstName
		c.Data["lastName"] = lastName
		c.Data["email"] = email
		c.Data["password"] = password
		c.Data["confirm_password"] = confirmPassword
		c.Data["phone"] = phone

		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, err.Key+" "+err.Message+"\n")
			}
			c.Data["Errors"] = errormap
			c.Layout = "layout.tpl"
			c.LayoutSections = make(map[string]string)
			c.LayoutSections["Nav"] = "navbar.tpl"
			c.LayoutSections["Footer"] = "footer.tpl"
			c.Data["Title"] = "Registration"
			c.TplName = "register.tpl"
			return
		}
		encPassword, _ := utils.HashPassword(password)
		m := models.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Password:  encPassword,
			Phone:     phone,
		}
		if m.AddUser() != 0 {
			c.Data["successMsg"] = "Account has been created successfully!"
			c.Layout = "layout.tpl"
			c.LayoutSections = make(map[string]string)
			c.LayoutSections["Nav"] = "navbar.tpl"
			c.LayoutSections["Footer"] = "footer.tpl"
			c.Data["Title"] = "Registration successfully done!"
			c.TplName = "success.tpl"
			c.Render()
		} else {
			c.Data["errorMsg"] = "Something went wrong. Please try again!"
			c.Layout = "layout.tpl"
			c.LayoutSections = make(map[string]string)
			c.LayoutSections["Nav"] = "navbar.tpl"
			c.LayoutSections["Footer"] = "footer.tpl"
			c.Data["Title"] = "Registration"
			c.TplName = "error.tpl"
			c.Render()
		}
	}
	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Nav"] = "navbar.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.Data["Title"] = "Registration"

	c.TplName = "register.tpl"
	c.Render()
}

// Logout used to destroy all the session values
func (c *UserController) Logout() {
	c.Data["Title"] = "Logout"
	c.DestroySession()
	c.Redirect(beego.URLFor("UserController.Login"), 302)
}
