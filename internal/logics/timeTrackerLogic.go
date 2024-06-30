package logics

import "gorm.io/gorm"

type (
	TimeTrackerLogic interface {
	}

	customTimeTrackerLogic struct {
		db *gorm.DB
	}
)

func NewTimeTrackerLogic(db *gorm.DB) TimeTrackerLogic {
	return &customTimeTrackerLogic{
		db: db,
	}
}
