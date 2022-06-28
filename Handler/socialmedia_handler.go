package handler

import (
	entity "DATABASECRUD/Entity"
	middleware "DATABASECRUD/Middleware"
	repo "DATABASECRUD/Repo"
	service "DATABASECRUD/Service"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type SocialMediaIface interface {
	SocilaMediaHandler(w http.ResponseWriter, r *http.Request)
}

type SocilaMediaHandler struct {
	db *sql.DB
}

func NewSocialMediaHandler(db *sql.DB) SocialMediaIface {
	return &SocilaMediaHandler{db: db}
}

func (h *SocilaMediaHandler) SocilaMediaHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	ctx := r.Context()
	user := middleware.ForUser(ctx)

	fmt.Println(user)
	fmt.Println(user.Id)
	socialmediaserv := service.NewSocialMediaSv()
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
		fmt.Println("GET")
		socialmedias := repo.QueryGetSocialMedia(h.db)
		jsonData, _ := json.Marshal(&socialmedias)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonData)
	case http.MethodPost:
		fmt.Println("POST")
		var newSocialMedia entity.SocialMedia
		json.NewDecoder(r.Body).Decode(&newSocialMedia)
		err := socialmediaserv.CekInputanSocialMedia(newSocialMedia.Name, newSocialMedia.Social_Media_Url)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		} else {
			fmt.Println("sudah dicek")
			user_id := user.Id
			response := repo.QueryPostSocialMedia(h.db, newSocialMedia, user_id)
			jsonData, _ := json.Marshal(&response)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(201)
			w.Write(jsonData)

		}
	case http.MethodPut:
		fmt.Println("PUT")
		if id != "" {
			var newSocialMedia entity.SocialMedia
			json.NewDecoder(r.Body).Decode(&newSocialMedia)
			err := socialmediaserv.CekInputanSocialMedia(newSocialMedia.Name, newSocialMedia.Social_Media_Url)
			if err != nil {
				w.Write([]byte(fmt.Sprint(err)))
			} else {
				fmt.Println("sudah dicek")
				response := repo.QueryUpdateSocialMedia(h.db, newSocialMedia, id)
				jsonData, _ := json.Marshal(&response)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(200)
				w.Write(jsonData)
			}

		} else {
			err = errors.New("id cannot empty")
			w.Write([]byte(fmt.Sprint(err)))
		}
	case http.MethodDelete:
		fmt.Println("DELETE")
		if id != "" {
			message := repo.QueryDeleteSocialMedia(h.db, id)
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
