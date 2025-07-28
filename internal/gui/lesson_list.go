package gui

import (
	"fmt"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"tutor-app/internal/domain"
	"tutor-app/internal/service"
)

type LessonListWidget struct {
	Widget       fyne.CanvasObject
	lessons      []domain.Lesson
	list         *widget.List
	service      *service.LessonService
	currentStudentID int
}

func NewLessonListWidget(service *service.LessonService) *LessonListWidget {
	w := &LessonListWidget{
		service: service,
	}

	// Список
	w.list = widget.NewList(
		func() int { return len(w.lessons) },
		func() fyne.CanvasObject {
			return widget.NewLabel("") // шаблон элемента
		},
		func(i int, o fyne.CanvasObject) {
			lesson := w.lessons[i]
			dateStr := lesson.Date.Format("02.01.2006")
			o.(*widget.Label).SetText(fmt.Sprintf("%s — %s — %.2f₽", dateStr, lesson.Topic, lesson.Price))
		},
	)

	border := container.NewBorder(widget.NewLabel("Уроки"), nil, nil, nil, w.list)
	w.Widget = border

	return w
}

// Метод для загрузки уроков по студенту
func (w *LessonListWidget) LoadForStudent(studentID int) {
	w.currentStudentID = studentID
	lessons, err := w.service.GetLessonsForStudent(studentID)
	if err != nil {
		w.lessons = []domain.Lesson{}
	} else {
		// Отсортируем по дате (по возрастанию)
		sort.Slice(lessons, func(i, j int) bool {
			return lessons[i].Date.Before(lessons[j].Date)
		})
		w.lessons = lessons
	}
	w.list.Refresh()
}
