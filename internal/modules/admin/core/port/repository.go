package port

import "aperture/internal/modules/admin/core/domain"

type AdminRepoMySQL interface {
	FindAllAdmin() ([]domain.AdminEntity, error)
	DetailAdmin(id uint) (*domain.AdminEntity, error)
	CreateNewAdmin(admin *domain.AdminEntity) error
	UpdateAdmin(id uint, admin *domain.AdminEntity) error
	DeleteAdmin(id uint) error
}
