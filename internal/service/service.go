package service

import (
	"net/url"
	"path"

	"github.com/evgenijkuznecov/green-api/internal/greenapi"
)

type Service struct {
	client *greenapi.Client
}

func New(client *greenapi.Client) *Service {
	return &Service{client: client}
}

func (s *Service) GetSettings(idInstance, apiToken string) ([]byte, error) {
	return s.client.GetSettings(idInstance, apiToken)
}

func (s *Service) GetStateInstance(idInstance, apiToken string) ([]byte, error) {
	return s.client.GetStateInstance(idInstance, apiToken)
}

func (s *Service) SendMessage(idInstance, apiToken, chatID, message string) ([]byte, error) {
	return s.client.SendMessage(idInstance, apiToken, chatID, message)
}

func (s *Service) SendFileByURL(idInstance, apiToken string, p greenapi.SendFileByURLPayload) ([]byte, error) {
	if p.FileName == "" {
		p.FileName = fileNameFromURL(p.URLFile)
	}
	return s.client.SendFileByURL(idInstance, apiToken, p)
}

func fileNameFromURL(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	name := path.Base(u.Path)
	if name == "." || name == "/" {
		return ""
	}
	return name
}
