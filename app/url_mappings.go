package app

import (
	"net/http"

	"github.com/psinthorn/gostack_items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.Create).Methods(http.MethodPost)
	router.HandleFunc("/items", controllers.Get).Methods(http.MethodGet)
}
