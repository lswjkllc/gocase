package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// 测试括号匹配
func TestBracket(t *testing.T) {
	assert.Equal(t, true, IsValidBrackets("(){}[]"))
	assert.Equal(t, true, IsValidBrackets("[{()}]"))
	assert.Equal(t, false, IsValidBrackets("([{()}]"))
	assert.Equal(t, false, IsValidBrackets("[{}}]"))
}

func BenchmarkBracket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidBrackets("[{}]")
	}
}
