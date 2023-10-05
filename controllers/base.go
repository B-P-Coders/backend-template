package controllers

import "gorm.io/gorm"

type Controller struct {
	Db *gorm.DB
}
