package json

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"
	"tutor-app/internal/domain"
)

type JSONLessonPool struct {
	filePath string
	mu       sync.Mutex
}

func NewJSONLessonPool(filePath string) *JSONLessonPool {
	return &JSONLessonPool{filePath: filePath}
}

func (p *JSONLessonPool) Create(studentID int, date time.Time, topic string, price float64, notes string) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	lessons, err := p.readAll()
	if err != nil {
		return 0, err
	}

	maxID := 0
	for _, l := range lessons {
		if l.ID > maxID {
			maxID = l.ID
		}
	}
	newID := maxID + 1

	lesson := domain.Lesson{
		ID:        newID,
		StudentID: studentID,
		Date:      date,
		Topic:     topic,
		Price:     price,
		IsPaid:    false,
		Notes:     notes,
	}

	lessons = append(lessons, lesson)

	if err := p.writeAll(lessons); err != nil {
		return 0, err
	}

	return newID, nil
}

func (p *JSONLessonPool) Get(id int) (domain.Lesson, error) {
	lessons, err := p.readAll()
	if err != nil {
		return domain.Lesson{}, err
	}
	for _, l := range lessons {
		if l.ID == id {
			return l, nil
		}
	}
	return domain.Lesson{}, errors.New("lesson not found")
}

func (p *JSONLessonPool) Update(updated domain.Lesson) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	lessons, err := p.readAll()
	if err != nil {
		return err
	}

	found := false
	for i, l := range lessons {
		if l.ID == updated.ID {
			lessons[i] = updated
			found = true
			break
		}
	}
	if !found {
		return errors.New("lesson not found")
	}

	return p.writeAll(lessons)
}

func (p *JSONLessonPool) Delete(id int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	lessons, err := p.readAll()
	if err != nil {
		return err
	}

	newList := make([]domain.Lesson, 0, len(lessons))
	found := false
	for _, l := range lessons {
		if l.ID == id {
			found = true
			continue
		}
		newList = append(newList, l)
	}
	if !found {
		return errors.New("lesson not found")
	}

	return p.writeAll(newList)
}

func (p *JSONLessonPool) List() ([]domain.Lesson, error) {
	return p.readAll()
}

func (p *JSONLessonPool) ListByStudent(studentID int) ([]domain.Lesson, error) {
	lessons, err := p.readAll()
	if err != nil {
		return nil, err
	}

	var result []domain.Lesson
	for _, l := range lessons {
		if l.StudentID == studentID {
			result = append(result, l)
		}
	}
	return result, nil
}

// Вспомогательные методы

func (p *JSONLessonPool) readAll() ([]domain.Lesson, error) {
	data, err := os.ReadFile(p.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []domain.Lesson{}, nil
		}
		return nil, err
	}

	var lessons []domain.Lesson
	if err := json.Unmarshal(data, &lessons); err != nil {
		return nil, err
	}
	return lessons, nil
}

func (p *JSONLessonPool) writeAll(lessons []domain.Lesson) error {
	data, err := json.MarshalIndent(lessons, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p.filePath, data, 0644)
}
