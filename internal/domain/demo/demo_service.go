package demo

import (
	"golang_sample/internal/abstract"
)

type Service interface {
	CreateDemo(demo *Demo) error
	GetDemo(id uint) (*Demo, error)
	GetAllDemos() ([]Demo, error)
	UpdateDemo(demo *Demo) error
	DeleteDemo(id uint) error
}

type demoService struct {
	*abstract.Service[Demo]
}

func NewDemoService(repo *abstract.Repository[Demo]) Service {
	return &demoService{
		Service: abstract.NewService(repo),
	}
}

func (s *demoService) CreateDemo(demo *Demo) error {
	return s.Create(demo)
}

func (s *demoService) GetDemo(id uint) (*Demo, error) {
	return s.GetByID(id)
}

func (s *demoService) GetAllDemos() ([]Demo, error) {
	return s.GetAll()
}

func (s *demoService) UpdateDemo(demo *Demo) error {
	return s.Update(demo)
}

func (s *demoService) DeleteDemo(id uint) error {
	return s.Delete(id)
}
