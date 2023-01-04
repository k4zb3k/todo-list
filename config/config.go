package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/k4zb3k/todo-list/internal/models"
	"github.com/k4zb3k/todo-list/pkg/logging"
)

var logger = logging.GetLogger()

func GetConfig() (*models.Config, error) {
	file, err := os.Open("/home/k4zb3k/Desktop/todo/config/config.json")
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var config models.Config

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &config, err
}