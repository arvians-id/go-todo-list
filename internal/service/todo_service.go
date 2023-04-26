package service

import (
	"context"
	"github.com/arvians-id/go-todo-list/helper"
	"github.com/arvians-id/go-todo-list/internal/model"
	"github.com/arvians-id/go-todo-list/internal/repository"
)

type TodoServiceContract interface {
	FindAll(ctx context.Context, activityGroupID int) ([]*model.Todo, error)
	FindByID(ctx context.Context, id int) (*model.Todo, error)
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	Update(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	Delete(ctx context.Context, id int) error
}

type TodoService struct {
	TodoRepository     repository.TodoRepository
	ActivityRepository repository.ActivityRepository
}

func NewTodoService(todoRepository *repository.TodoRepository, activityRepository *repository.ActivityRepository) *TodoService {
	return &TodoService{
		TodoRepository:     *todoRepository,
		ActivityRepository: *activityRepository,
	}
}

func (service *TodoService) FindAll(ctx context.Context, activityGroupID int) ([]*model.Todo, error) {
	return service.TodoRepository.FindAll(ctx, activityGroupID)
}

func (service *TodoService) FindByID(ctx context.Context, id int) (*model.Todo, error) {
	return service.TodoRepository.FindByID(ctx, id)
}

func (service *TodoService) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	_, err := service.ActivityRepository.FindByID(ctx, todo.ActivityGroupID)
	if err != nil {
		return nil, err
	}

	return service.TodoRepository.Create(ctx, todo)
}

func (service *TodoService) Update(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	todoCheck, err := service.TodoRepository.FindByID(ctx, todo.TodoID)
	if err != nil {
		return nil, err
	}

	todoCheck.Title = todo.Title
	todoCheck.Priority = todo.Priority
	todoCheck.IsActive = todo.IsActive
	todoCheck.UpdatedAt, _ = helper.TimeNow()

	err = service.TodoRepository.Update(ctx, todoCheck)
	if err != nil {
		return nil, err
	}

	return todoCheck, nil
}

func (service *TodoService) Delete(ctx context.Context, id int) error {
	_, err := service.TodoRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return service.TodoRepository.Delete(ctx, id)
}
