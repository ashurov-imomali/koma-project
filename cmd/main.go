package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main/internal/handler"
	"main/internal/repository"
	"main/internal/router"
	"main/internal/service"
	"net/http"
)

func main() {
	db, err := gorm.Open(postgres.Open("some Settings"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	repos := repository.NewRepos(db)
	srv := service.NewService(repos)
	handler := handler.NewHandler(srv)
	newRoute := router.NewRouter(handler)

	err = http.ListenAndServe("localhost:8080", newRoute)
	if err != nil {
		log.Println(err)
		return
	}
}
