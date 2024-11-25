package crudusersevice_test

import "crud/model"

type mockStorageUserCrud struct {
	deleteErr error
	getErr    error
	updateErr error
	createErr error
	listErr   error
	model     model.Usermodel
	modelarr  []model.Usermodel
}

func (m *mockStorageUserCrud) Create(model.Usermodel) error {
	return m.createErr
}

func (m *mockStorageUserCrud) Get(key string) (model.Usermodel, error) {
	return m.model, m.getErr
}

func (m *mockStorageUserCrud) List() ([]model.Usermodel, error) {
	return m.modelarr, m.listErr
}

func (m *mockStorageUserCrud) Update(new model.Usermodel) error {
	return m.updateErr
}

func (m *mockStorageUserCrud) Delete(a model.Usermodel) error {
	return m.deleteErr
}
