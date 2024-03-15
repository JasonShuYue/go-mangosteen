package cmd

import (
	"go-mangosteen/internal/database"
	"go-mangosteen/internal/router"
	"log"

	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use: "mangosteen",
	}

	srvCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			RunServer()
		},
	}

	createCmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			database.MysqlCreateTables()
		},
	}

	dbCmd := &cobra.Command{
		Use: "db",
	}

	mgrtCmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			database.PgMigrate()
		},
	}

	crudCmd := &cobra.Command{
		Use: "crud",
		Run: func(cmd *cobra.Command, args []string) {
			database.PgCrud()
		},
	}

	database.PgConnect()
	rootCmd.AddCommand(srvCmd, dbCmd)
	dbCmd.AddCommand(createCmd, mgrtCmd, crudCmd)

	defer database.PgClose()

	rootCmd.Execute()

}

func RunServer() {

	r := router.New()
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("next!!!!!")
}
