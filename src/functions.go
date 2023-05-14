package main

import (
	"database/sql"
	"log"
	"modules/internal/db"
	"context"
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"

)

var loadStruct struct {
	User string `json:"user"`
	Email string `json:"email"`
	Is_Admin bool `json: is_admin`
	Password string `json:"password"`
}

var loginStruct struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func DatabaseConnect() *db.Queries{
	dbconn, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/CRUD_LOGIN")
	if err != nil {
		log.Fatal(err)
	}
	database := db.New(dbconn);
	return database;
}

func Logon(user string, email string, isAdmin bool, password string, createdAt string) {

	database := DatabaseConnect();
	newUser, err_ := database.GetUsers(context.Background(), db.GetUsersParams{Username: user, Email: email})
	if err_ != nil {
		fmt.Println(err_)
	}
	if (strings.EqualFold(newUser.Email, email)) {
		fmt.Println("A user with that email already exists")
	} else {
		database.InsertUser(context.Background(), db.InsertUserParams{Username: user, Email: email, IsAdmin: isAdmin, UserPassword: password, CreatedAt: createdAt});
		fmt.Println("Logon sucessfuly!")
	}
	
}

func VerifyAdmin(login string) bool {
	database := DatabaseConnect();
	user, err_ := database.IsAdmin(context.Background(), db.IsAdminParams{Username: login, Email: login})
	if err_ != nil {
		fmt.Println(err_)
	} 

	if!(user.Email == login || user.Username == login) {
		return false
	}

	return true
}


func VerifyUser(login, password string) bool{
	database := DatabaseConnect();
	user, err_ := database.GetUsers(context.Background(), db.GetUsersParams{Username: login, Email: login})
	if err_ != nil {
		fmt.Println(err_)
	}

	if(user.Email == login || user.Username == login) {
		if (user.UserPassword == password) {
			return true;
		} else {
			return false
		}
	} else {
		return false;
	}

	return false;
}
