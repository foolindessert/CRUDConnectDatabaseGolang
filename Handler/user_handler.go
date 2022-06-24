package handler

import (
	entity "DATABASECRUD/Entity"
	_ "context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	service "DATABASECRUD/Service"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserHandlerInterface interface {
	UsersHandler(w http.ResponseWriter, r *http.Request)
}

type RegisterHandlerInterface interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
}

type LoginHandlerInterface interface {
	LoginUser(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	db *sql.DB
}

type RegisterHandler struct {
	db *sql.DB
}

type LoginHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) UserHandlerInterface {
	return &UserHandler{db: db}
}

func UserRegisterHandler(db *sql.DB) RegisterHandlerInterface {
	return &RegisterHandler{db: db}
}

func UserLoginHandler(db *sql.DB) LoginHandlerInterface {
	return &LoginHandler{db: db}
}

var (
	db *sql.DB

	err error
)

// LoginUser implements LoginHandlerInterface
func (h *LoginHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newUser entity.User
		var validasiUser *entity.User

		json.NewDecoder(r.Body).Decode(&newUser)
		tempPassword := newUser.Password
		newPassword := []byte(newUser.Password)
		_, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		validasiUser = &newUser
		serv := service.NewUserSvc()

		// newUser.Password = string(hashedPassword)
		// fmt.Println(newUser.Password)
		sqlStatment := `select * from public.users where email = $1`

		err = h.db.QueryRow(sqlStatment, newUser.Email).
			Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password, &newUser.Age, &newUser.CreatedAt, &newUser.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(newUser)
		validasiUser, err = serv.Login(validasiUser, tempPassword)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(fmt.Sprint(err)))
		} else {
			var token entity.Token
			token.TokenJwt = serv.GetToken(uint(newUser.Id), newUser.Email, newUser.Password)
			jsonData, _ := json.Marshal(&token)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(jsonData)
		}

	}
}

// RegisterUser implements RegisterHandlerInterface
func (h *RegisterHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	//cek pos

	if r.Method == "POST" {
		var newUser entity.User
		var validasiUser *entity.User
		json.NewDecoder(r.Body).Decode(&newUser)

		newPassword := []byte(newUser.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		validasiUser = &newUser
		serv := service.NewUserSvc()
		validasiUser, err = serv.Register(validasiUser)
		if err != nil {
			panic(err)
		}
		// newUser.Password = string(newPassword)
		// fmt.Println(newUser.Password)
		newUser.Password = string(hashedPassword)
		sqlStatment := `insert into users
		(username,email,password,age,created_date,updated_date)
		values ($1,$2,$3,$4,$5,$5) Returning id` //sesuai dengan nama table
		err = h.db.QueryRow(sqlStatment,
			newUser.Username,
			newUser.Email,
			newUser.Password,
			newUser.Age,
			time.Now(),
		).Scan(&newUser.Id)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(newUser)
		// fmt.Println(newUser.Id)
		response_Register := entity.ResponseRegister{
			Age:      newUser.Age,
			Email:    newUser.Email,
			Id:       newUser.Id,
			Username: newUser.Username,
		}
		jsonData, _ := json.Marshal(&response_Register)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(201)
		w.Write(jsonData)

	}
}

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
		h.deleteUserHandler(w, r)

	}
}

//user
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
	json.NewDecoder(r.Body).Decode(&newUser)
	sqlStatment := `insert into users
	(username,email,password,age,Created_date,Updated_date)
	values ($1,$2,$3,$4,$5,$5)` //sesuai dengan nama table
	res, err := h.db.Exec(sqlStatment,
		newUser.Username,
		newUser.Email,
		newUser.Password,
		newUser.Age,
		time.Now(),
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

func (h *UserHandler) updateUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" { // get by id
		var newUser entity.User
		var validasiUser *entity.User
		params := mux.Vars(r)
		id := params["id"]
		serv := service.NewUserSvc()

		json.NewDecoder(r.Body).Decode(&newUser)
		fmt.Println(newUser)
		var tempUser entity.User
		sqlstatment3 := `select * from users where id= $1`
		err := h.db.QueryRow(sqlstatment3, id).
			Scan(&tempUser.Id, &tempUser.Username, &tempUser.Email, &tempUser.Password, &tempUser.Age, &tempUser.CreatedAt, &tempUser.UpdatedAt)
		// count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}
		fmt.Println(tempUser)
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		fmt.Println(reqToken)
		if err := serv.CheckToken(reqToken, uint(tempUser.Id), tempUser.Email, tempUser.Password); err != nil {
			panic(err)
		}
		validasiUser = &newUser

		validasiUser, err = serv.UpdateUser(validasiUser)
		if err != nil {
			panic(err)
		}

		// json.NewDecoder(r.Body).Decode(&newUser)
		// fmt.Println(r.Body)
		fmt.Println(newUser)
		sqlstatment := `
		update users set username = $1, email = $2, updated_date = $3
		where id = $4;`

		_, err = h.db.Exec(sqlstatment,
			newUser.Username,
			newUser.Email,
			time.Now(),
			id,
		)
		if err != nil {
			fmt.Println("error update")
			panic(err)
		}
		sqlstatment2 := `select * from users where id= $1`
		err = h.db.QueryRow(sqlstatment2, id).
			Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password, &newUser.Age, &newUser.CreatedAt, &newUser.UpdatedAt)
		// count, err := res.RowsAffected()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(newUser)
		responseUpdateUser := entity.ResponseUpdateUser{
			Id:        newUser.Id,
			Email:     newUser.Email,
			Username:  newUser.Username,
			Age:       newUser.Age,
			UpdatedAt: time.Now(),
		}
		jsonData, _ := json.Marshal(&responseUpdateUser)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonData)
		// w.Write([]byte(fmt.Sprint("User  update ", count)))
		return
	}
}

func (h *UserHandler) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	serv := service.NewUserSvc()
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	temp_id := serv.VerivyToken(reqToken)
	fmt.Println(temp_id)
	// if temp_id != nil{}
	sqlstament := `DELETE from users where id = $1;`
	_, err := h.db.Exec(sqlstament, temp_id)

	if err != nil {
		panic(err)
	}
	message := entity.Message{
		Message: "Your account has been successfully deleted",
	}
	jsonData, _ := json.Marshal(&message)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)

}
