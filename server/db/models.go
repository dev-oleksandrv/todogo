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
	SpaceID     int       	`gorm:"not null"`                 
    Space       Space      	`gorm:"constraint:OnDelete:CASCADE;" json:"-"` 
	CreatedAt   time.Time	`gorm:"autoCreateTime"`
	UpdatedAt 	time.Time	`gorm:"autoUpdateTime"`
}

type Space struct {
	ID    		int     	`gorm:"primaryKey;autoIncrement"`
	Title 		string  	`gorm:"size:255;not null"`       
    Lists 		[]List  	`gorm:"foreignKey:SpaceID"`  
	Users		[]User		`gorm:"many2many:user_spaces"`
	CreatedAt   time.Time	`gorm:"autoCreateTime"`
	UpdatedAt 	time.Time	`gorm:"autoUpdateTime"`
}

type User struct {
	ID			int 		`gorm:"primaryKey;autoIncrement"`
	Email		string		`gorm:"unique;not null"`
	Password	string		`gorm:"not null" json:"-"`
	Spaces		[]Space		`gorm:"many2many:user_spaces"`
}