package repository

import (
	"github.com/k4zb3k/todo-list/pkg/logging"
	"gorm.io/gorm"
)

var logger = logging.GetLogger()

type Repository struct {
	Connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{Connection: conn}
}

//=============================================