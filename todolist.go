package main

import (
	"github.com/goadesign/goa"
	"github.com/anegawa-j/todoAppGo/app"
)

// TodolistController implements the todolist resource.
type TodolistController struct {
	*goa.Controller
}

// NewTodolistController creates a todolist controller.
func NewTodolistController(service *goa.Service) *TodolistController {
	return &TodolistController{Controller: service.NewController("TodolistController")}
}

// Show runs the show action.
func (c *TodolistController) Show(ctx *app.ShowTodolistContext) error {
	// TodolistController_Show: start_implement

	// Put your logic here

	res := &app.GoaExampleTodolist{}
	return ctx.OK(res)
	// TodolistController_Show: end_implement
}
