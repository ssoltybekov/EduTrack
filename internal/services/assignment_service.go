package services

import (
    "time"

    "edutrack/internal/db"
    "edutrack/internal/dto"
    "edutrack/internal/models"
    "edutrack/internal/pkg/errors"

    "gorm.io/gorm"
)

type AssignmentService struct{}

func NewAssignmentService() *AssignmentService {
    return &AssignmentService{}
}

func (s *AssignmentService) GetAllByLesson(lessonID uint) ([]dto.AssignmentOutputDTO, error) {
    var assignments []models.Assignment
    err := db.DB.Where("lesson_id = ?", lessonID).Find(&assignments).Error
    if err != nil {
        return nil, err
    }
    var output []dto.AssignmentOutputDTO
    for _, a := range assignments {
        output = append(output, dto.AssignmentOutputDTO{
            ID:          a.ID,
            Title:       a.Title,
            Description: a.Description,
            Deadline:    a.Deadline.Format("2006-01-02"),
            LessonID:    a.LessonID,
        })
    }
    return output, nil
}

func (s *AssignmentService) GetById(id uint) (*dto.AssignmentOutputDTO, error) {
    var assignment models.Assignment
    err := db.DB.First(&assignment, id).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.ErrNotFound
        }
        return nil, err
    }
    return &dto.AssignmentOutputDTO{
        ID:          assignment.ID,
        Title:       assignment.Title,
        Description: assignment.Description,
        Deadline:    assignment.Deadline.Format("2006-01-02"),
        LessonID:    assignment.LessonID,
    }, nil
}

func (s *AssignmentService) Create(input *dto.AssignmentInputDTO, lessonID uint, teacherID uint) (*dto.AssignmentOutputDTO, error) {
    var lesson models.Lesson
    if err := db.DB.First(&lesson, lessonID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.ErrNotFound
        }
        return nil, err
    }
    if lesson.UserID != teacherID {
        return nil, errors.ErrForbidden
    }
    deadline, err := time.Parse("2006-01-02", input.Deadline)
    if err != nil {
        return nil, errors.ErrInvalidInput
    }
    assignment := models.Assignment{
        Title:       input.Title,
        Description: input.Description,
        Deadline:    deadline,
        LessonID:    lessonID,
    }
    if err := db.DB.Create(&assignment).Error; err != nil {
        return nil, err
    }
    return &dto.AssignmentOutputDTO{
        ID:          assignment.ID,
        Title:       assignment.Title,
        Description: assignment.Description,
        Deadline:    assignment.Deadline.Format("2006-01-02"),
        LessonID:    assignment.LessonID,
    }, nil
}

func (s *AssignmentService) Update(id uint, input *dto.AssignmentInputDTO, teacherID uint) (*dto.AssignmentOutputDTO, error) {
    var assignment models.Assignment
    if err := db.DB.First(&assignment, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.ErrNotFound
        }
        return nil, err
    }
    var lesson models.Lesson
    if err := db.DB.First(&lesson, assignment.LessonID).Error; err != nil {
        return nil, err
    }
    if lesson.UserID != teacherID {
        return nil, errors.ErrForbidden
    }
    deadline, err := time.Parse("2006-01-02", input.Deadline)
    if err != nil {
        return nil, errors.ErrInvalidInput
    }
    assignment.Title = input.Title
    assignment.Description = input.Description
    assignment.Deadline = deadline
    if err := db.DB.Save(&assignment).Error; err != nil {
        return nil, err
    }
    return &dto.AssignmentOutputDTO{
        ID:          assignment.ID,
        Title:       assignment.Title,
        Description: assignment.Description,
        Deadline:    assignment.Deadline.Format("2006-01-02"),
        LessonID:    assignment.LessonID,
    }, nil
}

func (s *AssignmentService) Delete(id uint, teacherID uint) error {
    var assignment models.Assignment
    if err := db.DB.First(&assignment, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.ErrNotFound
        }
        return err
    }
    var lesson models.Lesson
    if err := db.DB.First(&lesson, assignment.LessonID).Error; err != nil {
        return err
    }
    if lesson.UserID != teacherID {
        return errors.ErrForbidden
    }
    return db.DB.Unscoped().Delete(&models.Assignment{}, id).Error
}