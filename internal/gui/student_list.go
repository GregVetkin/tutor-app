package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"tutor-app/internal/domain"
	"tutor-app/internal/service"
)

type StudentListWidget struct {
	Widget      fyne.CanvasObject
	students    []domain.Student
	list        *widget.List
	onSelect    func(student domain.Student)
	service     *service.StudentService
}

func NewStudentListWidget(studentService *service.StudentService, onSelect func(student domain.Student)) *StudentListWidget {
	w := &StudentListWidget{
		service:  studentService,
		onSelect: onSelect,
	}

	w.loadStudents()

	// Виджет-список
	w.list = widget.NewList(
		func() int {
			return len(w.students)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(w.students[i].Name)
		},
	)

	// Обработка выбора
	w.list.OnSelected = func(id widget.ListItemID) {
		selected := w.students[id]
		if w.onSelect != nil {
			w.onSelect(selected)
		}
	}

	w.Widget = container.NewBorder(
		widget.NewLabel("Ученики"), nil, nil, nil, w.list,
	)

	return w
}

func (w *StudentListWidget) loadStudents() {
	students, err := w.service.ListStudents()
	if err != nil {
		// В реальном приложении — показать ошибку
		w.students = []domain.Student{}
	} else {
		w.students = students
	}
}

// Позволяет перезагрузить студентов (например, после добавления)
func (w *StudentListWidget) Refresh() {
	w.loadStudents()
	w.list.Refresh()
}
