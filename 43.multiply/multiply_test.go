package frytea_leetcode

import (
	"fmt"
	"testing"
)

type question43 struct {
	para43
	ans43
}

// para 是参数
// one 代表第一个参数
type para43 struct {
	num1 string
	num2 string
}

// ans 是答案
// one 代表第一个答案
type ans43 struct {
	one string
}

func Test_Problem43(t *testing.T) {

	qs := []question43{

		{
			para43{"2", "3"},
			ans43{"6"},
		},

		{
			para43{"123", "456"},
			ans43{"56088"},
		},
		{
			para43{"498828660196", "840477629533"},
			ans43{"419254329864656431168468"},
		},
		{

			para43{"0", "840477629533"},
			ans43{"0"},
		},
		{

			para43{"498828660196", "0"},
			ans43{"0"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 43------------------------\n")
	for _, q := range qs {
		a, p := q.ans43, q.para43
		r := multiply(p.num1, p.num2)
		fmt.Printf(" 【output】:%v/%v 【input】:%v * %v\n", r, a.one, p.num1, p.num2)
		if r != a.one {
			t.Errorf("【verify】:%v != %v", r, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
