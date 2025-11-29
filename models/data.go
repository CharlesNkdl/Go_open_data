package models

import "gorm.io/gorm"

type GovData struct {
	gorm.Model
	ExternalID string  `json:"id_externe" gorm:"uniqueIndex"`
	Region     string  `json:"region"`
	Amount     float64 `json:"montant"`
	Label      string  `json:"libelle"`
	Year       int     `json:"annee"`
}
