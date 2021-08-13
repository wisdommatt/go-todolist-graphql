package todo

import "gorm.io/gorm"

// Repository is the interface that describes a todo repository object.
type Repository interface {
	Save(todo *Todo) error
	GetTodos(status string) (todos []Todo, err error)
}

// TodoRepo is the default implementation for the repository interface.
type TodoRepo struct {
	db *gorm.DB
}

// NewRepository returns a new repository object.
func NewRepository(db *gorm.DB) *TodoRepo {
	return &TodoRepo{
		db: db,
	}
}

// Save saves a todo to the database.
func (repo *TodoRepo) Save(todo *Todo) error {
	repo.db.Create(todo)
	return nil
}

// GetTodos retrieves todos from the database by filter.
func (repo *TodoRepo) GetTodos(status string) (todos []Todo, err error) {
	if status != "" {
		repo.db.Where("status = ?", status).Find(&todos)
		return todos, nil
	}
	repo.db.Find(&todos)
	return todos, nil
}
