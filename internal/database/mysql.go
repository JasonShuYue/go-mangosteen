package database

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

func MysqlCreateTables() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users(
		id INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(50) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createTableSQL)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully create users table")

}

func MysqlClose() {
	DB.Close()
	log.Println("database close successfully")
}
