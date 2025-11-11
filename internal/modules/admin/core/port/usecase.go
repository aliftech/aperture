package port

import (
	"aperture/internal/modules/admin/adapter/web/dto"
	"aperture/internal/modules/admin/core/domain"
)

type AdminUsecase interface {
	GetAllAdmin() ([]domain.AdminEntity, error)
	GetAdminByID(id uint) (*domain.AdminEntity, error)
	CreateAdmin(req dto.AdminRequest) (*domain.AdminEntity, error)
	UpdateAdmin(id uint, req dto.AdminRequest) (*domain.AdminEntity, error)
	DeleteAdmin(id uint) (*domain.AdminEntity, error)
}
