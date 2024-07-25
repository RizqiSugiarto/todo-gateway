package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/digisata/auth-service/bootstrap"
	"github.com/digisata/auth-service/domain"
	"github.com/digisata/auth-service/helper"
	"github.com/digisata/auth-service/pkg/jwtio"
	memcachedRepo "github.com/digisata/auth-service/repository/memcached"
	mongoRepo "github.com/digisata/auth-service/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserUseCase struct {
	jwt     *jwtio.JSONWebToken
	cfg     *bootstrap.Config
	ur      UserRepository
	cr      CacheRepository
	timeout time.Duration
}

var _ UserRepository = (*mongoRepo.UserRepository)(nil)
var _ CacheRepository = (*memcachedRepo.CacheRepository)(nil)

func NewUserUseCase(jwt *jwtio.JSONWebToken, cfg *bootstrap.Config, ur UserRepository, cr CacheRepository, timeout time.Duration) *UserUseCase {
	return &UserUseCase{
		jwt:     jwt,
		cfg:     cfg,
		ur:      ur,
		cr:      cr,
		timeout: timeout,
	}
}

func (uc UserUseCase) generateToken(user domain.User) (domain.AuthResponse, error) {
	var (
		res        domain.AuthResponse
		customerID string
	)

	if user.CustomerID != nil {
		customerID = *user.CustomerID
	}

	payload := jwtio.Payload{
		ID:         user.ID.Hex(),
		CustomerID: &customerID,
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
	}

	now := time.Now()

	accessToken, err := uc.jwt.CreateAccessToken(payload, uc.cfg.Jwt.AccessTokenSecret, now)
	if err != nil {
		return res, err
	}

	err = uc.cr.Set(domain.CacheItem{
		Key: helper.HashKey(accessToken),
		Exp: int32(now.Add(time.Minute * time.Duration(uc.cfg.Jwt.AccessTokenExpiryMinute)).Unix()),
	})
	if err != nil {
		return res, status.Error(codes.Internal, err.Error())
	}

	refreshToken, err := uc.jwt.CreateRefreshToken(payload, uc.cfg.Jwt.RefreshTokenSecret, now)
	if err != nil {
		return res, err
	}

	err = uc.cr.Set(domain.CacheItem{
		Key: helper.HashKey(refreshToken),
		Exp: int32(now.Add(time.Minute * time.Duration(uc.cfg.Jwt.RefreshTokenExpiryMinute)).Unix()),
	})
	if err != nil {
		return res, status.Error(codes.Internal, err.Error())
	}

	res = domain.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return res, nil
}

