package app

import (
	"log"
	"net/http"

	"github.com/kmchary/microservices-1/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
