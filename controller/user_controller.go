package controller

import (
	"context"

	"github.com/digisata/auth-service/domain"
	"github.com/digisata/auth-service/pkg/constants"
	"github.com/digisata/auth-service/stubs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

func (c AuthController) CreateUser(ctx context.Context, req *stubs.CreateUserRequest) (*stubs.AuthBaseResponse, error) {
	user := domain.User{
		ID:         primitive.NewObjectID(),
		CustomerID: req.CustomerId,
		Name:       req.GetName(),
		Email:      req.GetEmail(),
		Password:   req.GetPassword(),
		Role:       int8(req.GetRole()),
		IsActive:   req.GetIsActive(),
		Note:       req.Note,
	}

	if user.Role == int8(constants.COMMITTEE) && user.CustomerID == nil {
		return nil, status.Error(codes.InvalidArgument, "please choose the customer")
	}

	err := c.UserUseCase.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &stubs.AuthBaseResponse{
		Message: "Success",
	}

	return res, nil
}

func (c AuthController) GetAllUser(ctx context.Context, req *stubs.GetAllUserRequest) (*stubs.GetAllUserResponse, error) {
	filter := domain.GetAllUserRequest{
		Search:   req.Search,
		IsActive: req.IsActive,
		Role:     req.Role,
	}

	users, err := c.UserUseCase.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	res := &stubs.GetAllUserResponse{
		Message: "Success",
	}
	for _, user := range users {
		data := &stubs.GetUserByIDResponse{
			Id:         user.ID.Hex(),
			CustomerId: user.CustomerID,
			Name:       user.Name,
			Email:      user.Email,
			Role:       int32(user.Role),
			IsActive:   user.IsActive,
			Note:       user.Note,
			CreatedAt:  int32(user.CreatedAt),
			UpdatedAt:  int32(user.UpdatedAt),
			DeletedAt:  int32(user.DeletedAt),
		}

		res.Data = append(res.Data, data)
	}

	return res, nil
}

func (c AuthController) GetUserByID(ctx context.Context, req *stubs.GetUserByIDRequest) (*stubs.AuthBaseResponse, error) {
	data, err := c.UserUseCase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	getUserByIDResponse := &stubs.GetUserByIDResponse{
		Id:         data.ID.Hex(),
		CustomerId: data.CustomerID,
		Name:       data.Name,
		Email:      data.Email,
		Role:       int32(data.Role),
		IsActive:   data.IsActive,
		Note:       data.Note,
		CreatedAt:  int32(data.CreatedAt),
		UpdatedAt:  int32(data.UpdatedAt),
		DeletedAt:  int32(data.DeletedAt),
	}

	anyData, err := anypb.New(getUserByIDResponse)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &stubs.AuthBaseResponse{
		Message: "Success",
		Data:    anyData,
	}

	return res, nil
}

func (c AuthController) UpdateUser(ctx context.Context, req *stubs.UpdateUserRequest) (*stubs.AuthBaseResponse, error) {
	idHex, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	user := domain.UpdateUser{
		ID:         idHex,
		CustomerID: req.CustomerId,
		Name:       req.Name,
		IsActive:   req.IsActive,
		Note:       req.Note,
	}

	err = c.UserUseCase.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &stubs.AuthBaseResponse{
		Message: "Success",
	}

	return res, nil
}

func (c AuthController) DeleteUser(ctx context.Context, req *stubs.DeleteUserRequest) (*stubs.AuthBaseResponse, error) {
	idHex, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	user := domain.DeleteUser{
		ID: idHex,
	}

	err = c.UserUseCase.Delete(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &stubs.AuthBaseResponse{
		Message: "Success",
	}

	return res, nil
}
