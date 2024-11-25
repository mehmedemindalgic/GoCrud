package crudusersevice_test

import (
	"crud/model"
	"crud/src/internal/service/crudusersevice"
	"fmt"
	"testing"
)

func Test_userService_Create(t *testing.T) {
	mockedStorage := &mockStorageUserCrud{}

	service := crudusersevice.New(crudusersevice.WriteDb(mockedStorage))

	{

		mockedStorage.deleteErr = fmt.Errorf("test")
		mockedStorage.getErr = fmt.Errorf("test")
		mockedStorage.updateErr = fmt.Errorf("test")
		mockedStorage.createErr = fmt.Errorf("test")
		mockedStorage.listErr = fmt.Errorf("test")

		err := service.Delete(model.Usermodel{})
		t.Run("Case 1", func(t *testing.T) {
			if err.Error() != "test" {
				t.Errorf("Expected error to be ")
			}
		})
		_, err = service.Get("1")
		t.Run("Case 2", func(t *testing.T) {
			if err.Error() != "test" {
				t.Errorf("Expected error to be ")
			}
		})
		err = service.Update(model.Usermodel{}, model.Usermodel{})
		t.Run("Case 3", func(t *testing.T) {
			if err.Error() != "test" {
				t.Errorf("Expected error to be ")
			}
		})
		err = service.Create(model.Usermodel{})
		t.Run("Case 4", func(t *testing.T) {
			if err.Error() != "test" {
				t.Errorf("Expected error to be ")
			}
		})
		_, err = service.List()
		t.Run("Case 5", func(t *testing.T) {
			if err.Error() != "test" {
				t.Errorf("Expected error to be ")
			}
		})

	}
	{
		err := service.Update(model.Usermodel{
			Name:  "test",
			Email: "test",
			Pass:  "test",
		}, model.Usermodel{
			Name:  "test",
			Email: "test",
			Pass:  "test",
		},
		)
		t.Run("Case 6", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected error to be ")
			}
		})

	}

}
