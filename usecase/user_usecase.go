package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/digisata/todo-gateway/bootstrap"
	"github.com/digisata/todo-gateway/domain"
	"github.com/digisata/todo-gateway/helper"
	"github.com/digisata/todo-gateway/pkg/jwtio"
	memcachedRepo "github.com/digisata/todo-gateway/repository/memcached"
	mongoRepo "github.com/digisata/todo-gateway/repository/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
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
		res domain.AuthResponse
	)

	payload := jwtio.Payload{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
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

	res = domain.AuthResponse{
		AccessToken: accessToken,
	}

	return res, nil
}

func (uc UserUseCase) Login(ctx context.Context, req domain.User) (domain.AuthResponse, error) {
	var res domain.AuthResponse
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	user, err := uc.ur.GetByEmail(ctx, req.Email)
	if err != nil {
		return res, status.Error(codes.InvalidArgument, "Incorrect email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return res, status.Error(codes.InvalidArgument, "Incorrect email or password")
	}

	if user.DeletedAt != 0 {
		return res, status.Error(codes.Unauthenticated, "Your account has been deleted")
	}

	res, err = uc.generateToken(user)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) Logout(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	accessToken, _ := uc.jwt.GetAccessToken(ctx)

	err := uc.cr.Delete(helper.HashKey(accessToken))
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
