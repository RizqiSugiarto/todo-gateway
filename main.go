package main

import (
	"context"
	"time"

	"github.com/digisata/todo-gateway/bootstrap"
	"github.com/digisata/todo-gateway/controller"
	"github.com/digisata/todo-gateway/domain"
	"github.com/digisata/todo-gateway/gateway"
	"github.com/digisata/todo-gateway/pkg/grpcclient"
	"github.com/digisata/todo-gateway/pkg/grpcserver"
	"github.com/digisata/todo-gateway/pkg/interceptors"
	"github.com/digisata/todo-gateway/pkg/jwtio"
	memcachedRepo "github.com/digisata/todo-gateway/repository/memcached"
	mongoRepo "github.com/digisata/todo-gateway/repository/mongo"
	stubs "github.com/digisata/todo-gateway/stubs"
	activityPB "github.com/digisata/todo-gateway/stubs/activity"
	taskPB "github.com/digisata/todo-gateway/stubs/task"
	textPB "github.com/digisata/todo-gateway/stubs/text"
	"github.com/digisata/todo-gateway/usecase"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	ctx := context.Background()

	// Setup app
	app, err := bootstrap.App()
	if err != nil {
		panic(err)
	}

	cfg := app.Cfg

	jwt := jwtio.NewJSONWebToken(&cfg.Jwt, app.MemcachedDB)

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()

	// Dependencies injection
	db := app.Mongo.Database(cfg.Mongo.DBName)
	defer app.CloseDBConnection()

	userRepository := mongoRepo.NewUserRepository(db, domain.USER_COLLECTION)
	cacheRepository := memcachedRepo.NewCacheRepository(app.MemcachedDB)
	timeout := time.Duration(cfg.ContextTimeout) * time.Second
	authController := &controller.AuthController{
		UserUseCase: usecase.NewUserUseCase(jwt, cfg, userRepository, cacheRepository, timeout),
	}

	// Setup GRPC server
	im := interceptors.NewInterceptorManager(jwt, sugar)
	grpcServer, err := grpcserver.NewGrpcServer(cfg.GrpcServer, im, sugar)
	if err != nil {
		panic(err)
	}

	stubs.RegisterAuthServiceServer(grpcServer, authController)
	grpc_health_v1.RegisterHealthServer(grpcServer.Server, health.NewServer())

	err = grpcServer.Run()
	if err != nil {
		panic(err)
	}
	defer grpcServer.Stop(ctx)

	// Setup GRPC client
	authService, err := grpcclient.NewGrpcClient(ctx, cfg.GRPCClient.AuthService, im)
	if err != nil {
		panic(err)
	}
	defer authService.Close()

	var todoService *grpc.ClientConn

	if cfg.TodoService {
		todoService, err = grpcclient.NewGrpcClient(ctx, cfg.GRPCClient.TodoService, nil)
		if err != nil {
			panic(err)
		}
		defer todoService.Close()
	}

	authController.ActivityUseCase = activityPB.NewActivityServiceClient(todoService)
	authController.TaskUseCase = taskPB.NewTaskServiceClient(todoService)
	authController.TextUseCase = textPB.NewTextServiceClient(todoService)

	// Setup gateway mux
	gatewayServer := gateway.NewGateway(cfg.Host, cfg.Port)
	err = stubs.RegisterAuthServiceHandler(ctx, gatewayServer.ServeMux, authService)
	if err != nil {
		panic(err)
	}

	err = gatewayServer.Run(ctx, cfg)
	if err != nil {
		panic(err)
	}

	<-ctx.Done()
}
