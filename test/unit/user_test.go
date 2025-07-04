package user_test

import (
	"base-gin/internal/domain/user/entity"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name      string
		userName  string
		email     string
		password  string
		wantError bool
	}{
		{
			name:      "valid user",
			userName:  "张三",
			email:     "zhangsan@example.com",
			password:  "password123",
			wantError: false,
		},
		{
			name:      "empty name",
			userName:  "",
			email:     "test@example.com",
			password:  "password123",
			wantError: true,
		},
		{
			name:      "invalid email",
			userName:  "张三",
			email:     "invalid-email",
			password:  "password123",
			wantError: true,
		},
		{
			name:      "short password",
			userName:  "张三",
			email:     "test@example.com",
			password:  "123",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := entity.NewUser(tt.userName, tt.email, tt.password)

			if tt.wantError {
				if err == nil {
					t.Errorf("期望错误，但没有得到错误")
				}
				return
			}

			if err != nil {
				t.Errorf("不期望错误，但得到错误: %v", err)
				return
			}

			if user.Name != tt.userName {
				t.Errorf("期望用户名 %s，得到 %s", tt.userName, user.Name)
			}

			if user.Email != tt.email {
				t.Errorf("期望邮箱 %s，得到 %s", tt.email, user.Email)
			}

			if user.Password != tt.password {
				t.Errorf("期望密码 %s，得到 %s", tt.password, user.Password)
			}
		})
	}
}

func TestUserValidate(t *testing.T) {
	user := &entity.User{
		ID:        1,
		Name:      "张三",
		Email:     "zhangsan@example.com",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := user.Validate(); err != nil {
		t.Errorf("有效用户验证失败: %v", err)
	}

	// 测试无效邮箱
	user.Email = "invalid-email"
	if err := user.Validate(); err == nil {
		t.Error("期望邮箱验证失败，但验证通过")
	}
}
