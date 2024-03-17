package data_test

import (
	"go-mangosteen/internal/database"
	"testing"
)

func BenchmarkCrud(b *testing.B) {
	database.Connect()
	database.CreateTables()
	database.Migrate()

	defer database.Close()

	for i := 0; i < b.N; i++ {
		database.Crud()
	}
}
