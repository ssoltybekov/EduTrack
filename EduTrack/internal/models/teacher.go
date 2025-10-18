package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name        string
	Email       string `gorm:"uniqueIndex"`
	Subject     string
	Assignments []Assignment `gorm:"foreignKey:TeacherID"`
}
