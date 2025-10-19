package services

import (
	"edutrack/internal/db"
	"edutrack/internal/models"
)

type StudentService struct{}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (s *StudentService) GetAll() ([]models.Student, error) {
	var students []models.Student
	err := db.DB.Find(students).Error
	if err != nil {
		return nil, err
	}
	return students, err
}

func (s *StudentService) GetById(id uint) (*models.Student, error) {
	var student models.Student
	err := db.DB.First(student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, err
}

func (s *StudentService) Create(student *models.Student) error {
	return db.DB.Create(student).Error
}

func (s *StudentService) Update(student *models.Student) error {
	return db.DB.Save(student).Error
}

func (s *StudentService) Delete(id uint) error {
	return db.DB.Delete(&models.Student{}, id).Error
}
