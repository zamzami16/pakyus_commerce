package entity

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	SellerId    uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"type:numeric;not null;default:0"`
	Stock       int       `gorm:"type:int;not null;default:0"`
	CategoryId  uuid.UUID `gorm:"type:uuid"`
	Sku         string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Weight      float64   `gorm:"type:numeric;not null;default:0"`
	Dimensions  string    `gorm:"type:varchar(100)"`
	Condition   string    `gorm:"type:varchar(50);not null"`
	IsActive    bool      `gorm:"type:boolean;not null;default:true"`
	CreatedAt   int64     `gorm:"autoCreateTime:milli"`
	UpdatedAt   int64     `gorm:"autoUpdateTime:milli"`

	Seller   User     `gorm:"foreignKey:SellerId;references:ID"`
	Category Category `gorm:"foreignKey:CategoryId;references:ID"`
}

func (p *Product) TableName() string {
	return "products"
}
