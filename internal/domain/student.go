package domain

type Student struct {
	ID   int
	Name string
}

type StudentPool interface {
	Create(name string) (int, error)
	Get(id int) (Student, error)
	Update(student Student) error
	Delete(id int) error
	List() ([]Student, error)
}
