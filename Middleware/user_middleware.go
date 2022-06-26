package middleware

import (
	entity "DATABASECRUD/Entity"
	service "DATABASECRUD/Service"
	"context"
	"fmt"
	"net/http"
	"strings"
)

var tempKey = &tempcontext{"user"}

type tempcontext struct {
	data string
}

func AuthCekToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serv := service.NewUserSvc()
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		// fmt.Println(reqToken)
		fmt.Println("Lolos token")
		temp_id, err := serv.VerivyToken(reqToken)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		}
		fmt.Print("nilai token id")
		fmt.Println(temp_id)
		user := entity.User{Id: int(temp_id)}

		ctx := context.WithValue(r.Context(), tempKey, &user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}

func ForUser(ctx context.Context) *entity.User {
	temp, _ := ctx.Value(tempKey).(*entity.User)
	return temp
}
