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
	serv := service.NewUserSvc()
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	temp_id := serv.VerivyToken(reqToken)
	fmt.Println(temp_id)
	socialmediaserv := service.NewSocialMediaSv()
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
		fmt.Println("GET")
		sqlStament := `
		select distinct on (s.id)s.id,s.name,s.social_media_url,s.user_id,u.created_date,u.updated_date,u.id,u.username,p.url from social_media s left join users u on s.user_id = u.id left join photos p on u.id = p.user_id `
		rows, err := h.db.Query(sqlStament)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		socialmedias := []*entity.ResponseSocialMediaGet{}
		for rows.Next() {
			var socialmedia entity.ResponseSocialMediaGet
			if serr := rows.Scan(&socialmedia.Id, &socialmedia.Name, &socialmedia.Social_Media_Url, &socialmedia.User_id, &socialmedia.CreatedAt, &socialmedia.UpdatedAt, &socialmedia.User.Id, &socialmedia.User.Username, &socialmedia.User.Url); serr != nil {
				fmt.Println("Scan error", serr)
			}
			socialmedias = append(socialmedias, &socialmedia)
		}
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
			sqlStament := `insert into social_media
			(name,social_media_url,user_id)
			values ($1,$2,$3) Returning id`
			// intId, err := strconv.Atoi(id)
			err = h.db.QueryRow(sqlStament, newSocialMedia.Name, newSocialMedia.Social_Media_Url, temp_id).Scan(&newSocialMedia.Id)
			if err != nil {
				fmt.Println(err)
			}

			response := entity.ResponseSocialMediaPost{}
			sqlstatment2 := `
			select s.id,s.name,s.social_media_url,s.user_id,u.created_date from social_media s left join users u on s.user_id = u.id where s.id = $1`
			err = h.db.QueryRow(sqlstatment2, newSocialMedia.Id).
				Scan(&response.Id, &response.Name, &response.Social_Media_Url, &response.User_id, &response.CreatedAt)
			if err != nil {
				panic(err)
			}

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
				sqlStament := `update social_media set name = $1, social_media_url= $2 where id = $3`
				//query.scan
				_, err = h.db.Exec(sqlStament,
					newSocialMedia.Name,
					newSocialMedia.Social_Media_Url,
					id,
				)
				if err != nil {
					fmt.Println("error update")
					panic(err)
				}
				response := entity.ResponseSocialMediaPut{}
				sqlstatment2 := `select s.id,s.name,s.social_media_url,s.user_id,u.updated_date from social_media s left join users u on s.user_id = u.id where s.id = $1`
				err = h.db.QueryRow(sqlstatment2, id).
					Scan(&response.Id, &response.Name, &response.Social_Media_Url, &response.User_id, &response.UpdatedAt)
				// count, err := res.RowsAffected()
				if err != nil {
					panic(err)
				}
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
			sqlstament := `DELETE from social_media where id = $1 and user_id = $2;`
			_, err := h.db.Exec(sqlstament, id, temp_id)

			if err != nil {
				panic(err)
			}
			message := entity.Message{
				Message: "Your SocialMedia has been successfully deleted",
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
