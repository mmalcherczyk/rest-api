package app

import (
	"github.com/mmalcherczyk/rest-api/api"
	"github.com/mmalcherczyk/rest-api/core/handlers"
	"github.com/mmalcherczyk/rest-api/core/models"
	"github.com/mmalcherczyk/rest-api/core/processors"
	"github.com/mmalcherczyk/rest-api/core/storages"
)

var users []models.User

func init() {
	users = []models.User{
		{ID: "1", Name: "Daniel", Lastname: "Malcherczyk", PhoneNumber: "secret", PIN: "1234"},
		{ID: "2", Name: "Michał", Lastname: "Malcherczyk", PhoneNumber: "secret", PIN: "4321"},
	}
}

func Serve() {
	storage := storages.NewUserStorage()
	storage.Users = users

	processor := processors.NewUserProcessor(storage)
	handler := handlers.NewUserHandler(processor)
	router := api.CreateRouter(handler)

	router.Run("localhost:9090")

}
