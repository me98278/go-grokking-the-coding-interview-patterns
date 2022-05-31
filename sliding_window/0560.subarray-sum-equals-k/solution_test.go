package problem0560

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	nums []int
	k    int
}

type ans struct {
	one int
}

func TestProblem0560(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{
				[]int{1, 1, 1},
				2,
			},
			ans{
				2,
			},
		},
		{
			para{
				[]int{2, 1, 5, 1, 3, 2},
				3,
			},
			ans{
				9,
			},
		},
		{
			para{
				[]int{2, 3, 4, 1, 5},
				2,
			},
			ans{
				7,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, SubArraySum(p.nums, p.k))
	}
}
