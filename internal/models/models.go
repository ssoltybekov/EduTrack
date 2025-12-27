package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleTeacher Role = "teacher"
	RoleStudent Role = "student"
)

type User struct {
	gorm.Model
	Name         string `json:"name" gorm:"not null"`
	Email        string `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string `json:"-"`
	Role         Role   `json:"role" gorm:"type:varchar(20);not null"`
	Subject      string `json:"subject,omitempty"`
	Group        string `json:"group,omitempty"`

	Lessons     []Lesson     `gorm:"foreignKey:UserID;references:ID"`
	Submissions []Submission `gorm:"foreignKey:UserID;references:ID"`
}

type Lesson struct {
	gorm.Model
	Title       string       `json:"title" gorm:"not null"`
	Description string       `json:"description"`
	VideoURL    string       `json:"video_url" gorm:"not null"`
	UserID      uint         `json:"teacher_id"`
	User        User         `gorm:"foreignKey:UserID;references:ID"`
	Assignments []Assignment `gorm:"foreignKey:LessonID;references:ID"`
}

// type Student struct {
// 	gorm.Model
// 	Name        string
// 	Email       string `gorm:"uniqueIndex"`
// 	Group       string
// 	Submissions []Submission `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// }

// type Teacher struct {
// 	gorm.Model
// 	Name         string
// 	Email        string `gorm:"uniqueIndex"`
// 	Subject      string
// 	PasswordHash string       `json:"-" gorm:"not null"`
// 	Assignments  []Assignment `gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// }

// type Assignment struct {
// 	gorm.Model
// 	Title       string
// 	Description string
// 	Deadline    time.Time
// 	TeacherID   uint
// 	Teacher     Teacher      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// 	Submissions []Submission `gorm:"foreignKey:AssignmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// }

type Assignment struct {
	gorm.Model
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Deadline    time.Time    `json:"deadline"`
	LessonID    uint         `json:"lesson_id" gorm:"not null"`
	Lesson      Lesson       `gorm:"foreignKey:LessonID;references:ID"`
	Submissions []Submission `gorm:"foreignKey:AssignmentID;references:ID"`
}

// type Analytics struct {
// 	gorm.Model
// 	StudentID            uint
// 	Student              Student
// 	AverageGrade         float64
// 	CompletedAssignments int
// 	ProgressLevel        string
// }

// type Submission struct {
// 	gorm.Model
// 	StudentID    uint
// 	Student      Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// 	AssignmentID uint
// 	Assignment   Assignment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// 	Grade        float64
// 	Feedback     string
// }

type Submission struct {
	gorm.Model
	Content     string    `json:"content" gorm:"type:text"`
	Grade       *float64  `json:"grade"`
	Feedback    string    `json:"feedback" gorm:"type:text"`
	SubmittedAt time.Time `json:"submitted_at" gorm:"default:current_timestamp"`

	UserID       uint `json:"student_id" gorm:"not null"`
	AssignmentID uint `json:"assignment_id" gorm:"not null"`

	User       User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Assignment Assignment `gorm:"foreignKey:AssignmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Analytics struct {
	gorm.Model
	UserID             uint    `json:"user_id"`
	AverageGrade       float64 `json:"average_grade"`
	CompletedLessons   int     `json:"completed_lessons"`
	ProgressPercentage float64 `json:"progress_percentage"`
}
