// Package jwtio is shared pkg of json web token
package jwtio

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/digisata/auth-service/helper"
	"github.com/digisata/auth-service/pkg/constants"
	"github.com/digisata/auth-service/pkg/memcached"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type (
	Config struct {
		AccessTokenExpiryMinute  int    `mapstructure:"access_token_expiry_minute"`
		RefreshTokenExpiryMinute int    `mapstructure:"refresh_token_expiry_minute"`
		AccessTokenSecret        string `mapstructure:"access_token_secret"`
		RefreshTokenSecret       string `mapstructure:"refresh_token_secret"`
	}

	Payload struct {
		ID         string
		CustomerID *string
		Name       string
		Email      string
		Role       int8
	}

	JSONWebToken struct {
		cfg         *Config
		memcachedDB *memcached.Database
	}

	JwtCustomClaims struct {
		ID         string  `json:"id"`
		CustomerID *string `json:"customer_id"`
		Name       string  `json:"name"`
		Role       int8    `json:"role"`
		jwt.RegisteredClaims
	}

	JwtCustomRefreshClaims struct {
		ID string `json:"id"`
		jwt.RegisteredClaims
	}
)

func NewJSONWebToken(cfg *Config, memcachedDB *memcached.Database) *JSONWebToken {
	return &JSONWebToken{
		cfg:         cfg,
		memcachedDB: memcachedDB,
	}
}

func (j JSONWebToken) CreateAccessToken(payload Payload, secret string, now time.Time) (string, error) {
	claims := &JwtCustomClaims{
		ID:         payload.ID,
		CustomerID: payload.CustomerID,
		Name:       payload.Name,
		Role:       payload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:  payload.Email,
			IssuedAt: jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}

	return t, nil
}

func (j JSONWebToken) CreateRefreshToken(payload Payload, secret string, now time.Time) (string, error) {
	claims := &JwtCustomRefreshClaims{
		ID: payload.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:  payload.Email,
			IssuedAt: jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}

	return rt, nil
}

func (j JSONWebToken) Verify(ctx context.Context, isRefreshToken bool) (jwt.MapClaims, error) {
	accessToken, err := j.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return j.validateToken(token, j.cfg.AccessTokenSecret)
	})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if !isRefreshToken {
		_, err = j.memcachedDB.Get(helper.HashKey(accessToken))
		if err != nil {
			if errors.Is(err, memcache.ErrCacheMiss) {
				return nil, status.Error(codes.Unauthenticated, constants.TOKEN_EXPIRED)
			}

			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, status.Error(codes.Unauthenticated, constants.FAILED_TO_EXTRACT)
	}

	return claims, nil
}

func (j JSONWebToken) VerifyRefreshToken(refreshToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return j.validateToken(token, j.cfg.RefreshTokenSecret)
	})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	_, err = j.memcachedDB.Get(helper.HashKey(refreshToken))
	if err != nil {
		if errors.Is(err, memcache.ErrCacheMiss) {
			return nil, status.Error(codes.Unauthenticated, constants.REFRESH_TOKEN_EXPIRED)
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, status.Error(codes.Unauthenticated, constants.FAILED_TO_EXTRACT)
	}

	return claims, nil
}

func ExtractValueFromToken[T string | int8](token *jwt.Token, key string) (T, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		var zero T
		return zero, status.Error(codes.Unauthenticated, constants.FAILED_TO_EXTRACT)
	}

	return claims[key].(T), nil
}

func (j JSONWebToken) validateToken(token *jwt.Token, secret string) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, status.Error(codes.Unauthenticated, constants.UNEXPECTED_SIGNING_METHOD)

	}

	return []byte(secret), nil
}

func (j JSONWebToken) GetAccessToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "Metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return "", status.Error(codes.Unauthenticated, "Authorization token is not provided")
	}

	split := strings.Split(values[0], " ")
	if len(split) != 2 {
		return "", status.Error(codes.Unauthenticated, "Invalid access token format")
	}

	if split[0] != "Bearer" {
		return "", status.Error(codes.Unauthenticated, "Invalid access token format")
	}

	return split[1], nil
}
