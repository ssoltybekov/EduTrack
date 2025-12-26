package dto

type SubmissionInputDTO struct {
	StudentID    uint    `json:"student_id" validate:"required"`
	AssignmentID uint    `json:"assignment_id" validate:"required"`
	Grade        float64 `json:"grade" validate:"omitempty,gte=0,lte=10"`
	Feedback     string  `json:"feedback" validate:"omitempty,max=1000"`
}

type SubmissionOutputDTO struct {
	ID           uint    `json:"id"`
	StudentID    uint    `json:"student_id"`
	AssignmentID uint    `json:"assignment_id"`
	Grade        float64 `json:"grade"`
	Feedback     string  `json:"feedback"`
}


