package unittest

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSortChar(t *testing.T) {
  tests := []struct {
    name      string
    request   string
    expected  string
  }{
    {
      name: "hai",
      request: "hai",
      expected: "ahi",
    },
    {
      name: "hello",
      request: "hello",
      expected: "ehllo",
    },
    {
      name: "bibit",
      request: "bibit",
      expected: "bbiit",
    },
  }

  for _, test := range tests {
    t.Run(test.name, func(t *testing.T)  {
      result := sortChar(test.request)
      assert.Equal(t, test.expected, result)
    })
  }
}

func TestGroupAnagram(t *testing.T) {
  tests := []struct {
    name      string
    request   []string
    expected  map[string][]string
  }{
    {
      name: "oneGroup",
      request: []string{"rumah", "murah"},
      expected: map[string][]string{"ahmru": {"rumah", "murah"}},
    },
    {
      name: "twoGroup",
      request: []string{"rumah", "murah", "kita", "atik", "tika"},
      expected: map[string][]string{"ahmru": {"rumah", "murah"}, "aikt": {"kita", "atik", "tika"}},
    },
    {
      name: "threeGroup",
      request: []string{"rumah", "murah", "kita", "atik", "tika", "ayam", "maya"},
      expected: map[string][]string{"ahmru": {"rumah", "murah"}, "aikt": {"kita", "atik", "tika"}, "aamy": {"ayam", "maya"}},
    },
  }

  for _, test := range tests {
    t.Run(test.name, func(t *testing.T)  {
      result := groupAnagram(test.request)
      assert.Equal(t, test.expected, result)
    })
  }
}