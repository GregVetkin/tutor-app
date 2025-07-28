package gui

import (
	"fmt"
	"sort"
	"strconv"
	"tutor-app/internal/app"
	"tutor-app/internal/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowAddLessonForm(appInstance *app.App, parent fyne.Window, onCreated func()) {
	students, err := appInstance.StudentService.ListStudents()
	if err != nil || len(students) == 0 {
		dialog.ShowInformation("Нет учеников", "Пожалуйста, добавьте учеников перед созданием урока.", parent)
		return
	}

	// Сортируем студентов по имени
	sort.Slice(students, func(i, j int) bool {
		return students[i].Name < students[j].Name
	})

	studentNames := make([]string, len(students))
	for i, s := range students {
		studentNames[i] = fmt.Sprintf("%s (ID: %d)", s.Name, s.ID)
	}

	selectedStudent := &students[0]
	studentSelect := widget.NewSelect(studentNames, func(chosen string) {
		for _, s := range students {
			if chosen == fmt.Sprintf("%s (ID: %d)", s.Name, s.ID) {
				selectedStudent = &s
				break
			}
		}
	})
	studentSelect.SetSelected(studentNames[0])

	topicEntry := widget.NewEntry()
	topicEntry.SetPlaceHolder("Тема урока")
	priceEntry := widget.NewEntry()
	priceEntry.SetPlaceHolder("Цена, например 1200")
	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("2025-07-28")
	noteEntry := widget.NewEntry()
	noteEntry.SetPlaceHolder("Примечание")

	topicEntry.Wrapping = fyne.TextWrapOff
	priceEntry.Wrapping = fyne.TextWrapOff
	dateEntry.Wrapping = fyne.TextWrapOff
	noteEntry.Wrapping = fyne.TextWrapOff

	dialog.ShowForm("Новый урок", "Создать", "Отмена",
		[]*widget.FormItem{
			widget.NewFormItem("Ученик", studentSelect),
			widget.NewFormItem("Тема", topicEntry),
			widget.NewFormItem("Цена", priceEntry),
			widget.NewFormItem("Дата (ГГГГ-ММ-ДД)", dateEntry),
			widget.NewFormItem("Примечание", noteEntry),
		},
		func(ok bool) {
			if !ok {
				return
			}

			if selectedStudent == nil {
				dialog.ShowError(fmt.Errorf("Ученик не выбран"), parent)
				return
			}

			date, err := utils.ParseDate(dateEntry.Text)
			if err != nil {
				dialog.ShowError(fmt.Errorf("Неверный формат даты"), parent)
				return
			}

			price, err := strconv.ParseFloat(priceEntry.Text, 64)
			if err != nil {
				dialog.ShowError(fmt.Errorf("Неверный формат цены"), parent)
				return
			}

			_, err = appInstance.LessonService.CreateLesson(selectedStudent.ID, date, topicEntry.Text, price, noteEntry.Text)
			if err != nil {
				dialog.ShowError(err, parent)
				return
			}

			if onCreated != nil {
				onCreated()
			}
		}, parent)
}
