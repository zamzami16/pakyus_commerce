package entity

import "github.com/google/uuid"

type Category struct {
	ID       uuid.UUID  `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string     `gorm:"column:name;uniqueIndex"`
	ParentId *uuid.UUID `gorm:"column:parent_id;type:uuid;default:null"`
	Products []Product  `gorm:"foreignKey:category_id;references:id"`
}

func (c *Category) TableName() string {
	return "categories"
}
