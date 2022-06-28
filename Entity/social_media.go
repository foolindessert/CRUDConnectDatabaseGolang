package entity

import "time"

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
