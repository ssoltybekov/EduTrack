package services

import (
	"edutrack/internal/db"
	"edutrack/internal/dto"
	"edutrack/internal/models"
	"edutrack/internal/pkg/errors"

	"gorm.io/gorm"
)

type LessonService struct{}

func NewLessonService() *LessonService {
	return &LessonService{}
}

func (s *LessonService) Create(input *dto.LessonInputDTO, teacherID uint) (*dto.LessonOutputDTO, error) {
	lesson := models.Lesson{
		Title:       input.Title,
		Description: input.Description,
		VideoURL:    input.VideoURL,
		UserID:      teacherID,
	}
	if err := db.DB.Create(&lesson).Error; err != nil {
		return nil, err
	}
	return s.mapToOutput(&lesson), nil
}

func (s *LessonService) GetAll() ([]dto.LessonOutputDTO, error) {
	var lessons []models.Lesson
	if err := db.DB.Preload("User").Find(&lessons).Error; err != nil {
		return nil, err
	}
	var out []dto.LessonOutputDTO
	for _, l := range lessons {
		out = append(out, *s.mapToOutput(&l))
	}
	return out, nil
}

func (s *LessonService) GetById(id uint) (*dto.LessonOutputDTO, error) {
	var lesson models.Lesson
	if err := db.DB.Preload("User").First(&lesson, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return s.mapToOutput(&lesson), nil
}

func (s *LessonService) Update(id uint, input *dto.LessonInputDTO, teacherID uint) (*dto.LessonOutputDTO, error) {
	var lesson models.Lesson
	if err := db.DB.First(&lesson, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	if lesson.UserID != teacherID {
		return nil, errors.ErrForbidden
	}

	lesson.Title = input.Title
	lesson.Description = input.Description
	lesson.VideoURL = input.VideoURL

	if err := db.DB.Save(&lesson).Error; err != nil {
		return nil, err
	}
	return s.mapToOutput(&lesson), nil
}

func (s *LessonService) Delete(id uint, teacherID uint) error {
	var lesson models.Lesson
	if err := db.DB.First(&lesson, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrNotFound
		}
		return err
	}

	if lesson.UserID != teacherID {
		return errors.ErrForbidden
	}

	return db.DB.Unscoped().Delete(&lesson, id).Error
}

func (s *LessonService) mapToOutput(lesson *models.Lesson) *dto.LessonOutputDTO {
	return &dto.LessonOutputDTO{
		ID:          lesson.ID,
		Title:       lesson.Title,
		Description: lesson.Description,
		VideoURL:    lesson.VideoURL,
		TeacherID:   lesson.UserID,
	}
}
