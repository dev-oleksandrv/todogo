package repository

import (
	"log"

	"github.com/dev-oleksandrv/db"
	"gorm.io/gorm"
)

type GORMTaskRepository struct {
	db *gorm.DB
}

func NewGORMTaskRepository(db *gorm.DB) *GORMTaskRepository {
	return &GORMTaskRepository{db}
}

func (r *GORMTaskRepository) Create(task db.Task) (db.Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		log.Printf("Error creating task: %v", result.Error)
		return task, result.Error
	}
	return task, nil
}

func (r *GORMTaskRepository) GetAll() ([]db.Task, error) {
	var tasks []db.Task
	result := r.db.Find(&tasks)
	if result.Error != nil {
		log.Printf("Error creating task: %v", result.Error)
		return nil, result.Error
	}
	return tasks, nil
}

func (r *GORMTaskRepository) Delete(id uint) error {
	var task db.Task
	if err := r.db.Find(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Task with ID %d not found", id)
			return err
		}
		log.Printf("Error fetching task: %v", err)
		return err
	}
	if err := r.db.Delete(&task).Error; err != nil {
		log.Printf("Error deleting task with ID %d: %v", id, err)
		return err
	}
	log.Printf("Task with ID %d deleted successfully", id)
	return nil
}