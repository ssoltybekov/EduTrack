package services

import (
	"edutrack/internal/db"
	"edutrack/internal/models"
)

type AssignmentService struct {}

func NewAssignmentService() *AssignmentService {
	return &AssignmentService{}
}

func (s *AssignmentService) GetAll() ([]models.Assignment, error) {
	var assignments []models.Assignment
	err := db.DB.Preload("Teacher").Preload("Submissions").Find(&assignments).Error
	if err != nil {
		return nil, err
	}

	return assignments, err
}

func (s *AssignmentService) GetById(id uint) (*models.Assignment, error) {
	var assignment models.Assignment
	err := db.DB.Preload("Teacher").Preload("Submissions").First(&assignment, id).Error
	if err != nil {
		return nil, err
	}
	return &assignment, err
}

func (s *AssignmentService) Create(assignment *models.Assignment) error {
	return db.DB.Create(assignment).Error
}

func (s *AssignmentService) Update(id uint ,updated *models.Assignment) (*models.Assignment, error) {
	var existing models.Assignment
	err := db.DB.First(existing, id).Error
	if err != nil {
		return nil, err
	}

	existing.Title = updated.Title
	existing.Description = updated.Description
	existing.Deadline = updated.Deadline
	existing.TeacherID = updated.TeacherID

	return &existing, nil
}

func (s *AssignmentService) Delete(id uint) error {
	return db.DB.Delete(&models.Assignment{}, id).Error
}
