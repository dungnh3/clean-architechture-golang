package service

import (
	"github.com/dungnh3/clean-architechture/app/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (s *Service) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		s.logger.Error("parse body failed", zap.Any("error", err))
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	if err := s.userSvc.Create(&user); err != nil {
		ctx.JSON(http.StatusForbidden, "forbidden")
		return
	}
	ctx.JSON(http.StatusOK, "success")
	return
}

func (s *Service) GetUsers(ctx *gin.Context) {

}

func (s *Service) GetUserByID(ctx *gin.Context) {

}

func (s *Service) UpdateUser(ctx *gin.Context) {

}
