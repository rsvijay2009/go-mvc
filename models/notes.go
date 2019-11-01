package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// AddNote to create a new note
func (n *Notes) AddNote() int64 {
	db := orm.NewOrm()
	pk, err := db.Insert(n)

	if err != nil {
		panic(err)
	}

	return pk
}

// GetNotesByUserID to get all the notes of a specific user
func GetNotesByUserID(uid int) []*Notes {
	db := orm.NewOrm()
	var notes []*Notes
	db.QueryTable(new(Notes)).Filter("user_id", uid).All(&notes)

	return notes
}

// GetNotesByID to get the specific note
func GetNotesByID(nid int) Notes {
	db := orm.NewOrm()
	var notes Notes
	db.QueryTable(new(Notes)).Filter("id", nid).One(&notes)

	return notes
}

// UpdateNote to update a note
func (n *Notes) UpdateNote() int64 {
	db := orm.NewOrm()
	rowsAffected, err := db.QueryTable("notes").Filter("id", n.Id).Update(orm.Params{
		"comments":   n.Comments,
		"updated_at": time.Now(),
	})

	if err != nil {
		panic(err)
	}

	return rowsAffected
}

// DeleteNote to update a note
func (n *Notes) DeleteNote() int64 {
	db := orm.NewOrm()
	rowsAffected, err := db.QueryTable("notes").Filter("id", n.Id).Delete()

	if err != nil {
		panic(err)
	}

	return rowsAffected
}
