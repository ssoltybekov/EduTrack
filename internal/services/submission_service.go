package services

import (
	"time"

	"edutrack/internal/db"
	"edutrack/internal/dto"
	"edutrack/internal/models"
	"edutrack/internal/pkg/errors"

	"gorm.io/gorm"
)

type SubmissionService struct{}

func NewSubmissionService() *SubmissionService {
	return &SubmissionService{}
}

func (s *SubmissionService) Create(input *dto.SubmissionInputDTO, assignmentID uint, studentID uint) (*dto.SubmissionOutputDTO, error) {
	var assignment models.Assignment
	if err := db.DB.First(&assignment, assignmentID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	var user models.User
	if err := db.DB.First(&user, studentID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	if user.Role != models.RoleStudent {
		return nil, errors.ErrInvalidInput 
	}

	
	if time.Now().After(assignment.Deadline) {
		return nil, errors.ErrInvalidInput
	}

	
	var count int64
	if err := db.DB.Model(&models.Submission{}).
		Where("assignment_id = ? AND user_id = ?", assignmentID, studentID).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.ErrInvalidInput 
	}

	submission := models.Submission{
		Content:      input.Content,
		UserID:       studentID,
		AssignmentID: assignmentID,
		SubmittedAt:  time.Now(),
	}

	if err := db.DB.Create(&submission).Error; err != nil {
		return nil, err
	}

	return s.mapToOutput(&submission), nil
}

func (s *SubmissionService) GetAllByAssignment(assignmentID uint) ([]dto.SubmissionOutputDTO, error) {
	var submissions []models.Submission
	if err := db.DB.Preload("User").Where("assignment_id = ?", assignmentID).Find(&submissions).Error; err != nil {
		return nil, err
	}
	var out []dto.SubmissionOutputDTO
	for _, sub := range submissions {
		out = append(out, *s.mapToOutput(&sub))
	}
	return out, nil
}

func (s *SubmissionService) GetById(id uint) (*dto.SubmissionOutputDTO, error) {
	var submission models.Submission
	if err := db.DB.Preload("User").First(&submission, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return s.mapToOutput(&submission), nil
}

func (s *SubmissionService) Grade(id uint, grade float64, feedback string, teacherID uint) (*dto.SubmissionOutputDTO, error) {
	var submission models.Submission
	if err := db.DB.Preload("Assignment.Lesson").First(&submission, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	if submission.Assignment.Lesson.UserID != teacherID {
		return nil, errors.ErrForbidden
	}

	submission.Grade = &grade
	submission.Feedback = feedback

	if err := db.DB.Save(&submission).Error; err != nil {
		return nil, err
	}

	return s.mapToOutput(&submission), nil
}

func (s *SubmissionService) mapToOutput(sub *models.Submission) *dto.SubmissionOutputDTO {
	return &dto.SubmissionOutputDTO{
		ID:           sub.ID,
		Content:      sub.Content,
		Grade:        sub.Grade,
		Feedback:     sub.Feedback,
		SubmittedAt:  sub.SubmittedAt.Format("2006-01-02 15:04"),
		StudentID:    sub.UserID,
		AssignmentID: sub.AssignmentID,
	}
}
