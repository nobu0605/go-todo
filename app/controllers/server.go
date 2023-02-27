package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"go-todo/app/models"
	"go-todo/config"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		// if ok, _ := sess.CheckSession(); !ok {
		// 	err = errors.New("Invalid session")
		// }
	}
	return
}

var validPath = regexp.MustCompile("^/todos/(edit|save|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		id, _ := strconv.Atoi(q[2])
		fmt.Println(id)
		fn(w, r, id)
	}
}

func setHeaderMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		frontUrl := os.Getenv("FRONT_URL")
		w.Header().Set("Access-Control-Allow-Origin", frontUrl)
        next.ServeHTTP(w, r)
    }
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	
	http.HandleFunc("/", setHeaderMiddleware(top))
	http.HandleFunc("/login", setHeaderMiddleware(login))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
