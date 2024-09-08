package list

import (
	"github.com/dev-oleksandrv/db"
	"github.com/dev-oleksandrv/repository"
)

type ListService struct {
	listRepo repository.GORMListRepository
}

func NewListService(listRepo repository.GORMListRepository) *ListService {
	return &ListService{ listRepo }
}

func (s *ListService) CreateList(list db.List) (db.List, error) {
	return s.listRepo.Create(list)
}

func (s *ListService) GetAllLists() ([]db.List, error) {
	return s.listRepo.GetAll()
}

func (s *ListService) DeleteList(id uint) error {
	return s.listRepo.Delete(id)
}