package main_test

import (
	"context"
	"dcard-golang-project/middlewares"
	"dcard-golang-project/models"
	"dcard-golang-project/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDBInit(t *testing.T) {

	if models.DB == nil {
		t.Error("TestDBInit failed: Failed to initialize database")
	}
}

func TestRedisInit(t *testing.T) {
	ctx := context.Background()
	_, err := models.Client.Ping(ctx).Result()
	if err != nil {
		t.Errorf("TestRedisInit failed: %v", err)
	}
}

func TestLoadLuaScript(t *testing.T) {
	if utils.LuaHash1 == "" {
		t.Error("TestLoadLuaScript failed: Failed to preload Lua script 1")
	}

	if utils.LuaHash2 == "" {
		t.Error("TestLoadLuaScript failed: Failed to preload Lua script 2")
	}
}

func TestErrorHandler(t *testing.T) {

	// 	/* 新建立一個 Engine 模擬一個錯誤來確定 ErrorHandler 是否正常運行 */

	router := gin.New()

	router.Use(middlewares.ErrorHandler())

	router.GET("/errortest", func(c *gin.Context) {
		utils.ErrorCheck()
	})

	req, _ := http.NewRequest("GET", "/errortest", nil)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != 400 {
		t.Errorf("TestErrorHandler failed, expected status code 400, got %d", resp.Code)
	}

	expectedErrorMessage := `{"error":"custom error for test"}`
	if body := resp.Body.String(); body != expectedErrorMessage {
		t.Errorf("TestErrorHandler failed, expected body %s, got %s", expectedErrorMessage, body)
	}
}
