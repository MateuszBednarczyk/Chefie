package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewJwtService(t *testing.T) {
	result := NewJwtService()
	assert.Equal(t, &jwtService{}, result)
}

func TestJwtService_GenerateTokens(t *testing.T) {
	serviceInstance := NewJwtService()
	result := serviceInstance.GenerateTokens("us")

	assert.True(t, len(result.Content) > 0, len(result.Content))
	assert.True(t, result.Code == 200, result.Code)
}

func TestJwtService_IsTokenValid(t *testing.T) {
	serviceInstance := NewJwtService()
	result := serviceInstance.IsTokenValid("1212121212")

	assert.True(t, result.Message != "Correct token", result.Message)
}
