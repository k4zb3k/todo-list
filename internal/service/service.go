package service

import (
	"github.com/k4zb3k/todo-list/internal/repository"

	"github.com/k4zb3k/todo-list/pkg/logging"
)

var logger = logging.GetLogger()

type Services struct {
	Repository *repository.Repository
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep}
}
//======================================================