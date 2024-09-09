package space

import (
	"errors"
	"log"

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

func (s *SpaceService) DeleteSpace(userID, id int) error {
	isAssociated, err := s.spaceRepo.IsUserAssociatedWithSpace(userID, id)
	
	if err != nil {
		log.Printf("Error checking association: %v", err)
		return err
	}

	if !isAssociated {
		log.Printf("User is not authorized to delete this space")
		return errors.New("user is not authorized to delete this space")
	}

	err = s.spaceRepo.Delete(id)
	if err != nil {
		log.Printf("Error deleting space: %v", err)
		return err
	}
	
	return nil
}