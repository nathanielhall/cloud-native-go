package todo

import (
	"github.com/nathanielhall/cloud-native-go/util/logger"
	"gorm.io/gorm"
)

type API struct {
	logger     *logger.Logger
	repository *Repository
}

func New(logger *logger.Logger, db *gorm.DB) *API {
	return &API{
		logger:     logger,
		repository: NewRepository(db),
	}
}