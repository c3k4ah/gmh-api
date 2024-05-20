package models

import (
	"gorm.io/gorm"
)

type Appartement struct {
	gorm.Model
	Titre         string       `json:"titre" gorm:"type:varchar(255)"`
	Description   string       `json:"description" gorm:"type:varchar(255)"`
	PrixLoyer     string       `json:"prix_loyer" gorm:"type:Double"`
	NombreChambre string       `json:"nombre_chambre" gorm:"type:Double"`
	Localisation  Localisation `json:"localisation" gorm:"foreignKey:LocalisationID"`
}
