package services

import (
	"edutrack/internal/db"
	"edutrack/internal/dto"
	"edutrack/internal/models"
	"edutrack/internal/pkg/errors"

	"gorm.io/gorm"
)

type TeacherService struct{}

func NewTeacherService() *TeacherService {
	return &TeacherService{}
}

func (s *TeacherService) GetAll() ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := db.DB.Find(&teachers).Error
	return teachers, err
}

func (s *TeacherService) GetById(id uint) (*dto.TeacherOutputDTO, error) {
	var teacher models.Teacher
	if err := db.DB.First(&teacher, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return &dto.TeacherOutputDTO{
		ID:      teacher.ID,
		Name:    teacher.Name,
		Email:   teacher.Email,
		Subject: teacher.Subject,
	}, nil
}

func (s *TeacherService) Create(input *dto.TeacherInputDTO) (*dto.TeacherOutputDTO, error) {
	teacher := models.Teacher{
		Name:    input.Name,
		Email:   input.Email,
		Subject: input.Subject,
	}
	if err := db.DB.Create(&teacher).Error; err != nil {
		return nil, err
	}
	return &dto.TeacherOutputDTO{
		ID:      teacher.ID,
		Name:    teacher.Name,
		Email:   teacher.Email,
		Subject: teacher.Subject,
	}, nil
}

func (s *TeacherService) Update(id uint, updated *dto.TeacherInputDTO) (*dto.TeacherOutputDTO, error) {
	var existing models.Teacher
	if err := db.DB.First(&existing, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	existing.Name = updated.Name
	existing.Email = updated.Email
	existing.Subject = updated.Subject
	if err := db.DB.Save(&existing).Error; err != nil {
		return nil, err
	}
	return &dto.TeacherOutputDTO{
		ID:      existing.ID,
		Name:    existing.Name,
		Email:   existing.Email,
		Subject: existing.Subject,
	}, nil
}

func (s *TeacherService) Delete(id uint) error {
	if err := db.DB.First(&models.Teacher{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrNotFound
		}
		return err
	}
	return db.DB.Unscoped().Delete(&models.Teacher{}, id).Error
}
