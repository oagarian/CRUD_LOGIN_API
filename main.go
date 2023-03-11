package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
	"reflect"
	"strings"
)

func handler(c *fiber.Ctx) error {
	fmt.Fprintf(c, "Arroba")
	return nil
}

type Account struct {
	User string `json:"User"`
	Email string `json:"Email"`
	Password string `json:"Password"`
}

var Users []Account

func logon(user, email, password string) {
	newUser := Account{
		User: user, 
		Email: email, 
		Password: password,
	}
	Users = append(Users, newUser)
}

func verifyUser(user, email, password string) bool {
	for _, u := range Users {
		values := reflect.ValueOf(u)
		typesOf := values.Type()
		for i := 0; i < values.NumField(); i++ {
			if (strings.EqualFold(typesOf.Field(i).Name, "User") || strings.EqualFold(typesOf.Field(i).Name, "Email")) {
				valueOfFields := fmt.Sprintf("%s", values.Field(i).Interface())
				var valueOfPassword string
				if (strings.EqualFold(typesOf.Field(i).Name, "Password")) {
					valueOfPassword = fmt.Sprintf("%s", values.Field(i).Interface())
				}
				if(!strings.EqualFold(valueOfFields, user)){
					if(valueOfPassword == password) {
						return false;
					}
				}
			} else {
				return false;
			}
		}

	}
	return true;
	
}

func main() {
	app := fiber.New()
	app.Get("/", handler)
	app.Listen(":8080")
}