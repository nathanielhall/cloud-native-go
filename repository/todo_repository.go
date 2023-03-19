package repository

import (
	"github.com/nathanielhall/cloud-native-go/model"
	"github.com/nathanielhall/cloud-native-go/util/logger"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
	lr *logger.Logger
}

func NewTodoRepository(db *gorm.DB, logger *logger.Logger) *todoRepository {
	return &todoRepository{
		db: db,
		lr: logger,
	}
}

func (repo *todoRepository) GetAll() (model.Todos, error) {
	todos := make([]*model.Todo, 0)
	if err := repo.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	if len(todos) == 0 {
		return nil, nil
	}

	repo.lr.Info().Msgf("LEN %v", len(todos))

	return todos, nil
}
