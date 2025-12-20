package dto

type SubmissionInputDTO struct {
	StudentID    uint    `json:"student_id"`
	AssignmentID uint    `json:"assignment_id"`
	Grade        float64 `json:"grade"`
	Feedback     string  `json:"feedback"`
}

type SubmissionOutputDTO struct {
	ID           uint    `json:"id"`
	StudentID    uint    `json:"student_id"`
	AssignmentID uint    `json:"assignment_id"`
	Grade        float64 `json:"grade"`
	Feedback     string  `json:"feedback"`
}