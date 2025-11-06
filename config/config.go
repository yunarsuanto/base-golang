package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

// Config struct for .env.yml
type Config struct {
	DefaultTimezone *time.Location
	DefaultOtp      string
	Server          ServerConfig
	Db              DBConfig
	Redis           RedisServer
	JwtConfig       JwtConfig
	SsoConfig       SsoConfig
	Firebase        FirebaseConfig
	Mailer          MailerConfig
	Form            FormConfig
}

// ServerConfig struct to handle server configuration
type ServerConfig struct {
	AppName                string
	FrontendUrl            string
	AppUrl                 string
	Addr                   string
	WriteTimeout           int
	ReadTimeout            int
	GraceFulTimeout        int
	MaxReceivedMessageSize int
	MaxImageHeight         int
	StorageProvider        string
	LocalStoragePath       string
	LogoUrl                string
	UrlFrontend            string
}

// DBConfig struct to handle database configuration
type DBConfig struct {
	Name            string
	Host            string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime int
}

// RedisServer struct to handle redis configuration
type RedisServer struct {
	Addr     string
	Password string
	Timeout  int
	MaxIdle  int
}

// JwtConfig struct to handle JWT configuration
type JwtConfig struct {
	Issuer               string
	Secret               string
	AccessTokenDuration  int // by minutes
	RefreshTokenDuration int // by months
	ClientId             string
}

type SsoConfig struct {
	Url         string
	ApiUrl      string
	AppId       string
	RedirectUrl string
}

type FirebaseConfig struct {
	ServiceKeyPath string
}

type MailerConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Sender   string
}

type FormConfig struct {
	FormApprovalDuration int // by days
}

// InitConfig function to init configuration, returns Config struct
func InitConfig(defaultTimezone *time.Location) Config {
	viper.SetConfigName(".env")
	if os.Getenv("ENV") != "" {
		viper.SetConfigName(".env-" + os.Getenv("ENV"))
	}

	viper.AddConfigPath(".")

	var configuration Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	if defaultTimezone != nil {
		configuration.DefaultTimezone = defaultTimezone
	} else {
		configuration.DefaultTimezone = time.Local
	}

	return configuration
}
