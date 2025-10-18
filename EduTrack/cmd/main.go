// 1) подключаем базу данных
// 2) настройка маршрутов
// 3) запуск сервера

package main

import (
	"edutrack/internal/db"
	"edutrack/internal/models"
	"fmt"
	"log"
)

func main() {
	db.Init()

	err := db.DB.AutoMigrate(&models.Student{}, &models.Teacher{}, &models.Analytics{}, &models.Submission{}, &models.Assignment{})

	if err != nil {
		log.Fatal("Ошибка миграции: ", err)
	}

	fmt.Println("Все таблицы успешно созданы")
}
