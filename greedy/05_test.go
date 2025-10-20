package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	//输入：s = "babad"
	//输出："bab"
	//解释："aba" 同样是符合题意的答案。
	s := "babad"
	ans := longestPalindrome(s)
	if ans != "bab" {
		t.Errorf("longestPalindrome(%v) = %v; want bab", s, ans)
	}
	//输入：s = "cbbd"
	//输出："bb"
	s = "cbbd"
	ans = longestPalindrome(s)
	if ans != "bb" {
		t.Errorf("longestPalindrome(%v) = %v; want bb", s, ans)
	}
}
