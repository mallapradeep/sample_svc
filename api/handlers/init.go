package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Init is used to hold the handlers
func Init() {
	router := mux.NewRouter()
	router.HandleFunc("/area", rectangleArea.Methods(http.MethodPost))
	http.ListenAndServe(":8080", router)
}
