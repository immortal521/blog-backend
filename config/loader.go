package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	globalConfig *Config
	once         sync.Once
	mu           sync.RWMutex
)

type Options struct {
	ConfigFile string
	ConfigType string
	EnvPrefix  string
	WatchFile  bool
	OnChange   func(cfg *Config)
}

func Load(opts ...Options) (*Config, error) {
	var err error
	log.Println("[config] loading...")
	log.Println("[config] opts:", opts)
	once.Do(func() {
		globalConfig, err = load(mergeOptions(opts...))
	})
	if err != nil {
		return nil, err
	}
	return globalConfig, nil
}

func MustLoad(opts ...Options) *Config {
	cfg, err := Load(opts...)
	if err != nil {
		log.Fatalf("[config] load failed: %v", err)
	}
	return cfg
}

// Get 获取全局配置（需先调用 Load）
func Get() *Config {
	mu.RLock()
	defer mu.RUnlock()
	if globalConfig == nil {
		panic("[config] not initialized, call Load() first")
	}
	return globalConfig
}

func load(opt Options) (*Config, error) {
	v := viper.New()

	// 文件配置
	if opt.ConfigFile != "" {
		v.SetConfigFile(opt.ConfigFile)
	} else {
		v.SetConfigName("config")
		v.SetConfigType(opt.ConfigType)
		v.AddConfigPath(".")
		v.AddConfigPath("./config")
		// v.AddConfigPath("/etc/myapp")
	}

	// 环境变量（优先级高于文件）
	v.SetEnvPrefix(opt.EnvPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// 读取文件
	if err := v.ReadInConfig(); err != nil {
		// 如果配置文件不存在且未强制指定，可以仅依赖环境变量
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("read config file: %w", err)
		}
		log.Println("[config] no config file found, using env vars and defaults")
	} else {
		log.Printf("[config] loaded from: %s", v.ConfigFileUsed())
	}

	// 解析到结构体
	cfg, err := decode(v)
	if err != nil {
		return nil, err
	}

	// 校验必填项
	if err := validate(cfg); err != nil {
		return nil, err
	}

	// 热重载
	if opt.WatchFile {
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			log.Printf("[config] file changed: %s", e.Name)
			newCfg, err := decode(v)
			if err != nil {
				log.Printf("[config] reload error: %v", err)
				return
			}
			mu.Lock()
			globalConfig = newCfg
			mu.Unlock()
			if opt.OnChange != nil {
				opt.OnChange(newCfg)
			}
		})
	}

	return cfg, nil
}

func decode(v *viper.Viper) (*Config, error) {
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &cfg, nil
}

func validate(cfg *Config) error {
	var errs []string

	if cfg.App.Name == "" {
		errs = append(errs, "app.name is required")
	}
	if cfg.Server.Port <= 0 || cfg.Server.Port > 65535 {
		errs = append(errs, "server.port must be between 1 and 65535")
	}
	if cfg.App.IsProd() && cfg.JWT.Secret == "" {
		errs = append(errs, "jwt.secret is required in production (set JWT_SECRET env var)")
	}
	if cfg.App.IsProd() && cfg.Database.Password == "" {
		errs = append(errs, "database.password is required in production (set DATABASE_PASSWORD env var)")
	}

	if len(errs) > 0 {
		return fmt.Errorf("config validation failed:\n  - %s", strings.Join(errs, "\n  - "))
	}
	return nil
}

func mergeOptions(opts ...Options) Options {
	opt := Options{
		ConfigFile: "config.yml",
		ConfigType: "yaml",
		EnvPrefix:  "APP",
	}
	if opts == nil {
		return opt
	}
	if len(opts) > 0 {
		o := opts[0]
		if o.ConfigFile != "" {
			opt.ConfigFile = o.ConfigFile
		}
		if o.ConfigType != "" {
			opt.ConfigType = o.ConfigType
		}
		if o.EnvPrefix != "" {
			opt.EnvPrefix = o.EnvPrefix
		}
		opt.WatchFile = o.WatchFile
		opt.OnChange = o.OnChange
		if envFile := os.Getenv("CONFIG_FILE"); envFile != "" && opt.ConfigFile == "" {
			opt.ConfigFile = envFile
		}
	}
	return opt
}
