package models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	StudentID    uint
	Student      Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AssignmentID uint
	Assignment   Assignment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Grade        float64
	Feedback     string
}
