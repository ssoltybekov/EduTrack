package dto

type StudentInputDTO struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Group string `json:"group"`
}

type StudentOutputDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Group string `json:"group"`
}
