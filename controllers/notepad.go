package controllers

import (
	"strconv"

	"github.com/rsvijay2009/go-mvc/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// NotePadController define the Base controller
type NotePadController struct {
	beego.Controller
}

// CreateNote used to create a new note
func (c *NotePadController) CreateNote() {
	m := models.Notes{}
	UserID := c.GetSession("uId")
	UserName := c.GetSession("username")

	if c.Ctx.Input.IsPost() {
		form := models.Notes{}
		c.ParseForm(&form)
		comments := c.GetString("comments")
		valid := validation.Validation{}
		valid.Required(comments, "Notes")

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
			c.Data["Title"] = "Create Note"
			c.Data["UserID"] = UserID
			c.Data["UserName"] = UserName
			c.TplName = "notepad.tpl"
			return
		}
		m = models.Notes{
			Comments: comments,
			UserId:   UserID.(int),
		}
		if m.AddNote() != 0 {
			c.Layout = "layout.tpl"
			c.LayoutSections = make(map[string]string)
			c.LayoutSections["Nav"] = "navbar.tpl"
			c.LayoutSections["Footer"] = "footer.tpl"
			c.TplName = "notepad.tpl"
			c.Data["UserID"] = UserID
			c.Data["UserName"] = UserName
			c.Redirect(beego.URLFor("NotePadController.CreateNote"), 302)
		} else {
			c.Data["errorMsg"] = "Something went wrong. Please try again!"
			c.Layout = "layout.tpl"
			c.LayoutSections = make(map[string]string)
			c.LayoutSections["Nav"] = "navbar.tpl"
			c.LayoutSections["Footer"] = "footer.tpl"
			c.Data["UserID"] = UserID
			c.Data["UserName"] = UserName
			c.Data["Title"] = "Create Note"
			c.TplName = "notepad.tpl"
			c.Render()
		}
	}
	if UserID != nil {
		m := models.Notes{
			UserId: UserID.(int),
		}
		notes := models.GetNotesByUserID(m.UserId)
		c.Data["Notes"] = notes
		c.Layout = "layout.tpl"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Nav"] = "navbar.tpl"
		c.LayoutSections["Footer"] = "footer.tpl"
		c.Data["Title"] = "Your Notes"
		c.Data["UserID"] = UserID
		c.Data["UserName"] = UserName

		c.TplName = "notepad.tpl"
		c.Render()
	}
	c.Redirect(beego.URLFor("UserController.Login"), 302)
}

// GetNotes used to get a note based on note id
func (c *NotePadController) GetNotes() {
	id := c.Ctx.Input.Param(":id")
	mode := c.Ctx.Input.Param(":mode")
	intID, _ := strconv.Atoi(id)

	UserID := c.GetSession("uId")
	UserName := c.GetSession("username")

	if UserID != nil {
		m := models.Notes{
			Id: intID,
		}
		note := models.GetNotesByID(m.Id)
		c.Data["Note"] = note
		c.Data["UserID"] = UserID
		c.Data["UserName"] = UserName
		c.Layout = "layout.tpl"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Nav"] = "navbar.tpl"
		c.LayoutSections["Footer"] = "footer.tpl"

		if mode == "edit" {
			c.TplName = "edit_notepad.tpl"
			c.Data["Title"] = "Edit note details"
		} else {
			c.TplName = "view_notepad.tpl"
			c.Data["Title"] = "View note details"
		}
		c.Render()
	}
}

// UpdateNotes used to update a note
func (c *NotePadController) UpdateNotes() {
	id := c.Ctx.Input.Param(":id")
	noteID, _ := strconv.Atoi(id)
	UserName := c.GetSession("username")

	n := models.Notes{}
	c.ParseForm(&n)
	comments := c.GetString("comments")

	valid := validation.Validation{}
	valid.Required(comments, "Notes")

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
		c.Data["Title"] = "Update note"
		c.Data["UserName"] = UserName
		c.TplName = "edit_notepad.tpl"
		return
	}
	n = models.Notes{
		Comments: comments,
		Id:       noteID,
	}
	if n.UpdateNote() != 0 {
		c.Layout = "layout.tpl"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Nav"] = "navbar.tpl"
		c.LayoutSections["Footer"] = "footer.tpl"
		c.Data["Title"] = "Update note"
		c.Data["UserName"] = UserName
		c.TplName = "notepad.tpl"
		c.Redirect(beego.URLFor("NotePadController.CreateNote"), 302)
	} else {
		c.Data["errorMsg"] = "Something went wrong. Please try again!"
		c.Layout = "layout.tpl"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Nav"] = "navbar.tpl"
		c.LayoutSections["Footer"] = "footer.tpl"
		c.Data["Title"] = "Update note"
		c.Data["UserName"] = UserName
		c.TplName = "notepad.tpl"
		c.Render()
	}
}

// DeleteNotes used to delete a note
func (c *NotePadController) DeleteNotes() {
	UserName := c.GetSession("username")
	noteID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	n := models.Notes{
		Id: noteID,
	}
	if n.DeleteNote() != 0 {
		c.Layout = "layout.tpl"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Nav"] = "navbar.tpl"
		c.LayoutSections["Footer"] = "footer.tpl"
		c.Data["Title"] = "Your notes"
		c.Data["UserName"] = UserName
		c.TplName = "notepad.tpl"
		c.Redirect(beego.URLFor("NotePadController.CreateNote"), 302)
	} else {
		c.Data["errorMsg"] = "Something went wrong. Please try again!"
		c.Layout = "layout.tpl"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Nav"] = "navbar.tpl"
		c.LayoutSections["Footer"] = "footer.tpl"
		c.Data["Title"] = "Your notes"
		c.Data["UserName"] = UserName
		c.TplName = "notepad.tpl"
		c.Render()
	}
}
