package cruduserstroage

import (
	"fmt"
)

func (k *storageusercrud) Create(a userModel) error {

	err := k.db.Create(&a).Error
	if err != nil {
		return fmt.Errorf("Stroage/Create: " + err.Error())
	}
	return nil
}

func (k *storageusercrud) Get(key string) (userModel, error) {

	var user userModel
	err := k.db.Where("id = ?", key).Find(&user).Error
	if err != nil {
		return user, fmt.Errorf("Stroage/Get: " + err.Error())
	}
	return user, nil

}

func (k *storageusercrud) List() ([]userModel, error) {

	var users []userModel
	err := k.db.Unscoped().Find(&users).Error // silienleri de getirir
	//err := k.db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("Stroage/List: " + err.Error())
	}
	return users, nil

}

func (k *storageusercrud) Update(new userModel) error {

	err := k.db.Updates(new).Error
	if err != nil {
		return fmt.Errorf("Stroage/Update: " + err.Error())
	}
	return nil
}

func (k *storageusercrud) Delete(a userModel) error {

	err := k.db.Delete(&a).Error
	if err != nil {
		return fmt.Errorf("Stroage/Delete: " + err.Error())
	}
	return nil
}
