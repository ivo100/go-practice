package simple

import (
	"flightpath/internal/graph"
	"log"
	"strings"
)

// dag is a DAG structure representing the Flights graph - vertices are airports, edges are segments
// we assume  there are no cycles and return flights / partitions
type dag struct {
	// Start is the starting point of the DAG
	Start string
	// Flights is map with key = leg origin, value = Leg (could be simplified to destination only)
	// we don't have generic dag - more of "linear" list / degenerated tree A -> B -> C -> D
	Flights map[string]graph.Leg
	// map code -> degrees ++ in, -- out
	Degrees map[string]int
}

// NewDAG creates a new DAG - note that it does not use typical adj list - just hash map
// it doesn't have even nodes - they are part of the the map source -> Leg
func NewDAG() graph.Flights {
	return &dag{
		Flights: make(map[string]graph.Leg),
		Degrees: make(map[string]int),
	}
}

// Sort performs topological sort of the DAG
func (d *dag) Sort() ([]string, error) {
	if err := d.findStart(); err != nil {
		return nil, err
	}
	result := make([]string, 0)
	k := d.Start
	L := len(d.Flights)
	// cycles will cause infinite loop
	for i := 0; i <= L+2; i++ {
		leg, ok := d.Flights[k]
		if !ok {
			result = append(result, k)
			break // end
		}
		result = append(result, leg.From)
		k = leg.To
	}
	if len(result) != L+1 {
		// may happen on partitions or cycles
		return result, graph.ErrInvalidArg
	}
	return result, nil
}

// AddLeg adds edge to the DAG graph
func (d *dag) AddLeg(from, to string) error {
	if len(from) != graph.AirportCodeLength {
		return graph.ErrInvalidArg
	}
	if len(to) != graph.AirportCodeLength {
		return graph.ErrInvalidArg
	}
	if from == to {
		return graph.ErrInvalidArg
	}
	from = strings.ToUpper(from)
	to = strings.ToUpper(to)
	//log.Printf("from %s, to %s", from, to)
	exists := func(code string) bool { _, ok := d.Degrees[code]; return ok }
	// this will be OK in a general dag but we don't expect forks
	if _, ok := d.Flights[from]; ok {
		log.Printf("Fork %s", from)
		return graph.ErrInvalidArg
	}
	leg := graph.Leg{From: from, To: to}
	d.Flights[from] = leg
	// calc in/out degree - we will use this later to find Start node
	if !exists(from) {
		d.Degrees[from] = 1
	}
	if !exists(to) {
		d.Degrees[to] = 0
	} else {
		d.Degrees[to]--
	}
	return nil
}

func (d *dag) findStart() error {
	for k, v := range d.Degrees {
		if v <= 0 {
			continue
		}
		if d.Start != "" {
			return graph.ErrInvalidArg
		}
		d.Start = k
	}
	return nil
}
