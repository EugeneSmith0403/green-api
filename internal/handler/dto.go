package handler

import "encoding/json"

type SendMessageRequest struct {
	ChatID  string `json:"chatId" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type SendFileRequest struct {
	ChatID   string `json:"chatId" binding:"required"`
	URLFile  string `json:"urlFile" binding:"required"`
	FileName string `json:"fileName"`
	Caption  string `json:"caption"`
}

type APIResponse struct {
	Success bool            `json:"success"`
	Action  string          `json:"action"`
	Data    json.RawMessage `json:"data,omitempty" swaggertype:"object"`
	Raw     string          `json:"raw,omitempty"`
	Error   string          `json:"error,omitempty"`
}
