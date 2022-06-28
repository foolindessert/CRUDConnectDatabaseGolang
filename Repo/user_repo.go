package repo

import (
	entity "DATABASECRUD/Entity"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var (
	db  *sql.DB
	err error
)

func QueryRegisterUser(db *sql.DB, newUser entity.User) entity.ResponseRegister {

	sqlStatment := `insert into users
			(username,email,password,age,created_date,updated_date)
			values ($1,$2,$3,$4,$5,$5) Returning id` //sesuai dengan nama table
	err := db.QueryRow(sqlStatment,
		newUser.Username,
		newUser.Email,
		newUser.Password,
		newUser.Age,
		time.Now(),
	).Scan(&newUser.Id)
	if err != nil {
		panic(err)
	} else {
		// fmt.Println(newUser)
		// fmt.Println(newUser.Id)
		response_Register := entity.ResponseRegister{
			Age:      newUser.Age,
			Email:    newUser.Email,
			Id:       newUser.Id,
			Username: newUser.Username,
		}
		return response_Register
	}

}

func QueryLoginUser(db *sql.DB, newUser entity.User) (entity.User, error) {
	sqlStatment := `select id,username,email,password,age,created_date,updated_date from public.users where email = $1`

	err := db.QueryRow(sqlStatment, newUser.Email).
		Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password, &newUser.Age, &newUser.CreatedAt, &newUser.UpdatedAt)

	if err != nil {
		return entity.User{}, errors.New("username cannot be empty")
	}
	return newUser, nil
}

func QueryUpdateUser(db *sql.DB, newUser entity.User, id string) entity.ResponseUpdateUser {
	sqlstatment := `
		update users set username = $1, email = $2, updated_date = $3
		where id = $4;`

	_, err = db.Exec(sqlstatment,
		newUser.Username,
		newUser.Email,
		time.Now(),
		id,
	)
	if err != nil {
		fmt.Println("error update")
		panic(err)

	}
	sqlstatment2 := `select * from users where id= $1`
	err = db.QueryRow(sqlstatment2, id).
		Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password, &newUser.Age, &newUser.CreatedAt, &newUser.UpdatedAt)

	if err != nil {
		panic(err)
	}
	fmt.Println(newUser)
	responseUpdateUser := entity.ResponseUpdateUser{
		Id:        newUser.Id,
		Email:     newUser.Email,
		Username:  newUser.Username,
		Age:       newUser.Age,
		UpdatedAt: time.Now(),
	}
	return responseUpdateUser
}

func QueryDeleteUser(db *sql.DB, newUser *entity.User) entity.Message {
	sqlstament := `DELETE from users where id = $1;`
	_, err := db.Exec(sqlstament, newUser.Id)

	if err != nil {
		panic(err)
	}
	message := entity.Message{
		Message: "Your account has been successfully deleted",
	}
	return message

}
