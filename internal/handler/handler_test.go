package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/evgenijkuznecov/green-api/internal/greenapi"
	"github.com/evgenijkuznecov/green-api/internal/handler"
	"github.com/evgenijkuznecov/green-api/internal/service"
)

const (
	testID    = "1101000001"
	testToken = "test-token"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func newRouter(upstreamURL string) *gin.Engine {
	h := handler.New(service.New(greenapi.NewClient(upstreamURL)))
	r := gin.New()
	h.RegisterRoutes(r)
	return r
}

func basePath(action string) string {
	return fmt.Sprintf("/api/green/instance/%s/%s/%s", testID, action, testToken)
}

func doGet(r *gin.Engine, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func doPost(r *gin.Engine, path string, body any) *httptest.ResponseRecorder {
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func decodeResponse(t *testing.T, w *httptest.ResponseRecorder) handler.APIResponse {
	t.Helper()
	var resp handler.APIResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	return resp
}

func upstreamJSON(payload string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(payload))
	}))
}

func upstreamError(status int, body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
}

func TestGetSettings_Success(t *testing.T) {
	upstream := upstreamJSON(`{"wid":"71234567890@c.us","webhookUrl":""}`)
	defer upstream.Close()

	w := doGet(newRouter(upstream.URL), basePath("settings"))

	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d: %s", w.Code, w.Body.String())
	}
	resp := decodeResponse(t, w)
	if !resp.Success {
		t.Errorf("want success=true, error: %s", resp.Error)
	}
	if resp.Action != "getSettings" {
		t.Errorf("want action=getSettings, got %s", resp.Action)
	}
	if resp.Data == nil {
		t.Error("want data != nil")
	}
}

func TestGetStateInstance_Success(t *testing.T) {
	upstream := upstreamJSON(`{"stateInstance":"authorized"}`)
	defer upstream.Close()

	w := doGet(newRouter(upstream.URL), basePath("state"))

	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
	resp := decodeResponse(t, w)
	if !resp.Success {
		t.Errorf("want success=true, error: %s", resp.Error)
	}
	if resp.Action != "getStateInstance" {
		t.Errorf("want action=getStateInstance, got %s", resp.Action)
	}
}

func TestSendMessage_Success(t *testing.T) {
	upstream := upstreamJSON(`{"idMessage":"BAE5F4886F6F2D05"}`)
	defer upstream.Close()

	w := doPost(newRouter(upstream.URL), basePath("send-message"), map[string]string{
		"chatId": "71234567890@c.us", "message": "Hello",
	})

	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
	resp := decodeResponse(t, w)
	if !resp.Success {
		t.Errorf("want success=true, error: %s", resp.Error)
	}
	if resp.Action != "sendMessage" {
		t.Errorf("want action=sendMessage, got %s", resp.Action)
	}
}

func TestSendFileByUrl_Success(t *testing.T) {
	upstream := upstreamJSON(`{"idMessage":"BAE5F4886F6F2D06"}`)
	defer upstream.Close()

	w := doPost(newRouter(upstream.URL), basePath("send-file"), map[string]string{
		"chatId": "71234567890@c.us", "urlFile": "https://example.com/file.pdf",
	})

	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
	resp := decodeResponse(t, w)
	if !resp.Success {
		t.Errorf("want success=true, error: %s", resp.Error)
	}
	if resp.Action != "sendFileByUrl" {
		t.Errorf("want action=sendFileByUrl, got %s", resp.Action)
	}
}

func TestGetSettings_UpstreamError(t *testing.T) {
	upstream := upstreamError(http.StatusUnauthorized, `{"error":"unauthorized"}`)
	defer upstream.Close()

	w := doGet(newRouter(upstream.URL), basePath("settings"))

	if w.Code != http.StatusBadGateway {
		t.Fatalf("want 502, got %d", w.Code)
	}
	resp := decodeResponse(t, w)
	if resp.Success {
		t.Error("want success=false")
	}
	if resp.Error == "" {
		t.Error("want non-empty error message")
	}
}

func TestValidationErrors(t *testing.T) {
	r := newRouter("http://unused")

	cases := []struct {
		name string
		path string
		body map[string]string
	}{
		{
			"send_message_missing_chatId",
			basePath("send-message"),
			map[string]string{"message": "hi"},
		},
		{
			"send_message_missing_message",
			basePath("send-message"),
			map[string]string{"chatId": "123@c.us"},
		},
		{
			"send_file_missing_urlFile",
			basePath("send-file"),
			map[string]string{"chatId": "123@c.us"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := doPost(r, tc.path, tc.body)
			if w.Code != http.StatusBadRequest {
				t.Fatalf("want 400, got %d: %s", w.Code, w.Body.String())
			}
			resp := decodeResponse(t, w)
			if resp.Success {
				t.Error("want success=false")
			}
			if resp.Error == "" {
				t.Error("want non-empty error")
			}
		})
	}
}
