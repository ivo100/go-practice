package sherlock

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSherlockAndAnagrams(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int32
	}{
		{"mom", "mom", 2},
		{"abba", "abba", 4},
		{"abcd", "abcd", 0},
		{"kkkk", "kkkk", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, SherlockAndAnagrams(tt.arg))
		})
	}
}

func Test_findAnagrams(t *testing.T) {
	//got := findAnagrams(tt.args.list)
}

func Test_substrWithLen(t *testing.T) {
	list := substrWithLen("mom", 2)
	require.Equal(t, 2, len(list))
}
