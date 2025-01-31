package main

import (
	"time"

	"github.com/google/uuid"
)

type Food struct{
	ID			int			`json:"id"`
	Name		string		`json:"name" validate:"required,min=2,max=100"`
	Price		float64	`json:"price" validate:"required"`
	Food_image 	string		`json:"food_image" validate:"required"`
	Created_at	time.Time	`json:"created_at"`
	Updated_at	time.Time	`json:"updated_at"`
	Food_id		uuid.UUID		`json:"food_id"`
	Menu_id		uuid.UUID		`json:"menu_id" validate:"required"`
}