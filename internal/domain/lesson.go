package domain

import "time"

type Lesson struct {
	ID        int
	StudentID int       // связь с учеником
	Date      time.Time
	Topic     string
	Price     float64
	IsPaid    bool
	Notes     string
}


type LessonPool interface {
	Create(studentID int, date time.Time, topic string, price float64, notes string) (int, error)
	Get(id int) (Lesson, error)
	Update(lesson Lesson) error
	Delete(id int) error
	List() ([]Lesson, error)
	ListByStudent(studentID int) ([]Lesson, error)
}
