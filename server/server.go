package server

import (
	"github.com/rahul0tripathi/gamur/apis/routes"
	"go.uber.org/zap"
	"net/http"
)

type server struct {
	router routes.Router
	logger *zap.SugaredLogger
}

type Server interface {
	Run()
}

func NewServer(r routes.Router, l *zap.SugaredLogger) Server {
	s := server{}
	s.router = r
	s.logger = l
	return s
}

func (s server) Run() {
	s.logger.Info("Starting up api server")
	err := http.ListenAndServe(":3005", s.router.GetBaseRouter())
	if err != nil {
		s.logger.Error(err)
	}
}
