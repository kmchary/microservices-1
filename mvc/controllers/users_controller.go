package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kmchary/microservices-1/mvc/services"
	"github.com/kmchary/microservices-1/mvc/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userIdValue := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt((userIdValue), 10, 64)
	if err != nil {
		apiErr := utils.ApplicationError{
			Message:    fmt.Sprintf("invalid user_id %v is passed\n", userIdValue),
			StatusCode: http.StatusBadRequest,
		}
		apiErrJSON, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(apiErrJSON)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		apiErrJSON, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(apiErrJSON)
		return
	}

	userJSON, _ := json.Marshal(user)
	res.Write(userJSON)
}
