package dao

import (
	"fmt"
	"github.com/jiuzi/sshm/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path"
)

func InitDB() *gorm.DB {
	gormConfig := &gorm.Config{}
	configDir, _ := os.UserConfigDir()
	sshmDbPath := path.Join(configDir, "sshm.db")
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s?_journal=WAL&_vacuum=incremental",
		sshmDbPath)), gormConfig)
	if err != nil {
		log.Fatalf("failed to connect database:%s", err.Error())
	}
	err = db.AutoMigrate(model.Machine{})
	if err != nil {
		log.Fatalf("failed to connect database:%s", err.Error())
	}
	return db
}
