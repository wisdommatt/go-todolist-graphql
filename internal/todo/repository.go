package todo

import "gorm.io/gorm"

// Repository is the interface that describes a todo repository object.
type Repository interface {
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
