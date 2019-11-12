package service

import (
	db2 "github.com/KodaiD/rest/db"
	"github.com/KodaiD/rest/entity"
	"github.com/gin-gonic/gin"
)

type Service struct {}

type User entity.User

func (s Service) GetUsersAll() ([]User, error) {
	db := db2.GetDB()
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (s Service) CreateUserModel(c *gin.Context) (User, error) {
	db := db2.GetDB()
	var u User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s Service) GetUserByID(id string) (User, error) {
	db := db2.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s Service) UpdateUserByID(id string, c *gin.Context) (User, error){
	db := db2.GetDB()
	var u  = User{}

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

func (s Service) DeleteUserByID(id string) error {
	db := db2.GetDB()
	var u  User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}