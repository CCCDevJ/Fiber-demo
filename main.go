package main

import (
	"fiberdemo/database"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

// RenderForm renders the HTML form.
func RenderForm(c *fiber.Ctx) error {
	session, err := Store.Get(c)
	if err != nil {
		println(err)
	}
	email := session.Get(EMAIL)
	fmt.Println(email)
	if email == nil {
		return c.Render("login", fiber.Map{})
	}

	books, err := database.GetAllBookList()
	if err != nil {
		log.Fatal("GetAllBookList ERROR => ", err)
	}
	return c.Render("form", fiber.Map{"bookList": books})
}

// RenderLoginForm renders the HTML form.
func RenderLoginForm(c *fiber.Ctx) error {
	// books, err := database.GetAllBookList()
	// if err != nil {
	// 	log.Fatal("GetAllBookList ERROR => ", err)
	// }
	return c.Render("login", fiber.Map{})
}

// Process processes the form submission for Login
func ProcessLoginCheck(c *fiber.Ctx) error {
	session, err := Store.Get(c)
	if err != nil {
		println(err)
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "devjethava909@gmail.com" && password == "Dev@123" {

		session.Set(EMAIL, email)

		fmt.Println(session.Get("email"))

		greeting := fmt.Sprintf("Hello, %s!", email)

		err = session.Save()
		if err != nil {
			fmt.Println(err)
		}
		return c.Render("greeting", fiber.Map{"Greeting": greeting})
	}
	// println(email, password)
	// greeting := fmt.Sprintf("Hello, %s!", name)
	return c.Render("login", fiber.Map{
		"Error": "Username/password incorrect",
	})
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

var (
	Store    *session.Store
	AUTH_KEY string = "authenticated"
	EMAIL    string = "email"
)

func main() {
	database.ConnectDb()

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	/* Sessions Config */
	Store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, for https
		Expiration: time.Hour * 1,
	})

	// Serve static files (HTML templates and stylesheets).
	app.Static("/", "./static")

	// Define routes.
	app.Get("/", RenderForm)
	app.Get("/login", RenderLoginForm)
	app.Post("/loginCheck", ProcessLoginCheck)
	app.Post("/submit", ProcessForm)
	app.Get("/api/v1/books", GetAllBook)

	// Start the Fiber app on port 8080.
	app.Listen(":3080")
}
