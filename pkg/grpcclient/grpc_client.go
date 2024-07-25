package grpcclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/digisata/auth-service/pkg/interceptors"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	backoffLinear  time.Duration = 100 * time.Millisecond
	backoffRetries uint          = 3
)

type Config struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	Network string `mapstructure:"network"`
	TlS     bool   `mapstructure:"tls"`
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func NewGrpcClient(ctx context.Context, cfg Config, im interceptors.InterceptorManager, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	log.Println("Creating gRPC client...")
	if cfg.TlS {
		creds, err := loadTLSCredentials()
		if err != nil {
			return nil, err
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	opts = append(
		opts,
		grpc.WithUnaryInterceptor(grpcMiddleware.ChainUnaryClient(
			grpcRetry.UnaryClientInterceptor([]grpcRetry.CallOption{
				grpcRetry.WithBackoff(grpcRetry.BackoffLinear(backoffLinear)),
				grpcRetry.WithCodes(codes.NotFound, codes.Aborted),
				grpcRetry.WithMax(backoffRetries),
			}...),
		)),
	)

	if im != nil {
		opts = append(
			opts,
			grpc.WithUnaryInterceptor(grpcMiddleware.ChainUnaryClient(
				im.ClientRequestLoggerInterceptor(),
			)),
		)
	}

	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Println("Connecting to gRPC server at:", address)

	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "grpc.DialContext")
	}

	log.Println("gRPC server connected on", address)

	return conn, nil
}
