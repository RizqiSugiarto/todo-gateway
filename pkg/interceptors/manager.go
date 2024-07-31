package interceptors

import (
	"context"

	"github.com/digisata/todo-gateway/pkg/constants"
	"github.com/digisata/todo-gateway/pkg/jwtio"
	"github.com/digisata/todo-gateway/pkg/tracing"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type InterceptorManager interface {
	Logger(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error)
	ClientRequestLoggerInterceptor() func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error
	AuthenticationInterceptor(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error)
}

type interceptorManager struct {
	logger           *zap.SugaredLogger
	jwtManager       *jwtio.JSONWebToken
	protectedMethods map[string]bool
}

func NewInterceptorManager(jwtManager *jwtio.JSONWebToken, logger *zap.SugaredLogger) *interceptorManager {
	return &interceptorManager{
		logger:           logger,
		jwtManager:       jwtManager,
		protectedMethods: protectedMethods(),
	}
}

func (im interceptorManager) Logger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	reply, err := handler(ctx, req)
	if err != nil {
		im.logger.Errorw(constants.ERROR,
			"method", info.FullMethod,
			"request", req,
			"error", err.Error(),
		)

		return reply, err
	}

	im.logger.Infow(constants.INFO,
		"method", info.FullMethod,
		"request", req,
		"error", nil,
	)

	return reply, err
}

func (im interceptorManager) ClientRequestLoggerInterceptor() func(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		im.logger.Infow(constants.INFO,
			"method", method,
			"request", req,
			"error", nil,
		)

		return err
	}
}

func (im interceptorManager) AuthenticationInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "Interceptors.AuthenticationInterceptor")
	defer span.End()

	if _, isProtected := im.protectedMethods[info.FullMethod]; !isProtected {
		return handler(ctx, req)
	}

	claims, err := im.jwtManager.Verify(ctx, info.FullMethod == constants.PATH+"RefreshToken")
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "claims", claims)

	return handler(ctx, req)
}

func protectedMethods() map[string]bool {
	return map[string]bool{
		// Auth
		constants.PATH + "Logout": true,

		// Activities
		constants.PATH + "CreateActivity": true,
		constants.PATH + "GetAllActivity": true,
		constants.PATH + "GetActivity":    true,
		constants.PATH + "UpdateActivity": true,
		constants.PATH + "DeleteActivity": true,

		// Tasks
		constants.PATH + "CreateTask":      true,
		constants.PATH + "GetAllTask":      true,
		constants.PATH + "GetTask":         true,
		constants.PATH + "UpdateTask":      true,
		constants.PATH + "BatchUpdateTask": true,
		constants.PATH + "DeleteTask":      true,

		//Text
		constants.PATH + "CreateText": true,
		constants.PATH + "GetAllText": true,
		constants.PATH + "GetText":    true,
		constants.PATH + "UpdateText": true,
		constants.PATH + "DeleteText": true,
	}
}
