package crudusersevice

import (
	"crud/model"
	"fmt"
	"strconv"
)

func (s *serviceusercrud) Create(a model.Usermodel) error {

	allusers, err := s.storage.List()
	if err != nil {
		return err
	}
	// user, err := s.storage.GetByEmail(a.Email)
	// if err == nil {
	// return fmt.Errorf("email is already exist")
	// }
	for _, user := range allusers {
		if user.Email == a.Email {
			return fmt.Errorf("email is already exist")
		}
	}
	err = s.storage.Create(a)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceusercrud) Get(key string) (model.Usermodel, error) {
	user, err := s.storage.Get(key)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *serviceusercrud) List() ([]model.Usermodel, error) {

	user, err := s.storage.List()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *serviceusercrud) Update(old model.Usermodel, new model.Usermodel) error {

	olduser, err := s.storage.Get(strconv.FormatUint(uint64(old.ID), 10))
	if err != nil {
		return err
	}
	if olduser.ID == 0 {
		return fmt.Errorf("user not found")
	}
	if olduser.ID != new.ID && new.Email != "" && new.Name != "" && new.Pass != "" {
		return fmt.Errorf("invalid data")
	}
	alluser, err := s.storage.List()
	if err != nil {
		return err
	}
	for _, user := range alluser {
		if user.Email == new.Email && user.ID != new.ID {
			return fmt.Errorf("email is already exist")
		}
	}
	err = s.storage.Update(new)
	if err != nil {
		return err
	}
	return nil

}

func (s *serviceusercrud) Delete(a model.Usermodel) error {

	user, err := s.storage.Get(strconv.FormatUint(uint64(a.ID), 10))
	if err != nil {
		return err
	}
	err = s.storage.Delete(user)
	if err != nil {
		return err
	}
	return nil

}
