package database

import (
	"base-gin/configs"
	"base-gin/internal/infrastructure/database/models"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	config *configs.DatabaseConfig
	gormDB *gorm.DB
}

func NewDB(config *configs.Config) *DB {
	db := &DB{
		config: &config.Database,
	}

	// 初始化GORM连接
	if err := db.initGORM(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 自动迁移数据库表
	if err := db.migrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Printf("数据库连接成功: %s", db.getDatabaseType())
	return db
}

func (db *DB) initGORM() error {
	var dialector gorm.Dialector
	var err error

	// 根据配置选择数据库驱动
	switch db.getDatabaseType() {
	case "sqlite":
		// 使用SQLite - 确保数据目录存在
		dbPath := "data/app.db"
		if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
			return fmt.Errorf("failed to create database directory: %v", err)
		}
		dialector = sqlite.Open(dbPath)
	case "postgres":
		// 使用PostgreSQL
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
			db.config.Host, db.config.Port, db.config.Username, db.config.Password, db.config.Database)
		dialector = postgres.Open(dsn)
	default:
		// 默认使用SQLite - 确保数据目录存在
		dbPath := "data/app.db"
		if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
			return fmt.Errorf("failed to create database directory: %v", err)
		}
		dialector = sqlite.Open(dbPath)
	}

	// 配置GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db.gormDB, err = gorm.Open(dialector, gormConfig)
	return err
}

func (db *DB) migrate() error {
	// 自动迁移所有模型
	if err := db.gormDB.AutoMigrate(&models.UserModel{}); err != nil {
		return err
	}

	// 为SQLite创建支持软删除的唯一索引
	if db.getDatabaseType() == "sqlite" {
		// 删除可能存在的旧索引
		// db.gormDB.Exec("DROP INDEX IF EXISTS idx_users_email")
		// db.gormDB.Exec("DROP INDEX IF EXISTS idx_users_email_active")

		// 创建支持软删除的部分唯一索引（只对未删除的记录）
		return db.gormDB.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email_active ON users(email) WHERE deleted_at IS NULL").Error
	} else {
		// PostgreSQL支持部分索引
		return db.gormDB.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email_active ON users(email) WHERE deleted_at IS NULL").Error
	}
}

func (db *DB) getDatabaseType() string {
	// 如果Host为空或为localhost且数据库名为sqlite，则使用SQLite
	if db.config.Host == "" || (db.config.Host == "localhost" && db.config.Database == "sqlite") {
		return "sqlite"
	}
	return "postgres"
}

// GetGormDB 获取GORM数据库实例
func (db *DB) GetGormDB() *gorm.DB {
	return db.gormDB
}

func (db *DB) GetConnectionString() string {
	switch db.getDatabaseType() {
	case "sqlite":
		return "data/app.db"
	default:
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			db.config.Host, db.config.Port, db.config.Username, db.config.Password, db.config.Database)
	}
}

func (db *DB) Close() error {
	if db.gormDB != nil {
		if sqlDB, err := db.gormDB.DB(); err == nil {
			return sqlDB.Close()
		}
	}
	log.Println("数据库连接已关闭")
	return nil
}
