package entity

import "time"

type Photo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponsePostPhoto struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponsePhotoGet struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     ReUser    `json:"User"`
}

type ResponsePuPhoto struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
