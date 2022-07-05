package handler

import (
	entity "DATABASECRUD/Entity"
	middleware "DATABASECRUD/Middleware"
	repo "DATABASECRUD/Repo"
	service "DATABASECRUD/Service"
	"DATABASECRUD/helper"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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

func (h *PhotoHandler) PhotoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	ctx := r.Context()
	user := middleware.ForUser(ctx)

	fmt.Println(user)
	fmt.Println(user.Id)
	switch r.Method {
	case http.MethodGet:
		fmt.Println("Get")
		photos := repo.QueryGetPhoto(h.db)
		jsonData, _ := json.Marshal(&photos)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
		w.WriteHeader(200)
	case http.MethodPost:
		fmt.Println("Post")

		//penampungan rbody
		var newPhotos entity.Photo
		json.NewDecoder(r.Body).Decode(&newPhotos)
		fmt.Println(newPhotos)
		//check validasi user
		photoserv := service.NewPhotoSvc()
		err := photoserv.CekInputanPhoto(newPhotos.Title, newPhotos.Url)
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
			//query insert
			user_id := user.Id
			response := repo.QueryPostPhoto(h.db, newPhotos, user_id)
			jsonData, _ := json.Marshal(&response)
			w.Header().Add("Content-Type", "application/json")
			w.Write(jsonData)
			w.WriteHeader(201)
		}

	case http.MethodPut:
		fmt.Println("Put")
		if id != "" {

			//penampungan rbody
			var newPhotos entity.Photo
			json.NewDecoder(r.Body).Decode(&newPhotos)
			//check validasi user
			photoserv := service.NewPhotoSvc()
			err := photoserv.CekInputanPhoto(newPhotos.Title, newPhotos.Url)
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
				response := repo.QueryUpdatePhoto(h.db, newPhotos, id)
				jsonData, _ := json.Marshal(&response)
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
				w.WriteHeader(200)
			}
		} else {
			err = errors.New("id cannot empty")
			w.Write([]byte(fmt.Sprint(err)))
		}
	case http.MethodDelete:
		fmt.Println("Delete")
		if id != "" {

			message := repo.QueryDeletePhoto(h.db, id)
			jsonData, _ := json.Marshal(&message)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(jsonData)
		} else {
			err = errors.New("id cannot empty")
			w.Write([]byte(fmt.Sprint(err)))
		}
	}
}
