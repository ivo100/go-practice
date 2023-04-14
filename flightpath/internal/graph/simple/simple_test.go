package simple

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_AddLeg(t *testing.T) {
	r := NewDAG()
	good := [][]string{
		{"SAN", "LAS"},
		{"LAS", "SFO"},
	}
	for _, b := range good {
		assert.NoError(t, r.AddLeg(b[0], b[1]))
	}
}

func Test_AddLegNegative(t *testing.T) {
	r := NewDAG()
	// detect errors
	badBoys := [][]string{
		{"SAN", "SAN"},
		{"SFO", ""},
		{"x", "bar"},
	}
	for _, b := range badBoys {
		assert.Error(t, r.AddLeg(b[0], b[1]))
	}
}

func Test_MainCase(t *testing.T) {
	r := NewDAG()
	// The main case to test [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
	good := [][]string{
		{"IND", "EWR"},
		{"SFO", "ATL"},
		{"GSO", "IND"},
		{"ATL", "GSO"},
	}
	for _, b := range good {
		assert.NoError(t, r.AddLeg(b[0], b[1]))
	}
	p, err := r.Sort()
	assert.NoError(t, err)
	require.Equal(t, 5, len(p))
	require.Equal(t, "SFO", p[0])
	require.Equal(t, "EWR", p[4])
}

func Test_Cycle(t *testing.T) {
	r := NewDAG()
	assert.NoError(t, r.AddLeg("LAS", "SFO"))
	assert.NoError(t, r.AddLeg("SFO", "JFK"))
	assert.NoError(t, r.AddLeg("JFK", "LAS"))
	_, err := r.Sort()
	assert.Error(t, err)
}

func Test_Partition(t *testing.T) {
	r := NewDAG()
	assert.NoError(t, r.AddLeg("LAS", "SFO"))
	assert.NoError(t, r.AddLeg("SFO", "JFK"))
	assert.NoError(t, r.AddLeg("SAN", "ATL"))
	assert.NoError(t, r.AddLeg("ATL", "IND"))
	_, err := r.Sort()
	assert.Error(t, err)
}
