package main

import (
	"kredit-api/db"
	"kredit-api/router"
	"log"
)

func main() {
	gormDB, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	err = router.Router(gormDB)
	if err != nil {
		log.Fatal(err)
	}
}
