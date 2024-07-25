package bootstrap

import (
	"fmt"
	"log"

	"github.com/digisata/auth-service/pkg/grpcclient"
	"github.com/digisata/auth-service/pkg/grpcserver"
	"github.com/digisata/auth-service/pkg/jwtio"
	"github.com/digisata/auth-service/pkg/memcached"
	"github.com/digisata/auth-service/pkg/mongo"
	"github.com/spf13/viper"
)

type (
	Config struct {
		AppEnv            string            `mapstructure:"app_env"`
		AllowedOrigins    string            `mapstructure:"allowed_origins"`
		PreferenceService bool              `mapstructure:"preference_service"`
		InvitationService bool              `mapstructure:"invitation_service"`
		Host              string            `mapstructure:"host"`
		Port              int               `mapstructure:"port"`
		ContextTimeout    int               `mapstructure:"context_timeout"`
		Jwt               jwtio.Config      `mapstructure:"jwt"`
		Mongo             mongo.Config      `mapstructure:"mongo"`
		Memcached         memcached.Config  `mapstructure:"memcached"`
		GrpcServer        grpcserver.Config `mapstructure:"grpc_server"`
		GRPCClient        GRPCClient        `mapstructure:"grpc_client"`
	}

	GRPCClient struct {
		AuthService       grpcclient.Config `mapstructure:"auth_service"`
		PreferenceService grpcclient.Config `mapstructure:"preference_service"`
		InvitationService grpcclient.Config `mapstructure:"invitation_service"`
	}
)

func LoadConfig() (*Config, error) {
	var cfg Config

	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("can't find the config file: %v", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("environment can't be loaded: %v", err)
	}

	log.Printf("The App is running in %s environment", cfg.AppEnv)

	return &cfg, nil
}
