package glib

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_MainCase(t *testing.T) {
	// The main case to test [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
	a := []string{"IND", "EWR", "SFO", "ATL", "GSO"}
	f := [][]string{
		{"IND", "EWR"},
		{"SFO", "ATL"},
		{"GSO", "IND"},
		{"ATL", "GSO"},
	}
	p, err := Sort(a, f)
	assert.NoError(t, err)
	require.Equal(t, 5, len(p))
	require.Equal(t, "SFO", p[0])
	require.Equal(t, "EWR", p[4])
}
