package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	handler "DATABASECRUD/Handler"
	middleware "DATABASECRUD/Middleware"
	"DATABASECRUD/conf"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB

	err error
)

const PORT = ":8080"

func main() {
	// mmemastikan db connect atau tidak
	db, err = sql.Open("postgres", ConnectDbPsql(conf.Host, conf.User, conf.Password, conf.Dbname, conf.Port))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Succesfully connected to database")

	// handler crud
	r := mux.NewRouter()
	userHandler := handler.NewUserHandler(db)
	registerHandler := handler.UserRegisterHandler(db)
	loginHandler := handler.UserLoginHandler(db)
	r.HandleFunc("/users/register", registerHandler.RegisterUser)
	r.HandleFunc("/users/login", loginHandler.LoginUser)
	r.Handle("/users/{id}", middleware.AuthCekToken(http.HandlerFunc(userHandler.UsersHandler))).Methods("PUT")
	r.Handle("/users", middleware.AuthCekToken(http.HandlerFunc(userHandler.UsersHandler))).Methods("DELETE")

	//handler photo
	photoHandler := handler.NewPhotoHandler(db)
	r.Handle("/photos", middleware.AuthCekToken(http.HandlerFunc(photoHandler.PhotoHandler))).Methods("POST")
	r.Handle("/photos", middleware.AuthCekToken(http.HandlerFunc(photoHandler.PhotoHandler))).Methods("GET")
	r.Handle("/photos/{id}", middleware.AuthCekToken(http.HandlerFunc(photoHandler.PhotoHandler))).Methods("PUT")
	r.Handle("/photos/{id}", middleware.AuthCekToken(http.HandlerFunc(photoHandler.PhotoHandler))).Methods("DELETE")
	//handler comment
	commentHandler := handler.NewCommentHandler(db)
	r.Handle("/comments", middleware.AuthCekToken(http.HandlerFunc(commentHandler.CommentHandler))).Methods("POST")
	r.Handle("/comments", middleware.AuthCekToken(http.HandlerFunc(commentHandler.CommentHandler))).Methods("GET")
	r.Handle("/comments/{id}", middleware.AuthCekToken(http.HandlerFunc(commentHandler.CommentHandler))).Methods("PUT")
	r.Handle("/comments/{id}", middleware.AuthCekToken(http.HandlerFunc(commentHandler.CommentHandler))).Methods("DELETE")
	//handler socialMedia
	socialmediaHandler := handler.NewSocialMediaHandler(db)
	r.Handle("/socialmedias", middleware.AuthCekToken(http.HandlerFunc(socialmediaHandler.SocilaMediaHandler))).Methods("POST")
	r.Handle("/socialmedias", middleware.AuthCekToken(http.HandlerFunc(socialmediaHandler.SocilaMediaHandler))).Methods("GET")
	r.Handle("/socialmedias/{id}", middleware.AuthCekToken(http.HandlerFunc(socialmediaHandler.SocilaMediaHandler))).Methods("PUT")
	r.Handle("/socialmedias/{id}", middleware.AuthCekToken(http.HandlerFunc(socialmediaHandler.SocilaMediaHandler))).Methods("DELETE")

	fmt.Println("Now listening on port 0.0.0.0" + PORT)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func ConnectDbPsql(host, user, password, name string, port int) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		name)
	return psqlInfo
}
