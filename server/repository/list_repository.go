package repository

import (
	"log"

	"github.com/dev-oleksandrv/db"
	"gorm.io/gorm"
)

type GORMListRepository struct {
	db *gorm.DB
}

func NewGORMListRepository(db *gorm.DB) *GORMListRepository {
	return &GORMListRepository{db}
}

func (r *GORMListRepository) Create(list db.List) (db.List, error) {
	result := r.db.Create(&list)
	if result.Error != nil {
		log.Printf("Error creating list: %v", result.Error)
		return list, result.Error
	}
	return list, nil
}

func (r *GORMListRepository) GetAll() ([]db.List, error) {
	var lists []db.List
	result := r.db.Find(&lists)
	if result.Error != nil {
		log.Printf("Error creating list: %v", result.Error)
		return nil, result.Error
	}
	return lists, nil
}

func (r *GORMListRepository) Delete(id uint) error {
	var list db.List
	if err := r.db.Find(&list, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("List with ID %d not found", id)
			return err
		}
		log.Printf("Error fetching list: %v", err)
		return err
	}
	if err := r.db.Delete(&list).Error; err != nil {
		log.Printf("Error deleting list with ID %d: %v", id, err)
		return err
	}
	log.Printf("List with ID %d deleted successfully", id)
	return nil
}