package repository

import (
	"pakyus_commerce/internal/entity"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressRepository struct {
	Repository[entity.Address]
	Log *logrus.Logger
}

func NewAddressRepository(log *logrus.Logger) *AddressRepository {
	return &AddressRepository{
		Log: log,
	}
}

func (r *AddressRepository) FindByIdAndContactId(tx *gorm.DB, address *entity.Address, id uuid.UUID, contactId uuid.UUID) error {
	return tx.Where("id = ? AND contact_id = ?", id, contactId).First(address).Error
}

func (r *AddressRepository) FindAllByContactId(tx *gorm.DB, contactId uuid.UUID) ([]entity.Address, error) {
	var addresses []entity.Address
	if err := tx.Where("contact_id = ?", contactId).Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}
