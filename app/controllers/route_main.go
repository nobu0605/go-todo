package controllers

import (
	"encoding/json"
	"go-todo/app/models"
	"io"
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

func createTodo(w http.ResponseWriter, r *http.Request) {
	// var value interface{}

	body, _ := io.ReadAll(r.Body)
	// int1, _ := strconv.Atoi(string(body))
	input := make(map[string]any)
	_ = json.Unmarshal(body, &input)

	err := models.CreateTodo(
		interface{}(input["title"]).(string), 
		interface{}(input["user_id"]).(float64), 
		interface{}(input["description"]).(string), 
		interface{}(input["status_id"]).(float64))
	if err != nil {
		log.Fatalln(err)
	}
	// resp := map[string]string{"status":"OK"}
	// json.NewEncoder(w).Encode(resp)
}

