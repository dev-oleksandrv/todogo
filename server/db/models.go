package db

import "time"

type Task struct {
	ID          int 		`gorm:"primaryKey;autoIncrement"`
	Content     string  	`gorm:"size:255;not null"`
	Description string		`gorm:"type:text"`
	IsChecked 	bool 		`gorm:"default:false"`
	CreatedAt   time.Time	`gorm:"autoCreateTime"`
	UpdatedAt 	time.Time	`gorm:"autoUpdateTime"`
}