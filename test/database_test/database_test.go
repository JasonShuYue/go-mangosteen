package data_test

import (
	"go-mangosteen/internal/database"
	"testing"
)

func BenchmarkCrud(b *testing.B) {
	database.PgConnect()
	database.PgCreateTables()
	database.PgMigrate()

	defer database.MysqlClose()

	for i := 0; i < b.N; i++ {
		database.PgCrud()
	}
}
