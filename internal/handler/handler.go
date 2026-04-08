package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/evgenijkuznecov/green-api/internal/greenapi"
	"github.com/evgenijkuznecov/green-api/internal/service"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

// GetSettings godoc
// @Summary      Get instance settings
// @Tags         green-api
// @Produce      json
// @Param        idInstance  path      string  true  "Instance ID"
// @Param        apiToken    path      string  true  "API token"
// @Success      200         {object}  APIResponse
// @Failure      502         {object}  APIResponse
// @Failure      500         {object}  APIResponse
// @Router       /api/green/instance/{idInstance}/settings/{apiToken} [get]
func (h *Handler) GetSettings(c *gin.Context) {
	id, token := instanceParams(c)
	data, err := h.svc.GetSettings(id, token)
	h.respond(c, "getSettings", data, err)
}

// GetStateInstance godoc
// @Summary      Get instance state
// @Tags         green-api
// @Produce      json
// @Param        idInstance  path      string  true  "Instance ID"
// @Param        apiToken    path      string  true  "API token"
// @Success      200         {object}  APIResponse
// @Failure      502         {object}  APIResponse
// @Failure      500         {object}  APIResponse
// @Router       /api/green/instance/{idInstance}/state/{apiToken} [get]
func (h *Handler) GetStateInstance(c *gin.Context) {
	id, token := instanceParams(c)
	data, err := h.svc.GetStateInstance(id, token)
	h.respond(c, "getStateInstance", data, err)
}

// SendMessage godoc
// @Summary      Send a text message
// @Tags         green-api
// @Accept       json
// @Produce      json
// @Param        idInstance  path      string              true  "Instance ID"
// @Param        apiToken    path      string              true  "API token"
// @Param        request     body      SendMessageRequest  true  "Message payload"
// @Success      200         {object}  APIResponse
// @Failure      400         {object}  APIResponse
// @Failure      502         {object}  APIResponse
// @Failure      500         {object}  APIResponse
// @Router       /api/green/instance/{idInstance}/send-message/{apiToken} [post]
func (h *Handler) SendMessage(c *gin.Context) {
	var req SendMessageRequest
	if !bindJSON(c, &req) {
		return
	}
	id, token := instanceParams(c)
	data, err := h.svc.SendMessage(id, token, req.ChatID, req.Message)
	h.respond(c, "sendMessage", data, err)
}

// SendFileByURL godoc
// @Summary      Send a file by URL
// @Tags         green-api
// @Accept       json
// @Produce      json
// @Param        idInstance  path      string           true  "Instance ID"
// @Param        apiToken    path      string           true  "API token"
// @Param        request     body      SendFileRequest  true  "File payload"
// @Success      200         {object}  APIResponse
// @Failure      400         {object}  APIResponse
// @Failure      502         {object}  APIResponse
// @Failure      500         {object}  APIResponse
// @Router       /api/green/instance/{idInstance}/send-file/{apiToken} [post]
func (h *Handler) SendFileByURL(c *gin.Context) {
	var req SendFileRequest
	if !bindJSON(c, &req) {
		return
	}
	id, token := instanceParams(c)
	data, err := h.svc.SendFileByURL(id, token, greenapi.SendFileByURLPayload{
		ChatID:   req.ChatID,
		URLFile:  req.URLFile,
		FileName: req.FileName,
		Caption:  req.Caption,
	})
	h.respond(c, "sendFileByUrl", data, err)
}

func instanceParams(c *gin.Context) (id, token string) {
	return c.Param("idInstance"), c.Param("apiToken")
}

func bindJSON(c *gin.Context, req any) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return false
	}
	return true
}

func (h *Handler) respond(c *gin.Context, action string, data []byte, err error) {
	if err != nil {
		var upstreamErr *greenapi.UpstreamError
		if errors.As(err, &upstreamErr) {
			resp := APIResponse{
				Success: false,
				Action:  action,
				Error:   fmt.Sprintf("upstream returned %d", upstreamErr.StatusCode),
			}
			if json.Valid([]byte(upstreamErr.Body)) {
				resp.Data = json.RawMessage(upstreamErr.Body)
			} else {
				resp.Raw = upstreamErr.Body
			}
			c.JSON(http.StatusBadGateway, resp)
			return
		}

		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Action:  action,
			Error:   err.Error(),
		})
		return
	}

	resp := APIResponse{
		Success: true,
		Action:  action,
	}
	if json.Valid(data) {
		resp.Data = json.RawMessage(data)
	} else {
		resp.Raw = string(data)
	}
	c.JSON(http.StatusOK, resp)
}
