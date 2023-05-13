package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"modules/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)


func register(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := loadStruct
	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}
	Logon(payload.User, payload.Email, payload.Password)
	
	jsonModel, err := json.Marshal(payload);
	if err != nil {
		log.Fatal(err)
	}
	return c.SendStatus(c.Response().StatusCode());

	return c.JSON(jsonModel);

	
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
		fmt.Println("User deleted!")
	} else {
		fmt.Println("Error!")
	}
	return c.SendStatus(c.Response().StatusCode());
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

	database := DatabaseConnect();
	if(VerifyUser(payload.User, payload.Password)) {
		database.UpdateUser(context.Background(), db.UpdateUserParams{Username: payload.NewUser, Email: payload.NewEmail, UserPassword: payload.NewPassword, Username_2: payload.User, Email_2: payload.Email})
		fmt.Println("User updated!")
	} else {
		fmt.Println("Error!")
	}

	
	return c.SendStatus(c.Response().StatusCode());

	jsonModel, err := json.Marshal(payload);
	if err != nil {
		log.Fatal(err)
	}
	
	return c.JSON(jsonModel);
}


func login(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := loginStruct

	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}
	if(VerifyUser(payload.Login, payload.Password)) {
		fmt.Println("Login sucessfully!")
	}

	return c.SendStatus(c.Response().StatusCode());
	jsonModel, err := json.Marshal(payload);
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(jsonModel);
}


func main() {
	app := fiber.New()
	app.Post("/register", register)
	app.Get("/login", login)
    app.Put("/update", updateUser)
	app.Delete("/delete", deleteUser)
	app.Listen(":8080")
}