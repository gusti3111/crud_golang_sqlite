package models

// buat model database menggunakan struct

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID      uint `gorm:"primaryKey;autoIncrement:true"`
	Name    string
	Content string
	Date    string
	IsDone  bool
}
