package main

import (
	"log"
	"net/http"

	"github.com/nandaarta/enhance-me/internal/driver"
	"github.com/nandaarta/enhance-me/internal/user"
)

func main() {
	driver.RunMigration()

	svc, err := user.NewService("admin", "admin")
	if err != nil {
		log.Fatal(err)
	}

	h := user.Handler{Svc: *svc}

	http.HandleFunc("/user", h.AddUser)

	log.Println("Http server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
