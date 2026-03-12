// Package config provider config
package config

import (
	"fmt"
	"time"
)

const (
	EnvProd = "production"
	EnvDev  = "development"
)

type Config struct {
	App      AppConfig      `mapstructure:"app" yaml:"app"`
	Server   ServerConfig   `mapstructure:"server" yaml:"server"`
	Database DatabaseConfig `mapstructure:"database" yaml:"database"`
	Redis    RedisConfig    `mapstructure:"redis" yaml:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt" yaml:"jwt"`
	Log      LogConfig      `mapstructure:"log" yaml:"log"`
	Email    EmailConfig    `mapstructure:"email" yaml:"email"`
	LLM      LLMConfig      `mapstructure:"llm" yaml:"llm"`
	Rustfs   RustfsConfig   `mapstructure:"rustfs" yaml:"rustfs"`
}

type AppConfig struct {
	Name        string   `mapstructure:"name" yaml:"name"`
	Version     string   `mapstructure:"version" yaml:"version"`
	Environment string   `mapstructure:"environment" yaml:"environment"`
	Debug       bool     `mapstructure:"debug" yaml:"debug"`
	Domain      string   `mapstructure:"domain" yaml:"domain"`
	CorsOrigins []string `mapstructure:"cors_origins" yaml:"cors_origins"`
}

func (a AppConfig) IsProd() bool {
	return a.Environment == EnvProd
}

func (a AppConfig) IsDev() bool {
	return a.Environment == EnvDev
}

type ServerConfig struct {
	Host             string        `mapstructure:"host" yaml:"host"`
	Port             int           `mapstructure:"port" yaml:"port"`
	ReadTimeout      time.Duration `mapstructure:"read_timeout" yaml:"read_timeout"`
	WriteTimeout     time.Duration `mapstructure:"write_timeout" yaml:"write_timeout"`
	IdleTimeout      time.Duration `mapstructure:"idle_timeout" yaml:"idle_timeout"`
	MaxHeaderBytes   int           `mapstructure:"max_header_bytes" yaml:"max_header_bytes"`
	GracefulShutdown time.Duration `mapstructure:"graceful_shutdown" yaml:"graceful_shutdown"`
}

func (s ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type DatabaseConfig struct {
	Host            string        `mapstructure:"host" yaml:"host"`
	Port            int           `mapstructure:"port" yaml:"port"`
	User            string        `mapstructure:"user" yaml:"user"`
	Password        string        `mapstructure:"password" yaml:"password"`
	Name            string        `mapstructure:"name" yaml:"name"`
	SSLMode         string        `mapstructure:"ssl_mode" yaml:"ssl_mode"`
	MaxOpenConns    int           `mapstructure:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" yaml:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time" yaml:"conn_max_idle_time"`
	Timeout         time.Duration `mapstructure:"timeout" yaml:"timeout"`
}

type RedisConfig struct {
	Host               string        `mapstructure:"host" yaml:"host"`
	Port               int           `mapstructure:"port" yaml:"port"`
	Password           string        `mapstructure:"password" yaml:"password"`
	DB                 int           `mapstructure:"db" yaml:"db"`
	PoolSize           int           `mapstructure:"pool_size" yaml:"pool_size"`
	MinIdleConns       int           `mapstructure:"min_idle_conns" yaml:"min_idle_conns"`
	DialTimeout        time.Duration `mapstructure:"dial_timeout" yaml:"dial_timeout"`
	ReadTimeout        time.Duration `mapstructure:"read_timeout" yaml:"read_timeout"`
	WriteTimeout       time.Duration `mapstructure:"write_timeout" yaml:"write_timeout"`
	PoolTimeout        time.Duration `mapstructure:"pool_timeout" yaml:"pool_timeout"`
	IdleTimeout        time.Duration `mapstructure:"idle_timeout" yaml:"idle_timeout"`
	IdleCheckFrequency time.Duration `mapstructure:"idle_check_frequency" yaml:"idle_check_frequency"`
}

type JWTConfig struct {
	Secret            string        `mapstructure:"secret" yaml:"secret"`
	AccessExpiration  time.Duration `mapstructure:"access_expiration" yaml:"access_expiration"`
	RefreshExpiration time.Duration `mapstructure:"refresh_expiration" yaml:"refresh_expiration"`
	Issuer            string        `mapstructure:"issuer" yaml:"issuer"`
}

type LogConfig struct {
	Level      string `mapstructure:"level" yaml:"level"`
	Format     string `mapstructure:"format" yaml:"format"`
	FilePath   string `mapstructure:"file_path" yaml:"file_path"`
	MaxSize    int    `mapstructure:"max_size" yaml:"max_size"`
	MaxBackups int    `mapstructure:"max_backups" yaml:"max_backups"`
	MaxAge     int    `mapstructure:"max_age" yaml:"max_age"`
	Compress   bool   `mapstructure:"compress" yaml:"compress"`
}

type EmailConfig struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	From     string `mapstructure:"from" yaml:"from"`
}

type LLMConfig struct {
	APIKey string `mapstructure:"apikey" yaml:"apikey"`
}

type RustfsConfig struct {
	Region          string `mapstructure:"region" yaml:"region"`
	AccessKeyID     string `mapstructure:"access_key_id" yaml:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key" yaml:"secret_access_key"`
	Endpoint        string `mapstructure:"endpoint" yaml:"endpoint"`
}
