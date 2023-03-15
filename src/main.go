package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"modules/internal/db"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)



func handler(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	
	fmt.Fprintf(c, "Arroba")
	return nil
}

var loadStruct struct {
	User string
	Email string
	Password string
}

var loginStruct struct {
	Login string
	Password string
}

type Account struct {
	User string `json:"user"`
	Email string `json:"email"`
	Password string `json:"password"`
}

var Users []Account

func databaseConnect() *db.Queries{
	dbconn, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/CRUD_LOGIN")
	if err != nil {
		log.Fatal(err)
	}
	database := db.New(dbconn);
	return database;
}

func logon(user, email, password string) {

	database := databaseConnect();
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


func verifyUser(login, password string) bool{
	database := databaseConnect();
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

func register(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := loadStruct
	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}
	logon(payload.User, payload.Email, payload.Password)
	return nil
	
}

func deleteUser(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := loginStruct
	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}

	database := databaseConnect();
	if(verifyUser(payload.Login, payload.Password)) {
		database.DeleteUser(context.Background(), db.DeleteUserParams{payload.Login, payload.Login})
		fmt.Println("User deleted!")
	} else {
		fmt.Println("Error!")
	}
	return nil
}

func updateUser(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := struct {
		User string
		Email string
		Password string
		NewUser string
		NewEmail string
		NewPassword string
	}{}
	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}

	database := databaseConnect();
	if(verifyUser(payload.User, payload.Password)) {
		database.UpdateUser(context.Background(), db.UpdateUserParams{payload.NewUser, payload.NewEmail, payload.NewPassword, payload.User, payload.Email})
		fmt.Println("User updated!")
	} else {
		fmt.Println("Error!")
	}
	return nil
}


func login(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := loginStruct

	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}
	if(verifyUser(payload.Login, payload.Password)) {
		fmt.Println("Login sucessfully!")
	}

	return nil
}


func main() {
	app := fiber.New()
	app.Get("/", handler)
	app.Post("/register", register)
	app.Get("/login", login)
    app.Put("/update", updateUser)
	app.Delete("/delete", deleteUser)
	app.Listen(":8080")
}