package repository

import (
	"log"

	"github.com/dev-oleksandrv/db"
	"gorm.io/gorm"
)

type GORMSpaceRepository struct {
	db *gorm.DB
}

func NewGORMSpaceRepository(db *gorm.DB) *GORMSpaceRepository {
	return &GORMSpaceRepository{db}
}

func (r *GORMSpaceRepository) Create(userID int,space db.Space) (db.Space, error) {
	var user db.User
	result := r.db.Create(&space)
	if result.Error != nil {
		log.Printf("Error creating space: %v", result.Error)
		return space, result.Error
	}
	if err := r.db.First(&user, userID).Association("Spaces").Append(&space); err != nil {
		log.Printf("Error connection space to user: %v", err)
		return space, err
	}
	return space, nil
}

func (r *GORMSpaceRepository) GetAllByUserID(userID int) ([]db.Space, error) {
	var spaces []db.Space
	var user db.User
	if err := r.db.Preload("Spaces").First(&user, userID).Error; err != nil {
		log.Printf("Error while preloading spaces: %v", err)
		return spaces, err
	}
	return user.Spaces, nil
}

func (r *GORMSpaceRepository) Delete(id uint) error {
	var space db.Space
	if err := r.db.Find(&space, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Space with ID %d not found", id)
			return err
		}
		log.Printf("Error fetching space: %v", err)
		return err
	}
	if err := r.db.Delete(&space).Error; err != nil {
		log.Printf("Error deleting space with ID %d: %v", id, err)
		return err
	}
	log.Printf("Space with ID %d deleted successfully", id)
	return nil
}