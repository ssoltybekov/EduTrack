package models

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	Title       string
	Description string
	Deadline    time.Time
	TeacherID   uint
	Teacher     Teacher      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // связь с Teacher
	Submissions []Submission `gorm:"foreignKey:AssignmentID"`                        // связь с Submission
}
