package service

import (
	"errors"
	"time"
	"tutor-app/internal/domain"
)

type LessonService struct {
	repo       domain.LessonPool
	studentRepo domain.StudentPool
}

func NewLessonService(r domain.LessonPool, sr domain.StudentPool) *LessonService {
	return &LessonService{
		repo:       r,
		studentRepo: sr,
	}
}

func (s *LessonService) CreateLesson(studentID int, date time.Time, topic string, price float64, notes string) (int, error) {
	_, err := s.studentRepo.Get(studentID)
	if err != nil {
		return 0, errors.New("такого ученика не существует")
	}

	if topic == "" {
		return 0, errors.New("тема не может быть пустой")
	}

	if price < 0 {
		return 0, errors.New("некорректная цена")
	}

	return s.repo.Create(studentID, date, topic, price, notes)
}

func (s *LessonService) MarkLessonAsPaid(id int) error {
	lesson, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	lesson.IsPaid = true
	return s.repo.Update(lesson)
}

func (s *LessonService) GetLessonsForStudent(studentID int) ([]domain.Lesson, error) {
	return s.repo.ListByStudent(studentID)
}

