package lit

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"

	"github.com/joeshaw/envdecode"
)

type AppConfig struct {
	Server serverConfig
	Db     dbConfig
}

type dbConfig struct {
	// DSN string `env:"DB_DSN,default=data/gorm.db"`
	Host     string `env:"DB_HOST,default=localhost"`
	Port     int    `env:"DB_PORT,default=9001"`
	Username string `env:"DB_USERNAME,default=manny"`
	Password string `env:"DB_PASSWORD,default=change-in-production"`
	DbName   string `env:"DB_NAME,default=brain"`

	SslMode  string `env:"DB_SSLMODE,default=disable"`
	TimeZone string `env:"DB_TZ,default=Asia/Almaty"`
}

type serverConfig struct {
	Host string `env:"SERVER_HOST,default=localhost"`
	Port int    `env:"SERVER_PORT,default=3001"`
}

func LoadAppConfig() (*AppConfig, error) {
	var c AppConfig
	if err := envdecode.StrictDecode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func (dbConfig *dbConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName, dbConfig.Port, dbConfig.SslMode, dbConfig.TimeZone)
}
