package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"fmt"

	"github.com/MartinRavanov72/go-rest-api/api/auth"
	"github.com/MartinRavanov72/go-rest-api/api/models"
	"github.com/MartinRavanov72/go-rest-api/api/responses"
	"github.com/MartinRavanov72/go-rest-api/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.CreateUser(server.DB)



	token, err := auth.CreateToken(user.ID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	cookie := http.Cookie{
    	Name: "jwt",
		Value: token, 
		Expires: time.Now().Add(365 * 24 * time.Hour),
		Secure: true,
		MaxAge: 50000,
		HttpOnly: true,
	}

	

	http.SetCookie(w, &cookie)

	
	

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
	co, err := r.Cookie("jwt")
    if err != nil {
        fmt.Printf("Cant find cookie :/\r\n")
        return
	}
	fmt.Printf(co.Value)
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	cookie := http.Cookie{
    	Name: "jwt",
		Value: token, 
		Expires: time.Now().Add(365 * 24 * time.Hour),
		Secure: true,
		MaxAge: 50000,
		HttpOnly: true,
	}

	

	http.SetCookie(w, &cookie)

	
	
	responses.JSON(w, http.StatusOK, cookie)
	co, err := r.Cookie("jwt")
    if err != nil {
        fmt.Printf("Cant find cookie :/\r\n")
        return
	}
	fmt.Printf(co.Value)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}

func (server *Server) Logout(w http.ResponseWriter, r *http.Request){

	cookie := http.Cookie{
    	Name: "jwt",
		Value: "",
		Expires: time.Unix(0, 0),
		HttpOnly: true,
	}
	

	http.SetCookie(w, &cookie)
	responses.JSON(w, http.StatusOK, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}


