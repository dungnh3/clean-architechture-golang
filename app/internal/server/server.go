package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/dungnh3/clean-architechture/app/api"
	"github.com/dungnh3/clean-architechture/app/internal/service"
	"github.com/dungnh3/clean-architechture/pkg/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	cfg    *server.Config
	router *gin.Engine
}

func NewServer(cfg *server.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Register(svc *service.Service) error {
	s.router = api.RegisterServiceHandler(svc)
	return nil
}

func (s *Server) Serve() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", s.cfg.Port),
		Handler: s.router,
	}
	errChan := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("error when start http server, listen: %v", err)
			errChan <- err
		}
	}()

	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			return err
		case <-quitChan:
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("server forced to shutdown: ", err)
				return err
			}
			return nil
		}
	}
}
