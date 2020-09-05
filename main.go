package main

import (
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

// Todo struct the data
type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{ID: 1, Name: "Walk the dog", Completed: false},
	{ID: 2, Name: "Walk the cat", Completed: false},
	{ID: 3, Name: "Walk the bat", Completed: false},
	{ID: 4, Name: "Walk the rat", Completed: false},
}

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello world")
	})
	setupAPIV1(app)

	app.Listen(3000)
}

func setupAPIV1(app *fiber.App) {
	v1 := app.Group("/v1")
	setupTodoRoutes(v1)
}

func setupTodoRoutes(grp fiber.Router) {
	todoRoutes := grp.Group("/todos")
	todoRoutes.Get("/", GetTodos)
	todoRoutes.Get("/:id", GetTodo)
	todoRoutes.Post("/", CreateTodo)
	todoRoutes.Delete("/:id", DeleteTodo)
	todoRoutes.Patch("/:id", UpdateTodo)
}

// UpdateTodo will update the todo item
func UpdateTodo(ctx *fiber.Ctx) {
	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}

	paramID := ctx.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}

	var body request
	err = ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse body",
		})
		return
	}

	var todo *Todo

	for _, t := range todos {
		if t.ID == id {
			todo = t
			break
		}
	}

	if todo == nil {
		ctx.Status(fiber.StatusNotFound)
		return
	}

	if body.Name != nil {
		todo.Name = *body.Name
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}
}

// DeleteTodo will delete the todo that specify id
func DeleteTodo(ctx *fiber.Ctx) {
	paramID := ctx.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[0:i], todos[i+1:]...)
			ctx.Status(fiber.StatusNoContent)
			return
		}
	}

	ctx.Status(fiber.StatusNotFound)
}

// CreateTodo will create the new todo
func CreateTodo(ctx *fiber.Ctx) {
	type request struct {
		Name string `json:"name"`
	}

	var body request

	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}

	todo := &Todo{
		ID:        len(todos) + 1,
		Name:      body.Name,
		Completed: false,
	}

	todos = append(todos, todo)

	ctx.Status(fiber.StatusCreated).JSON(todos)
}

// GetTodo will send each todo by its id
func GetTodo(ctx *fiber.Ctx) {
	paramID := ctx.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	for _, todo := range todos {
		if todo.ID == id {
			ctx.Status(fiber.StatusOK).JSON(todo)
			return
		}
	}

}

// GetTodos will send the all todos
func GetTodos(ctx *fiber.Ctx) {
	ctx.Status(fiber.StatusOK).JSON(todos)
}
