package services

import (
	"edutrack/internal/db"
	"edutrack/internal/dto"
	"edutrack/internal/models"
	"edutrack/internal/pkg/errors"
	"log"

	"gorm.io/gorm"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAll() ([]dto.UserOutputDTO, error) {
	var users []models.User
	err := db.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	var output []dto.UserOutputDTO
	for _, u := range users {
		output = append(output, dto.UserOutputDTO{
			ID:      u.ID,
			Name:    u.Name,
			Email:   u.Email,
			Role:    string(u.Role),
			Subject: u.Subject,
			Group:   u.Group,
		})
	}
	return output, nil
}

func (s *UserService) GetById(id uint) (*dto.UserOutputDTO, error) {
	var user models.User
	err := db.DB.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return &dto.UserOutputDTO{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Role:    string(user.Role),
		Subject: user.Subject,
		Group:   user.Group,
	}, nil
}

func (s *UserService) Create(input *dto.UserInputDTO) (*dto.UserOutputDTO, error) {
	user := models.User{
		Name:    input.Name,
		Email:   input.Email,
		Role:    models.Role(input.Role),
		Subject: input.Subject,
		Group:   input.Group,
	}
	if err := db.DB.Create(&user).Error; err != nil {
		log.Println("Ошибка создания пользователя:", err) 
		return nil, err
	}
	return &dto.UserOutputDTO{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Role:    string(user.Role),
		Subject: user.Subject,
		Group:   user.Group,
	}, nil
}

func (s *UserService) Update(id uint, input *dto.UserInputDTO) (*dto.UserOutputDTO, error) {
	var existing models.User
	if err := db.DB.First(&existing, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	if string(existing.Role) != input.Role {
		return nil, errors.ErrInvalidInput
	}
	existing.Name = input.Name
	existing.Email = input.Email
	existing.Subject = input.Subject
	existing.Group = input.Group
	if err := db.DB.Save(&existing).Error; err != nil {
		return nil, err
	}
	return &dto.UserOutputDTO{
		ID:      existing.ID,
		Name:    existing.Name,
		Email:   existing.Email,
		Role:    string(existing.Role),
		Subject: existing.Subject,
		Group:   existing.Group,
	}, nil
}

func (s *UserService) Delete(id uint) error {
	if err := db.DB.First(&models.User{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrNotFound
		}
		return err
	}
	return db.DB.Unscoped().Delete(&models.User{}, id).Error
}
