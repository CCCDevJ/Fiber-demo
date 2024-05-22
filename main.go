package main

import (
	"fiberdemo/database"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// RenderForm renders the HTML form.
func RenderForm(c *fiber.Ctx) error {
	books, err := database.GetAllBookList()
	if err != nil {
		log.Fatal("GetAllBookList ERROR => ", err)
	}
	return c.Render("form", fiber.Map{"bookList": books})
}

// ProcessForm processes the form submission.
func ProcessForm(c *fiber.Ctx) error {
	name := c.FormValue("name")
	greeting := fmt.Sprintf("Hello, %s!", name)
	return c.Render("greeting", fiber.Map{"Greeting": greeting})
}

func GetAllBook(c *fiber.Ctx) error {

	books, err := database.GetAllBookList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	// log.Println("book data All: ", books)
	return c.Status(fiber.StatusOK).JSON(books)
}

func main() {
	database.ConnectDb()
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	// Serve static files (HTML templates and stylesheets).
	app.Static("/", "./static")

	// Define routes.
	app.Get("/", RenderForm)
	app.Post("/submit", ProcessForm)
	app.Get("/api/v1/books", GetAllBook)

	// Start the Fiber app on port 8080.
	app.Listen(":3080")
}
