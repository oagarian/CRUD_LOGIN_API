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
	User string
	Email string
	Password string
}

var loginStruct struct {
	Login string
	Password string
}

func DatabaseConnect() *db.Queries{
	dbconn, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/CRUD_LOGIN")
	if err != nil {
		log.Fatal(err)
	}
	database := db.New(dbconn);
	return database;
}

func Logon(user, email, password string) {

	database := DatabaseConnect();
	newUser, err_ := database.GetUsers(context.Background(), db.GetUsersParams{user, email})
	if err_ != nil {
		fmt.Println(err_)
	}
	if (strings.EqualFold(newUser.Email, email)) {
		fmt.Println("A user with that email already exists")
	} else {
		fmt.Println(newUser.Email)
		fmt.Println(email)
		database.InsertUser(context.Background(), db.InsertUserParams{user, email, password});
		fmt.Println("Logon sucessfuly!")
	}
	
}


func VerifyUser(login, password string) bool{
	database := DatabaseConnect();
	user, err_ := database.GetUsers(context.Background(), db.GetUsersParams{login, login})
	if err_ != nil {
		fmt.Println(err_)
	}

	if(user.Email == login || user.Username == login) {
		if (user.UserPassword == password) {
			return true;
		} else {
			fmt.Println("Login failed!")
			return false
		}
	} else {
		fmt.Println("Login failed!")
		return false;
	}
	return false;
}
