package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-todo-list/helper"
	"github.com/arvians-id/go-todo-list/internal/model"
)

type TodoRepositoryContract interface {
	FindAll(ctx context.Context, activityGroupID int) ([]*model.Todo, error)
	FindByID(ctx context.Context, id int) (*model.Todo, error)
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	Update(ctx context.Context, todo *model.Todo) error
	Delete(ctx context.Context, id int) error
}

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (repository *TodoRepository) FindAll(ctx context.Context, activityGroupID int) ([]*model.Todo, error) {
	query := "SELECT * FROM todos WHERE activity_group_id = ?"
	if activityGroupID == 0 {
		activityGroupID = 1
		query = "SELECT * FROM todos WHERE ?"
	}
	stmt, err := repository.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, activityGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.TodoID, &todo.ActivityGroupID, &todo.Title, &todo.Priority, &todo.IsActive, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	return todos, nil
}

func (repository *TodoRepository) FindByID(ctx context.Context, id int) (*model.Todo, error) {
	query := `SELECT * FROM todos WHERE todo_id = ?`
	var todo model.Todo
	err := repository.DB.QueryRowContext(ctx, query, id).Scan(&todo.TodoID, &todo.ActivityGroupID, &todo.Title, &todo.Priority, &todo.IsActive, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (repository *TodoRepository) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	query := `INSERT INTO todos (activity_group_id, title, priority, is_active) VALUES (?, ?, ?, ?)`
	result, err := repository.DB.ExecContext(ctx, query, todo.ActivityGroupID, todo.Title, todo.Priority, todo.IsActive)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return todo, err
	}

	todo.TodoID = int(id)

	todo.CreatedAt, _ = helper.TimeNow()
	todo.UpdatedAt, _ = helper.TimeNow()

	return todo, nil
}

func (repository *TodoRepository) Update(ctx context.Context, todo *model.Todo) error {
	query := `UPDATE todos SET title = ?, is_active = ?, priority = ?, updated_at = ? WHERE todo_id = ?`
	_, err := repository.DB.ExecContext(ctx, query, todo.Title, todo.IsActive, todo.Priority, todo.UpdatedAt, todo.TodoID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *TodoRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM todos WHERE todo_id = ?`
	_, err := repository.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
