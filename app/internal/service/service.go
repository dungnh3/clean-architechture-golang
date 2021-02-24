package service

import (
	"github.com/dungnh3/clean-architechture/app/config"
	"github.com/dungnh3/clean-architechture/app/internal/repository"
	"github.com/dungnh3/clean-architechture/app/internal/usecase"
	"github.com/dungnh3/clean-architechture/pkg/database"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	logger  *zap.Logger
	userSvc usecase.UserUsecase
}

func NewService(cfg *config.Config) (*Service, error) {
	logger, _ := zap.NewProduction()
	db, err := mustConnectDB(cfg.Database, logger)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserSQLStore(db, logger)
	userSvc := usecase.NewUserUsecase(userRepo)
	svc := &Service{
		logger:  logger,
		userSvc: userSvc,
	}
	return svc, nil
}

func mustConnectDB(cfg *database.Config, logger *zap.Logger) (*gorm.DB, error) {
	dsn := cfg.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxIdleTime(60 * time.Second)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	err = db.Raw("SELECT 1").Error
	if err != nil {
		panic(err)
	}
	logger.Info("Connected with database server",
		zap.Any("host", cfg.Host),
		zap.Any("port", cfg.Port),
		zap.Any("database", cfg.Database),
	)
	return db, nil
}
