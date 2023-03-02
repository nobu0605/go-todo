package controllers

import (
	"encoding/json"
	"go-todo/app/models"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	todos,err := models.GetTodos()
	if err != nil {
		log.Fatalln(err)
	}
	
	json.NewEncoder(w).Encode(todos)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("authorization")
	if auth == "" { return }

	sess := models.Session{UUID: auth}
	user, err := sess.GetUserBySession()

	if err != nil {
		log.Fatalln(err)
	}
	
	json.NewEncoder(w).Encode(user)
}

func getStatuses(w http.ResponseWriter, r *http.Request) {
	statuses,err := models.GetStatuses()
	if err != nil {
		log.Fatalln(err)
	}
	
	json.NewEncoder(w).Encode(statuses)
}

