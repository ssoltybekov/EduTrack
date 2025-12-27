package dto

type AssignmentInputDTO struct {
    Title       string `json:"title" validate:"required,min=5,max=200"`
    Description string `json:"description" validate:"omitempty,max=1000"`
    Deadline    string `json:"deadline" validate:"required"` 
}

// Output для ответа клиенту
type AssignmentOutputDTO struct {
    ID          uint   `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Deadline    string `json:"deadline"` 
    LessonID    uint   `json:"lesson_id"`
}