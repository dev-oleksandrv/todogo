package db

import "time"

type Task struct {
	ID          int 		`gorm:"primaryKey;autoIncrement"`
	Content     string  	`gorm:"size:255;not null"`
	Description string		`gorm:"type:text"`
	IsChecked 	bool 		`gorm:"default:false"`
	Position	int			`gorm:"default:0;not null"`
	ListID      int       	`gorm:"not null"`                 
    List        List      	`gorm:"constraint:OnDelete:CASCADE;" json:"-"` 
	CreatedAt   time.Time	`gorm:"autoCreateTime"`
	UpdatedAt 	time.Time	`gorm:"autoUpdateTime"`
}

type List struct {
	ID    		int     	`gorm:"primaryKey;autoIncrement"`
    Title 		string  	`gorm:"size:255;not null"` 
    Emoji 		*string 	`gorm:"size:10"`                  
    Tasks 		[]Task  	`gorm:"foreignKey:ListID"`    
	CreatedAt   time.Time	`gorm:"autoCreateTime"`
	UpdatedAt 	time.Time	`gorm:"autoUpdateTime"`
}