package main

import (
	"fmt"
    "log"
	"github.com/gofiber/fiber/v2"

)

func handler(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	
	fmt.Fprintf(c, "Arroba")
	return nil
}

type Account struct {
	User string `json:"user"`
	Email string `json:"email"`
	Password string `json:"password"`
}

var Users []Account

func logon(user, email, password string) {
	newUser := Account{
		User: user, 
		Email: email, 
		Password: password,
	}
	Users = append(Users, newUser)
	fmt.Println("Logon sucessfuly!")
}

func register(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := struct {
        User  string 
        Email string 
		Password string 
    }{}

	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}

	logon(payload.User, payload.Email, payload.Password)
	fmt.Println(payload.User)
	fmt.Println(payload.Email)
	fmt.Println(payload.Password)

	return nil
	
}


func verifyUser(user, email, password string) bool {
	
}

func login(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := struct {
        User  string 
        Email string 
		Password string 
    }{}

	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}

	if (verifyUser(payload.User, payload.Email, payload.Password)) {
		fmt.Println("Login sucessfuly! - pt2")
	} else {
		fmt.Println("Login failed!")
	}
	return nil
}


func main() {
	app := fiber.New()
	app.Get("/", handler)
	app.Post("/register", register)
	app.Get("/login")
	app.Listen(":8080")
}