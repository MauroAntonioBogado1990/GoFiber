package main

import (
	

	"github.com/gofiber/fiber/v2"
)

func main() {
	
	//creacion de la var
	appli := fiber.New()
	
	//usamos metodo get para escuhar
	appli.Get("/user", func(a *fiber.Ctx) error {
		return a.SendString("Hola Usuario de API")
	})
	//puerto en donde va a escuchar
	appli.Listen(":3000")
}
