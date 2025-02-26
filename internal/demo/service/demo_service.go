package service

import (
	"golang_sample/internal/demo/model"
	"golang_sample/internal/demo/repository"
)

type DemoService interface {
	CreateDemo(demo *model.Demo) error
	GetDemo(id uint) (*model.Demo, error)
	GetAllDemos() ([]model.Demo, error)
	UpdateDemo(demo *model.Demo) error
	DeleteDemo(id uint) error
}

type demoService struct {
	repo repository.DemoRepository
}

func NewDemoService(r repository.DemoRepository) DemoService {
	return &demoService{repo: r}
}

func (s *demoService) CreateDemo(demo *model.Demo) error {
	return s.repo.Create(demo)
}

func (s *demoService) GetDemo(id uint) (*model.Demo, error) {
	return s.repo.GetByID(id)
}

func (s *demoService) GetAllDemos() ([]model.Demo, error) {
	return s.repo.GetAll()
}

func (s *demoService) UpdateDemo(demo *model.Demo) error {
	return s.repo.Update(demo)
}

func (s *demoService) DeleteDemo(id uint) error {
	return s.repo.Delete(id)
}
