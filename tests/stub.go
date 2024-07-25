package tests

import "fmt"

type APIClient interface {
	GetData() (string, error)
}

type RealAPIClient struct{}

func (c *RealAPIClient) GetData() (string, error) {
	// 實際的 API 調用
	return "real data", nil
}

type Service struct {
	client APIClient
}

func (s *Service) Process() (string, error) {
	data, err := s.client.GetData()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Processed: %s", data), nil
}
