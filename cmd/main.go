package main

import (
	"os"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k4zb3k/todo-list/config"
	"github.com/k4zb3k/todo-list/internal/db"
	"github.com/k4zb3k/todo-list/internal/repository"
	"github.com/k4zb3k/todo-list/internal/server"
	"github.com/k4zb3k/todo-list/internal/service"
	"github.com/k4zb3k/todo-list/pkg/logging"
)

var logger = logging.GetLogger()

func main() {
	err := execute()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

func execute() error {
	router := mux.NewRouter()

	connection, err := db.GetDbConnection()
	if err != nil {
		logger.Error("Error connection to DB", err)
	}
	
	newRepository := repository.NewRepository(connection)

	service := service.NewServices(newRepository)

	newServer := server.NewServer(router, service)

	newServer.Init()

	getConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("GetConfig is crashed", err)
	}

	address := net.JoinHostPort(getConfig.Host, getConfig.Port)
	srv := http.Server{
		Addr:    address,
		Handler: router,
	}
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("error in ListenAndServe", err)
	}

	return nil
} 