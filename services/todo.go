package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"todoApi/models"
)

type TodoServiceMDB struct {
	db *mongo.Client
}

func NewTodoServiceviceMDB(db *mongo.Client) *TodoServiceMDB {
	return &TodoServiceMDB{
		db: db,
	}
}

type ToDoService interface {
	FindAll(ctx context.Context) ([]models.ToDo, error)
	Add(ctx context.Context, todo models.ToDo) error
	Delete(ctx context.Context, todo models.ToDo) error
	UpdateStatus(ctx context.Context, id string, status string) error
}

func (t *TodoServiceMDB) FindAll(ctx context.Context) ([]models.ToDo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TodoServiceMDB) Add(ctx context.Context, todo models.ToDo) error {
	//TODO implement me
	panic("implement me")
}

func (t *TodoServiceMDB) Delete(ctx context.Context, todo models.ToDo) error {
	//TODO implement me
	panic("implement me")
}

func (t *TodoServiceMDB) UpdateStatus(ctx context.Context, id string, status string) error {
	//TODO implement me
	panic("implement me")
}
