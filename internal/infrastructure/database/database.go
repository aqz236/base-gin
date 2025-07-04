package database

import (
	"base-gin/configs"
	"fmt"
	"log"
)

type DB struct {
	config *configs.DatabaseConfig
}

func NewDB(config *configs.Config) *DB {
	db := &DB{
		config: &config.Database,
	}

	// 在真实项目中，这里会初始化数据库连接
	log.Printf("数据库配置: Host=%s, Port=%d, Database=%s",
		db.config.Host, db.config.Port, db.config.Database)

	return db
}

func (db *DB) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.config.Host, db.config.Port, db.config.Username, db.config.Password, db.config.Database)
}

func (db *DB) Close() error {
	// 在真实项目中，这里会关闭数据库连接
	log.Println("数据库连接已关闭")
	return nil
}
