package problem0340

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
	s string
	k int
}

type ans struct {
	length int
}

func TestProblem0340(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{
				"araaci",
				2,
			},
			ans{
				4,
			},
		},
		{
			para{
				"araaci",
				1,
			},
			ans{
				2,
			},
		},
		{
			para{
				"cbbebi",
				3,
			},
			ans{
				5,
			},
		},
		{
			para{
				"cbbebi",
				10,
			},
			ans{
				6,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		i := LengthOfLongestSubstringKDistrinct(p.s, p.k)
		fmt.Println("Result ", i)
		fmt.Println("Expected  ", a.length)
		ast.Equal(a.length, i)

	}
}
