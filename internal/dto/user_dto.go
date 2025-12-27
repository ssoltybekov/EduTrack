package dto

type UserInputDTO struct {
    Name     string `json:"name" validate:"required,min=2"`
    Email    string `json:"email" validate:"required,email"`
    Role     string `json:"role" validate:"required,oneof=teacher student"`
    Subject  string `json:"subject,omitempty"`
    Group    string `json:"group,omitempty"`
}

type UserOutputDTO struct {
    ID      uint   `json:"id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Role    string `json:"role"`
    Subject string `json:"subject,omitempty"`
    Group   string `json:"group,omitempty"`
}