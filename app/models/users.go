package models

import (
	"log"
	"time"

	"go-todo/utils/errorhandler"
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



func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		dbError := errorhandler.WrapDBError(err)
		return dbError
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)

	if err != nil {
		dbError := errorhandler.WrapDBError(err)
		return user,dbError
	}

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	var now = time.Now()
	session = Session{}
	cmd1 := `insert into sessions (
		uuid, 
		email, 
		user_id, 
		expiration_date,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, now.Add(1 * time.Hour) ,now)
	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at
	 from sessions where user_id = ? and email = ?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at
	 from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	session := Session{}
	cmd1 := `select user_id from sessions where uuid = ?`
   	err = Db.QueryRow(cmd1, sess.UUID).Scan(&session.UserID)
	
	user = User{}
	cmd2 := `select id, uuid, name, email, created_at FROM users
	where id = ?`
	err = Db.QueryRow(cmd2, session.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)

	return user, err
}
