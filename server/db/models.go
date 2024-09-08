package db

import "time"

type Task struct {
	ID          int 		`gorm:"primaryKey;autoIncrement"`
	Content     string  	`gorm:"size:255;not null"`
	Description string		`gorm:"type:text"`
	IsChecked 	bool 		`gorm:"default:false"`
	Position	int			`gorm:"default:0;not null"`
	ListID      int       	`gorm:"not null"`                 // Foreign key for List (many-to-one relation)
    List        List      	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Set up the relation
	CreatedAt   time.Time	`gorm:"autoCreateTime"`
	UpdatedAt 	time.Time	`gorm:"autoUpdateTime"`
}

type List struct {
	ID    int     `gorm:"primaryKey;autoIncrement"`
    Title string  `gorm:"size:255;not null"` 
    Emoji *string `gorm:"size:10"`                  
    Tasks []Task  `gorm:"foreignKey:ListID"`    
}