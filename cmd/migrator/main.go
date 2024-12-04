package main

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"os"
)

type Config struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	DBName  string `yaml:"db_name"`
	User    string `yaml:"username"`
	SSLMode string `yaml:"ssl_mode"`
	IsDrop  bool   `yaml:"is_drop"`
}

func main() {
	cfg := MustLoad()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, os.Getenv("DB_PASSWORD"), cfg.DBName, cfg.SSLMode))

	if err != nil {
		panic(err)
	}

	var migrationsPath string

	flag.StringVar(&migrationsPath, "migration-path", "", "path to migrations folder")
	flag.Parse()

	if cfg.IsDrop {
		if err = goose.Down(db.DB, migrationsPath); err != nil {
			panic(err)
		}
	}

	if err = goose.Up(db.DB, migrationsPath); err != nil {
		panic(err)
	}
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config path does not exist")
	}

	var config Config
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("Error read config: " + err.Error())
	}
	return &config
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "config path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_MIGRATION_PATH")
	}
	return res
}
