package problem0209

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

func TestProblem0209(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{
				[]int{2, 3, 1, 2, 4, 3},
				7,
			},
			ans{
				2,
			},
		},
		{
			para{
				[]int{1, 4, 4},
				4,
			},
			ans{
				1,
			},
		},
		{
			para{
				[]int{1, 1, 1, 1, 1, 1, 1, 1},
				11,
			},
			ans{
				0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		i := MinSubArrayLen(p.k, p.nums)
		fmt.Println("Result ", i)
		fmt.Println("Expected  ", a.one)
		ast.Equal(a.one, i)
	}
}
