package service

import (
	"context"
	"github.com/arvians-id/go-todo-list/helper"
	"github.com/arvians-id/go-todo-list/internal/model"
	"github.com/arvians-id/go-todo-list/internal/repository"
)

type ActivityServiceContract interface {
	FindAll(ctx context.Context) ([]*model.Activity, error)
	FindByID(ctx context.Context, id int) (*model.Activity, error)
	Create(ctx context.Context, activity *model.Activity) (*model.Activity, error)
	Update(ctx context.Context, activity *model.Activity) (*model.Activity, error)
	Delete(ctx context.Context, id int) error
}

type ActivityService struct {
	ActivityRepository repository.ActivityRepository
}

func NewActivityService(activityRepository *repository.ActivityRepository) *ActivityService {
	return &ActivityService{ActivityRepository: *activityRepository}
}

func (service *ActivityService) FindAll(ctx context.Context) ([]*model.Activity, error) {
	return service.ActivityRepository.FindAll(ctx)
}

func (service *ActivityService) FindByID(ctx context.Context, id int) (*model.Activity, error) {
	return service.ActivityRepository.FindByID(ctx, id)
}

func (service *ActivityService) Create(ctx context.Context, activity *model.Activity) (*model.Activity, error) {
	return service.ActivityRepository.Create(ctx, activity)
}

func (service *ActivityService) Update(ctx context.Context, activity *model.Activity) (*model.Activity, error) {
	activityCheck, err := service.ActivityRepository.FindByID(ctx, activity.ActivityID)
	if err != nil {
		return nil, err
	}

	activityCheck.Title = activity.Title
	activityCheck.UpdatedAt, _ = helper.TimeNow()
	err = service.ActivityRepository.Update(ctx, activityCheck)
	if err != nil {
		return nil, err
	}

	return activityCheck, nil
}

func (service *ActivityService) Delete(ctx context.Context, id int) error {
	_, err := service.ActivityRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return service.ActivityRepository.Delete(ctx, id)
}
