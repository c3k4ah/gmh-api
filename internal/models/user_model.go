package models

type User struct {
	// gorm.Model
	UserID    uint   `json:"user_id" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"type:varchar(255)"`
	FirstName string `json:"first_name" gorm:"type:varchar(255)"`
	LastName  string `json:"last_name" gorm:"type:varchar(255)"`
	Pseudo    string `json:"pseudo" gorm:"varchar(255)"`
	Password  string `json:"password" gorm:"varchar(255)"`
}
