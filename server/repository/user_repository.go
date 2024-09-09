package repository

import (
	"log"

	"github.com/dev-oleksandrv/db"
	"gorm.io/gorm"
)

type GORMUserRepository struct {
	db *gorm.DB
}

func NewGORMUserRepository(db *gorm.DB) *GORMUserRepository {
	return &GORMUserRepository{db}
}

func (r *GORMUserRepository) Create(user db.User) (db.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		log.Printf("Error creating list: %v", result.Error)
		return user, result.Error
	}
	return user, nil
}

func (r *GORMUserRepository) GetUserByEmail(email string) db.User {
	var user db.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error == nil {
		log.Printf("User with email %v was not found", email)
	}
	return user
}

func (r *GORMUserRepository) GetUserByID(userID int) *db.User {
	var user db.User
	result := r.db.First(user, userID)
	if result.Error == nil {
		log.Printf("User with id %v was not found", userID)
		return nil
	}
	return &user
}
