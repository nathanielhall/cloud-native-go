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

type Form struct {
	Name 		  string `json:"name" form:"required,max=255"`
	Description   string `json:"description"`
	Priority      int    `json:"priority"`
	Status        string `json:"status"`
}

func (f *Form) ToModel() *Todo {
	return &Todo{
		ID: uuid.New(),
		Name: f.Name,
		Description: f.Description,
		Priority: f.Priority,
		Status: f.Status,
	}
}

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