package controller

import (
	"context"

	activityPB "github.com/digisata/todo-gateway/stubs/activity"
)

func (c AuthController) CreateActivity(ctx context.Context, req *activityPB.CreateActivityRequest) (*activityPB.ActivityBaseResponse, error) {
	res, err := c.ActivityUseCase.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) GetActivity(ctx context.Context, req *activityPB.GetActivityByIDRequest) (*activityPB.ActivityBaseResponse, error) {
	res, err := c.ActivityUseCase.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) GetAllActivity(ctx context.Context, req *activityPB.GetAllActivityRequest) (*activityPB.GetAllActivityResponse, error) {
	res, err := c.ActivityUseCase.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) UpdateActivity(ctx context.Context, req *activityPB.UpdateActivityByIDRequest) (*activityPB.ActivityBaseResponse, error) {
	res, err := c.ActivityUseCase.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) DeleteActivity(ctx context.Context, req *activityPB.DeleteActivityByIDRequest) (*activityPB.ActivityBaseResponse, error) {
	res, err := c.ActivityUseCase.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
