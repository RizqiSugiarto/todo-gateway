package controller

import (
	"context"

	"github.com/digisata/auth-service/domain"
	"github.com/digisata/auth-service/stubs"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c AuthController) GetProfile(ctx context.Context, req *emptypb.Empty) (*stubs.AuthBaseResponse, error) {
	claims := ctx.Value("claims")

	data, err := c.ProfileUseCase.GetByID(ctx, claims.(jwt.MapClaims)["id"].(string))
	if err != nil {
		return nil, err
	}

	getProfileResponse := &stubs.GetProfileResponse{
		Id:         data.ID.Hex(),
		CustomerId: data.CustomerID,
		Name:       data.Name,
		Email:      data.Email,
		CreatedAt:  int32(data.CreatedAt),
		UpdatedAt:  int32(data.UpdatedAt),
		DeletedAt:  int32(data.DeletedAt),
	}

	anyData, err := anypb.New(getProfileResponse)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &stubs.AuthBaseResponse{
		Message: "Success",
		Data:    anyData,
	}

	return res, nil
}

func (c AuthController) ChangePassword(ctx context.Context, req *stubs.ChangePasswordRequest) (*stubs.AuthBaseResponse, error) {
	changePasswordRequest := domain.ChangePasswordRequest{
		OldPassword: req.GetOldPassword(),
		NewPassword: req.GetNewPassword(),
	}

	err := c.ProfileUseCase.ChangePassword(ctx, changePasswordRequest)
	if err != nil {
		return nil, err
	}

	res := &stubs.AuthBaseResponse{
		Message: "Success",
	}

	return res, nil
}
