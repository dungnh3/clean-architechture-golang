package main

import (
	"github.com/dungnh3/clean-architechture/app/config"
	"github.com/dungnh3/clean-architechture/app/internal/server"
	"github.com/dungnh3/clean-architechture/app/internal/service"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
	"os"
)

var (
	cfg    *config.Config
	logger *zap.Logger
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	var err error

	cfg, err = config.Load()
	if err != nil {
		return err
	}

	logger, err = zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()
	logger.Info("start", zap.Any("cfg", cfg))

	rootCmd := &cobra.Command{
		Use:   "server",
		Short: "start server",
		RunE:  serverAction,
	}
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
	return nil
}

func serverAction(cmd *cobra.Command, args []string) error {

	svr := server.NewServer(cfg.Server)

	svc, err := service.NewService(cfg)
	if err != nil {
		return err
	}

	if err := svr.Register(svc); err != nil {
		logger.Error("register sercver failed",
			zap.Any("error", err))
		return err
	}

	logger.Info("starting http server..",
		zap.Any("host", cfg.Server.Host),
		zap.Any("port", cfg.Server.Port))
	if err := svr.Serve(); err != nil {
		logger.Error("start server failed",
			zap.Any("error", err))
		return err
	}
	return nil
}
