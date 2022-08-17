package app

import (
	"fmt"
	"github.com/adnanmhd/go-porto-homies/app/controller/http/v1"
	"github.com/adnanmhd/go-porto-homies/app/repo"
	"github.com/adnanmhd/go-porto-homies/app/usecase"
	"os"
	"os/signal"
	"syscall"

	"github.com/adnanmhd/go-porto-homies/config"
	"github.com/adnanmhd/go-porto-homies/pkg/httpserver"
	"github.com/adnanmhd/go-porto-homies/pkg/logger"
	"github.com/adnanmhd/go-porto-homies/pkg/mysql"
	"github.com/gin-gonic/gin"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	// Logger
	l := logger.New(cfg.Log.Level)

	// Repository
	db, err := mysql.NewConnection(cfg.MySQL)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - mysql.NewConnection: %w", err))
	}
	defer db.Close()

	// Use case
	propertyUseCase := usecase.New(
		repo.New(db),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, propertyUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
