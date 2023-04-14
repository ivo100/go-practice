package graph

import (
	"fmt"
	"strings"
)

// Leg is one segment of the Flights path
type Leg struct {
	From string
	To   string
}

func (l Leg) String() string {
	return fmt.Sprintf(`[%q,%q]`, l.From, l.To)
}

// Flights represents our flights graph - nodes are airports identified by their 3-letter IATA code
// edges are origin/destination pairs
// flights graph is ephemeral - it contains information only about single user itinerary in one direction
// return flights are EXCLUDED
type Flights interface {
	// AddLeg is used to build flights graph by adding a flight segment (aka leg)
	// it may return ErrInvalidArg on invalid argument
	// we assume DAG - no cycles (return flights), partitions
	AddLeg(from, to string) error

	// Sort performs topological sorting of the graph and returns airport codes of the legs
	// as a linear sequence, e.g. SAN, SFO, LAS
	// it is possible to have more than one possible sort order of a DAG
	// in case of errors (cycles etc) - result will be nil slice and error will be ErrInvalidArg
	Sort() ([]string, error)
}

// NodesFromEdges is small utility function that validates list of flight segments for correctness
// it also calculates the vertices from edges to match the input given to our service
// nodes are identified by 3-letter IANA codes - any violations are considered errors
func NodesFromEdges(p [][]string) (nodes []string, err error) {
	err = ErrInvalidArg
	if p == nil || len(p) == 0 {
		return
	}
	deg := make(map[string]int)
	for _, v := range p {
		ok, src, dst := valid(v)
		if !ok {
			return
		}
		deg[src]++
		deg[dst]--
	}
	nodes = make([]string, 0, len(deg))
	for k := range deg {
		nodes = append(nodes, k)
	}
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			p[i][j] = scrub(p[i][j])
		}
	}
	return nodes, nil
}

func scrub(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	return s
}

func valid(q []string) (ok bool, src string, dst string) {
	ok = false
	if len(q) != 2 {
		return
	}
	src = scrub(q[0])
	dst = scrub(q[1])
	if len(src) != AirportCodeLength {
		return
	}
	if len(dst) != AirportCodeLength {
		return
	}
	ok = src != dst
	return
}
