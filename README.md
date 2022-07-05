# Final Project Hacktiv8 Golang Traine

Final Project Scalable Web Service with Golang Hacktiv8


* #### Endpoint List : 
    * ##### User : 
        * #### Register
            
            [POST]```http://localhost:8080/users/register```
            
            body :

            ```json
            {
                "age": 23,
                "email":"desril@gmail.com",
                "password":"desril",
                "username":"desril"
            }
            ```

            response
            ```json
            {
                "data": {
                    "id": 1,
                    "age": 21,
                    "email": "desril@gmail.com",
                    "password": "$2a$10$cvpW1zR8RXkG5VBosoBJ/./kXKaO7pKXmzaLfUgsE6rU61TxqEJvi",
                    "username": "desril",
                    "date": "2022-06-27T13:06:37.558+07:00"
                }
            }
            ```
            
        * #### Login
            [POST]```http://localhost:8080/users/login```
            
            body :

            ```json
            {
                "email": "desril@mailz.com",
                "password": "desril"
            }
            ```

            response
            ```json
            {
                "data": {
                   "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRlc3JpbEBtYWlsei5jb20iLCJpZCI6MX0._CTowajg3nwUKU4qn4qJzP02pOmQBzzYcaZTEinmQs8"
                }
            }
            ```
            
       * #### Update User
            [PUT]```http://localhost:8080/users/1```
            
            Authorization :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```

            contoh :
            ```
            {
                Authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjoyNH0.XpOW4v9hpneBw9gnsVAGli_zDqj7VmMLW6ZHL80MauQ"
            }
            ```

            body :

            ```json
            {
                "email" : "desril12@mailz.com",
                "username":"desril12"
            }  
            ```

            response
            ```json
            {
                "status": "ok",
                "data": {
                          "id": 1,
                          "email": "desril12@mailz.com",
                          "username": "desril12",
                          "age": 22,
                          "updated_at": "2022-06-29T22:01:39.6471884+07:00"
                        }
            }
            ```
        
        * #### Delete User
            [DELETE]```http://localhost:8080/users```
            
            Authorization :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```

            token didapatkan ketika melalui proses login

            contoh :
            ```
            {
                Authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjoyNH0.XpOW4v9hpneBw9gnsVAGli_zDqj7VmMLW6ZHL80MauQ"
            }
            ```            

            response
            ```json
            {
                  {
                          "message": "Your account has been successfully deleted"
                  }
            }
            ```
            
   * ##### Photos : 
        * #### Post Photos
            
            [POST]```http://localhost:8080/photos```
            
            Authorization :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```
            
            body :

            ```json
            {
                "title":"Hard Work 2",
                "caption":"self motivation",
                "photo_url":"self/toself"
            }
            ```

            response
            ```json
            {
                    "id": 21,
                    "title": "Hard Work 2",
                    "caption": "self motivation",
                    "photo_url": "self/toself",
                    "user_id": 6,
                    "created_at": "2022-06-30T08:57:56.5840838+07:00"
            }
            ```
            
        * #### Get Photos
            [GET]```http://localhost:8080/photos```
            
            Authorization :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```

            response
            ```json
            [
                {
                     "id": 1,
                     "title": "walau",
                     "caption": "hore",
                     "photo_url": "kkkk/kkkk",
                     "user_id": 6,
                     "created_at": "2022-06-24T00:00:00Z",
                     "updated_at": "2022-06-24T00:00:00Z",
                     "User": {
                        "email": "a3@gmail.com",
                        "username": "a3"
                     }
                },
            ]
            ```
            
       * #### Update Phtotos
            [PUT]```http://localhost:8080/photos/{id}```
            
            Authorization :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```

            body :

            ```json
            {
                "title":"why 2",
                "caption":"renungan",
                "photo_url":"sad/tosad" 
            }  
            ```

            response
            ```json
            {
                "id": 19,
               "title": "why 2",
               "caption": "renungan",
               "photo_url": "sad/tosad",
               "user_id": 6,
               "updated_at": "2022-06-30T00:00:00Z"
            }
            ```
        
        * #### Delete Photos
            [DELETE]```http://localhost:8080/photos/{id}```
            
            Authorization :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```           

            response
            ```json
            {
                  {
                          "message": "Your photo has been successfully deleted"
                  }
            }
            ```
  
   * ##### Comment : 
        * #### Post Comments
            
            [POST]```http://localhost:8080/comments```
            
            Authorization :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```
            
            body :

            ```json
            {
                "message":"FIGHTING!!!!",
                "photo_id": 19
            }
            ```

            response
            ```json
            {
                     "id": 10,
                     "message": "FIGHTING!!!!",
                     "photo_id": 19,
                     "user_id": 8,
                     "created_at": "2022-06-30T09:07:30.2262943+07:00"
            }
            ```
            
        * #### Get Comment
            [GET]```http://localhost:8080/comments```
            
            Authorization :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```

            response
            ```json
            [
                {
                    "id": 3,
                    "user_id": 6,
                    "photo_id": 4,
                    "message": "kurang bagus",
                    "created_at": "2022-06-24T00:00:00Z",
                    "updated_at": "2022-06-27T00:00:00Z",
                    "User": {
                        "id": 6,
                        "email": "a3@gmail.com",
                        "username": "a3"
                    },
                    "Photo": {
                        "id": 4,
                        "title": "asdasd",
                        "caption": "hore 11",
                        "photo_url": "ddddd",
                        "user_id": 6
                    }
                },
            ]
            ```
            
       * #### Update Comments
            [PUT]```http://localhost:8080/comments/{id}```
            
            Authorization :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```

            body :

            ```json
            {
                "message":"bagus" 
            }  
            ```

            response
            ```json
            {
                  "id": 8,
                  "title": "why 2",
                  "caption": "renungan",
                  "photo_url": "sad/tosad",
                  "user_id": 8,
                  "updated_at": "2022-06-30T00:00:00Z"
            }
            ```
        
        * #### Delete Comment
            [DELETE]```http://localhost:8080/comments/{id}```
            
            Authorization :

            ```
            {
                Authorization: "Bearer {{token}}"
            }
            ```           

            response
            ```json
            {
                  {
                          "message": "Your comment has been successfully deleted"
                  }
            }
            ```
