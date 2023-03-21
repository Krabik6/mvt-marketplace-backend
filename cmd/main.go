package main

import (
	"log"
	"mvt-marketplace-backend/handler"
	"mvt-marketplace-backend/repository"
	"net/http"
)

func main() {
	db, err := repository.NewPostgresDB(
		repository.Config{
			Host:     "db",
			Port:     "5432",
			Username: "postgres",
			Password: "qwerty",
			DBName:   "postgres",
			SSLMode:  "disable",
		})

	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)

	handlers := handler.NewHandler(repos)

	server := http.Server{
		Handler: handlers.InitRouts(),
		Addr:    ":8000",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
