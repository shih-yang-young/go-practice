package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubAPIClient struct{}

func (s *StubAPIClient) GetData() (string, error) {
	return "stub data", nil
}

func TestService_Process(t *testing.T) {
	stubClient := &StubAPIClient{}
	service := &Service{client: stubClient}

	result, err := service.Process()
	assert.NoError(t, err)
	assert.Equal(t, "Processed: stub data", result)
}
