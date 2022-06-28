package repo

import (
	entity "DATABASECRUD/Entity"
	"database/sql"
	"time"
)

func QueryGetComment(db *sql.DB) []*entity.ResponseCommentGet {
	sqlStament := `
	select c.id, c.message,c.photo_id,c.user_id,c.updated_date,c.created_date,u.id,u.email,u.username,p.id,p.title,p.caption,p.url,p.user_id from comment c left join photos p on c.photo_id = p.id left join users u on c.user_id = u.id`
	rows, err := db.Query(sqlStament)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	comments := []*entity.ResponseCommentGet{}
	for rows.Next() {
		var comment entity.ResponseCommentGet
		if serr := rows.Scan(&comment.Id, &comment.Message, &comment.Photo_id, &comment.User_id, &comment.UpdatedAt, &comment.CreatedAt, &comment.User.Id, &comment.User.Email, &comment.User.Username, &comment.Photo.Id, &comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.Url, &comment.Photo.User_id); serr != nil {
			panic(serr)
		}
		comments = append(comments, &comment)
	}
	return comments
}

func QueryPostComment(db *sql.DB, newComment entity.Commment, user_id int) entity.ResponseCommentPost {
	sqlStament := `insert into comment
	(user_id,photo_id,message,created_date,updated_date)
	values ($1,$2,$3,$4,$4) Returning id`
	// intId, err := strconv.Atoi(id)
	err = db.QueryRow(sqlStament,
		user_id,
		newComment.Photo_id,
		newComment.Message,
		time.Now()).Scan(&newComment.Id)
	if err != nil {
		panic(err)
	}
	response := entity.ResponseCommentPost{
		Id:        newComment.Id,
		Message:   newComment.Message,
		Photo_id:  newComment.Photo_id,
		User_id:   int(user_id),
		CreatedAt: time.Now(),
	}
	return response
}

func QueryUpdateComment(db *sql.DB, newComment entity.Commment, id string) entity.ResponseUpdateComment {
	sqlStament := `update comment set message = $1, updated_date =$2 where id = $3`
	//query.scan
	_, err = db.Exec(sqlStament,
		newComment.Message,
		time.Now(),
		id,
	)
	if err != nil {
		panic(err)
	}
	response := entity.ResponseUpdateComment{}
	sqlstatment2 := `select c.id,p.title,p.caption,p.url,c.user_id,c.updated_date from comment c left join photos p on c.photo_id = p.id where c.id= $1`
	err = db.QueryRow(sqlstatment2, id).
		Scan(&response.Id, &response.Title, &response.Caption, &response.Url, &response.User_id, &response.UpdatedAt)
	if err != nil {
		panic(err)
	}
	return response
}

func QueryDeleteComment(db *sql.DB, id string) entity.Message {
	sqlstament := `DELETE from comment where id = $1 ;`
	_, err := db.Exec(sqlstament, id)

	if err != nil {
		panic(err)
	}
	message := entity.Message{
		Message: "Your Comment has been successfully deleted",
	}
	return message
}
