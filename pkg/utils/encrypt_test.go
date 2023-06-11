package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"asd.asd", "dvFbab4Wnj"},
		{"sad.asd", "H4iUYyizdo"},
		{"google.com", "9Txmx1CPWV"},
		{"qwee.qweqwe", "vfS0QpXBCN"},
	}
	for _, c := range tests {
		got := Encrypt(c.in)
		assert.Equal(t, got, c.want)
	}
}
