package main

import (
	"context"
	"time"

	"github.com/digisata/auth-service/bootstrap"
	"github.com/digisata/auth-service/controller"
	"github.com/digisata/auth-service/domain"
	"github.com/digisata/auth-service/gateway"
	"github.com/digisata/auth-service/pkg/grpcclient"
	"github.com/digisata/auth-service/pkg/grpcserver"
	"github.com/digisata/auth-service/pkg/interceptors"
	"github.com/digisata/auth-service/pkg/jwtio"
	memcachedRepo "github.com/digisata/auth-service/repository/memcached"
	mongoRepo "github.com/digisata/auth-service/repository/mongo"
	stubs "github.com/digisata/auth-service/stubs"
	invitationCategoryPB "github.com/digisata/auth-service/stubs/invitation-category"
	socialPB "github.com/digisata/auth-service/stubs/social"
	"github.com/digisata/auth-service/usecase"
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
	profileRepository := mongoRepo.NewProfileRepository(db, domain.USER_COLLECTION)
	cacheRepository := memcachedRepo.NewCacheRepository(app.MemcachedDB)
	timeout := time.Duration(cfg.ContextTimeout) * time.Second
	authController := &controller.AuthController{
		UserUseCase:    usecase.NewUserUseCase(jwt, cfg, userRepository, cacheRepository, timeout),
		ProfileUseCase: usecase.NewProfileUseCase(jwt, cfg, profileRepository, cacheRepository, timeout),
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

	var (
		preferenceService *grpc.ClientConn
		invitationService *grpc.ClientConn
	)

	if cfg.PreferenceService {
		preferenceService, err = grpcclient.NewGrpcClient(ctx, cfg.GRPCClient.PreferenceService, nil)
		if err != nil {
			panic(err)
		}
		defer preferenceService.Close()
	}

	if cfg.InvitationService {
		invitationService, err = grpcclient.NewGrpcClient(ctx, cfg.GRPCClient.InvitationService, nil)
		if err != nil {
			panic(err)
		}
		defer invitationService.Close()
	}

	authController.SocialUseCase = socialPB.NewSocialServiceClient(preferenceService)
	authController.InvitationCategoryUseCase = invitationCategoryPB.NewInvitationCategoryServiceClient(invitationService)

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
