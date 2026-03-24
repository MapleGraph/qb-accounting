package config

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"qb-accounting/internal/helpers"

	qbconfig "github.com/MapleGraph/qb-core/v2/config"
	"github.com/joho/godotenv"
)

// Config holds all configuration for the accounting service
type Config struct {
	Server     ServerConfig
	Redis      RedisConfig
	Database   DatabaseConfig
	ClickHouse ClickHouseConfig
	Logger     LoggerConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port                 string
	GrpcPort             string
	Mode                 string // gin mode: debug, test, release
	GlobalServiceAddress string // Deprecated: Use GrpcServices instead
	GrpcServices         GrpcServicesConfig
}

// GrpcServicesConfig holds configuration for all gRPC services
type GrpcServicesConfig struct {
	SetupService        GrpcServiceConfig
	EmployeeService     GrpcServiceConfig
	NotificationService GrpcServiceConfig
	CatalogueService    GrpcServiceConfig
}

// GrpcServiceConfig holds configuration for a single gRPC service
type GrpcServiceConfig struct {
	Address        string
	ConnectTimeout string // Duration string (e.g., "5s", "10s")
	RequestTimeout string // Duration string (e.g., "5s", "10s")
	Enabled        bool
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type ClickHouseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type LoggerConfig struct {
	BufferSize     int
	FlushTimeout   time.Duration
	ServiceName    string
	ServiceVersion string
	Environment    string
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("No .env file found, using defaults or system env vars: %v", err)
	}
	return &Config{
		Server: ServerConfig{
			Port:                 helpers.GetEnv("SERVER_PORT", "8086"),
			GrpcPort:             helpers.GetEnv("SERVER_GRPC_PORT", "50056"),
			Mode:                 helpers.GetEnv("MODE", "debug"),
			GlobalServiceAddress: helpers.GetEnv("GLOBAL_SERVICE_ADDR", "localhost:50052"),
			GrpcServices: GrpcServicesConfig{
				SetupService: GrpcServiceConfig{
					Address:        helpers.GetEnv("GRPC_SETUP_SERVICE_ADDR", "localhost:50054"),
					ConnectTimeout: helpers.GetEnv("GRPC_SETUP_CONNECT_TIMEOUT", "5s"),
					RequestTimeout: helpers.GetEnv("GRPC_SETUP_REQUEST_TIMEOUT", "10s"),
					Enabled:        helpers.GetEnvBool("GRPC_SETUP_ENABLED", true),
				},
				EmployeeService: GrpcServiceConfig{
					Address:        helpers.GetEnv("GRPC_EMPLOYEE_SERVICE_ADDR", "localhost:50053"),
					ConnectTimeout: helpers.GetEnv("GRPC_EMPLOYEE_CONNECT_TIMEOUT", "5s"),
					RequestTimeout: helpers.GetEnv("GRPC_EMPLOYEE_REQUEST_TIMEOUT", "10s"),
					Enabled:        helpers.GetEnvBool("GRPC_EMPLOYEE_ENABLED", true),
				},
				NotificationService: GrpcServiceConfig{
					Address:        helpers.GetEnv("GRPC_NOTIFICATION_SERVICE_ADDR", "localhost:50056"),
					ConnectTimeout: helpers.GetEnv("GRPC_NOTIFICATION_CONNECT_TIMEOUT", "5s"),
					RequestTimeout: helpers.GetEnv("GRPC_NOTIFICATION_REQUEST_TIMEOUT", "10s"),
					Enabled:        helpers.GetEnvBool("GRPC_NOTIFICATION_ENABLED", true),
				},
				CatalogueService: GrpcServiceConfig{
					Address:        helpers.GetEnv("GRPC_CATALOGUE_SERVICE_ADDR", "localhost:50053"),
					ConnectTimeout: helpers.GetEnv("GRPC_CATALOGUE_CONNECT_TIMEOUT", "5s"),
					RequestTimeout: helpers.GetEnv("GRPC_CATALOGUE_REQUEST_TIMEOUT", "10s"),
					Enabled:        helpers.GetEnvBool("GRPC_CATALOGUE_ENABLED", true),
				},
			},
		},
		Redis: RedisConfig{
			Addr:     helpers.GetEnv("REDIS_ADDR", "localhost:6379"),
			Password: helpers.GetEnv("REDIS_PASSWORD", "password"),
			DB:       helpers.GetEnvInt("REDIS_DB", 0),
		},
		Database: DatabaseConfig{
			Host:     helpers.GetEnv("DB_HOST", "localhost"),
			Port:     helpers.GetEnv("DB_PORT", "5432"),
			User:     helpers.GetEnv("DB_USER", "postgres"),
			Password: helpers.GetEnv("DB_PASSWORD", "password"),
			DBName:   helpers.GetEnv("DB_NAME", "qb_accounting"),
			SSLMode:  helpers.GetEnv("DB_SSL_MODE", "disable"),
		},
		ClickHouse: ClickHouseConfig{
			Host:     helpers.GetEnv("CLICKHOUSE_HOST", "localhost"),
			Port:     helpers.GetEnv("CLICKHOUSE_PORT", "8123"),
			User:     helpers.GetEnv("CLICKHOUSE_USER", "default"),
			Password: helpers.GetEnv("CLICKHOUSE_PASSWORD", "clickhouse"),
			DBName:   helpers.GetEnv("CLICKHOUSE_DATABASE", "qb_accounting"),
		},
		Logger: LoggerConfig{
			BufferSize:     helpers.GetEnvInt("LOG_BUFFER_SIZE", 1000),
			FlushTimeout:   helpers.GetEnvDuration("LOG_FLUSH_INTERVAL", 5*time.Second),
			ServiceName:    helpers.GetEnv("SERVICE_NAME", "qb-accounting"),
			ServiceVersion: helpers.GetEnv("SERVICE_VERSION", "1.0.0"),
			Environment:    helpers.GetEnv("ENVIRONMENT", "debug"),
		},
	}
}

