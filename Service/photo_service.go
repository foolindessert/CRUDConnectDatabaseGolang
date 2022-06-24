package service

import "errors"

type PhotoIface interface {
	CekInputanPhoto(title, url string) error
}
type PhotoSvc struct{}

func NewPhotoSvc() PhotoIface {
	return &PhotoSvc{}
}

func (u *PhotoSvc) CekInputanPhoto(title, url string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	if url == "" {
		return errors.New("username cannot be empty")
	}
	return nil
}
