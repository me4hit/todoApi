package usecases

import (
	"context"
	"todoApi/models"
	"todoApi/services"
)

type todoUsecase struct {
	TodoService services.ToDoService
}

type ToDoUseCase interface {
	GetAllToDo(ctx context.Context) ([]models.ToDo, error)
	AddNewTodo(ctx context.Context, todo models.ToDo) error
	DeleteToDo(ctx context.Context, id string) error
	UpdateStatusToDo(ctx context.Context, id string, status string) error
}

func (t *todoUsecase) GetAllToDo(ctx context.Context) ([]models.ToDo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoUsecase) AddNewTodo(ctx context.Context, todo models.ToDo) error {
	//TODO implement me
	panic("implement me")
}

func (t *todoUsecase) DeleteToDo(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (t *todoUsecase) UpdateStatusToDo(ctx context.Context, id string, status string) error {
	//TODO implement me
	panic("implement me")
}
