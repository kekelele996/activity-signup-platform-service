package service

import (
	"log/slog"
	"os"
	"testing"

	"gbevent/internal/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newEvDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&model.User{}, &model.Activity{}, &model.Registration{}, &model.CheckInRecord{}, &model.Notification{}, &model.Comment{}, &model.Favorite{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func evLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
}
