package demo

import (
	"golang_sample/internal/abstract"
	"gorm.io/gorm"
)

type Repository struct {
	*abstract.Repository[Demo]
}

func NewDemoRepository(db *gorm.DB) *Repository {
	return &Repository{
		Repository: abstract.NewRepository[Demo](db),
	}
}
