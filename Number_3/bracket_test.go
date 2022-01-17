package unittest

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestFindFirstStringInBracket(t *testing.T) {
  tests := []struct {
    name      string
    request   string
    expected  string
  }{
    {
      name: "normal",
      request: "(Bibit)",
      expected: "Bibit",
    },
    {
      name: "twoBracket",
      request: "(Bibit)(Stockbit)",
      expected: "Bibit",
    },
    {
      name: "nestedBracket",
      request: "(Nama saya (Bibit))",
      expected: "Nama saya (Bibit)",
    },
    {
      name: "tooMany_)",
      request: "((Bibit))))))",
      expected: "(Bibit)",
    },
    {
      name: "tooMany_(",
      request: "(((((((((((Bibit))",
      expected: "",
    },
    {
      name: "startWith_)",
      request: "))))(Bibit)",
      expected: "Bibit",
    },
    {
      name: "nullString",
      request: "",
      expected: "",
    },
  }

  for _, test := range tests {
    t.Run(test.name, func(t *testing.T)  {
      result := findFirstStringInBracket(test.request)
      assert.Equal(t, test.expected, result)
    })
  }
}