package handlers

import (
	"github.com/makhmudovazeez/effective_mobile/internal/logics"
	"gorm.io/gorm"
)

type (
	UserHandler interface {
	}

	customUserHandler struct {
		db    *gorm.DB
		logic logics.UserLogic
	}
)

func NewUserHandler(db *gorm.DB) UserHandler {
	return &customUserHandler{
		db:    db,
		logic: logics.NewUserLogic(db),
	}
}
