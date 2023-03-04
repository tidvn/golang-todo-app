package controllers

import (
	"strconv"
	"todo-app/models"

	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TodoController operations for Todo
type TodoController struct {
	beego.Controller
}

func (c *TodoController) Create() {

	status, _ := strconv.Atoi(c.GetString("Status"))
	P := &models.Todo{
		Status: status,
	}
	db := models.TodoData()
	id, err := db.AddTodo(P)

	if err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = id
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *TodoController) GetOne() {

}

func (c *TodoController) GetAll() {
	db := models.TodoData()
	list, err := db.GetAll()
	if err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = list
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

func (c *TodoController) Edit() {
	db := models.TodoData()
	idstr := c.Ctx.Input.Param(":id")
	title := c.GetString("Title")
	content := c.GetString("Content")
	status, _ := strconv.Atoi(c.GetString("Status"))
	if status > 2 || status < 0 {
		status = 0
	}
	id, _ := primitive.ObjectIDFromHex(idstr)
	u := &models.Todo{
		ID:      id,
		Title:   title,
		Content: content,
		Status:  status,
	}
	t, err := db.UpdateTodo(u)
	if err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = t
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()

}

func (c *TodoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	db := models.TodoData()
	err := db.DeleteTodo(idStr)
	if err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()

}
