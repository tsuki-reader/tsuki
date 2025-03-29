package models

import (
	"gorm.io/gorm"
)

// BaseModel includes common fields and query methods
type Base struct {
	gorm.Model
}

func (b *Base) Where(conditions map[string]interface{}) *gorm.DB {
	return DATABASE.Where(conditions)
}

func (b *Base) Find(id uint, out interface{}) error {
	return DATABASE.First(out, id).Error
}

func (b *Base) Preload(associations ...string) *gorm.DB {
	query := DATABASE
	for _, association := range associations {
		query = query.Preload(association)
	}
	return query
}
