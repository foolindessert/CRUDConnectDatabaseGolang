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

type CommentHandlerInterface interface {
	CommentHandler(w http.ResponseWriter, r *http.Request)
}

type CommentHandler struct {
	db *sql.DB
}

func NewCommentHandler(db *sql.DB) CommentHandlerInterface {
	return &CommentHandler{db: db}
}

func (h *CommentHandler) CommentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	ctx := r.Context()
	user := middleware.ForUser(ctx)

	fmt.Println(user)
	fmt.Println(user.Id)
	commentserv := service.NewCommentSvc()
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
		comments := repo.QueryGetComment(h.db)
		jsonData, _ := json.Marshal(&comments)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonData)

	case http.MethodPost:
		fmt.Println("POST")
		var newComment entity.Commment
		json.NewDecoder(r.Body).Decode(&newComment)
		err := commentserv.CekInputanComment(newComment.Message)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		} else {
			fmt.Println("Comment ada isi")
			user_id := user.Id
			response := repo.QueryPostComment(h.db, newComment, user_id)
			jsonData, _ := json.Marshal(&response)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(201)
			w.Write(jsonData)

		}
	case http.MethodPut:
		fmt.Println("PUT")
		var newComment entity.Commment
		json.NewDecoder(r.Body).Decode(&newComment)
		err := commentserv.CekInputanComment(newComment.Message)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		} else {
			fmt.Println("Comment ada isi")
			if id != "" {
				response := repo.QueryUpdateComment(h.db, newComment, id)
				jsonData, _ := json.Marshal(&response)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(200)
				w.Write(jsonData)

			} else {
				err = errors.New("id cannot empty")
				w.Write([]byte(fmt.Sprint(err)))
			}
		}
	case http.MethodDelete:
		fmt.Println("DELETE")
		if id != "" {
			message := repo.QueryDeleteComment(h.db, id)
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
