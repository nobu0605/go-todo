package models

import (
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

type Status struct {
	ID        int `json:"id"`
	Name   string `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type StatusWithTodos struct {
	color  string `json:"color"`
	Name   string `json:"name"`
	Todos []Todo `json:"todos"`
}

func GetTodos()(statusesWithTodos map[int]map[string]interface{}, err error){
	cmd := `select DISTINCT
	statuses.id as status_id,
	statuses.name as status
	from statuses
	left join todos
	on statuses.id = todos.status_id
	order by statuses.id asc`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	columns, err2 := rows.Columns()
    if err2 != nil {
        log.Fatalln(err)
    }
    count := len(columns)
	values := make([]interface{}, count)
    scanArgs := make([]interface{}, count)
    for i := range values {
        scanArgs[i] = &values[i]
    }
	
	statusData := make(map[int]map[string]interface{})

	index := 0
    for rows.Next() {
		status := make(map[string]interface{})
        err := rows.Scan(scanArgs...)
        if err != nil {
            log.Fatalln(err)
        }
        for i, v := range values {
			var todos = []Todo{}
			if columns[i] == "status_id" && v != nil {
				cmd := `select 
				todos.id,
				todos.title,
				todos.description,
				statuses.name as status
				from todos
				left join statuses
				on todos.status_id = statuses.id
				where todos.id in
				(select id 
				from todos where status_id = ?)`
			
				todoRows, err := Db.Query(cmd,v)
				
				if err != nil {
					// 対象のステータスに該当のTodoがない場合、continue
					continue
				}
				for todoRows.Next() {
					var todo Todo
					err = todoRows.Scan(
						&todo.ID,
						&todo.Title,
						&todo.Description,
						&todo.Status)
					if err != nil {
						log.Fatalln(err)
					}
					todos = append(todos, todo)
				}
				status["todos"] = todos
				status["status_id"] = v
			}else{
				status["name"] = v
				status["color"] = "white"
			}
        }
		statusData[index] = status
		index++
    }
	    
	return statusData, err
}

func CreateTodo(title string, user_id float64, description string, status_id float64)(err error){
	cmd := `insert into todos (
		title,
		user_id,
		description,
		status_id,
		created_at ) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, title, user_id, description, status_id, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetStatuses()(statuses []Status, err error){
	cmd := `select 
	id,
	name
	from statuses`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	
	for rows.Next() {
		var status Status
		err = rows.Scan(&status.ID,
			&status.Name)
		if err != nil {
			log.Fatalln(err)
		}
		statuses = append(statuses, status)
	}
	rows.Close()

	return statuses, err
}
