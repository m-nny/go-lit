package lit

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"

	"github.com/joeshaw/envdecode"
)

type AppConfig struct {
	Http HttpConfig
	Db   DbConfig
}

type DbConfig struct {
	// DSN string `env:"DB_DSN,default=data/gorm.db"`
	Host     string `env:"DB_HOST,default=localhost"`
	Port     int    `env:"DB_PORT,default=9001"`
	Username string `env:"DB_USERNAME,default=manny"`
	Password string `env:"DB_PASSWORD,default=change-in-production"`
	DbName   string `env:"DB_NAME,default=brain"`

	SslMode  string `env:"DB_SSLMODE,default=disable"`
	TimeZone string `env:"DB_TZ,default=Asia/Almaty"`
}

type HttpConfig struct {
	Host  string `env:"SERVER_HOST,default=localhost"`
	Port  int    `env:"SERVER_PORT,default=3001"`
	Debug bool   `env:"SERVER_DEBUG,default=true"`
}

func LoadAppConfig() (*AppConfig, error) {
	var c AppConfig
	if err := envdecode.StrictDecode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func (dConfig *DbConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dConfig.Host, dConfig.Username, dConfig.Password, dConfig.DbName, dConfig.Port, dConfig.SslMode, dConfig.TimeZone)
}

func (hConfig *HttpConfig) Address() string {
	return fmt.Sprintf("%s:%d", hConfig.Host, hConfig.Port)
}
