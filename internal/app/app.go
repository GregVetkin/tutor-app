package app

import (
	"tutor-app/internal/repository/json"
	"tutor-app/internal/service"
)

type App struct {
	StudentService *service.StudentService
	LessonService  *service.LessonService
}

func NewApp() *App {
	// Пути к JSON-файлам
	studentPath := "data/students.json"
	lessonPath := "data/lessons.json"

	// Пулы (хранилища)
	studentPool := json.NewJSONStudentPool(studentPath)
	lessonPool := json.NewJSONLessonPool(lessonPath)

	// Сервисы
	studentService := service.NewStudentService(studentPool)
	lessonService := service.NewLessonService(lessonPool, studentPool)

	return &App{
		StudentService: studentService,
		LessonService:  lessonService,
	}
}
