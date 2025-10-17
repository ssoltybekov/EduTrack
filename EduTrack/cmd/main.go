// 1) подключаем базу данных
// 2) настройка маршрутов
// 3) запуск сервера

package main

import (
	"edutrack/internal/db"
	"edutrack/internal/models"
)

func main() {
	db.Init()

	db.DB.AutoMigrate(&models.Student{})

	
}