package logics

import "gorm.io/gorm"

type (
	UserLogic interface {
	}

	customUserLogic struct {
		db *gorm.DB
	}
)

func NewUserLogic(db *gorm.DB) UserLogic {
	return &customUserLogic{
		db: db,
	}
}
