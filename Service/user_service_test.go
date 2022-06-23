package service

import (
	entity "DATABASECRUD/Entity"
	"reflect"
	"testing"
	"time"

	_ "github.com/golang-jwt/jwt/v4"
)

func TestNewPhotoSvc(t *testing.T) {
	tests := []struct {
		name string
		want PhotoIface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPhotoSvc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPhotoSvc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserSvc(t *testing.T) {
	tests := []struct {
		name string
		want UserIface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserSvc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserSvc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSvc_Register(t *testing.T) {
	type args struct {
		user *entity.User
	}
	tests := []struct {
		name    string
		u       *UserSvc
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			u:    &UserSvc{},
			args: args{
				user: &entity.User{
					Id:        1,
					Username:  "asd",
					Email:     "asd@gmail.com",
					Password:  "asdasdasdasd",
					Age:       23,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			want: &entity.User{
				Id:        1,
				Username:  "asd",
				Email:     "asd@gmail.com",
				Password:  "asdasdasdasd",
				Age:       23,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.Register(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserSvc.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserSvc.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSvc_Login(t *testing.T) {
	type args struct {
		user         *entity.User
		tempPassword string
	}
	tests := []struct {
		name    string
		u       *UserSvc
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			u:    &UserSvc{},
			args: args{
				user:         &entity.User{},
				tempPassword: "",
			},
			want:    &entity.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.Login(tt.args.user, tt.args.tempPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserSvc.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserSvc.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSvc_UpdateUser(t *testing.T) {
	type args struct {
		user *entity.User
	}
	tests := []struct {
		name    string
		u       *UserSvc
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			u:       &UserSvc{},
			args:    args{},
			want:    &entity.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.UpdateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserSvc.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserSvc.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSvc_GetToken(t *testing.T) {
	type args struct {
		id       uint
		email    string
		password string
	}
	tests := []struct {
		name string
		u    *UserSvc
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.GetToken(tt.args.id, tt.args.email, tt.args.password); got != tt.want {
				t.Errorf("UserSvc.GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSvc_CheckToken(t *testing.T) {
	type args struct {
		compareToken string
		id           uint
		email        string
		password     string
	}
	tests := []struct {
		name    string
		u       *UserSvc
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.CheckToken(tt.args.compareToken, tt.args.id, tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("UserSvc.CheckToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserSvc_VerivyToken(t *testing.T) {
	type args struct {
		TempToken string
	}
	tests := []struct {
		name string
		u    *UserSvc
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.VerivyToken(tt.args.TempToken); got != tt.want {
				t.Errorf("UserSvc.VerivyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
