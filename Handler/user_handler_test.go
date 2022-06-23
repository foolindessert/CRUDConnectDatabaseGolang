package handler

import (
	_ "context"
	"database/sql"
	"net/http"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestNewUserHandler(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want UserHandlerInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRegisterHandler(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want RegisterHandlerInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserRegisterHandler(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRegisterHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLoginHandler(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want LoginHandlerInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserLoginHandler(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserLoginHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginHandler_LoginUser(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *LoginHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.LoginUser(tt.args.w, tt.args.r)
		})
	}
}

func TestRegisterHandler_RegisterUser(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *RegisterHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.RegisterUser(tt.args.w, tt.args.r)
		})
	}
}

func TestUserHandler_UsersHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *UserHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.UsersHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestUserHandler_getUsersHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *UserHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.getUsersHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestUserHandler_getUsersByIDHandler(t *testing.T) {
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		id string
	}
	tests := []struct {
		name string
		h    *UserHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.getUsersByIDHandler(tt.args.w, tt.args.r, tt.args.id)
		})
	}
}

func TestUserHandler_createUsersHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *UserHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.createUsersHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestUserHandler_updateUserHandler(t *testing.T) {
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		id string
	}
	tests := []struct {
		name string
		h    *UserHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.updateUserHandler(tt.args.w, tt.args.r, tt.args.id)
		})
	}
}

func TestUserHandler_deleteUserHandler(t *testing.T) {
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		id string
	}
	tests := []struct {
		name string
		h    *UserHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.deleteUserHandler(tt.args.w, tt.args.r)
		})
	}
}
