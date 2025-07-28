package main

import (
	"fmt"
	"time"
	"tutor-app/internal/app"
)

func main() {
	app := app.NewApp()

	studentID, err := app.StudentService.CreateStudent("Иван Иванов")
	if err != nil {
		panic(err)
	}
	fmt.Println("Создан студент с ID:", studentID)

	lessonID, err := app.LessonService.CreateLesson(
		studentID,
		time.Now(),
		"Разбор Future Perfect",
		1200.0,
		"Практика разговорной части",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Создан урок с ID:", lessonID)

	lessons, err := app.LessonService.GetLessonsForStudent(studentID)
	if err != nil {
		panic(err)
	}
	for _, l := range lessons {
		fmt.Printf("- %s (%s) | %.0f руб. | Оплачено: %v\n",
			l.Topic, l.Date.Format("02.01.2006 15:04"), l.Price, l.IsPaid)
	}
}
