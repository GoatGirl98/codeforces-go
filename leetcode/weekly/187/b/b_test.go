// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[1,0,0,0,1,0,0,1]`, `2`, 
			`true`, 
		},
		{
			`[1,0,0,1,0,1]`, `2`, 
			`false`, 
		},
		{
			`[1,1,1,1,1]`, `0`, 
			`true`, 
		},
		{
			`[0,1,0,1]`, `1`, 
			`true`, 
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, kLengthApart, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-187/problems/check-if-all-1s-are-at-least-length-k-places-away/