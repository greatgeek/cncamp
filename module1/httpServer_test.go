package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestHttpServer 单元测试
func TestHealthzHandler(t *testing.T) {
	// 设置环境变量 VERSION
	VERSION := "1.0.0"
	os.Setenv("VERSION", VERSION)

	// 创建一个请求来传递给我们的处理程序
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 创建一个 ResponseRecorder 来记录响应
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(healthzHandler)

	// 创建一个请求并用将 ResponseRecorder 放入处理程序中
	handler.ServeHTTP(recorder, req)

	// 检查状态码是否为 200
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 检查返回的 header 是否符合预期
	expected := VERSION
	if recorder.Result().Header.Get("VERSION") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Result().Header.Get("VERSION"), expected)
	}
}
