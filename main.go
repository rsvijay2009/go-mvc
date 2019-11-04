package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/rsvijay2009/go-mvc/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	mysqlUsername, _ := os.LookupEnv("MYSQL_USER_NAME")
	mysqlPassword, _ := os.LookupEnv("MYSQL_ROOT_PASSWORD")
	dbContainer, _ := os.LookupEnv("DB_CONTAINER_NAME")
	dbName, _ := os.LookupEnv("MYSQL_DATABASE_NAME")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	err := orm.RegisterDataBase("default", "mysql", mysqlUsername+":"+mysqlPassword+"@tcp("+dbContainer+":3306)/"+dbName+"?charset=utf8&parseTime=True")

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
