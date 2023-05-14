package main

import (
	"context"
	"fmt"
	"log"
	"modules/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"time"
)


func register(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := loadStruct
	currentTime := time.Now().Format("2006-01-02")
	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}

	Logon(payload.User, payload.Email, payload.Is_Admin, payload.Password, currentTime)
	
	return c.SendStatus(c.Response().StatusCode());
	
}

func deleteUser(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := loginStruct
	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}

	database := DatabaseConnect();
	if(VerifyUser(payload.Login, payload.Password)) {
		database.DeleteUser(context.Background(), db.DeleteUserParams{Username: payload.Login, Email: payload.Login})
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("FAILED")
	}
	return c.SendStatus(c.Response().StatusCode());
}

func updateUser(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")

	payload := struct {
		User string `json:"user"`
		Password string `json:"password"`
		NewUser string `json:"new_user"`
		NewEmail string `json:"new_email"`
		NewPassword string `json:"new_password"`
	}{} 

	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}

	database := DatabaseConnect();
	if(VerifyUser(payload.User, payload.Password)) {
		fmt.Println(payload)
		database.UpdateUser(context.Background(), db.UpdateUserParams{Username: payload.NewUser, Email: payload.NewEmail, UserPassword: payload.NewPassword, Username_2: payload.User, Email_2: payload.User})
				
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("FAILED")
	}

	return c.SendStatus(c.Response().StatusCode());
}


func login(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := loginStruct

	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}
	if(!VerifyUser(payload.Login, payload.Password)) {
		return c.Status(fiber.StatusBadRequest).SendString("FAILED")
	}

	
	return c.SendStatus(c.Response().StatusCode());
	
}


func main() {
	app := fiber.New()
	app.Post("/register", register)
	app.Get("/login", login)
    	app.Put("/update", updateUser)
	app.Delete("/delete", deleteUser)
	app.Listen(":8080")
}