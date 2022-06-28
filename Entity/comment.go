package entity

import "time"

type Commment struct {
	Id        int       `json:"id"`
	User_id   int       `json:"user_id"`
	Photo_id  int       `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ResponseCommentPost struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	Photo_id  int       `json:"photo_id"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseCommentGet struct {
	Id        int                     `json:"id"`
	User_id   int                     `json:"user_id"`
	Photo_id  int                     `json:"photo_id"`
	Message   string                  `json:"message"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
	User      ResponseCommentUserGet  `json:"User"`
	Photo     ResponseCommentPhotoGet `json:"Photo"`
}

type ResponseCommentUserGet struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ResponseCommentPhotoGet struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Caption string `json:"caption"`
	Url     string `json:"photo_url"`
	User_id int    `json:"user_id"`
}

type ResponseUpdateComment struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
