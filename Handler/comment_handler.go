package handler

import (
	entity "DATABASECRUD/Entity"
	middleware "DATABASECRUD/Middleware"
	service "DATABASECRUD/Service"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

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
		sqlStament := `
		select c.id, c.message,c.photo_id,c.user_id,c.updated_date,c.created_date,u.id,u.email,u.username,p.id,p.title,p.caption,p.url,p.user_id from comment c left join photos p on c.photo_id = p.id left join users u on c.user_id = u.id`
		rows, err := h.db.Query(sqlStament)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		}
		defer rows.Close()
		comments := []*entity.ResponseCommentGet{}
		for rows.Next() {
			var comment entity.ResponseCommentGet
			if serr := rows.Scan(&comment.Id, &comment.Message, &comment.Photo_id, &comment.User_id, &comment.UpdatedAt, &comment.CreatedAt, &comment.User.Id, &comment.User.Email, &comment.User.Username, &comment.Photo.Id, &comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.Url, &comment.Photo.User_id); serr != nil {
				w.Write([]byte(fmt.Sprint(serr)))
			}
			comments = append(comments, &comment)
		}
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
			sqlStament := `insert into comment
			(user_id,photo_id,message,created_date,updated_date)
			values ($1,$2,$3,$4,$4) Returning id`
			// intId, err := strconv.Atoi(id)
			err = h.db.QueryRow(sqlStament,
				user.Id,
				newComment.Photo_id,
				newComment.Message,
				time.Now(),
			).Scan(&newComment.Id)
			if err != nil {
				w.Write([]byte(fmt.Sprint(err)))
			}
			response := entity.ResponseCommentPost{
				Id:        newComment.Id,
				Message:   newComment.Message,
				Photo_id:  newComment.Photo_id,
				User_id:   int(user.Id),
				CreatedAt: time.Now(),
			}

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
				sqlStament := `update comment set message = $1, updated_date =$2 where id = $3`
				//query.scan
				_, err = h.db.Exec(sqlStament,
					newComment.Message,
					time.Now(),
					id,
				)
				if err != nil {
					fmt.Println("error update")
					w.Write([]byte(fmt.Sprint(err)))
				}
				response := entity.ResponseUpdateComment{}
				sqlstatment2 := `select c.id,p.title,p.caption,p.url,c.user_id,c.updated_date from comment c left join photos p on c.photo_id = p.id where c.id= $1`
				err = h.db.QueryRow(sqlstatment2, id).
					Scan(&response.Id, &response.Title, &response.Caption, &response.Url, &response.User_id, &response.UpdatedAt)
				// count, err := res.RowsAffected()
				if err != nil {
					w.Write([]byte(fmt.Sprint(err)))
				}
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
			sqlstament := `DELETE from comment where id = $1 and user_id = $2;`
			_, err := h.db.Exec(sqlstament, id, user.Id)

			if err != nil {
				w.Write([]byte(fmt.Sprint(err)))
			}
			message := entity.Message{
				Message: "Your Comment has been successfully deleted",
			}
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
