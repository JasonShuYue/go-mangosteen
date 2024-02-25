package cmd

import (
	"go-mangosteen/internal/datebase"
	"go-mangosteen/internal/router"
	"log"
)

func RunServer() {
	datebase.Connect()
	datebase.CreateTables()
	defer datebase.Close()

	r := router.New()
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("next!!!!!")
}
