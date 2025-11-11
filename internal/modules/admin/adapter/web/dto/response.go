package dto

import (
	"aperture/pkg/models"
	"time"
)

type AdminResponse struct {
	ID        uint               `json:"id"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
	Email     string             `json:"email"`
	Status    models.AdminStatus `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
