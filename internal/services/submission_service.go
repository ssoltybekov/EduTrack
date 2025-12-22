package services

import (
	"edutrack/internal/db"
	"edutrack/internal/dto"
	"edutrack/internal/models"
	"errors"
)

type SubmissionService struct{}

func NewSubmissionService() *SubmissionService {
	return &SubmissionService{}
}

func (s *SubmissionService) GetAll() ([]models.Submission, error) {
	var submissions []models.Submission
	err := db.DB.Preload("Student").Preload("Assignment").Find(submissions).Error
	if err != nil {
		return nil, err
	}

	return submissions, err
}

func (s *SubmissionService) GetById(id uint) (*models.Submission, error) {
	var submission models.Submission
	err := db.DB.Preload("Student").Preload("Assignment").First(submission, id).Error
	if err != nil {
		return nil, err
	}
	return &submission, err
}

func (s *SubmissionService) Create(input *dto.SubmissionInputDTO) (*dto.SubmissionOutputDTO, error) {
	submission := models.Submission{
		StudentID:    input.StudentID,
		AssignmentID: input.AssignmentID,
		Grade:        input.Grade,
		Feedback:     input.Feedback,
	}

	if err := db.DB.Create(&submission).Error; err != nil {
		return nil, err
	}

	return &dto.SubmissionOutputDTO{
		ID:           submission.ID,
		StudentID:    submission.StudentID,
		AssignmentID: submission.AssignmentID,
		Grade:        submission.Grade,
		Feedback:     submission.Feedback,
	}, nil
}

func (s *SubmissionService) Update(id uint, input *dto.SubmissionInputDTO) (*dto.SubmissionOutputDTO, error) {
	var submission models.Submission

	if err := db.DB.First(&submission, id).Error; err != nil {
		return nil, errors.New("submission not found")
	}

	submission.StudentID = input.StudentID
	submission.AssignmentID = input.AssignmentID
	submission.Grade = input.Grade
	submission.Feedback = input.Feedback

	if err := db.DB.Save(&submission).Error; err != nil {
		return nil, err
	}

	return &dto.SubmissionOutputDTO{
		ID:           submission.ID,
		StudentID:    submission.StudentID,
		AssignmentID: submission.AssignmentID,
		Grade:        submission.Grade,
		Feedback:     submission.Feedback,
	}, nil
}

func (s *SubmissionService) Delete(id uint) error {
	return db.DB.Delete(&models.Submission{}, id).Error
}
