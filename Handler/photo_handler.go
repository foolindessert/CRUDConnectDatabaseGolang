package handler

import (
	entity "DATABASECRUD/Entity"
	service "DATABASECRUD/Service"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type PhotoHandlerInterface interface {
	PhotoHandler(w http.ResponseWriter, r *http.Request)
}

type PhotoHandler struct {
	db *sql.DB
}

func NewPhotoHandler(db *sql.DB) PhotoHandlerInterface {
	return &PhotoHandler{db: db}
}

// var (
// 	db *sql.DB

// 	err error
// )

func (h *PhotoHandler) PhotoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	switch r.Method {
	case http.MethodGet:
		fmt.Println("Get")
	case http.MethodPost:
		fmt.Println("Post")
		serv := service.NewUserSvc()
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		temp_id := serv.VerivyToken(reqToken)
		fmt.Println(temp_id)
		//penampungan rbody
		var newPhotos entity.Photo
		json.NewDecoder(r.Body).Decode(&newPhotos)
		//query insert
		sqlStament := `insert into photos
		(title,caption,url,user_id,created_date,updated_date)
		values ($1,$2,$3,$4,$5,$5) Returning id`
		//query.scan
		err = h.db.QueryRow(sqlStament,
			newPhotos.Title,
			newPhotos.Caption,
			newPhotos.Url,
			temp_id,
			time.Now(),
		).Scan(&newPhotos.Id)
		if err != nil {
			fmt.Println(err)
		}
		response := entity.ResponsePostPhoto{
			Id:        newPhotos.Id,
			Title:     newPhotos.Title,
			Caption:   newPhotos.Caption,
			Url:       newPhotos.Url,
			User_id:   int(temp_id),
			CreatedAt: time.Now(),
		}

		jsonData, _ := json.Marshal(&response)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
		w.WriteHeader(201)

	case http.MethodPut:
		fmt.Println("Put")
	case http.MethodDelete:
		fmt.Println("Delete")
	}
}
