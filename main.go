package main

import (
	"bestwallet/setup"
	"log"
)

func main() {
	db, err := setup.SetupDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := setup.SetupRouter(db)

	router.Run(":8080")
}
