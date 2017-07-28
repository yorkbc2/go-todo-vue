package main 


import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
	"encoding/json"
)

type User struct {
	ID int 
	Login string
	UKEY string
}

type CheckByKey struct {
	Check bool 
	UserID int
}

type Todo struct {
	ID int `json: "id"`
	Content string `json: "content"`
	Checked int `json: "checked"`
}

func CreateApiServer(db *sql.DB) {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/todo/create", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		userKey := r.PostFormValue("ukey")
		content := r.PostFormValue("content")

		check := CheckUserByKey(db, userKey)

		insert, err := db.Prepare("INSERT INTO `todos` SET id=?,user_id=?,content=?,checked=?")
		CheckError(err)

		smtm, err := insert.Exec(nil, check.UserID, content, 0)
		CheckError(err)

		id, err := smtm.LastInsertId()
		CheckError(err)

		var todo = Todo{ID: int(id), Content: content, Checked: 0}

		data, err := json.Marshal(todo)
		CheckError(err)

		w.Write(data)

		}).Methods("POST")

	router.HandleFunc("/signin", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		login := string(r.PostFormValue("login"))
		password := r.PostFormValue("password")

		login = strings.Trim(login, " ")


		rows,err := db.Query("SELECT * FROM `users` WHERE login='"+login+"'")
		CheckError(err)

		var forCheckup bool
		var passwordForCheckup []byte

		for rows.Next() {
			var passwordFromDatabase []byte
			var id int
			var ukey []byte

			err = rows.Scan(&id, &login, &passwordFromDatabase, &ukey)
			CheckError(err)

			if len(passwordFromDatabase) >= 1 {

				forCheckup = true
				passwordForCheckup = passwordFromDatabase

			} else {
				forCheckup = false
			}
			
		}

		if forCheckup == true {

			if err := bcrypt.CompareHashAndPassword(passwordForCheckup, []byte(password)); err != nil {

				w.Write([]byte("0"))

			} else {
				w.Write([]byte("1"))
			}

		}}).Methods("POST")

	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		login := r.PostFormValue("login")

		password := r.PostFormValue("password")

		rows,err := db.Query("SELECT * FROM `users` WHERE login='"+login+"'")
		CheckError(err)

		var forCheckup bool
		var user User

		for rows.Next() {
			var passwordFromDatabase []byte
			var id int
			var ukey []byte
			var dbLogin string

			err = rows.Scan(&id, &dbLogin, &passwordFromDatabase, &ukey)
			CheckError(err)

			if len(passwordFromDatabase) >= 1 {

				forCheckup = true

			} else {
				forCheckup = false
			}
			
		}

		if forCheckup == true {
			w.Write([]byte("0"))
			
		} else {

			hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			CheckError(err)

			keyStamp := time.Now().UnixNano() / int64(time.Millisecond)

			ukey, err := bcrypt.GenerateFromPassword([]byte(string(keyStamp)), bcrypt.DefaultCost)
			CheckError(err)


			insert, err := db.Prepare("INSERT users SET id=?,login=?,password=?,ukey=?")
			CheckError(err)

			smtm, err := insert.Exec(nil, login, string(hash), string(ukey))
			CheckError(err)

			id, err := smtm.LastInsertId()
			CheckError(err)

			user = User{ID: int(id), Login: login, UKEY: string(ukey)}

			data, _ := json.Marshal(user)

			w.Write(data)
		}



		}).Methods("POST")

	router.HandleFunc("/todos", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var u_id string 
		var allTodos []Todo

		u_id = r.PostFormValue("id")

		todos,err := db.Query("SELECT * FROM `todos` WHERE user_id='" + u_id + "' ORDER BY id DESC")
		CheckError(err)

		for todos.Next() {
			var id int
			var content string 
			var checked int 
			var user_id int 

			err = todos.Scan(&id, &user_id, &content, &checked)

			allTodos = append(allTodos, Todo{ID: id, Content: content, Checked: checked})
		}

		data, err := json.Marshal(allTodos)
		CheckError(err)

		w.Write(data)

	}).Methods("POST")

	fmt.Println("Server is running on port "+port)
	log.Fatal(http.ListenAndServe(port, router))
}

func CheckUserByKey(db *sql.DB, ukey string) CheckByKey {
	rows,err := db.Query("SELECT * FROM `users` WHERE ukey='"+ukey+"'")
	CheckError(err)

	var check bool
	var UserID int 

	for rows.Next() {
		var passwordFromDatabase []byte
		var id int
		var ukey []byte
		var dbLogin string

		err = rows.Scan(&id, &dbLogin, &passwordFromDatabase, &ukey)
		CheckError(err)


		if len(ukey) > 1 {
			check = true
			UserID = id
		} else {
			check = false
			UserID = 0
		}

	}

	return CheckByKey{UserID: UserID, Check: check}

}