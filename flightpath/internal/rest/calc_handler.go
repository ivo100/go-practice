package rest

import (
	"flightpath/internal/graph"
	"flightpath/internal/graph/glib"
	"flightpath/internal/graph/kahn"
	"flightpath/internal/graph/simple"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"strings"
)

// FindRoute is the main handler method of the /calculdate service endpoint
// It binds request body which is array of [from, to] airport codes (echo will bind it for us) which are DAG edges
// Then it will do some validation and extract nodes from edges, check for trivial requests before deciding which graph implementation to use
// for topological sorting of the DAG
func FindRoute(c echo.Context) (err error) {
	p := new(Payload)
	if err = c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Logger().Debugf("FindRoute request %v", p)
	// validate and extract nodes
	nodes, err := graph.NodesFromEdges(*p)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// we can handle some simple requests like 1 and 2 legs (that may look like cycles) w/o going to sort)
	resp, err := handleTrivialRequest(p)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(resp) > 0 {
		return c.JSON(http.StatusOK, resp)
	}

	result, err := Sort(nodes, *p)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// we need only first and last from the whole path
	l := len(result)
	if l == 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}
	resp = append(resp, result[0])
	resp = append(resp, result[l-1])
	c.Logger().Debugf("FindRoute result %v", resp)
	return c.JSON(http.StatusOK, resp)
}

// Sort is the main "business" method of the service
func Sort(nodes []string, edges [][]string) ([]string, error) {
	// env. var. USE_GRAPH if set can select which implementation to use
	// if default if not set is to use graph library
	// one of: simple, graphlib, kahn
	impl := os.Getenv("USE_GRAPH")
	if impl == "simple" {
		log.Printf("Using simple algorithm")
		return doSimple(nodes, edges)
	}
	if impl == "kahn" {
		log.Printf("Using Kahn's algorithm")
		return doKahn(nodes, edges)
	}
	log.Printf("Using graph library")
	return glib.Sort(nodes, edges)
}

func handleTrivialRequest(p *Payload) ([]string, error) {
	if p == nil || len(*p) == 0 {
		return nil, graph.ErrInvalidArg
	}
	for _, v := range *p {
		if !valid(v) {
			return nil, graph.ErrInvalidArg
		}
	}
	l := len(*p)
	// single flight
	if l == 1 {
		q := (*p)[0]
		return []string{strings.ToUpper(q[0]), strings.ToUpper(q[1])}, nil
	}
	// 2 legs
	// check for roundtrip, exact match
	// same as single flight
	if l == 2 {
		q := (*p)[0]
		r := (*p)[1]
		if q[0] == r[1] && q[1] == r[0] {
			return []string{strings.ToUpper(q[0]), strings.ToUpper(q[1])}, nil
		}
	}
	// not trivial - do sorting
	return nil, nil
}

func doSimple(nodes []string, edges [][]string) ([]string, error) {
	_ = nodes
	r := simple.NewDAG()
	for i := 0; i < len(edges); i++ {
		if err := r.AddLeg(edges[i][0], edges[i][1]); err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}
	return r.Sort()
}

func doKahn(nodes []string, edges [][]string) ([]string, error) {
	return kahn.Sort(nodes, edges)
}

func scrub(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	return s
}

func valid(q []string) bool {
	if len(q) != 2 {
		return false
	}
	s := scrub(q[0])
	d := scrub(q[1])
	if len(s) != graph.AirportCodeLength {
		return false
	}
	if len(d) != graph.AirportCodeLength {
		return false
	}
	return s != d
}
