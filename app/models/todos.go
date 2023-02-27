package models

import (
	"fmt"
	"log"
	"time"
)

type Todo struct {
	ID        int `json:"id"`
	Title   string `json:"title"`
	UserID    int `json:"userId"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetTodos()(todos []Todo, err error){
	cmd := `select 
	todos.id,
	todos.title,
	todos.description,
	statuses.name as status,
	todos.created_at
	from todos
	left join statuses
	on todos.status_id = statuses.id`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	
	fmt.Println(rows)
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Status,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func CreateTodo()(err error){
	cmd := `insert into todos (
		title,
		user_id,
		description,
		status_id,
		created_at ) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, "content", 1, "description", 1, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
