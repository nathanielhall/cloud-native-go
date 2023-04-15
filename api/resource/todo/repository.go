package todo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAll() (Todos, error) {
	todos := make([]*Todo, 0)
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	if len(todos) == 0 {
		return nil, nil
	}
	return todos, nil
}

func (r *Repository) Create(todo *Todo) (*Todo, error) {
	if err := r.db.Create(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *Repository) GetById(id uuid.UUID) (*Todo, error) {
	todo := &Todo{}
	if err := r.db.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *Repository) Update(todo *Todo) error {
	if err := r.db.First(&Todo{}, todo.ID).Updates(todo).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(id uuid.UUID) error {
	todo := &Todo{}
	if err := r.db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}