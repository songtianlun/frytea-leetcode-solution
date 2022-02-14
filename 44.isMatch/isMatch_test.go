package frytea_leetcode

import (
	"fmt"
	"testing"
)

type question44 struct {
	para44
	ans44
}

// para 是参数
// one 代表第一个参数
type para44 struct {
	one string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans44 struct {
	one bool
}

func Test_Problem44(t *testing.T) {

	qs := []question44{

		{
			para44{"aa", "a"},
			ans44{false},
		},
		{
			para44{"aa", "*"},
			ans44{true},
		},
		{
			para44{"cb", "?a"},
			ans44{false},
		},
		{
			para44{"adceb", "*a*b"},
			ans44{true},
		},
		{
			para44{"ab", "?*"},
			ans44{true},
		},
		{
			para44{"a", "aa"},
			ans44{false},
		},
		{
			para44{"", "******"},
			ans44{true},
		},
		{
			para44{"abcabczzzde", "*abc???de*"},
			ans44{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 44------------------------\n")
	for _, q := range qs {
		a, p := q.ans44, q.para44
		r := isMatch(p.one, p.two)
		fmt.Printf(" 【output】:%v/%v 【input】:%v %v\n", r, a.one, p.one, p.two)
		if r != a.one {
			t.Errorf("【verify】:%v != %v", r, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
