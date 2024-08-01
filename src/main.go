package main

import (
	"log"
	"sync"
	"z7pz/golangapi/src/discord"
	"z7pz/golangapi/src/rest"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	// TODO setup enviroment variables
	go func() {
		defer wg.Done()
		client := discord.Discord{Token: "my token"}
		client.Init(db)
	}()

	go func() {
		defer wg.Done()
		client := rest.Rest{Port: "8080"}
		client.Init(db)
	}()
	wg.Wait()
}
