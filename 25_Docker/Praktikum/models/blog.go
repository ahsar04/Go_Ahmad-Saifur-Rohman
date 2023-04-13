package models

import (
	"github.com/jinzhu/gorm"
)

type Blog struct {
	gorm.Model
	User_id   uint    `json:"user_id"`
    User      User    `gorm:"foreignKey:user_id" json:"user"`
	Judul     string `json:"judul" form:"judul"`
	Konten    string `json:"konten" form:"konten"`
}