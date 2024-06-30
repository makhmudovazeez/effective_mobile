package handlers

import (
	"github.com/makhmudovazeez/effective_mobile/internal/logics"
	"gorm.io/gorm"
)

type (
	TimeTrackerHandler interface {
	}

	customTimeTrackerHandler struct {
		db    *gorm.DB
		logic logics.TimeTrackerLogic
	}
)

func NewTimeTrackerHandler(db *gorm.DB) TimeTrackerHandler {
	return &customTimeTrackerHandler{
		db:    db,
		logic: logics.NewTimeTrackerLogic(db),
	}
}
