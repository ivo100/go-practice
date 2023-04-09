package main

import (
	"github.com/dominikbraun/graph"
	"log"
)

func main() {
	g := graph.New(graph.StringHash, graph.Directed())

	_ = g.AddVertex("pants")
	_ = g.AddVertex("socks")
	_ = g.AddVertex("shoes")

	_ = g.AddEdge("pants", "shoes")
	_ = g.AddEdge("socks", "shoes")
	_ = g.AddEdge("socks", "pants")

	//_ = graph.BFS(g, "a", func(value string) bool {
	//	log.Print(value)
	//	return false
	//})

	t, err := graph.TopologicalSort(g)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("%v", t)
}
