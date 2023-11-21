package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"test/ent"
	"test/models"
)

//go:embed tailwind.css
var Styles []byte

var Todos = map[int]models.Todo{
	1: {
		ID:      1,
		Content: "Buy milk",
	},
	2: {
		ID:      2,
		Content: "Buy eggs",
	},
	3: {
		ID:      3,
		Content: "Buy bread",
	},
	4: {
		ID:      4,
		Content: "Buy jam",
	},
}

func main() {
	Host := os.Getenv("POSTGRES_HOST")
	Port := os.Getenv("POSTGRES_PORT")
	User := os.Getenv("POSTGRES_USER")
	DbName := os.Getenv("POSTGRES_DB")
	Password := os.Getenv("POSTGRES_PASSWORD")
	Connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", Host, Port, User, DbName, Password)

	client, err := ent.Open("postgres", Connection)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)

	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// 	app := fiber.New()

	// 	app.Get("/", func(c *fiber.Ctx) error {
	// 		c.Set("Content-Type", "text/html")
	// 		return templates.Index("Lois", Todos).Render(c.UserContext(), c)
	// 	})

	// 	app.Get("/style.css", func(c *fiber.Ctx) error {
	// 		c.Set("Content-Type", "text/css")
	// 		return c.Send(Styles)
	// 	})

	// 	// /edit/:id
	// 	app.Get("/edit/:id", func(c *fiber.Ctx) error {
	// 		id, err := c.ParamsInt("id", -1)
	// 		if err != nil {
	// 			return c.SendStatus(http.StatusBadRequest)
	// 		}

	// 		todo, ok := Todos[id]
	// 		if !ok {
	// 			c.Status(http.StatusNotFound)
	// 			// TODO: templ
	// 			c.Set("Content-Type", "text/html")
	// 			return c.SendString("<b>Not found</b>")
	// 		}

	// 		c.Set("Content-Type", "text/html")
	// 		return templates.EditTodo(todo).Render(c.UserContext(), c)
	// 	})
	// 	app.Post("/edit/:id", func(c *fiber.Ctx) error {
	// 		id, err := c.ParamsInt("id", -1)
	// 		if err != nil {
	// 			return c.SendStatus(http.StatusBadRequest)
	// 		}

	// 		todo, ok := Todos[id]
	// 		if !ok {
	// 			c.Status(http.StatusNotFound)
	// 			// TODO: templ
	// 			c.Set("Content-Type", "text/html")
	// 			return c.SendString("<b>Not found</b>")
	// 		}

	// 		content := c.FormValue("content")

	// 		Todos[id] = models.Todo{
	// 			ID:      todo.ID,
	// 			Content: content,
	// 		}

	// 		return refreshHtmx(c)
	// 	})

	// 	app.Post("/add/", func(c *fiber.Ctx) error {
	// 		content := c.FormValue("content")

	// 		// "INSERT INTO todo (content) VALUES (?)"
	// 		// pls database
	// 		id := len(Todos) + 1
	// 		Todos[id] = models.Todo{
	// 			ID:      id,
	// 			Content: content,
	// 		}

	// 		return refreshHtmx(c)
	// 	})

	// 	app.Delete("/delete/:id", func(c *fiber.Ctx) error {
	// 		id, err := c.ParamsInt("id", -1)
	// 		if err != nil {
	// 			return c.SendStatus(http.StatusBadRequest)
	// 		}

	// 		todo, ok := Todos[id]
	// 		if !ok {
	// 			c.Status(http.StatusNotFound)
	// 			c.Set("Content-Type", "text/html")
	// 			return c.SendString("<b>Not found</b>")
	// 		}

	// 		// FIXME: replace with SQL
	// 		// "DELETE FROM todo WHERE ID = ?"
	// 		var newTodos = make(map[int]models.Todo)
	// 		for k, v := range Todos {
	// 			if k != todo.ID {
	// 				newTodos[k] = v
	// 			}
	// 		}
	// 		Todos = newTodos

	// 		return refreshHtmx(c)
	// 	})

	// 	fmt.Println("Listening on port 80")
	// 	if err := app.Listen(":80"); err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}
	// }

	// // https://htmx.org/docs/#response-headers
	//
	//	func refreshHtmx(c *fiber.Ctx) error {
	//		c.Set("HX-Refresh", "true")
	//		return c.Send(nil)
}
