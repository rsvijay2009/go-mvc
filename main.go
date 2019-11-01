package main

import (
	"fmt"
	"log"

	_ "github.com/rsvijay2009/go-mvc/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(db_container:3306)/webapp?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Database connection failed!")
		panic(err)
	} else {
		fmt.Println("Database connected successfully!")
	}
}
func main() {
	beego.Run()
}
