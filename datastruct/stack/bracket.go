package main

/*
1、括号匹配
*/

func IsValidBrackets(s string) bool {
	if len(s) == 0 {
		return true
	}
	m := map[byte]byte{'}': '{', ')': '(', ']': '['}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '{' || b == '(' || b == '[' {
			stack = append(stack, b)
		} else if len(stack) != 0 && stack[len(stack)-1] == m[b] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
