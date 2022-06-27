package handler

import (
	entity "DATABASECRUD/Entity"
	middleware "DATABASECRUD/Middleware"
	_ "context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
			w.Write([]byte(fmt.Sprint(err)))

		}
		validasiUser = &newUser
		serv := service.NewUserSvc()

		// newUser.Password = string(hashedPassword)
		// fmt.Println(newUser.Password)
		sqlStatment := `select * from public.users where email = $1`

		err = h.db.QueryRow(sqlStatment, newUser.Email).
			Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password, &newUser.Age, &newUser.CreatedAt, &newUser.UpdatedAt)
		if err != nil {
			w.Write([]byte(fmt.Sprint(errors.New("email not register"))))

		} else {
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
			w.Write([]byte(fmt.Sprint(err)))

		}
		validasiUser = &newUser
		serv := service.NewUserSvc()
		validasiUser, err = serv.Register(validasiUser)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))

		} else {
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
				w.Write([]byte(fmt.Sprint(err)))
			} else {
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

	}
}

func (h *UserHandler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	switch r.Method {
	case http.MethodPut:
		//users/{id}
		h.updateUserHandler(w, r, id)
	case http.MethodDelete:
		//users/{id}
		h.deleteUserHandler(w, r)

	}
}
func (h *UserHandler) updateUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" { // get by id
		ctx := r.Context()
		user := middleware.ForUser(ctx)

		fmt.Println(user)
		fmt.Println(user.Id)
		var newUser entity.User
		json.NewDecoder(r.Body).Decode(&newUser)
		fmt.Println(newUser)
		var validasiUser *entity.User
		validasiUser = &newUser
		serv := service.NewUserSvc()
		validasiUser, err = serv.UpdateUser(validasiUser)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		} else {
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
				w.Write([]byte(fmt.Sprint(errors.New("user id don't exists"))))

			} else {
				sqlstatment2 := `select * from users where id= $1`
				err = h.db.QueryRow(sqlstatment2, id).
					Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password, &newUser.Age, &newUser.CreatedAt, &newUser.UpdatedAt)
				// count, err := res.RowsAffected()
				if err != nil {
					w.Write([]byte(fmt.Sprint(errors.New("user id don't exists"))))
				} else {
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
					return
				}
			}
		}

	}
}

func (h *UserHandler) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForUser(ctx)

	fmt.Println(user)
	fmt.Println(user.Id)
	// if temp_id != nil{}
	sqlstament := `DELETE from users where id = $1;`
	_, err := h.db.Exec(sqlstament, user.Id)

	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))

	}
	message := entity.Message{
		Message: "Your account has been successfully deleted",
	}
	jsonData, _ := json.Marshal(&message)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)

}
