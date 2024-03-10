package datebase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlHost     = "localhost"
	mysqlPort     = 3306
	mysqlUser     = "mangosteen"
	mysqlPassword = "123456"
	mysqlDbname   = "mangosteen_dev"
)

func MysqlConnect() {
	conStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlDbname)
	db, err := sql.Open("mysql", conStr)

	if err != nil {
		log.Fatalln("sql.Open err:", err)
	}

	DB = db

	err = db.Ping()

	if err != nil {
		log.Fatalln("db.Ping error:", err)
	}

	log.Println("Successfully connect Mysql")
}

func MysqlClose() {
	DB.Close()
	log.Println("database close successfully")
}
