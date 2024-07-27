package controller

import (
	"context"

	"github.com/digisata/todo-gateway/domain"
	"github.com/digisata/todo-gateway/stubs"
	activityPB "github.com/digisata/todo-gateway/stubs/activity"
	taskPB "github.com/digisata/todo-gateway/stubs/task"
	textPB "github.com/digisata/todo-gateway/stubs/text"
	"github.com/digisata/todo-gateway/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthController struct {
	stubs.UnimplementedAuthServiceServer
	UserUseCase     UserUseCase
	TaskUseCase     taskPB.TaskServiceClient
	ActivityUseCase activityPB.ActivityServiceClient
	TextUseCase     textPB.TextServiceClient
}

var _ UserUseCase = (*usecase.UserUseCase)(nil)

func (c AuthController) Login(ctx context.Context, req *stubs.LoginRequest) (*stubs.AuthBaseResponse, error) {
	payload := domain.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	data, err := c.UserUseCase.Login(ctx, payload)
	if err != nil {
		return nil, err
	}

	loginResponse := &stubs.LoginResponse{
		AccessToken: data.AccessToken,
	}

	anyData, err := anypb.New(loginResponse)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &stubs.AuthBaseResponse{
		Message: "Success",
		Data:    anyData,
	}

	return res, nil
}

func (c AuthController) Logout(ctx context.Context, req *emptypb.Empty) (*stubs.AuthBaseResponse, error) {
	err := c.UserUseCase.Logout(ctx)
	if err != nil {
		return nil, err
	}

	res := &stubs.AuthBaseResponse{
		Message: "Success",
	}

	return res, nil
}
