package repository

import (
	"database/sql"

	"github.com/nathanielhall/cloud-native-go/model"
	"github.com/nathanielhall/cloud-native-go/util/logger"
)

// repo := TodoRepository.New(db, logger)

// todos, err := repo.GetAll()
// todo, err := repo.Get(id)
// err := repo.Save(todo)

type todoRepository struct {
	db     *sql.DB
	lr *logger.Logger
}

func NewTodoRepository(db *sql.DB, logger *logger.Logger) *todoRepository {
	return &todoRepository{
		db:     db,
		lr: logger,
	}
}

func (repo *todoRepository) GetAll() ([]model.Todo, error) {
	var todo model.Todo
	var todos []model.Todo

	rows, err := repo.db.Query("select * from todos")
	if err != nil {
		repo.lr.Fatal().Err(err).Msg("Error TodoRepository.GetAll")
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Priority, &todo.Status); err != nil {
			repo.lr.Fatal().Err(err).Msg("Error when trying to get Todo by ID")
			return nil, err
		}
		repo.lr.Info().Msg("log rows next")
	}

	return todos, nil
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
