package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"todoApi/services"
)

type TodoHandler struct {
	todoService services.ToDoService
}

func NewToDoHandler(db *mongo.Client) *TodoHandler {
	return &TodoHandler{
		todoService: services.NewTodoServiceviceMDB(db),
	}
}

func (h *TodoHandler) GetToDo(ctx *gin.Context) {
	return
}

func (h *TodoHandler) GetAllToDo(ctx *gin.Context) {
	ctx.JSONP(200, "GetAllToDo")
	return
}

func (h *TodoHandler) AddNewToDo(ctx *gin.Context) {
	return
}

func (h *TodoHandler) DeleteToDo(ctx *gin.Context) {
	return
}

func (h *TodoHandler) UpdateToDo(ctx *gin.Context) {
	return
}
