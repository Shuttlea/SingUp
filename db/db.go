package db

import (
	"fmt"
	"log/slog"
	"singup/config"
	sw "singup/go"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	slog.Debug("ConnectDB", "host", cfg.HostDB, "user", cfg.UserDB, "database", cfg.DataBase, "port", cfg.PortDB, "sslmode", cfg.SslmodeDB, "password", cfg.PasswordDB)
	dsn := fmt.Sprintf("host=%s user=%s database=%s port=%s sslmode=%s password=%s", cfg.HostDB, cfg.UserDB, cfg.DataBase, cfg.PortDB, cfg.SslmodeDB, cfg.PasswordDB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Can't open database", "Error", err)
		panic(err)
	}
	err = db.AutoMigrate(&sw.User{})
	if err != nil {
		slog.Error("Can't open database", "Error", err)
		panic(err)
	}
	return db
}

func CloseDBConnection(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
