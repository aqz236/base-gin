package models

import (
	"base-gin/internal/domain/user/entity"
	"time"

	"gorm.io/gorm"
)

// UserModel GORM数据模型，用于数据库操作
type UserModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"type:varchar(50);not null" json:"name"`
	Email     string         `gorm:"type:varchar(100);not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定表名
func (UserModel) TableName() string {
	return "users"
}

// ToEntity 将GORM模型转换为领域实体
func (m *UserModel) ToEntity() *entity.User {
	return &entity.User{
		ID:        int(m.ID),
		Name:      m.Name,
		Email:     m.Email,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// FromEntity 从领域实体创建GORM模型
func (m *UserModel) FromEntity(user *entity.User) {
	if user.ID > 0 {
		m.ID = uint(user.ID)
	}
	m.Name = user.Name
	m.Email = user.Email
	m.Password = user.Password
	m.CreatedAt = user.CreatedAt
	m.UpdatedAt = user.UpdatedAt
}

// NewUserModelFromEntity 从领域实体创建新的GORM模型
func NewUserModelFromEntity(user *entity.User) *UserModel {
	model := &UserModel{}
	model.FromEntity(user)
	return model
}
