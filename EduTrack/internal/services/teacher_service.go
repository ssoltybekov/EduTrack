package services

import (
	"edutrack/internal/db"
	"edutrack/internal/models"
)

// // чтение учителей который в teacher.go отправлю

// // через /teachers Get вызывет эту функцию чтобы все записы из бд получить

// func GetAllTeachers() ([]models.Teacher, error) {
// 	var teachers []models.Teacher
// 	err := db.DB.Find(&teachers).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return teachers, nil
// }

// // теперь через айди находим тичеров чтобы конкретный тьютор был

// func GetTeacherById(id uint) (*models.Teacher, error) {
// 	var teacher models.Teacher
// 	err := db.DB.First(&teacher, id).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &teacher, err
// }

// // вставляем структуру

// func CreateTeacher(teacher *models.Teacher) error {
// 	return db.DB.Create(&teacher).Error
// }

// func UpdateTeacher(teacher *models.Teacher) error {
// 	return db.DB.Save(&teacher).Error
// }

// func DeleteTeacher(id uint) error {
// 	return db.DB.Delete(&models.Teacher{}, id).Error
// }

// теперь то же самое на chi

type TeacherService struct{}

func NewTeacherService() *TeacherService {
	return &TeacherService{}
}

func (s *TeacherService) GetAll() ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := db.DB.Find(&teachers).Error
	if err != nil {
		return nil, err
	}
	return teachers, err
}

func (s *TeacherService) GetById(id uint) (*models.Teacher, error) {
	var teacher *models.Teacher
	err := db.DB.First(&teacher, id).Error
	if err != nil {
		return nil, err
	}
	return teacher, err
}

func (s *TeacherService) Create(teacher *models.Teacher) error {
	return db.DB.Create(&teacher).Error
}

func (s *TeacherService) Update(teacher *models.Teacher) error {
	return db.DB.Save(&teacher).Error
}

func (s *TeacherService) Delete(id uint) error {
	return db.DB.Delete(&models.Teacher{}, id).Error
}
