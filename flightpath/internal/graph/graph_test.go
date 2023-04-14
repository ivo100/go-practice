package graph

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNodesFromEdges(t *testing.T) {
	good := [][]string{
		{"IND", "EWR"},
		{"sfo", "ATL"},
		{"GSO", "IND"},
		{"ATL", "gso"},
	}
	p, err := NodesFromEdges(good)
	assert.NoError(t, err)
	require.Equal(t, 5, len(p))
	require.Contains(t, p, "SFO")
	require.Contains(t, p, "GSO")
	require.NotContains(t, p, "gso")
}
