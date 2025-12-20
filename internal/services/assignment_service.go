package services

import (
	"edutrack/internal/db"
	"edutrack/internal/dto"
	"edutrack/internal/models"
	"time"
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

func (s *AssignmentService) Create(input *dto.AssignmentInputDTO) (*dto.AssignmentOutputDTO, error) {
	deadline, err := time.Parse("2006-01-02", input.Deadline)
	if err != nil {
		return nil, err
	}

	assignment := models.Assignment {
		Title:  	 input.Title,
		Description: input.Description,
		Deadline:	 deadline,
		TeacherID: 	 input.TeacherID,   
	}

	if err := db.DB.Create(&assignment).Error; err != nil {
		return nil, err
	}

	return &dto.AssignmentOutputDTO{
		ID:          assignment.ID,
		Title:       assignment.Title,
		Description: assignment.Description,
		Deadline:    assignment.Deadline.Format("2006-01-02"),
		TeacherID:   assignment.TeacherID,
	}, nil
}

func (s *AssignmentService) Update(id uint ,updated *dto.AssignmentInputDTO) (*dto.AssignmentOutputDTO, error) {
	var existing models.Assignment
	err := db.DB.First(existing, id).Error
	if err != nil {
		return nil, err
	}

	deadline, err := time.Parse("2006-01-02", updated.Deadline)
	if err != nil {
		return nil, err
	}

	existing.Title = updated.Title
	existing.Description = updated.Description
	existing.Deadline = deadline
	existing.TeacherID = updated.TeacherID

	return &dto.AssignmentOutputDTO{
        ID:          existing.ID,
        Title:       existing.Title,
        Description: existing.Description,
        Deadline:    existing.Deadline.Format("2006-01-02"),
        TeacherID:   existing.TeacherID,
    }, nil
}

func (s *AssignmentService) Delete(id uint) error {
	return db.DB.Unscoped().Delete(&models.Assignment{}, id).Error
}
