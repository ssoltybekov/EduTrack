package dto

type LessonInputDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	VideoURL    string `json:"video_url" validate:"required,url"`
}

type LessonOutputDTO struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoURL    string `json:"video_url"`
	TeacherID   uint   `json:"teacher_id"`
}
