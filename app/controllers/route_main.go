package controllers

import (
	"encoding/json"
	"fmt"
	"go-todo/app/models"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	// _, err := session(w, r)
	// frontUrl := os.Getenv("FRONT_URL")

	// w.Header().Set("Access-Control-Allow-Headers", frontUrl)
	// w.Header().Set("Access-Control-Allow-Origin", frontUrl)
	todos,err := models.GetTodos()
	if err != nil {
		log.Fatalln(err)
	}
	
	fmt.Println("todos route",todos)
	json.NewEncoder(w).Encode(todos)
	// } else {
	// 	http.Redirect(w, r, "/todos", 302)
	// }
}

