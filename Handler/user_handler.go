package handler

import (
	entity "DATABASECRUD/Entity"
	middleware "DATABASECRUD/Middleware"
	repo "DATABASECRUD/Repo"
	"DATABASECRUD/helper"
	_ "context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

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
	db  *sql.DB
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
		newUser, err = repo.QueryLoginUser(h.db, newUser)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			dataErr := helper.APIResponseFailed("Failed to Decode", http.StatusInternalServerError, false)
			jsonData, _ := json.Marshal(&dataErr)
			_, errWrite := w.Write(jsonData)
			if errWrite != nil {
				return
			}

		} else {
			fmt.Println(newUser)
			validasiUser, err = serv.Login(validasiUser, tempPassword)
			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				dataErr := helper.APIResponseFailed("Failed to Decode", http.StatusInternalServerError, false)
				jsonData, _ := json.Marshal(&dataErr)
				_, errWrite := w.Write(jsonData)
				if errWrite != nil {
					return
				}

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
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			dataErr := helper.APIResponseFailed("Failed to Decode", http.StatusInternalServerError, false)
			jsonData, _ := json.Marshal(&dataErr)
			_, errWrite := w.Write(jsonData)
			if errWrite != nil {
				return
			}

		}
		validasiUser = &newUser
		serv := service.NewUserSvc()
		validasiUser, err = serv.Register(validasiUser)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			dataErr := helper.APIResponseFailed("Failed to Decode", http.StatusInternalServerError, false)
			jsonData, _ := json.Marshal(&dataErr)
			_, errWrite := w.Write(jsonData)
			if errWrite != nil {
				return
			}

		} else {
			// newUser.Password = string(newPassword)
			// fmt.Println(newUser.Password)
			newUser.Password = string(hashedPassword)
			response_Register := repo.QueryRegisterUser(h.db, newUser)
			jsonData, _ := json.Marshal(&response_Register)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(201)
			w.Write(jsonData)
		}

	}
}

func QueryRegisterUser(dB *sql.DB, newUser entity.User) {
	panic("unimplemented")
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
			panic(err)
		}
		responseUpdateUser := repo.QueryUpdateUser(h.db, newUser, id)
		jsonData, _ := json.Marshal(&responseUpdateUser)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonData)
		return

	}
}

func (h *UserHandler) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := middleware.ForUser(ctx)

	fmt.Println(user)
	fmt.Println(user.Id)
	// if temp_id != nil{}

	message := repo.QueryDeleteUser(h.db, user)
	jsonData, _ := json.Marshal(&message)
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
