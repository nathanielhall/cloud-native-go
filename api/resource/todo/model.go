package todo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority 	int `json:"priority"`
	Status 		string `json:"status"`
}
type Todo struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Name 		  string 
	Description   string 
	Priority      int    
	Status        string 
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type Todos []*Todo

func (t *Todo) ToDto() *DTO {
	return &DTO{
		ID:  t.ID.String(),
		Name: t.Name,
		Description: t.Description,
		Priority: t.Priority,
		Status: t.Status,
	}
}

func (todos Todos) ToDto() []*DTO {
	dtos := make([]*DTO, len(todos))
	for i, v := range todos {
		dtos[i] = v.ToDto()
	}
	return dtos
}