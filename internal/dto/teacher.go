package dto

type TeacherInputDTO struct {
    Name    string `json:"name" validate:"required"`
    Email   string `json:"email" validate:"required,email"`
    Subject string `json:"subject"`
}

type TeacherOutputDTO struct {
    ID      uint   `json:"id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Subject string `json:"subject"`
}
