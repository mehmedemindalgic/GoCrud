package crudusersevice

import (
	"crud/model"
	"crud/src/internal/storage/postgres/cruduserstroage"
)

type ServiceUserCrud interface {
	Create(model.Usermodel) error
	Get(string) (model.Usermodel, error)
	List() ([]model.Usermodel, error)
	Update(model.Usermodel, model.Usermodel) error
	Delete(model.Usermodel) error
}
type serviceusercrud struct {
	storage cruduserstroage.StorageUserCrud
}

type serviceOption func(*serviceusercrud)

func WriteDb(storage cruduserstroage.StorageUserCrud) serviceOption {
	return func(s *serviceusercrud) {
		s.storage = storage
	}
}

func New(options ...serviceOption) ServiceUserCrud {
	k := &serviceusercrud{}
	for _, o := range options {
		o(k)
	}
	return k
}
