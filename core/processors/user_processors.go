package processors

import (
	"errors"

	"github.com/mmalcherczyk/rest-api/core/models"
	"github.com/mmalcherczyk/rest-api/core/storages"
)

type UserProcessor struct {
	storage *storages.Storage
}

func NewUserProcessor(storage *storages.Storage) *UserProcessor {
	return &UserProcessor{storage: storage}
}

func (proc *UserProcessor) ListUsers() ([]models.User, error) {
	return proc.storage.ReturnUsers(), nil
}

func (proc *UserProcessor) FindUserByID(id string) (models.User, error) {
	usr := proc.storage.ReturnUsersByID(id)
	if usr.ID != id {
		return usr, errors.New("user not found")
	}
	return usr, nil

}

func (proc *UserProcessor) CreateUser(usr models.User) error {
	return proc.storage.AppendUser(usr)

}
