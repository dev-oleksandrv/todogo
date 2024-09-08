package space

import (
	"github.com/dev-oleksandrv/db"
	"github.com/dev-oleksandrv/repository"
)

type SpaceService struct {
	spaceRepo repository.GORMSpaceRepository
}

func NewSpaceService(spaceRepo repository.GORMSpaceRepository) *SpaceService {
	return &SpaceService{ spaceRepo }
}

func (s *SpaceService) CreateSpace(space db.Space) (db.Space, error) {
	return s.spaceRepo.Create(space)
}

func (s *SpaceService) GetAllSpaces() ([]db.Space, error) {
	return s.spaceRepo.GetAll()
}

func (s *SpaceService) DeleteSpace(id uint) error {
	return s.spaceRepo.Delete(id)
}