// ToQBCoreConfig converts internal Config to qb-core ServiceConfig
func (c *Config) ToQBCoreConfig() *qbconfig.ServiceConfig {
	// Build PostgreSQL DatabaseURL from components
	userInfo := url.UserPassword(c.Database.User, c.Database.Password)
	databaseURL := fmt.Sprintf(
		"postgres://%s@%s:%s/%s?sslmode=%s",
		userInfo.String(),
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
		c.Database.SSLMode,
	)

	// Build ClickHouse address array
	clickhouseAddr := fmt.Sprintf("%s:%s", c.ClickHouse.Host, c.ClickHouse.Port)

	grpcConfigs := make(map[string]*qbconfig.GRPCClientConfig)
	if c.Server.GrpcServices.SetupService.Enabled {
		connectTimeout, _ := time.ParseDuration(c.Server.GrpcServices.SetupService.ConnectTimeout)
		requestTimeout, _ := time.ParseDuration(c.Server.GrpcServices.SetupService.RequestTimeout)
		grpcConfigs["setup"] = &qbconfig.GRPCClientConfig{
			Address:        c.Server.GrpcServices.SetupService.Address,
			ConnectTimeout: connectTimeout,
			RequestTimeout: requestTimeout,
		}
	}
	if c.Server.GrpcServices.EmployeeService.Enabled {
		connectTimeout, _ := time.ParseDuration(c.Server.GrpcServices.EmployeeService.ConnectTimeout)
		requestTimeout, _ := time.ParseDuration(c.Server.GrpcServices.EmployeeService.RequestTimeout)
		grpcConfigs["employee"] = &qbconfig.GRPCClientConfig{
			Address:        c.Server.GrpcServices.EmployeeService.Address,
			ConnectTimeout: connectTimeout,
			RequestTimeout: requestTimeout,
		}
	}
	if c.Server.GrpcServices.NotificationService.Enabled {
		connectTimeout, _ := time.ParseDuration(c.Server.GrpcServices.NotificationService.ConnectTimeout)
		requestTimeout, _ := time.ParseDuration(c.Server.GrpcServices.NotificationService.RequestTimeout)
		grpcConfigs["notification"] = &qbconfig.GRPCClientConfig{
			Address:        c.Server.GrpcServices.NotificationService.Address,
			ConnectTimeout: connectTimeout,
			RequestTimeout: requestTimeout,
		}
	}
	if c.Server.GrpcServices.CatalogueService.Enabled {
		connectTimeout, _ := time.ParseDuration(c.Server.GrpcServices.CatalogueService.ConnectTimeout)
		requestTimeout, _ := time.ParseDuration(c.Server.GrpcServices.CatalogueService.RequestTimeout)
		grpcConfigs["catalogue"] = &qbconfig.GRPCClientConfig{
			Address:        c.Server.GrpcServices.CatalogueService.Address,
			ConnectTimeout: connectTimeout,
			RequestTimeout: requestTimeout,
		}
	}

	serviceConfig := &qbconfig.ServiceConfig{
		ServiceName: c.Logger.ServiceName,
		Environment: c.Logger.Environment,
		Databases: map[string]*qbconfig.DatabaseConfig{
			"main": {
				WriteConnection: qbconfig.PostgresConfig{
					DatabaseURL: databaseURL,
				},
			},
		},
		Redis: &qbconfig.RedisConfig{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       c.Redis.DB,
		},
		ClickHouse: &qbconfig.ClickHouseConfig{
			Addr:     []string{clickhouseAddr},
			Database: c.ClickHouse.DBName,
			Username: c.ClickHouse.User,
			Password: c.ClickHouse.Password,
		},
		GRPC: grpcConfigs,
	}

	return serviceConfig
}
