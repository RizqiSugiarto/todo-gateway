// Package gateway is described reusable package for create gateway server
package gateway

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"strings"
	"time"

	"github.com/digisata/auth-service/api"
	"github.com/digisata/auth-service/bootstrap"
	"github.com/digisata/auth-service/docs"
	"github.com/digisata/auth-service/pkg/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	MaxHeaderBytes = 1 << 20
	ReadTimeOut    = 10 * time.Second
	WriteTimeOut   = 10 * time.Second
)

type Gateway struct {
	*runtime.ServeMux
	Host           string
	Port           int
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

func NewGateway(host string, port int, opts ...runtime.ServeMuxOption) *Gateway {
	opts = append(opts,
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)
	gwMux := runtime.NewServeMux(opts...)

	return &Gateway{
		ServeMux:       gwMux,
		Host:           host,
		Port:           port,
		MaxHeaderBytes: MaxHeaderBytes,
		ReadTimeout:    ReadTimeOut,
		WriteTimeout:   WriteTimeOut,
	}
}

func (gw *Gateway) swaggerUIHandler() (http.Handler, error) {
	err := mime.AddExtensionType(".svg", "image/svg+xml")
	if err != nil {
		return nil, errors.Wrap(err, "mime.AddExtensionType")
	}
	subFS, err := fs.Sub(docs.SwaggerUI, "swagger-ui")
	if err != nil {
		return nil, errors.Wrap(err, "fs.Sub")
	}
	return http.FileServer(http.FS(subFS)), nil
}

func (gw *Gateway) Run(ctx context.Context, cfg *bootstrap.Config) error {
	mux := http.NewServeMux()

	if cfg.AppEnv != "production" {
		sw, err := gw.swaggerUIHandler()
		if err != nil {
			return errors.Wrap(err, "gw.swaggerUIHandler")
		}

		fileServer := http.FileServer(http.FS(api.FS))

		mux.Handle("/static/", http.StripPrefix("/static", fileServer))
		mux.Handle("/", sw)
	}

	gwServer := &http.Server{
		Addr: fmt.Sprintf(":%d", gw.Port),
		Handler: middleware.CORS(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if strings.HasPrefix(request.URL.Path, "/auth-service/api/v1") {
				gw.ServeMux.ServeHTTP(writer, request)
				return
			}
			mux.ServeHTTP(writer, request)
		}), cfg.AppEnv, cfg.AllowedOrigins),
		ReadTimeout:    gw.ReadTimeout,
		WriteTimeout:   gw.WriteTimeout,
		MaxHeaderBytes: gw.MaxHeaderBytes,
	}

	go func() {
		<-ctx.Done()
		gwServer.Shutdown(ctx)
	}()

	log.Printf("App listening to port: %d", cfg.Port)
	err := gwServer.ListenAndServe()
	if err != nil {
		return errors.Wrap(err, "gwServer.ListenAndServe")
	}

	return nil
}
