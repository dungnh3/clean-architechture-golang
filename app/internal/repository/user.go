package repository

import (
	"context"
	"github.com/dungnh3/clean-architechture/app/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindUserByID(ctx context.Context, id uint64) (*model.User, error)
	FindUsers(ctx context.Context) ([]*model.User, error)
	UpdateEmail(ctx context.Context, id uint64, email string) error
}

type UserSQlStore struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewUserSQLStore(db *gorm.DB, logger *zap.Logger) *UserSQlStore {
	return &UserSQlStore{
		logger: logger,
		db:     db,
	}
}

func (r *UserSQlStore) Create(user *model.User) error {
	r.logger.Info("create user", zap.Any("user", user))
	return nil
}

func (u *UserSQlStore) FindUserByID(ctx context.Context, id uint64) (*model.User, error) {
	panic("implement me")
}

func (u *UserSQlStore) FindUsers(ctx context.Context) ([]*model.User, error) {
	panic("implement me")
}

func (u *UserSQlStore) UpdateEmail(ctx context.Context, id uint64, email string) error {
	panic("implement me")
}
