package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MartinRavanov72/go-rest-api/api/auth"
	"github.com/MartinRavanov72/go-rest-api/api/models"
	"github.com/MartinRavanov72/go-rest-api/api/responses"
	"github.com/MartinRavanov72/go-rest-api/api/utils/formaterror"
)

func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	
	co, err := r.Cookie("jwt")
    if err != nil {
        fmt.Printf("Cant find cookie :/\r\n")
        return
	}
	
	uid, err := auth.ExtractTokenID(co.Value)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != post.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	postCreated, err := post.CreatePost(server.DB)
	error := post.Validate()
	if error != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, error)
		return
	}
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, postCreated.ID))
	responses.JSON(w, http.StatusCreated, postCreated)
}

func (server *Server) GetPosts(w http.ResponseWriter, r *http.Request) {

	post := models.Post{}

	posts, err := post.GetAllPosts(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
	
}

