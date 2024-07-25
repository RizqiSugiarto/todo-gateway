package grpcserver

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/digisata/auth-service/pkg/constants"
	"github.com/digisata/auth-service/pkg/interceptors"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

type (
	Config struct {
		Host    string `mapstructure:"host"`
		Port    int    `mapstructure:"port"`
		Network string `mapstructure:"network"`
		TlS     bool   `mapstructure:"tls"`
	}

	GrpcServer struct {
		*grpc.Server
		logger   *zap.SugaredLogger
		Listener net.Listener
		Host     string
		Port     int
		Network  string
	}
)

const (
	maxConnectionIdle time.Duration = 300
	gRPCTimeout       time.Duration = 15
	maxConnectionAge  time.Duration = 300
	gRPCTime          time.Duration = 600

	serverCertFile   = "cert/server-cert.pem"
	serverKeyFile    = "cert/server-key.pem"
	clientCACertFile = "cert/ca-cert.pem"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	pemClientCA, err := os.ReadFile(clientCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func NewGrpcServer(cfg Config, im interceptors.InterceptorManager, logger *zap.SugaredLogger, opts ...grpc.ServerOption) (*GrpcServer, error) {
	log.Println("Starting gRPC server...")
	if cfg.TlS {
		creds, err := loadTLSCredentials()
		if err != nil {
			return nil, err
		}

		opts = append(opts, grpc.Creds(creds))
	} else {
		opts = append(opts, grpc.Creds(insecure.NewCredentials()))
	}

	opts = append(
		opts,
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Second,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Second,
			Time:              gRPCTime * time.Second,
		}),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcCtxtags.UnaryServerInterceptor(),
			grpcPrometheus.UnaryServerInterceptor,
			grpcRecovery.UnaryServerInterceptor(),
			im.Logger,
			im.AuthenticationInterceptor,
			im.AuthorizationInterceptor,
		)),
	)

	server := grpc.NewServer(opts...)
	grpcPrometheus.Register(server)

	return &GrpcServer{
		logger:  logger,
		Server:  server,
		Network: cfg.Network,
		Host:    cfg.Host,
		Port:    cfg.Port,
	}, nil
}

func (grpcServer *GrpcServer) Run() error {
	listener, err := net.Listen(grpcServer.Network, fmt.Sprintf(":%d", grpcServer.Port))
	if err != nil {
		return errors.Wrap(err, "net.Listen")
	}

	grpcServer.Listener = listener

	go func() {
		log.Println("gRPC server listening on", grpcServer.Port)
		if err := grpcServer.Server.Serve(grpcServer.Listener); err != nil {
			grpcServer.logger.Fatalw(constants.FATAL,
				"error", err.Error(),
			)
		}
	}()

	return nil
}

func (grpcServer GrpcServer) Stop(ctx context.Context) {
	if err := grpcServer.Listener.Close(); err != nil {
		grpcServer.logger.Fatalw(constants.FATAL,
			"error", err.Error(),
		)
	}

	go func() {
		defer grpcServer.Server.GracefulStop()
		<-ctx.Done()
	}()
}
