package main

import (
	"fmt"
	"os"

	"github.com/B-P-Coders/backend-template/config"
	"github.com/B-P-Coders/backend-template/database"
	"github.com/B-P-Coders/backend-template/router"
)

func main() {
	config.Load()

	db, err := database.Setup()
	if err != nil {
		fmt.Println(fmt.Errorf("Error setting up database: %s", err.Error()))
		os.Exit(2)
	}
	r := router.Setup(db)
	r.Listen(fmt.Sprintf(":%d", config.Conf.SrvConfig.Port))
}
