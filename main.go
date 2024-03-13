package main

import (
	"fmt"
	"log"
	"tokenbased-auth/database"
	"tokenbased-auth/internal"
	"tokenbased-auth/service"
)

func init() {
	var err error
	// DatabaseConnect will connect to the database
	internal.DB, err = database.DatabaseConnect()
	if err != nil {
		log.Fatal(err)
	}

	internal.TokenSignatureKey, err = database.GetTokenSignatureKey()
	if err != nil {
		log.Fatalln("Error getting token signature key", err)
	}
}

func main() {
	// InitRouter will return a router instance
	router := service.InitRouter()

	err := router.Run(internal.Port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is running on port: ", internal.Port)
}
