package dto


type SubmissionInputDTO struct {
    Content string `json:"content" validate:"required,min=10,max=5000"` 
}

type SubmissionGradeInputDTO struct {
    Grade    float64 `json:"grade" validate:"required,gte=0,lte=10"`
    Feedback string  `json:"feedback,omitempty" validate:"max=2000"`
}

type SubmissionOutputDTO struct {
    ID           uint     `json:"id"`
    Content      string   `json:"content"`
    Grade        *float64 `json:"grade,omitempty"`
    Feedback     string   `json:"feedback,omitempty"`
    SubmittedAt  string   `json:"submitted_at"` 
    StudentID    uint     `json:"student_id"`
    AssignmentID uint     `json:"assignment_id"`
}