package main

import (
	"github.com/gofiber/fiber/v2"
)

// creamos un struct para trabajar
type Usuario struct {
	Nombre   string
	Apellido string
}

// funcion
func handleUsuario(c *fiber.Ctx) error {
	usuario := Usuario{
		Nombre:   "Mauro",
		Apellido: "Bogado",
	}
	return c.Status(fiber.StatusOK).JSON(usuario)
}

// funcion de creacion
func CrearUsuario(c *fiber.Ctx) error {
	//usamos la varible para asignar
	usuario := Usuario{}
	//parseamos los valores
	if err := c.BodyParser(&usuario); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(usuario)
}
func main() {

	//creacion de la var
	appli := fiber.New()

	//usamos metodo get para escuhar
	appli.Get("/", func(a *fiber.Ctx) error {
		return a.SendString("Hola Usuario de API en Fiber")
	})
	appli.Get("/usuario", handleUsuario)
	//usamos el metodo post
	appli.Post("/usuarios", CrearUsuario)

	//puerto en donde va a escuchar
	appli.Listen(":3000")
}
