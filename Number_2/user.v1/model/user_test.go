package model

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/google/uuid"
)

func TestIsValid(t *testing.T) {
  tests := []struct {
    name      string
    request   *User
    expected  bool
  }{
    {
      name: "PassTest",
      request: &User{
		  ID:uuid.MustParse("83064af3-bb81-4514-a6d4-afba340825cd"),
		  Name:"Hasbi",
		  Email:"hsbdeveloper97@gmail.com",
		  Password:"6b8b5b14be25345f0f29975a47447391571af2b266a8a2c76e5b07418b4f5c96",
	  },
      expected: true,
    },
	{
		name: "nilID",
		request: &User{
			Name:"Hasbi",
			Email:"hsbdeveloper97@gmail.com",
			Password:"6b8b5b14be25345f0f29975a47447391571af2b266a8a2c76e5b07418b4f5c96",
		},
		expected: false,
	  },
	  {
		name: "nilName",
		request: &User{
			ID:uuid.MustParse("83064af3-bb81-4514-a6d4-afba340825cd"),
			Email:"hsbdeveloper97@gmail.com",
			Password:"6b8b5b14be25345f0f29975a47447391571af2b266a8a2c76e5b07418b4f5c96",
		},
		expected: false,
	  },
	  {
		name: "nilEmail",
		request: &User{
			ID:uuid.MustParse("83064af3-bb81-4514-a6d4-afba340825cd"),
			Name:"Hasbi",
			Password:"6b8b5b14be25345f0f29975a47447391571af2b266a8a2c76e5b07418b4f5c96",
		},
		expected: false,
	  },
	  {
		name: "nilPassword",
		request: &User{
			ID:uuid.MustParse("83064af3-bb81-4514-a6d4-afba340825cd"),
			Name:"Hasbi",
			Email:"hsbdeveloper97@gmail.com",
		},
		expected: false,
	  },
  }

  for _, test := range tests {
    t.Run(test.name, func(t *testing.T)  {
      result := test.request.isValid()
      assert.Equal(t, test.expected, result)
    })
  }
}