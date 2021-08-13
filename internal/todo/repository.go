package todo

import "gorm.io/gorm"

// Repository is the interface that describes a todo repository object.
type Repository interface {
	Save(todo *Todo) error
	GetTodoById(id int) (todo Todo, err error)
	GetTodos(status string) (todos []Todo, err error)
	DeleteById(id int) error
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
	db := repo.db.Create(todo)
	return db.Error
}

// GetTodoById retrieves a todo from the database by id.
func (repo *TodoRepo) GetTodoById(id int) (todo Todo, err error) {
	db := repo.db.Where("id = ?", id).First(&todo)
	return todo, db.Error
}

// GetTodos retrieves todos from the database by filter.
func (repo *TodoRepo) GetTodos(status string) (todos []Todo, err error) {
	if status != "" {
		db := repo.db.Where("status = ?", status).Find(&todos)
		return todos, db.Error
	}
	db := repo.db.Find(&todos)
	return todos, db.Error
}

// DeleteById deletes a todo from the database by id.
func (repo *TodoRepo) DeleteById(id int) error {
	db := repo.db.Delete(&Todo{ID: id})
	return db.Error
}
