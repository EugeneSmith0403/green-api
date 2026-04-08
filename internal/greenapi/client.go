package greenapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

type UpstreamError struct {
	StatusCode int
	Body       string
}

func (e *UpstreamError) Error() string {
	return fmt.Sprintf("upstream returned %d: %s", e.StatusCode, e.Body)
}

type SendMessagePayload struct {
	ChatID  string `json:"chatId"`
	Message string `json:"message"`
}

type SendFileByURLPayload struct {
	ChatID   string `json:"chatId"`
	URLFile  string `json:"urlFile"`
	FileName string `json:"fileName,omitempty"`
	Caption  string `json:"caption,omitempty"`
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) GetSettings(idInstance, apiToken string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/waInstance%s/getSettings/%s", c.baseURL, idInstance, apiToken)
	return c.get(endpoint)
}

func (c *Client) GetStateInstance(idInstance, apiToken string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", c.baseURL, idInstance, apiToken)
	return c.get(endpoint)
}

func (c *Client) SendMessage(idInstance, apiToken, chatID, message string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", c.baseURL, idInstance, apiToken)
	return c.post(endpoint, SendMessagePayload{ChatID: chatID, Message: message})
}

func (c *Client) SendFileByURL(idInstance, apiToken string, p SendFileByURLPayload) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s", c.baseURL, idInstance, apiToken)
	return c.post(endpoint, p)
}

func (c *Client) get(endpoint string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	return c.do(req)
}

func (c *Client) post(endpoint string, body any) ([]byte, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshal body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(req)
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, &UpstreamError{StatusCode: resp.StatusCode, Body: string(body)}
	}

	return body, nil
}
