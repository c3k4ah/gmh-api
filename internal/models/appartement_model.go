package models

import (
	"time"
)

type Appartement struct {
	ID            uint         `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     *time.Time   `json:"deleted_at" gorm:"index"`
	Titre         string       `json:"titre" gorm:"type:varchar(255)"`
	Description   string       `json:"description" gorm:"type:varchar(255)"`
	PrixLoyer     string       `json:"prix_loyer" gorm:"type:Double"`
	NombreChambre string       `json:"nombre_chambre" gorm:"type:Double"`
	Localisation  Localisation `json:"localisation" gorm:"foreignKey:LocalisationID"`
}
