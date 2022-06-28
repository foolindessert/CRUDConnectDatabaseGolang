package repo

import (
	entity "DATABASECRUD/Entity"
	"database/sql"
)

func QueryGetSocialMedia(db *sql.DB) []*entity.ResponseSocialMediaGet {
	sqlStament := `
	select distinct on (s.id)s.id,s.name,s.social_media_url,s.user_id,u.created_date,u.updated_date,u.id,u.username,p.url from social_media s left join users u on s.user_id = u.id left join photos p on u.id = p.user_id `
	rows, err := db.Query(sqlStament)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	socialmedias := []*entity.ResponseSocialMediaGet{}
	for rows.Next() {
		var socialmedia entity.ResponseSocialMediaGet
		if serr := rows.Scan(&socialmedia.Id, &socialmedia.Name, &socialmedia.Social_Media_Url, &socialmedia.User_id, &socialmedia.CreatedAt, &socialmedia.UpdatedAt, &socialmedia.User.Id, &socialmedia.User.Username, &socialmedia.User.Url); serr != nil {
			panic(serr)
		}
		socialmedias = append(socialmedias, &socialmedia)
	}
	return socialmedias
}

func QueryPostSocialMedia(db *sql.DB, newSocialMedia entity.SocialMedia, user_id int) entity.ResponseSocialMediaPost {
	sqlStament := `insert into social_media
			(name,social_media_url,user_id)
			values ($1,$2,$3) Returning id`
	// intId, err := strconv.Atoi(id)
	err = db.QueryRow(sqlStament, newSocialMedia.Name, newSocialMedia.Social_Media_Url, user_id).Scan(&newSocialMedia.Id)
	if err != nil {
		panic(err)
	}

	response := entity.ResponseSocialMediaPost{}
	sqlstatment2 := `
			select s.id,s.name,s.social_media_url,s.user_id,u.created_date from social_media s left join users u on s.user_id = u.id where s.id = $1`
	err = db.QueryRow(sqlstatment2, newSocialMedia.Id).
		Scan(&response.Id, &response.Name, &response.Social_Media_Url, &response.User_id, &response.CreatedAt)
	if err != nil {
		panic(err)
	}
	return response
}

func QueryUpdateSocialMedia(db *sql.DB, newSocialMedia entity.SocialMedia, id string) entity.ResponseSocialMediaPut {
	sqlStament := `update social_media set name = $1, social_media_url= $2 where id = $3`
	//query.scan
	_, err = db.Exec(sqlStament,
		newSocialMedia.Name,
		newSocialMedia.Social_Media_Url,
		id)
	if err != nil {
		panic(err)
	}
	response := entity.ResponseSocialMediaPut{}
	sqlstatment2 := `select s.id,s.name,s.social_media_url,s.user_id,u.updated_date from social_media s left join users u on s.user_id = u.id where s.id = $1`
	err = db.QueryRow(sqlstatment2, id).
		Scan(&response.Id, &response.Name, &response.Social_Media_Url, &response.User_id, &response.UpdatedAt)
	// count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return response
}

func QueryDeleteSocialMedia(db *sql.DB, id string) entity.Message {
	sqlstament := `DELETE from social_media where id = $1 ;`
	_, err := db.Exec(sqlstament, id)

	if err != nil {
		panic(err)
	}
	message := entity.Message{
		Message: "Your SocialMedia has been successfully deleted",
	}
	return message
}
