package server

import (
	"fmt"
	"net/http"

	"github.com/k4zb3k/todo-list/pkg/logging"
)

var logger = logging.GetLogger()

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}