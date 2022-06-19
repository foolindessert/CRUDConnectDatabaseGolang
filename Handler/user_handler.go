package handler

import (
	entity "DATABASECRUD/Entity"
	_ "context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type UserHandlerInterface interface {
	UsersHandler(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) UserHandlerInterface {
	return &UserHandler{db: db}
}

var (
	db *sql.DB

	err error
)

func (h *UserHandler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodGet:
		//users/{id}
		if id != "" { // get by id
			h.getUsersByIDHandler(w, r, id)
		} else { // get all
			//users
			h.getUsersHandler(w, r)
		}
	case http.MethodPost:
		//users
		h.createUsersHandler(w, r)
	case http.MethodPut:
		//users/{id}
		h.updateUserHandler(w, r, id)
	case http.MethodDelete:
		//users/{id}
		h.deleteUserHandler(w, r, id)
	}
}

func (h *UserHandler) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []*entity.User{}
	sqlStatment := `SELECT  * from users` //sesuai dengan nama table

	rows, err := h.db.Query(sqlStatment)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user entity.User
		if serr := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt); serr != nil {
			fmt.Println("Scan error", serr)
		}
		users = append(users, &user)
	}
	jsonData, _ := json.Marshal(&users)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *UserHandler) getUsersByIDHandler(w http.ResponseWriter, r *http.Request, id string) {
	users := []*entity.User{}
	sqlStatment := `SELECT  * from users where id = ` + `'` + id + `'` //sesuai dengan nama table
	rows, err := h.db.Query(sqlStatment)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user entity.User
		if serr := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt); serr != nil {
			fmt.Println("Scan error", serr)
		}
		users = append(users, &user)
	}
	jsonData, _ := json.Marshal(&users)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *UserHandler) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	var newUser entity.User
	users := entity.User{}
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()
	sqlStatment := `insert into users
	(username,email,password,age,createdat,updatedat) 
	values ($1,$2,$3,$4,$5,$6)
	returning * ` //sesuai dengan nama table
	err := h.db.QueryRow(sqlStatment,
		newUser.Username,
		newUser.Email,
		newUser.Password,
		newUser.Age,
		newUser.CreatedAt,
		newUser.UpdatedAt,
	).Scan(
		&users.Username,
		&users.Email,
		&users.Password,
		&users.Age,
		&users.CreatedAt,
		&users.UpdatedAt)

	if err != nil {
		panic(err)
	}
	w.Write([]byte(fmt.Sprint("User  created")))
	return
}

func (h *UserHandler) updateUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" { // get by id
		var newUser entity.User
		json.NewDecoder(r.Body).Decode(&newUser)
		newUser.CreatedAt = time.Now()
		newUser.UpdatedAt = time.Now()
		sqlstatment := `
		update users set username = $1, email = $2, password = $3, createdat = $4, updatedat = $5 
		where id = '` + id + `';`

		res, err := h.db.Exec(sqlstatment,
			newUser.Username,
			newUser.Email,
			newUser.Password,
			newUser.CreatedAt,
			newUser.UpdatedAt,
		)

		if err != nil {
			panic(err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}

		w.Write([]byte(fmt.Sprint("User  update ", count)))
		return
	}
}

func (h *UserHandler) deleteUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	sqlstament := `DELETE from users where id = $1;`
	if idInt, err := strconv.Atoi(id); err == nil {
		res, err := h.db.Exec(sqlstament, idInt)
		if err != nil {
			panic(err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(fmt.Sprint("Delete user rows ", count)))
		return
	}

}
