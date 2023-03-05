package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"go-todo/app/models"
	"go-todo/config"
	"go-todo/utils/errorhandler"
)


func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	auth := request.Header.Get("authorization")
	if auth == "" { return }

	sess = models.Session{UUID: auth}
	if ok, _ := sess.CheckSession(); !ok {
		err = errors.New("Invalid session")
	}
	return
}


func setHeaderMiddleware(w http.ResponseWriter) {
	frontUrl := os.Getenv("FRONT_URL")
	w.Header().Set("Access-Control-Allow-Origin", frontUrl)
	w.Header().Set("Access-Control-Allow-Headers", "authorization, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS")
}

func privateRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaderMiddleware(w)
		
		sess, err := session(w, r)
		fmt.Println("sess.ID",sess.ID)
		if err != nil {
			errorhandler.MakeErrResponse(err,w,401)
		}else if sess.ID == 0{
			// Do nothing
		}else{
			next.ServeHTTP(w, r)
		} 
    }
}

func publicRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaderMiddleware(w)
		next.ServeHTTP(w, r)
    }
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	
	// Public Route
	http.HandleFunc("/signup", publicRoute(signup))
	http.HandleFunc("/logout", publicRoute(logout))
	http.HandleFunc("/authenticate", publicRoute(authenticate))

	// Private Route
	http.HandleFunc("/", privateRoute(top))
	http.HandleFunc("/getUser", privateRoute(getUser))
	http.HandleFunc("/getStatuses", privateRoute(getStatuses))
	http.HandleFunc("/createTodo", privateRoute(createTodo))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
