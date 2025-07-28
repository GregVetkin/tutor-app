package service

import (
	"errors"
	"tutor-app/internal/domain"
)

type StudentService struct {
	repo domain.StudentPool
}

func NewStudentService(r domain.StudentPool) *StudentService {
	return &StudentService{repo: r}
}

func (s *StudentService) CreateStudent(name string) (int, error) {
	if name == "" {
		return 0, errors.New("имя не может быть пустым")
	}
	return s.repo.Create(name)
}

func (s *StudentService) GetStudent(id int) (domain.Student, error) {
	return s.repo.Get(id)
}

func (s *StudentService) UpdateStudent(id int, newName string) error {
	student, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	student.Name = newName
	return s.repo.Update(student)
}

func (s *StudentService) DeleteStudent(id int) error {
	return s.repo.Delete(id)
}

func (s *StudentService) ListStudents() ([]domain.Student, error) {
	return s.repo.List()
}
