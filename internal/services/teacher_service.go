package services

import (
	"edutrack/internal/db"
	"edutrack/internal/dto"
	"edutrack/internal/models"
)

type TeacherService struct{}

func NewTeacherService() *TeacherService {
	return &TeacherService{}
}

func (s *TeacherService) GetAll() ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := db.DB.Find(&teachers).Error
	if err != nil {
		return nil, err
	}
	
	return teachers, err
}

func (s *TeacherService) GetById(id uint) (*dto.TeacherOutputDTO, error) {
	var teacher models.Teacher
	err := db.DB.First(&teacher, id).Error
	if err != nil {
		return nil, err
	}
	return &dto.TeacherOutputDTO{
		ID: teacher.ID,
		Name: teacher.Name,
		Email: teacher.Email,
		Subject: teacher.Subject,
	}, err
}

func (s *TeacherService) Create(input *dto.TeacherInputDTO) (*dto.TeacherOutputDTO, error) {
	teacher := models.Teacher{
		Name: input.Name,
		Email: input.Email,
		Subject: input.Subject,
	}
	
	if err := db.DB.Create(&teacher).Error; err != nil {
		return nil, err
	}

	output := &dto.TeacherOutputDTO {
		ID: teacher.ID,
		Name: teacher.Name,
		Email: teacher.Email,
		Subject: teacher.Subject,
	}
	return output, nil
}

func (s *TeacherService) Update(id uint, updated *dto.TeacherInputDTO) (*dto.TeacherOutputDTO, error) {
	var existing models.Teacher
	if err := db.DB.First(&existing).Error; err != nil {
		return nil, err
	}

	existing.Name = updated.Name
	existing.Email = updated.Email
	existing.Subject = updated.Subject

	if err := db.DB.Save(&existing).Error; err != nil {
		return nil, err
	}

	return &dto.TeacherOutputDTO{
		ID: existing.ID,
		Name: existing.Name,
		Email: existing.Email,
		Subject: existing.Subject,
	}, nil
}

func (s *TeacherService) Delete(id uint) error {
	return db.DB.Unscoped().Delete(&models.Teacher{}, id).Error
}
