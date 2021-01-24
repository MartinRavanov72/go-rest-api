package controllers

import (
	"net/http"

	"github.com/MartinRavanov72/go-rest-api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome")
}