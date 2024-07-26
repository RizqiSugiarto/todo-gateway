package controller

import (
	"context"

	taskPB "github.com/digisata/todo-gateway/stubs/task"
)

func (c AuthController) CreateTask(ctx context.Context, req *taskPB.CreateTaskRequest) (*taskPB.TaskBaseResponse, error) {
	res, err := c.TaskUseCase.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) GetTask(ctx context.Context, req *taskPB.GetTaskByIDRequest) (*taskPB.GetTaskByIDResponse, error) {
	res, err := c.TaskUseCase.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) GetAllTask(ctx context.Context, req *taskPB.GetAllTaskByActivityIDRequest) (*taskPB.GetAllTaskByActivityIDResponse, error) {
	res, err := c.TaskUseCase.GetAllByUserID(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) UpdateTask(ctx context.Context, req *taskPB.UpdateTaskByIDRequest) (*taskPB.TaskBaseResponse, error) {
	res, err := c.TaskUseCase.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) BatchUpdateTask(ctx context.Context, req *taskPB.BatchUpdateTaskRequest) (*taskPB.TaskBaseResponse, error) {
	res, err := c.TaskUseCase.BatchUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) DeleteTask(ctx context.Context, req *taskPB.DeleteTaskByIDRequest) (*taskPB.TaskBaseResponse, error) {
	res, err := c.TaskUseCase.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
