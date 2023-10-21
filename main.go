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
		fmt.Println(fmt.Errorf("error setting up database: %s", err.Error()))
		os.Exit(2)
	}
	r := router.Setup(db)
	r.Listen(fmt.Sprintf("127.0.0.1:%d", config.Conf.SrvConfig.Port))
}
