package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name        string
	Email       string `gorm:"uniqueIndex"`
	Group       string
	Submissions []Submission `gorm:"foreignKey:StudentID"`
}

type Teacher struct {
	gorm.Model
	Name        string
	Email       string `gorm:"uniqueIndex"`
	Subject     string
	Assignments []Assignment `gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Assignment struct {
	gorm.Model
	Title       string
	Description string
	Deadline    time.Time
	TeacherID   uint
	Teacher     Teacher      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Submissions []Submission `gorm:"foreignKey:AssignmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`                        
}

type Analytics struct {
	gorm.Model
	StudentID            uint
	Student              Student
	AverageGrade         float64
	CompletedAssignments int
	ProgressLevel        string
}

type Submission struct {
	gorm.Model
	StudentID    uint
	Student      Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AssignmentID uint
	Assignment   Assignment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Grade        float64
	Feedback     string
}
