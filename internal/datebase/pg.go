package datebase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)

func PgConnect() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalln("sql.Open err:", err)
	}

	DB = db
	err = db.Ping()

	if err != nil {
		log.Fatalln("db.Ping error:", err)
	}

	log.Println("Successfully connect to db")
}

func PgMigrate() {
	// 给 User 添加手机字段
	_, err := DB.Exec(`ALTER TABLE users ADD COLUMN phone VARCHAR(50)`)

	if err != nil {
		log.Println(err)
	}

	log.Println("successfully add phone column to users table")

	_, err = DB.Exec(`ALTER TABLE users ADD COLUMN address VARCHAR(200)`)

	if err != nil {
		log.Println(err)
	}

	log.Println("successfully add address column to users table")

	// 新增 Items 表，字段为 id, amount, happened_at, created_at, updated_at
	DB.Exec(`CREATE TABLE IF NOT EXISTS items(
		id SERIAL PRIMARY KEY,
		amount INT NOT NULL,
		happend_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

	)`)

}

func PgCreateTables() {
	// 创建 User 表
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		log.Panicln("DB.Exec err:", err)
	}

	log.Println("Successfully create user table")
}

func PgClose() {
	DB.Close()
	log.Println("Successfully close DB")
}
