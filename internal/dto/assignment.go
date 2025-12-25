package dto

type AssignmentInputDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"omitempty,max=1000"`
	Deadline    string `json:"deadline" validate:"required"`
	TeacherID   uint   `json:"teacher_id" validate:"required"`
}

type AssignmentOutputDTO struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	TeacherID   uint   `json:"teacher_id"`
}
