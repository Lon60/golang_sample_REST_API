package repository

import (
	"errors"

	"golang_sample/internal/demo/model"
	"gorm.io/gorm"
)

type DemoRepository interface {
	Create(demo *model.Demo) error
	GetByID(id uint) (*model.Demo, error)
	GetAll() ([]model.Demo, error)
	Update(demo *model.Demo) error
	Delete(id uint) error
}

type demoRepository struct {
	db *gorm.DB
}

func NewDemoRepository(db *gorm.DB) DemoRepository {
	return &demoRepository{db: db}
}

func (r *demoRepository) Create(demo *model.Demo) error {
	return r.db.Create(demo).Error
}

func (r *demoRepository) GetByID(id uint) (*model.Demo, error) {
	var demo model.Demo
	result := r.db.First(&demo, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &demo, result.Error
}

func (r *demoRepository) GetAll() ([]model.Demo, error) {
	var demos []model.Demo
	result := r.db.Find(&demos)
	return demos, result.Error
}

func (r *demoRepository) Update(demo *model.Demo) error {
	return r.db.Save(demo).Error
}

func (r *demoRepository) Delete(id uint) error {
	return r.db.Delete(&model.Demo{}, id).Error
}
