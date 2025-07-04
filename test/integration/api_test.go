package integration_test

import (
	"base-gin/wire"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserAPI(t *testing.T) {
	// 初始化应用
	app, cleanup, err := wire.InitializeApp()
	if err != nil {
		t.Fatalf("初始化应用失败: %v", err)
	}
	defer cleanup()

	// 测试获取所有用户
	t.Run("GetAllUsers", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/users", nil)
		w := httptest.NewRecorder()
		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("期望状态码 %d，得到 %d", http.StatusOK, w.Code)
		}
	})

	// 测试获取单个用户
	t.Run("GetUser", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
		w := httptest.NewRecorder()
		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("期望状态码 %d，得到 %d", http.StatusOK, w.Code)
		}
	})

	// 测试创建用户
	t.Run("CreateUser", func(t *testing.T) {
		user := map[string]interface{}{
			"name":     "测试用户",
			"email":    "test@example.com",
			"password": "password123",
		}

		body, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("期望状态码 %d，得到 %d", http.StatusCreated, w.Code)
		}
	})

	// 测试健康检查
	t.Run("HealthCheck", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()
		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("期望状态码 %d，得到 %d", http.StatusOK, w.Code)
		}
	})
}
