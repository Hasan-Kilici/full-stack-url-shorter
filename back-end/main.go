package main

import (
	"urlShorter/Database"
	"github.com/gofiber/fiber/v2"
	"urlShorter/Routers"
)

func main(){
	Database.Connect()
	app := fiber.New()

	Routers.Initalize(app)

	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
}