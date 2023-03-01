package controllers

import (
	"encoding/json"
	"go-todo/app/models"
	"go-todo/utils/errorhandler"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	user := models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		PassWord: r.PostFormValue("password"),
	}
	if err := user.CreateUser(); err != nil {
		log.Println(err)
		errorhandler.MakeErrResponse(err,w,400)
		return
	}
	json.NewEncoder(w).Encode("Status OK")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		errorhandler.MakeErrResponse(err,w,400)
		return
	}

	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
			dbError := errorhandler.WrapDBError(err)
			errorhandler.MakeErrResponse(dbError,w,400)
			return
		}
		resp := map[string]string{"status":"OK","uuid":session.UUID}
		json.NewEncoder(w).Encode(resp)
	} else {
		errorhandler.MakeErrResponse(errorhandler.ErrPaswordUnmatch,w,400)
	}
}

func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(writer, request, "/login", 302)
}
