package cruduserstroage

import (
	"crud/model"

	"gorm.io/gorm"
)

type userModel = model.Usermodel

type StorageUserCrud interface {
	Create(userModel) error
	Get(id string) (userModel, error)
	List() ([]userModel, error)
	Update(new userModel) error
	Delete(userModel) error
}
type storageusercrud struct {
	db *gorm.DB
}

type UserCrudOption func(*storageusercrud)

func ConnectDb(db *gorm.DB) UserCrudOption {
	return func(s *storageusercrud) {
		s.db = db
	}
}
func New(options ...UserCrudOption) StorageUserCrud {
	k := &storageusercrud{}
	for _, o := range options {
		o(k)
	}
	return k
}
