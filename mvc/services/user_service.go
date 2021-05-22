package services

import (
	"github.com/kmchary/microservices-1/mvc/domain"
	"github.com/kmchary/microservices-1/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
