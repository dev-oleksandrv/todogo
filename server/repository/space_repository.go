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

func (r *GORMSpaceRepository) Delete(id int) error {
	var space db.Space
	if err := r.db.First(&space, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&space).Error; err != nil {
		return err
	}

	return nil
}

func (r *GORMSpaceRepository) IsUserAssociatedWithSpace(userID, spaceID int) (bool, error) {
	var count int64
	err := r.db.Model(&db.Space{}).Where("id = ?", spaceID).
		Joins("JOIN user_spaces ON spaces.id = user_spaces.space_id").
		Where("user_spaces.user_id = ?", userID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
