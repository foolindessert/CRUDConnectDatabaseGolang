package service

import "errors"

type SocialmediaIfac interface {
	CekInputanSocialMedia(name, social_media_url string) error
}
type SocialmediaSvc struct{}

// CekInputanSocialMedia implements SocialmediaIfac
func (*SocialmediaSvc) CekInputanSocialMedia(name, social_media_url string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if social_media_url == "" {
		return errors.New("social_media_url cannot be empty")
	}
	return nil
}

func NewSocialMediaSv() SocialmediaIfac {
	return &SocialmediaSvc{}
}
