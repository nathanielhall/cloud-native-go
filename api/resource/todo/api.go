package todo

import (
	"github.com/go-playground/validator/v10"
	"github.com/nathanielhall/cloud-native-go/util/logger"
	"gorm.io/gorm"
)

type API struct {
	logger     *logger.Logger
	validator  *validator.Validate
	repository *Repository
}

func New(logger *logger.Logger, validator *validator.Validate, db *gorm.DB) *API {
	return &API{
		logger:     logger,
		validator: validator,
		repository: NewRepository(db),
	}
}