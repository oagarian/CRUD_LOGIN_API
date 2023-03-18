package main

import (
	"context"
	"fmt"
	"log"
	"modules/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)


func handler(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	
	fmt.Fprintf(c, "Arroba")
	return nil
}


func register(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := loadStruct
	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}
	Logon(payload.User, payload.Email, payload.Password)
	return nil
	
}

func deleteUser(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	payload := loginStruct
	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}

	database := DatabaseConnect();
	if(VerifyUser(payload.Login, payload.Password)) {
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

	database := DatabaseConnect();
	if(VerifyUser(payload.User, payload.Password)) {
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
	if(VerifyUser(payload.Login, payload.Password)) {
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