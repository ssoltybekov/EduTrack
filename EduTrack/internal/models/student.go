package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name        string
	Email       string `gorm:"uniqueIndex"`
	Group       string
	Submissions []Submission `gorm:"foreignKey:StudentID"`
}
