package entity

import "github.com/google/uuid"

type Address struct {
	ID         uuid.UUID `gorm:"column:id;primaryKey;type:uuid"`
	ContactId  uuid.UUID `gorm:"column:contact_id;type:uuid"`
	Street     string    `gorm:"column:street"`
	City       string    `gorm:"column:city"`
	Province   string    `gorm:"column:province"`
	PostalCode string    `gorm:"column:postal_code"`
	Country    string    `gorm:"column:country"`
	CreatedAt  int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt  int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	Contact    Contact   `gorm:"foreignKey:contact_id;references:id"`
}

func (a *Address) TableName() string {
	return "addresses"
}
