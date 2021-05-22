package domain

import (
	"fmt"
	"net/http"

	"github.com/kmchary/microservices-1/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {
			Id:        uint64(123),
			FirstName: "John",
			LastName:  "Miller",
			Email:     "john.miller@test.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user with userId %v not found", userId),
		StatusCode: http.StatusNotFound,
	}

}
