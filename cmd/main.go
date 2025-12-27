// 1) подключаем базу данных
// 2) настройка маршрутов
// 3) запуск сервера

package main

import (
	"edutrack/internal/db"
	"edutrack/internal/models"
	"edutrack/internal/pkg/validator" 
	"edutrack/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db.Init()

	err := db.DB.AutoMigrate(&models.User{}, &models.Lesson{}, &models.Assignment{}, &models.Submission{}, &models.Analytics{})
	if err != nil {
		log.Fatal("Ошибка миграции: ", err)
	}

	validator.Init() 

	r := routes.Routes()

	fmt.Println("Все таблицы успешно созданы")
	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", r)
}
