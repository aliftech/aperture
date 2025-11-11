package domain

import (
	"aperture/pkg/models"
	"time"
)

type AdminEntity struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Status    models.AdminStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
