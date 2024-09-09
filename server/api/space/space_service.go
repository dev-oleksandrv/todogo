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

func (s *SpaceService) CreateSpace(userID int, space db.Space) (db.Space, error) {
	return s.spaceRepo.Create(userID, space)
}

func (s *SpaceService) GetAllSpacesByUserID(userID int) ([]db.Space, error) {
	return s.spaceRepo.GetAllByUserID(userID) 
}

func (s *SpaceService) DeleteSpace(id uint) error {
	return s.spaceRepo.Delete(id)
}