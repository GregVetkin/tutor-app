package json

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"tutor-app/internal/domain"
)

type JSONStudentPool struct {
	filePath string
	mu       sync.Mutex
}

func NewJSONStudentPool(filePath string) *JSONStudentPool {
	return &JSONStudentPool{filePath: filePath}
}

func (p *JSONStudentPool) Create(name string) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	students, err := p.readAll()
	if err != nil {
		return 0, err
	}

	// Генерация нового ID
	maxID := 0
	for _, s := range students {
		if s.ID > maxID {
			maxID = s.ID
		}
	}
	newID := maxID + 1

	student := domain.Student{ID: newID, Name: name}
	students = append(students, student)

	if err := p.writeAll(students); err != nil {
		return 0, err
	}

	return newID, nil
}

func (p *JSONStudentPool) Get(id int) (domain.Student, error) {
	students, err := p.readAll()
	if err != nil {
		return domain.Student{}, err
	}
	for _, s := range students {
		if s.ID == id {
			return s, nil
		}
	}
	return domain.Student{}, errors.New("student not found")
}

func (p *JSONStudentPool) Update(updated domain.Student) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	students, err := p.readAll()
	if err != nil {
		return err
	}

	found := false
	for i, s := range students {
		if s.ID == updated.ID {
			students[i] = updated
			found = true
			break
		}
	}
	if !found {
		return errors.New("student not found")
	}

	return p.writeAll(students)
}

func (p *JSONStudentPool) Delete(id int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	students, err := p.readAll()
	if err != nil {
		return err
	}

	newList := make([]domain.Student, 0, len(students))
	found := false
	for _, s := range students {
		if s.ID == id {
			found = true
			continue
		}
		newList = append(newList, s)
	}
	if !found {
		return errors.New("student not found")
	}

	return p.writeAll(newList)
}

func (p *JSONStudentPool) List() ([]domain.Student, error) {
	return p.readAll()
}

// Вспомогательные методы

func (p *JSONStudentPool) readAll() ([]domain.Student, error) {
	data, err := os.ReadFile(p.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []domain.Student{}, nil
		}
		return nil, err
	}

	var students []domain.Student
	if err := json.Unmarshal(data, &students); err != nil {
		return nil, err
	}
	return students, nil
}

func (p *JSONStudentPool) writeAll(students []domain.Student) error {
	data, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p.filePath, data, 0644)
}
