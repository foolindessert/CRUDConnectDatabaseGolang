package service

import "errors"

type CommentIface interface {
	CekInputanComment(message string) error
}

type CommentSvc struct{}

func NewCommentSvc() CommentIface {
	return &CommentSvc{}
}

func (u *CommentSvc) CekInputanComment(message string) error {
	if message == "" {
		return errors.New("message cannot be empty")
	}
	return nil
}
