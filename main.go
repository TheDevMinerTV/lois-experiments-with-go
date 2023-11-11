package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"test/models"
	"test/templates"

	"github.com/gofiber/fiber/v2"
)

//go:embed style.css
var Styles []byte

var Todos = []models.Todo{
	{
		ID:      1,
		Content: "Buy milk",
	},
	{
		ID:      2,
		Content: "Buy eggs",
	},
	{
		ID:      3,
		Content: "Buy bread",
	},
	{
		ID:      4,
		Content: "Buy jam",
	},
}

func getTodoByID(route string, id string) *models.Todo {
	for _, t := range Todos {
		if fmt.Sprintf("%d", t.ID) == id {
			return &t
		}
	}

	return nil
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return templates.Index("Lois", Todos).Render(c.UserContext(), c)
	})

	app.Get("/style.css", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/css")
		return c.Send(Styles)
	})

	// /edit/:id
	app.Get("/edit/:id", func(c *fiber.Ctx) error {
		id := c.Params("id", "-1")
		todo := getTodoByID("/edit/:id", id)
		if todo == nil {
			c.Status(http.StatusNotFound)
			// TODO: templ
			c.Set("Content-Type", "text/html")
			return c.SendString("<b>Not found</b>")
		}

		c.Set("Content-Type", "text/html")
		return templates.EditTodo(*todo).Render(c.UserContext(), c)
	})

	app.Get("/save/", func(c *fiber.Ctx) error {
		id := c.Params("id", "-1")
		todo := getTodoByID("/save/:id", id)
		if todo == nil {
			c.Status(http.StatusNotFound)
			// TODO: templ
			c.Set("Content-Type", "text/html")
			return c.SendString("<b>Not found</b>")
		}

		c.Set("Content-Type", "text/html")
		return templates.Todo(*todo).Render(c.UserContext(), c)
	})

	app.Delete("/delete/:id", func(c *fiber.Ctx) error {
		id := c.Params("id", "-1")
		todo := getTodoByID("/delete/:id", id)
		if todo == nil {
			c.Status(http.StatusNotFound)
			c.Set("Content-Type", "text/html")
			return c.SendString("<b>Not found</b>")
		}

		// FIXME: replace with SQL
		// "DELETE FROM todo WHERE ID = ?"
		var newTodos []models.Todo
		for _, t := range Todos {
			if t.ID != todo.ID {
				newTodos = append(newTodos, t)
			}
		}
		Todos = newTodos

		return refreshHtmx(c)
	})

	fmt.Println("Listening on port 80")
	if err := app.Listen(":80"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// https://htmx.org/docs/#response-headers
func refreshHtmx(c *fiber.Ctx) error {
	c.Set("HX-Refresh", "true")
	return c.Send(nil)
}
