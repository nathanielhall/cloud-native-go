package repository

import (
	"github.com/nathanielhall/cloud-native-go/model"
	"github.com/nathanielhall/cloud-native-go/util/logger"
	"gorm.io/gorm"
)

// repo := TodoRepository.New(db, logger)

// todos, err := repo.GetAll()
// todo, err := repo.Get(id)
// err := repo.Save(todo)

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

func (repo *todoRepository) GetAll() ([]model.Todo, error) {


	repo.lr.Debug().Msg("GetAll")

	var todos []model.Todo
	result := repo.db.Find(&todos)

	repo.lr.Info().Msgf("Starting server %v", result.RowsAffected)
	return nil, nil
}

// func (repo TodoRepository) Get(id int32) (model.Todo, error) {
// 	var t model.Todo

// 	err = db.QueryRow("select * from todos where id = ?", id).Scan(&t)

// 	defer stmt.Close()

// 	if err != nil {
// 		logger.Fatal().Err().Msg("Error TodoRepository.Get")
// 	}

// 	return t, nil
// }
