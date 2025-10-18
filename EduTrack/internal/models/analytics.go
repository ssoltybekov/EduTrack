package models

import "gorm.io/gorm"

type Analytics struct {
	gorm.Model
	StudentID uint
	Student   Student
	AverageGrade float64
	CompletedAssignments int
	ProgressLevel string 
}