func (uc UserUseCase) LoginAdmin(ctx context.Context, req domain.User) (domain.AuthResponse, error) {
	var res domain.AuthResponse
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	user, err := uc.ur.GetByEmail(ctx, req.Email)
	if err != nil {
		return res, status.Error(codes.InvalidArgument, "Incorrect email or password")
	}

	if user.Role != int8(domain.ADMIN) {
		return res, status.Error(codes.InvalidArgument, "Incorrect email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return res, status.Error(codes.InvalidArgument, "Incorrect email or password")
	}

	if !user.IsActive || user.DeletedAt != 0 {
		return res, status.Error(codes.Unauthenticated, "Your account has been deleted")
	}

	res, err = uc.generateToken(user)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) LoginCustomer(ctx context.Context, req domain.User) (domain.AuthResponse, error) {
	var res domain.AuthResponse
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	user, err := uc.ur.GetByEmail(ctx, req.Email)
	if err != nil {
		return res, status.Error(codes.InvalidArgument, "Incorrect email or password")
	}

	if user.Role != int8(domain.CUSTOMER) {
		return res, status.Error(codes.InvalidArgument, "Incorrect email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return res, status.Error(codes.InvalidArgument, "Incorrect email or password")
	}

	if !user.IsActive || user.DeletedAt != 0 {
		return res, status.Error(codes.Unauthenticated, "Your account has been deleted")
	}

	res, err = uc.generateToken(user)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) LoginCommittee(ctx context.Context, req domain.User) (domain.AuthResponse, error) {
	var res domain.AuthResponse
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	user, err := uc.ur.GetByEmail(ctx, req.Email)
	if err != nil {
		return res, status.Error(codes.InvalidArgument, "Incorrect email or password")
	}

	if user.Role != int8(domain.COMMITTEE) {
		return res, status.Error(codes.InvalidArgument, "Incorrect email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return res, status.Error(codes.InvalidArgument, "Incorrect email or password")
	}

	if !user.IsActive || user.DeletedAt != 0 {
		return res, status.Error(codes.Unauthenticated, "Your account has been deleted")
	}

	res, err = uc.generateToken(user)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) RefreshToken(ctx context.Context, req domain.RefreshTokenRequest) (domain.RefreshTokenResponse, error) {
	var res domain.RefreshTokenResponse
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	claims, err := uc.jwt.VerifyRefreshToken(req.RefreshToken)
	if err != nil {
		return res, err
	}

	user, err := uc.ur.GetByID(ctx, claims["id"].(string))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return res, status.Error(codes.NotFound, fmt.Sprintf("User with id %s not found", claims["id"].(string)))
		}

		return res, status.Error(codes.Internal, err.Error())
	}

	var customerID string

	if user.CustomerID != nil {
		customerID = *user.CustomerID
	}

	payload := jwtio.Payload{
		ID:         user.ID.Hex(),
		CustomerID: &customerID,
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
	}

	now := time.Now()

	newAccessToken, err := uc.jwt.CreateAccessToken(payload, uc.cfg.Jwt.AccessTokenSecret, now)
	if err != nil {
		return res, err
	}

	err = uc.cr.Set(domain.CacheItem{
		Key: helper.HashKey(newAccessToken),
		Exp: int32(now.Add(time.Minute * time.Duration(uc.cfg.Jwt.AccessTokenExpiryMinute)).Unix()),
	})
	if err != nil {
		return res, status.Error(codes.Internal, err.Error())
	}

	md, _ := metadata.FromIncomingContext(ctx)
	bearerAccessToken := md["authorization"]
	oldAccessToken := strings.Split(bearerAccessToken[0], " ")[1]

	err = uc.cr.Delete(helper.HashKey(oldAccessToken))
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		return res, status.Error(codes.Internal, err.Error())
	}

	res.AccessToken = newAccessToken

	return res, nil
}

func (uc UserUseCase) Create(ctx context.Context, req domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	_, err := uc.ur.GetByEmail(ctx, req.Email)
	if err == nil {
		return status.Error(codes.InvalidArgument, "User already exists with the given email")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	req.Password = string(encryptedPassword)
	err = uc.ur.Create(ctx, req)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (uc UserUseCase) GetAll(ctx context.Context, req domain.GetAllUserRequest) ([]domain.User, error) {
	var res []domain.User
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	res, err := uc.ur.GetAll(ctx, req)
	if err != nil {
		return res, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (uc UserUseCase) GetByID(ctx context.Context, userID string) (domain.User, error) {
	var res domain.User
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	res, err := uc.ur.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return res, status.Error(codes.NotFound, fmt.Sprintf("User with id %s not found", userID))
		}

		return res, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (uc UserUseCase) Update(ctx context.Context, req domain.UpdateUser) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	userID := req.ID.Hex()

	_, err := uc.ur.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return status.Error(codes.NotFound, fmt.Sprintf("User with id %s not found", userID))
		}

		return status.Error(codes.Internal, err.Error())
	}

	err = uc.ur.Update(ctx, req)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (uc UserUseCase) Delete(ctx context.Context, req domain.DeleteUser) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	userID := req.ID.Hex()

	_, err := uc.ur.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return status.Error(codes.NotFound, fmt.Sprintf("User with id %s not found", userID))
		}

		return status.Error(codes.Internal, err.Error())
	}

	err = uc.ur.Delete(ctx, req)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (uc UserUseCase) Logout(ctx context.Context, refreshToken string) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	accessToken, _ := uc.jwt.GetAccessToken(ctx)

	err := uc.cr.Delete(helper.HashKey(accessToken))
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		return status.Error(codes.Internal, err.Error())
	}

	err = uc.cr.Delete(helper.HashKey(refreshToken))
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
