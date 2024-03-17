package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID        uint    // Standard field for the primary key
	Name      string  // A regular string field
	Email     *string // A pointer to a string, allowing for null values
	Phone     *string
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}

type Item struct {
	ID        int
	UserID    int
	Amount    int
	HappendAt time.Time
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}

type Tag struct {
	ID   int
	Name string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)

var models = []any{&User{}, &Item{}, &Tag{}}

func Connect() {
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d sslmode=disable`,
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicln(err)
	}

	DB = db
}

func Migrate() {
	DB.Migrator().AutoMigrate(models...)
}

func CreateTables() {
	for _, model := range models {
		err := DB.Migrator().CreateTable(model)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("Successfully create table user")
}

func DeleteTable() {
	DB.Migrator().DropTable(&User{})
	DB.Migrator().DropTable("users")
}

func Close() {
	sqlDB, err := DB.DB()

	if err != nil {
		log.Panicln(err)
	}

	sqlDB.Close()
}

func Crud() {

}
