package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type data struct {
	DataTesting string
	Expected    bool
}

func TestSolution(t *testing.T) {
	dataSingleWord := []*data{
		{
			DataTesting: "SAIPPUAKIVIKAUPPIAS",
			Expected:    true,
		},
		{
			DataTesting: "Aibohphobia",
			Expected:    true,
		},
		{
			DataTesting: "Anna",
			Expected:    true,
		},
		{
			DataTesting: "Civic",
			Expected:    true,
		},
	}

	for _, test := range dataSingleWord {
		result := Solution(test.DataTesting)
		assert.Equal(t, test.Expected, result)
	}

	dataMultipleWord := []*data{
		{
			DataTesting: "My gym",
			Expected:    true,
		},
		{
			DataTesting: "No lemon, no melon",
			Expected:    true,
		},
	}

	for _, test := range dataMultipleWord {
		result := Solution(test.DataTesting)
		assert.Equal(t, test.Expected, result)
	}

}
