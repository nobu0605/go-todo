package models

import (
	"time"
)

type User struct {
	ID        int `json:"id"`
	UUID      string `json:"uuId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	PassWord  string `json:"passWord"`
	CreatedAt time.Time `json:"createdAt"`
	Todos     []Todo `json:"todos"`
}

type Session struct {
	ID        int `json:"id"`
	UUID      string `json:"uuId"`
	Email     string `json:"email"`
	UserID    int `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
