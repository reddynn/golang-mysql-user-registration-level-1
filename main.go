package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Dbconnect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1)/userdb")

	if err != nil {
		panic(err.Error())

	}

	return db

}
func main() {
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/signin", Signin)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

}

func Homepage(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("welcome"))
}

func Signup(w http.ResponseWriter, r *http.Request) {
	// rurl := "localhost:3000"
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	var user string
	dbs := Dbconnect()
	defer dbs.Close()
	err := dbs.QueryRow("select username from users where username=?", username).Scan(&user)
	switch {
	case err == sql.ErrNoRows:
		hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "unable to create your account", 500)
			return
		}
		_, err = dbs.Query("insert into users(username,password) values(?,?)", username, hashedpassword)
		if err != nil {
			http.Error(w, "server error unable to create user", 500)
			fmt.Println(err)
			return
		}
	case err != nil:
		http.Error(w, "server error, unable to create you account", 500)
		return
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

func Signin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	// var dbusername string
	var dbpassword string
	db := Dbconnect()
	defer db.Close()
	err := db.QueryRow("select password from users where username=?", username).Scan(&dbpassword)

	switch {

	case err != nil:
		http.Error(w, "unauthorized no user with this username", http.StatusUnauthorized)
		return
	default:
		err := bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(password))
		if err != nil {
			http.Error(w, "unauthorized password wrong", http.StatusUnauthorized)
			break
		}
		w.Write([]byte("hello" + username))
	
	}

}
