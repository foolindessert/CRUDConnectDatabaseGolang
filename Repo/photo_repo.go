package repo

import (
	entity "DATABASECRUD/Entity"
	"database/sql"
	"fmt"
	"time"
)

func QueryGetPhoto(db *sql.DB) []*entity.ResponsePhotoGet {
	sqlStament := `
	select p.id,p.title,p.caption,p.url,p.user_id,p.created_date,p.updated_date,u.email,u.username from photos p left join users u on p.user_id = u.id`
	rows, err := db.Query(sqlStament)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	photos := []*entity.ResponsePhotoGet{}
	for rows.Next() {
		var photo entity.ResponsePhotoGet
		if serr := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.Url, &photo.User_id, &photo.CreatedAt, &photo.UpdatedAt, &photo.Users.Email, &photo.Users.Username); serr != nil {
			fmt.Println("Scan error", serr)
		}
		photos = append(photos, &photo)
	}
	return photos
}

func QueryPostPhoto(db *sql.DB, newPhotos entity.Photo, user_id int) entity.ResponsePostPhoto {
	sqlStament := `insert into photos
	(title,caption,url,user_id,created_date,updated_date)
	values ($1,$2,$3,$4,$5,$5) Returning id`
	//query.scan
	err = db.QueryRow(sqlStament,
		newPhotos.Title,
		newPhotos.Caption,
		newPhotos.Url,
		user_id,
		time.Now(),
	).Scan(&newPhotos.Id)
	if err != nil {
		panic(err)
	}
	response := entity.ResponsePostPhoto{
		Id:        newPhotos.Id,
		Title:     newPhotos.Title,
		Caption:   newPhotos.Caption,
		Url:       newPhotos.Url,
		User_id:   int(user_id),
		CreatedAt: time.Now(),
	}
	return response
}

func QueryUpdatePhoto(db *sql.DB, newPhotos entity.Photo, id string) entity.ResponsePuPhoto {
	sqlStament := `update photos set title = $1, caption = $2 , url = $3, updated_date =$4 where id = $5`
	//query.scan
	_, err = db.Exec(sqlStament,
		newPhotos.Title,
		newPhotos.Caption,
		newPhotos.Url,
		time.Now(),
		id)
	if err != nil {
		fmt.Println("error update")
		panic(err)
	}
	sqlstatment2 := `select id,title,caption,url,user_id,created_date,updated_date from photos where id= $1`
	err = db.QueryRow(sqlstatment2, id).
		Scan(&newPhotos.Id, &newPhotos.Title, &newPhotos.Caption, &newPhotos.Url, &newPhotos.User_id, &newPhotos.CreatedAt, &newPhotos.UpdatedAt)

	if err != nil {
		panic(err)
	}

	response := entity.ResponsePuPhoto{
		Id:        newPhotos.Id,
		Title:     newPhotos.Title,
		Caption:   newPhotos.Caption,
		Url:       newPhotos.Url,
		User_id:   newPhotos.User_id,
		UpdatedAt: newPhotos.UpdatedAt,
	}

	return response

}

func QueryDeletePhoto(db *sql.DB, id string,) entity.Message{
	sqlstament := `DELETE from photos where id = $1  ;`
	_, err := db.Exec(sqlstament, id)

	if err != nil {
		panic(err)
	}
	message := entity.Message{
		Message: "Your photo has been successfully deleted",
	}
	return message
}
