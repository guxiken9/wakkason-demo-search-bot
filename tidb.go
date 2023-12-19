package main

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Memory struct {
	ID            int       `json:"memory_id" gorm:"primaryKey;column:memory_id"`
	Title         string    `json:"title" gorm:"column:title"`
	Memory        string    `json:"memory" gorm:"column:memory"`
	Image         string    `json:"image" gorm:"column:image"`
	PhotoOrignURL string    `json:"photo_origin_url" gorm:"column:photo_origin_url"`
	PhotoURL      string    `json:"photo_url" gorm:"column:photo_url"`
	CreatedBy     int       `json:"created_by" gorm:"column:created_by"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
}

type TiDBMessage struct {
	ID            int       `json:"message_id" gorm:"primaryKey;column:message_id"`
	ToUser        int       `json:"to_user" gorm:"column:to_User"`
	FromUser      int       `json:"from_user" gorm:"column:from_User"`
	Title         string    `json:"title" gorm:"column:title"`
	Message       string    `json:"message" gorm:"column:message"`
	Image         string    `json:"image" gorm:"column:image"`
	PhotoOrignURL string    `json:"photo_origin_url" gorm:"column:photo_origin_url"`
	PhotoURL      string    `json:"photo_url" gorm:"column:photo_url"`
	ScheduledTime time.Time `json:"scheduled_time" gorm:"column:scheduled_time"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
}

func Search(key string) ([]Memory, error) {

	var m []Memory
	db, err := CreateDB()
	if err != nil {
		return nil, err
	}

	db.Limit(1).Where("Memory LIKE ?", "%"+key+"%").Find(&m)

	// TiDB上からLIKE検索
	return m, nil
}

func CreateDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(getDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDSN() string {
	tidbHost := getEnvWithDefault("TIDB_HOST", "127.0.0.1")
	tidbPort := getEnvWithDefault("TIDB_PORT", "4000")
	tidbUser := getEnvWithDefault("TIDB_USER", "root")
	tidbPassword := getEnvWithDefault("TIDB_PASSWORD", "")
	tidbDBName := getEnvWithDefault("TIDB_DB_NAME", "test")
	useSSL := getEnvWithDefault("USE_SSL", "true")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&tls=%s",
		tidbUser, tidbPassword, tidbHost, tidbPort, tidbDBName, useSSL)
}

func getEnvWithDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
