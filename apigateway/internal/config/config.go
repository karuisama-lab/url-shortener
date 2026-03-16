package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Postgres   `env:"POSTGRES" yaml:"postgres"`
}

type Postgres struct {
	DSN            string `env:"DSN" env-default:"" yaml:"dsn"`
	Host           string `env:"HOST" yaml:"host"`
	Port           int    `env:"PORT" yaml:"port"`
	Login          string `env:"LOGIN" yaml:"login"`
	Password       string `env:"PASSWORD" yaml:"password"`
	DbName         string `env:"DB_NAME" yaml:"db_name"`
	SslMode        string `env:"SSL_MODE" yaml:"ssl_mode"`
	MigratePath    string `env:"MIGRATE_PATH" yaml:"migrate_path"`
	MigrationTable string `env:"MIGRATION_TABLE" yaml:"migration_table"`
	SchemeName     string `env:"SCHEME_NAME" env-default:"public" yaml:"scheme_name"`

	//===== НАСТРОЙКИ НА ПУЛ =====
	MinConns          int32         `env:"MIN_CONNS" env-default:"5" yaml:"min_conns"` //TODO: почему тут правильнее делать *int32? как это должно работать? в каких случаях так делать?
	MaxConns          *int32        `env:"MAX_CONNS" env-default:"25" yaml:"max_conns"`
	MaxIdleTime       time.Duration `env:"MAX_IDLE_TIME" env-default:"30m" yaml:"max_idle_time"`
	MaxLifetime       time.Duration `env:"MAX_LIFETIME" env-default:"90m" yaml:"max_lifetime"`
	HealthCheckPeriod time.Duration `env:"HEALTH_CHECK_PERIOD" env-default:"10s" yaml:"health_check_period"`

	//===== НАСТРОЙКИ НА ОТДЕЛЬНО СОЕДИНЕНИЕ =====
	ConnectTimeout time.Duration `env:"CONNECT_TIMEOUT" env-default:"2s" yaml:"connect_timeout"`
	ReadTimeout    time.Duration `env:"READ_TIMEOUT" env-default:"15s" yaml:"read_timeout"`
	WriteTimeout   time.Duration `env:"WRITE_TIMEOUT" env-default:"30s" yaml:"write_timeout"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"iddle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	ConfigPath := os.Getenv("CONFIG_PATH")
	log.Printf("config path has been read")
	if ConfigPath == "" {
		log.Fatalf("can not read config path")
	}

	var cfg *Config

	err := cleanenv.ReadConfig(ConfigPath, &cfg)
	if err != nil {
		log.Fatalf("can not fulfill config file")
	}

	return cfg
}
