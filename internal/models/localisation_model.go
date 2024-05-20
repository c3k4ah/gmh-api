package models

type Localisation struct {
	// gorm.Model
	LocalisationID uint   `json:"localisation_id" gorm:"primaryKey"`
	Adresse        string `json:"adresse" gorm:"type:varchar(255)"`
	DetailAdresse  string `json:"detail_adresse" gorm:"type:varchar(255)"`
	Ville          string `json:"ville" gorm:"type:varchar(255)"`
	Longitude      string `json:"longitude" gorm:"type:Double"`
	Latitude       string `json:"latitude" gorm:"type:Double"`
}
