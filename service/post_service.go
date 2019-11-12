package service

import (
	db2 "github.com/KodaiD/rest/db"
	"github.com/KodaiD/rest/entity"
	"github.com/gin-gonic/gin"
)

type Post entity.Post

func (s Service) GetPostsAll() ([]Post, error) {
	db := db2.GetDB()
	var p []Post

	if err := db.Find(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (s Service) CreatePostModel(c *gin.Context) (Post, error) {
	db := db2.GetDB()
	var p Post

	if err := c.BindJSON(&p); err != nil {
		return p, err
	}

	if err := db.Create(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (s Service) GetPostByID(id string) (Post, error) {
	db := db2.GetDB()
	var p Post

	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (s Service) UpdatePostByID(id string, c *gin.Context) (Post, error){
	db := db2.GetDB()
	var p  Post

	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}

	if err := c.BindJSON(&p); err != nil {
		return p, err
	}

	db.Save(&p)

	return p, nil
}

func (s Service) DeletePostByID(id string) error {
	db := db2.GetDB()
	var p Post

	if err := db.Where("id = ?", id).Delete(&p).Error; err != nil {
		return err
	}

	return nil
}