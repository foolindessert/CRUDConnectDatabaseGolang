package entity

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseRegister struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type Token struct {
	TokenJwt string `json:"token"`
}

type ResponseUpdateUser struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       string    `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ReUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Message struct {
	Message string `json:"message"`
}

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

type Commment struct {
	Id        int       `json:"id"`
	User_id   int       `json:"user_id"`
	Photo_id  int       `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

type SocialMedia struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
	User_id          int    `json:"user_id"`
}

type ResponseSocialMediaPost struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Social_Media_Url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
}

type ResponseSocialMediaPut struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Social_Media_Url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type SocialMedias struct {
	SocialMedia ResponseSocialMediaGet `json:"social_media"`
}

type ResponseSocialMediaGet struct {
	Id               int                        `json:"id"`
	Name             string                     `json:"name"`
	Social_Media_Url string                     `json:"social_media_url"`
	User_id          int                        `json:"user_id"`
	CreatedAt        time.Time                  `json:"created_at"`
	UpdatedAt        time.Time                  `json:"updated_at"`
	User             ResponseUserSocialMediaGet `json:"User"`
}

type ResponseUserSocialMediaGet struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Url      string `json:"profile_image_url"`
}
