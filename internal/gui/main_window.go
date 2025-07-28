package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"tutor-app/internal/app"
	"tutor-app/internal/domain"
)

func ShowMainWindow(appInstance *app.App, fyneApp fyne.App) {
	window := fyneApp.NewWindow("Tutor App")
	window.Resize(fyne.NewSize(1000, 600))

	// Состояние выбранного ученика
	var selectedStudent *domain.Student

	// Виджет списка уроков
	lessonList := NewLessonListWidget(appInstance.LessonService)

	// Виджет списка студентов
	studentList := NewStudentListWidget(appInstance.StudentService, func(student domain.Student) {
		selectedStudent = &student
		lessonList.LoadForStudent(student.ID)
	})

	// Кнопка "Добавить ученика"
	addStudentButton := widget.NewButton("Добавить ученика", func() {
		nameEntry := widget.NewEntry()
		nameEntry.SetPlaceHolder("Имя ученика")

		dialog.ShowForm("Новый ученик",
			"Создать", "Отмена",
			[]*widget.FormItem{
				widget.NewFormItem("Имя", nameEntry),
			},
			func(ok bool) {
				if ok && nameEntry.Text != "" {
					_, err := appInstance.StudentService.CreateStudent(nameEntry.Text)
					if err == nil {
						studentList.Refresh()
					} else {
						dialog.ShowError(err, window)
					}
				}
			},
			window,
		)
	})

	// Кнопка "Добавить урок"
	addLessonButton := widget.NewButton("Добавить урок", func() {
		ShowAddLessonForm(appInstance, window, func() {
			if selectedStudent != nil {
				lessonList.LoadForStudent(selectedStudent.ID)
			}
		})
	})

	// Верхняя панель
	topBar := container.NewHBox(
		layout.NewSpacer(),
		addLessonButton,
		addStudentButton,
	)

	// Контейнер с двумя панелями: список учеников и уроков
	split := container.NewHSplit(
		studentList.Widget,
		lessonList.Widget,
	)
	split.Offset = 0.3

	// Основное содержимое окна
	mainContent := container.NewBorder(
		topBar, nil, nil, nil,
		split,
	)

	window.SetContent(mainContent)
	window.ShowAndRun()
}
