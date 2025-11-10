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

func (s *TeacherService) GetById(id uint) (*models.Teacher, error) {
	var teacher models.Teacher
	err := db.DB.First(&teacher, id).Error
	if err != nil {
		return nil, err
	}
	return &teacher, err
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

func (s *TeacherService) Update(teacher *models.Teacher) error {
	return db.DB.Save(teacher).Error
}

func (s *TeacherService) Delete(id uint) error {
	return db.DB.Unscoped().Delete(&models.Teacher{}, id).Error
}
