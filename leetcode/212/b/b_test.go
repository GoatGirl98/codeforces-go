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
			`[4,6,5,9,3,7]`, `[0,0,2]`, `[2,3,5]`, 
			`[true,false,true]`,
		},
		{
			`[-12,-9,-3,-12,-6,15,20,-25,-20,-15,-10]`, `[0,1,6,4,8,7]`, `[4,4,9,7,9,10]`,
			`[false,true,false,false,true,true]`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, checkArithmeticSubarrays, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-212/problems/arithmetic-subarrays/
