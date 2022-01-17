package service

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/google/uuid"
)

func TestCreateToken(t *testing.T) {
	request := uuid.MustParse("83064af3-bb81-4514-a6d4-afba340825cd")
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjoiODMwNjRhZjMtYmI4MS00NTE0LWE2ZDQtYWZiYTM0MDgyNWNkIn0.i4zS0xvjXoxEk_FkhNBu3UafQh9DJyYnmUMeUF2lW6I"
	result, _ := CreateToken(request)
    assert.Equal(t, expected, result)
}