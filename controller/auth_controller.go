package controller

import (
	"context"

	"github.com/digisata/auth-service/domain"
	"github.com/digisata/auth-service/stubs"
	invitationCategoryPB "github.com/digisata/auth-service/stubs/invitation-category"
	socialPB "github.com/digisata/auth-service/stubs/social"
	"github.com/digisata/auth-service/usecase"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthController struct {
	stubs.UnimplementedAuthServiceServer
	UserUseCase               UserUseCase
	ProfileUseCase            ProfileUseCase
	SocialUseCase             socialPB.SocialServiceClient
	InvitationCategoryUseCase invitationCategoryPB.InvitationCategoryServiceClient
}

var _ UserUseCase = (*usecase.UserUseCase)(nil)
var _ ProfileUseCase = (*usecase.ProfileUseCase)(nil)

func (c AuthController) Verify(ctx context.Context, req *emptypb.Empty) (*stubs.AuthBaseResponse, error) {
	claims := ctx.Value("claims")
	var customerID *string

	customerIDClaim, ok := claims.(jwt.MapClaims)["customer_id"]
	if ok {
		value := customerIDClaim.(string)
		customerID = &value
	}

	verifyResponse := &stubs.VerifyResponse{
		Name:       claims.(jwt.MapClaims)["name"].(string),
		Id:         claims.(jwt.MapClaims)["id"].(string),
		Role:       int32(claims.(jwt.MapClaims)["role"].(float64)),
		CustomerId: customerID,
	}

	anyData, err := anypb.New(verifyResponse)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &stubs.AuthBaseResponse{
		Message: "Success",
		Data:    anyData,
	}

	return res, nil
}

func (c AuthController) LoginAdmin(ctx context.Context, req *stubs.LoginRequest) (*stubs.AuthBaseResponse, error) {
	payload := domain.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	data, err := c.UserUseCase.LoginAdmin(ctx, payload)
	if err != nil {
		return nil, err
	}

	loginResponse := &stubs.LoginResponse{
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
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

func (c AuthController) LoginCustomer(ctx context.Context, req *stubs.LoginRequest) (*stubs.AuthBaseResponse, error) {
	payload := domain.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	data, err := c.UserUseCase.LoginCustomer(ctx, payload)
	if err != nil {
		return nil, err
	}

	loginResponse := &stubs.LoginResponse{
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
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

func (c AuthController) LoginCommittee(ctx context.Context, req *stubs.LoginRequest) (*stubs.AuthBaseResponse, error) {
	payload := domain.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	data, err := c.UserUseCase.LoginCommittee(ctx, payload)
	if err != nil {
		return nil, err
	}

	loginResponse := &stubs.LoginResponse{
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
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

func (c AuthController) RefreshToken(ctx context.Context, req *stubs.RefreshTokenRequest) (*stubs.AuthBaseResponse, error) {
	refreshTokenRequest := domain.RefreshTokenRequest{
		RefreshToken: req.GetRefreshToken(),
	}

	data, err := c.UserUseCase.RefreshToken(ctx, refreshTokenRequest)
	if err != nil {
		return nil, err
	}

	refreshTokenResponse := &stubs.RefreshTokenResponse{
		AccessToken: data.AccessToken,
	}

	anyData, err := anypb.New(refreshTokenResponse)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &stubs.AuthBaseResponse{
		Message: "Success",
		Data:    anyData,
	}

	return res, nil
}

func (c AuthController) Logout(ctx context.Context, req *stubs.LogoutRequest) (*stubs.AuthBaseResponse, error) {
	err := c.UserUseCase.Logout(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	res := &stubs.AuthBaseResponse{
		Message: "Success",
	}

	return res, nil
}
