package store_util

import (
	"fmt"

	"gorm.io/gorm"
)

type GormImpl struct {
	Db *gorm.DB
}

func (u GormImpl) First(field string, value string, entity interface{}) error {
	result := u.Db.First(entity, fmt.Sprintf("%s = ?", field), value)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil
		}
	}

	return result.Error
}

func (u GormImpl) Create(entity interface{}) error {
	result := u.Db.Create(entity)

	return result.Error
}
