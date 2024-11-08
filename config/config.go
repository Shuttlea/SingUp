package config

import (
	"log/slog"

	"github.com/gofor-little/env"
)

func InitConfig() (*Config, error) {
	cfg := Config{}
	err := env.Load("config/config.env")
	if err != nil {
		return nil, err
	}
	slog.Info("Config file loaded")
	cfg.HostAddr, err = env.MustGet("HOST_ADDR")
	if err != nil {
		return nil, err
	}
	cfg.HostDB, err = env.MustGet("DATABASE_HOST")
	if err != nil {
		return nil, err
	}
	cfg.UserDB, err = env.MustGet("DATABASE_USER")
	if err != nil {
		return nil, err
	}
	// cfg.PasswordDB,err = env.MustGet("DATABASE_PASSWORD")
	// if err!= nil{
	// 	return nil, err
	// }
	cfg.DataBase, err = env.MustGet("DATABASE_DB")
	if err != nil {
		return nil, err
	}
	cfg.PortDB, err = env.MustGet("DATABASE_PORT")
	if err != nil {
		return nil, err
	}
	cfg.SslmodeDB, err = env.MustGet("DATABASE_SSLMODE")
	if err != nil {
		return nil, err
	}
	cfg.LogLevel, err = env.MustGet("LOG_LEVEL")
	if err != nil {
		return nil, err
	}
	cfg.MailBox, err = env.MustGet("MAIL_BOX")
	if err != nil {
		return nil, err
	}
	cfg.MailUser, err = env.MustGet("MAIL_USER")
	if err != nil {
		return nil, err
	}
	cfg.MailPass, err = env.MustGet("MAIL_PASS")
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

type Config struct {
	HostAddr   string `env:"HOST_ADDR"`
	HostDB     string `env:"DATABASE_HOST"`
	UserDB     string `env:"DATABASE_USER"`
	PasswordDB string `env:"DATABASE_PASSWORD"`
	DataBase   string `env:"DATABASE_DB"`
	PortDB     string `env:"DATABASE_PORT"`
	SslmodeDB  string `env:"DATABASE_SSLMODE"`
	LogLevel   string `env:"LOG_LEVEL"`
	MailBox    string `env:"MAIL_BOX"`
	MailUser   string `env:"MAIL_USER"`
	MailPass   string `env:"MAIL_PASS"`
}
