package main

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	// 1. 测试用例1
	n := 3
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	result := generateParenthesis(n)
	if len(result) != len(expected) {
		t.Errorf("generateParenthesis(%d) = %v; want %v", n, result, expected)
	}
	// 2. 测试用例2
	//n = 1
	//expected = []string{"()"}
	//result = generateParenthesis(n)
	//if len(result) != len(expected) {
	//	t.Errorf("generateParenthesis(%d) = %v; want %v", n, result, expected)
	//}
	//// 3. 测试用例3
	//n = 2
	//expected = []string{"(())", "()()"}
	//result = generateParenthesis(n)
	//if len(result) != len(expected) {
	//	t.Errorf("generateParenthesis(%d) = %v; want %v", n, result, expected)
	//}
}
