package services

import (
	"edutrack/internal/db"
	"edutrack/internal/dto"
	"edutrack/internal/models"
	"edutrack/internal/pkg/errors"

	"gorm.io/gorm"
)

type StudentService struct{}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (s *StudentService) GetAll() ([]models.Student, error) {
	var students []models.Student
	err := db.DB.Find(&students).Error
	return students, err
}

func (s *StudentService) GetById(id uint) (*dto.StudentOutputDTO, error) {
	var student models.Student
	if err := db.DB.First(&student, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return &dto.StudentOutputDTO{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email,
		Group: student.Group,
	}, nil
}

func (s *StudentService) Create(input *dto.StudentInputDTO) (*dto.StudentOutputDTO, error) {
	student := models.Student{
		Name:  input.Name,
		Email: input.Email,
		Group: input.Group,
	}
	if err := db.DB.Create(&student).Error; err != nil {
		return nil, err
	}
	return &dto.StudentOutputDTO{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email,
		Group: student.Group,
	}, nil
}

func (s *StudentService) Update(id uint, updated *dto.StudentInputDTO) (*dto.StudentOutputDTO, error) {
	var existing models.Student
	if err := db.DB.First(&existing, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	existing.Name = updated.Name
	existing.Email = updated.Email
	existing.Group = updated.Group
	if err := db.DB.Save(&existing).Error; err != nil {
		return nil, err
	}
	return &dto.StudentOutputDTO{
		ID:    existing.ID,
		Name:  existing.Name,
		Email: existing.Email,
		Group: existing.Group,
	}, nil
}

func (s *StudentService) Delete(id uint) error {
	if err := db.DB.First(&models.Student{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrNotFound
		}
		return err
	}
	return db.DB.Unscoped().Delete(&models.Student{}, id).Error
}
