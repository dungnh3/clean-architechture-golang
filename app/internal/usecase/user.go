package usecase

import (
	"github.com/dungnh3/clean-architechture/app/internal/model"
	"github.com/dungnh3/clean-architechture/app/internal/repository"
)

type UserUsecase interface {
	Create(user *model.User) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}

func (uc *userUsecase) Create(user *model.User) error {
	return uc.userRepo.Create(user)
}
