package services

import (
	"edutrack/internal/db"
	"edutrack/internal/models"
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

func (s *SubmissionService) Create(submission *models.Submission) error {
	return db.DB.Create(submission).Error
}

func (s *SubmissionService) Update(submission *models.Submission) error {
	return db.DB.Save(submission).Error
}

func (s *SubmissionService) Delete(id uint) error {
	return db.DB.Delete(&models.Submission{}, id).Error
}
