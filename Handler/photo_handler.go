package handler

import (
	entity "DATABASECRUD/Entity"
	service "DATABASECRUD/Service"
	"database/sql"
	"encoding/json"
	"errors"
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
		serv := service.NewUserSvc()
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		temp_id := serv.VerivyToken(reqToken)
		fmt.Println(temp_id)
		sqlStament := `
		select p.id,p.title,p.caption,p.url,p.user_id,p.created_date,p.updated_date,u.email,u.username from photos p left join users u on p.user_id = u.id`
		rows, err := h.db.Query(sqlStament)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		photos := []*entity.ResponsePhotoGet{}
		for rows.Next() {
			var photo entity.ResponsePhotoGet
			if serr := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.Url, &photo.User_id, &photo.CreatedAt, &photo.UpdatedAt, &photo.Users.Email, &photo.Users.Username); serr != nil {
				fmt.Println("Scan error", serr)
			}
			photos = append(photos, &photo)
		}
		jsonData, _ := json.Marshal(&photos)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
		w.WriteHeader(200)
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
		//check validasi user
		photoserv := service.NewPhotoSvc()
		err := photoserv.CekInputanPhoto(newPhotos.Title, newPhotos.Url)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		} else {
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
		}

	case http.MethodPut:
		fmt.Println("Put")
		if id != "" {
			serv := service.NewUserSvc()
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
			temp_id := serv.VerivyToken(reqToken)
			fmt.Println(temp_id)
			//penampungan rbody
			var newPhotos entity.Photo
			json.NewDecoder(r.Body).Decode(&newPhotos)
			//check validasi user
			photoserv := service.NewPhotoSvc()
			err := photoserv.CekInputanPhoto(newPhotos.Title, newPhotos.Url)
			if err != nil {
				w.Write([]byte(fmt.Sprint(err)))
			} else {
				sqlStament := `update photos set title = $1, caption = $2 , url = $3, updated_date =$4 where id = $5`
				//query.scan
				_, err = h.db.Exec(sqlStament,
					newPhotos.Title,
					newPhotos.Caption,
					newPhotos.Url,
					time.Now(),
					id,
				)
				if err != nil {
					fmt.Println("error update")
					panic(err)
				}
				sqlstatment2 := `select * from photos where id= $1`
				err = h.db.QueryRow(sqlstatment2, id).
					Scan(&newPhotos.Id, &newPhotos.Title, &newPhotos.Caption, &newPhotos.Url, &newPhotos.User_id, &newPhotos.CreatedAt, &newPhotos.UpdatedAt)
				// count, err := res.RowsAffected()
				if err != nil {
					panic(err)
				}

				response := entity.ResponsePuPhoto{
					Id:        newPhotos.Id,
					Title:     newPhotos.Title,
					Caption:   newPhotos.Caption,
					Url:       newPhotos.Url,
					User_id:   newPhotos.User_id,
					UpdatedAt: newPhotos.UpdatedAt,
				}

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
			//cek token
			serv := service.NewUserSvc()
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
			temp_id := serv.VerivyToken(reqToken)
			fmt.Println(temp_id)
			sqlstament := `DELETE from photos where id = $1 and user_id = $2;`
			_, err := h.db.Exec(sqlstament, id, temp_id)

			if err != nil {
				panic(err)
			}
			message := entity.Message{
				Message: "Your photo has been successfully deleted",
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
