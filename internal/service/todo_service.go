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
	TodoRepository repository.TodoRepository
}

func NewTodoService(todoRepository *repository.TodoRepository) *TodoService {
	return &TodoService{
		TodoRepository: *todoRepository,
	}
}

func (service *TodoService) FindAll(ctx context.Context, activityGroupID int) ([]*model.Todo, error) {
	return service.TodoRepository.FindAll(ctx, activityGroupID)
}

func (service *TodoService) FindByID(ctx context.Context, id int) (*model.Todo, error) {
	return service.TodoRepository.FindByID(ctx, id)
}

func (service *TodoService) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	return service.TodoRepository.Create(ctx, todo)
}

func (service *TodoService) Update(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	todoCheck, err := service.TodoRepository.FindByID(ctx, todo.TodoID)
	if err != nil {
		return nil, err
	}

	if todo.Priority != "" {
		todoCheck.Priority = todo.Priority
	}

	oldTitle := todoCheck.Title
	if todoCheck.Title != todo.Title {
		todoCheck.Title = todo.Title
	}
	if todo.Title == "" {
		todoCheck.Title = oldTitle
	}

	todoCheck.UpdatedAt, _ = helper.TimeNow()

	if todoCheck.IsActive != todo.IsActive {
		todoCheck.IsActive = todo.IsActive
	}

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
