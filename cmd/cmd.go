package cmd

import "go-mangosteen/internal/router"

func RunServer() {
	r := router.New()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
