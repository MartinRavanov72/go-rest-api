package controllers

import (
	// "encoding/json"
	// "errors"
	"fmt"
	// "io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/MartinRavanov72/go-rest-api/api/auth"
	"github.com/MartinRavanov72/go-rest-api/api/models"
	"github.com/MartinRavanov72/go-rest-api/api/responses"
	// "github.com/MartinRavanov72/go-rest-api/api/utils/formaterror"
)

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	co, err := r.Cookie("jwt")
    if err != nil {
        fmt.Printf("Cant find cookie :/\r\n")
        return
	}
	fmt.Printf(co.Value)

	user := models.User{}

	users, err := user.GetAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)

}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.GetUserByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

func (server *Server) GetCurrentUser(w http.ResponseWriter, r *http.Request) {

	co, err := r.Cookie("jwt")
    if err != nil {
        fmt.Printf("Cant find cookie :/\r\n")
        return
	}
	
	uid, err := auth.ExtractTokenID(co.Value)


	user := models.User{}
	userGotten, err := user.GetUserByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	// vars := mux.Vars(r)
	// uid, err := strconv.ParseUint(vars["id"], 10, 32)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	// user := models.User{}
	// err = json.Unmarshal(body, &user)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	// tokenID, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// if tokenID != uint32(uid) {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }
	// user.Prepare()
	// err = user.Validate("update")
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	// updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
	// if err != nil {
	// 	formattedError := formaterror.FormatError(err.Error())
	// 	responses.ERROR(w, http.StatusInternalServerError, formattedError)
	// 	return
	// }
	// responses.JSON(w, http.StatusOK, updatedUser)
}


