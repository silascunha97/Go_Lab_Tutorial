package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID    string  `json:"id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
