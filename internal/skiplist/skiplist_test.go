package skiplist

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func intToAlpha(n int) string {
	// For uppercase letters (A-Z)
	return string(rune('A' + n - 1))
}

func TestSkiplist(t *testing.T) {

	sls := NewSkiplist()
	for i := 1; i <= 8; i++ {
		sls.Insert(intToAlpha(i*3), strconv.Itoa(i*3))
	}

	tests := []struct {
		name     string
		funcType string
		argKey   string
		expVal   string
		expErr   bool
	}{
		{
			name:     "Successfully find a key and its value",
			funcType: "FIND",
			argKey:   "F",
			expVal:   "6",
			expErr:   false,
		},
		{
			name:     "Key not found",
			funcType: "FIND",
			argKey:   "G",
			expVal:   "",
			expErr:   true,
		},
		{
			name:     "Successfully delete a key and its value",
			funcType: "DELETE",
			argKey:   "F",
			expVal:   "",
			expErr:   false,
		},
		{
			name:     "Key deleted, not found",
			funcType: "FIND",
			argKey:   "F",
			expVal:   "6",
			expErr:   true,
		},
		{
			name:     "Key not found to delete",
			funcType: "DELETE",
			argKey:   "F",
			expVal:   "6",
			expErr:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var result string
			var err error
			switch tc.funcType {
			case "FIND":
				result, err = sls.Find(tc.argKey)
			case "DELETE":
				err = sls.Delete(tc.argKey)
			default:
				t.Fatalf("Unknown method %s", tc.funcType)
			}
			if !tc.expErr {
				assert.Equal(t, tc.expVal, result)
			}
			if tc.expErr {
				assert.Error(t, err)
			}
		})
	}
}
