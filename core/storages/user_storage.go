package storages

import (
	"github.com/mmalcherczyk/rest-api/core/models"
)

type Storage struct {
	Users []models.User
}

func NewUserStorage() *Storage {
	return &Storage{}
}

func (storage *Storage) ReturnUsers() []models.User {
	return storage.Users
}

func (storage *Storage) ReturnUsersByID(id string) models.User {
	var result models.User
	for _, usr := range storage.Users {
		if usr.ID == id {
			result = usr
		}
	}
	return result
}

func (storage *Storage) AppendUser(usr models.User) error {
	storage.Users = append(storage.Users, usr)
	return nil

}
