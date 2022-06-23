package middleware

import (
	entity "DATABASECRUD/Entity"
	service "DATABASECRUD/Service"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type AuthIFace interface {
	AuthCekToken(next http.Handler) http.Handler
}

type AuthToken struct {
	db *sql.DB
}

func NewUserSvc(db *sql.DB) AuthIFace {
	return &AuthToken{db: db}
}

var (
	db *sql.DB

	err error
)

func (u *AuthToken) AuthCekToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authhandler := strings.Split(r.Header.Get("Authorization"), " "); len(authhandler) == 2 && authhandler[0] == "Bearer" {
			var newUser entity.User
			var validasiUser *entity.User
			params := mux.Vars(r)
			id := params["id"]
			serv := service.NewUserSvc()

			json.NewDecoder(r.Body).Decode(&newUser)
			fmt.Println(newUser)
			var tempUser entity.User
			sqlstatment3 := `select * from users where id= $1`
			err := u.db.QueryRow(sqlstatment3, id).
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
			next.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
