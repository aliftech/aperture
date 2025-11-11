package mysql

import (
	"aperture/internal/modules/admin/core/domain"
	"aperture/internal/modules/admin/core/port"
	"aperture/pkg/db"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(dbConn *db.Connection) port.AdminRepoMySQL {
	return &adminRepository{db: dbConn.DB} // ✅ pointer
}

func (a *adminRepository) FindAllAdmin() ([]domain.AdminEntity, error) {
	var admins []domain.AdminEntity
	if err := a.db.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (a *adminRepository) DetailAdmin(id uint) (*domain.AdminEntity, error) {
	var admin *domain.AdminEntity
	if err := a.db.Where("id = ?", id).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (a *adminRepository) CreateNewAdmin(admin *domain.AdminEntity) error {
	return a.db.Create(admin).Error
}

func (a *adminRepository) UpdateAdmin(id uint, admin *domain.AdminEntity) error {
	result := a.db.Model(&domain.AdminEntity{}).
		Where("id = ?", id).
		Updates(domain.AdminEntity{
			Firstname: admin.Firstname,
			Lastname:  admin.Lastname,
			Email:     admin.Email,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (a *adminRepository) DeleteAdmin(id uint) error {
	result := a.db.Where("id = ?", id).Delete(&domain.AdminEntity{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
