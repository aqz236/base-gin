package logging

import (
	"base-gin/configs"
	"log"
	"os"
)

type Logger struct {
	config *configs.LogConfig
}

func NewLogger(config *configs.Config) *Logger {
	logger := &Logger{
		config: &config.Log,
	}

	// 在真实项目中，这里会配置日志输出
	log.Printf("日志配置: Level=%s, File=%s",
		logger.config.Level, logger.config.File)

	return logger
}

func (l *Logger) Info(message string) {
	log.Printf("[INFO] %s", message)
}

func (l *Logger) Error(message string) {
	log.Printf("[ERROR] %s", message)
}

func (l *Logger) Debug(message string) {
	if l.config.Level == "debug" {
		log.Printf("[DEBUG] %s", message)
	}
}

func (l *Logger) SetOutput(file string) error {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(f)
	return nil
}
