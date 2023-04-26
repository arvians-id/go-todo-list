package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-todo-list/helper"
	"github.com/arvians-id/go-todo-list/internal/model"
)

type ActivityRepositoryContract interface {
	FindAll(ctx context.Context) ([]*model.Activity, error)
	FindByID(ctx context.Context, id int) (*model.Activity, error)
	Create(ctx context.Context, activity *model.Activity) (*model.Activity, error)
	Update(ctx context.Context, activity *model.Activity) error
	Delete(ctx context.Context, id int) error
}

type ActivityRepository struct {
	DB *sql.DB
}

func NewActivityRepository(db *sql.DB) *ActivityRepository {
	return &ActivityRepository{DB: db}
}

func (repository *ActivityRepository) FindAll(ctx context.Context) ([]*model.Activity, error) {
	query := `SELECT * FROM activities`
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []*model.Activity
	for rows.Next() {
		var activity model.Activity
		err := rows.Scan(&activity.ActivityID, &activity.Title, &activity.Email, &activity.CreatedAt, &activity.UpdatedAt)
		if err != nil {
			return nil, err
		}
		activities = append(activities, &activity)
	}

	return activities, nil
}

func (repository *ActivityRepository) FindByID(ctx context.Context, id int) (*model.Activity, error) {
	query := `SELECT * FROM activities WHERE activity_id = ?`
	var activity model.Activity
	err := repository.DB.QueryRowContext(ctx, query, id).Scan(&activity.ActivityID, &activity.Title, &activity.Email, &activity.CreatedAt, &activity.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func (repository *ActivityRepository) Create(ctx context.Context, activity *model.Activity) (*model.Activity, error) {
	query := `INSERT INTO activities (title, email) VALUES (?, ?)`
	result, err := repository.DB.ExecContext(ctx, query, activity.Title, activity.Email)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	activity.ActivityID = int(id)

	activity.CreatedAt, _ = helper.TimeNow()
	activity.UpdatedAt, _ = helper.TimeNow()

	return activity, nil
}

func (repository *ActivityRepository) Update(ctx context.Context, activity *model.Activity) error {
	query := `UPDATE activities SET title = ?, updated_at = ? WHERE activity_id = ?`
	_, err := repository.DB.ExecContext(ctx, query, activity.Title, activity.UpdatedAt, activity.ActivityID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ActivityRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM activities WHERE activity_id = ?`
	_, err := repository.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
