package main

import (
	"fmt"
	"log"
	"strings"
)

type Stack struct {
	a []interface{}
}

func NewStack() *Stack {
	s := &Stack{}
	s.a = make([]interface{}, 0)
	return s
}

func (s *Stack) Push(item interface{}) {
	s.a = append(s.a, item)
}

func (s *Stack) Top() (item interface{}) {
	if s.IsEmpty() {
		panic("empty stack")
	}
	item = s.a[len(s.a)-1]
	return item
}

func (s *Stack) Pop() (item interface{}) {
	item = s.Top()
	s.a = s.a[:len(s.a)-1]
	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.a) == 0
}

type Queue struct {
	a []interface{}
}

func NewQueue() *Queue {
	s := &Queue{}
	s.a = make([]interface{}, 0)
	return s
}

func (s *Queue) Enqueue(item interface{}) {
	s.a = append(s.a, item)
}

func (s *Queue) First() (item interface{}) {
	if s.IsEmpty() {
		panic("empty Queue")
	}
	item = s.a[0]
	return item
}

func (s *Queue) Dequeue() (item interface{}) {
	item = s.First()
	s.a = s.a[1:]
	return item
}

func (s *Queue) IsEmpty() bool {
	return len(s.a) == 0
}

type Node struct {
	V       int
	visited bool
}

type AdjList struct {
	Node     Node
	Adjacent []int
}

type Graph struct {
	Lists []AdjList
}

func NewGraph(n int) *Graph {
	g := &Graph{}
	g.Lists = make([]AdjList, n)
	for i := 0; i < n; i++ {
		g.Lists[i] = AdjList{Node: Node{V: i}, Adjacent: make([]int, 0)}
	}
	return g
}

func (g *Graph) AddEdge(i, j int) {
	g.Lists[i].Adjacent = append(g.Lists[i].Adjacent, j)
}

func (g *Graph) String() string {
	var sb strings.Builder
	for i := 0; i < len(g.Lists); i++ {
		sb.WriteString(fmt.Sprintf("\nNode %v, Edges: ", g.Lists[i].Node))
		for _, v := range g.Lists[i].Adjacent {
			sb.WriteString(fmt.Sprintf("%v ", v))
		}
	}
	return sb.String()
}

func testStack() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	a := s.Pop()
	if a.(int) != 2 {
		panic("stack pop 2")
	}
	a = s.Top()
	if a.(int) != 1 {
		panic("stack top 1")
	}
	if s.IsEmpty() {
		panic("stack empty 1")
	}
	a = s.Pop()
	if a.(int) != 1 {
		panic("stack top 1")
	}
	if !s.IsEmpty() {
		panic("stack not empty 2")
	}
}

func testQueue() {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	a := q.First()
	if a.(int) != 1 {
		panic("queue first 1")
	}
	a = q.Dequeue()
	if a.(int) != 1 {
		panic("q deq 1")
	}
	if q.IsEmpty() {
		panic("q empty 1")
	}
	a = q.Dequeue()
	if a.(int) != 2 {
		panic("q deq 2")
	}
	if !q.IsEmpty() {
		panic("q not empty 2")
	}
}

/*
Graph g(5); // Total 5 vertices in graph

	g.addEdge(1, 0);
	g.addEdge(0, 2);
	g.addEdge(2, 1);
	g.addEdge(0, 3);
	g.addEdge(1, 4);
*/
func testDFS() {
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)
	log.Printf("Graph %v", g.String())
	g.RecursiveDFS(0)
}

func testBFS() {
	g := NewGraph(5)
	g.AddEdge(1, 0)
	g.AddEdge(0, 3)
	g.AddEdge(0, 2)
	g.AddEdge(2, 1)
	g.AddEdge(1, 4)

	log.Printf("Graph %v", g.String())
	g.BFS(0)
	// expected: 0 3 2 1 4

}

// RecursiveDFS /*
func (g *Graph) RecursiveDFS(i int) {
	log.Printf("RecursiveDFS - visit %v", g.Lists[i].Node.V)
	g.Lists[i].Node.visited = true
	for _, adj := range g.Lists[i].Adjacent {
		if !g.Lists[adj].Node.visited {
			g.RecursiveDFS(adj)
		}
	}
}

/*
Breadth First Search (BFS)

There are many ways to traverse graphs. BFS is the most commonly used approach.

BFS is a traversing algorithm where you should start
traversing from a selected node (source or starting node) and traverse the graph
layerwise thus exploring the neighbour nodes (nodes which are directly connected
to source node). You must then move towards the next-level neighbour nodes.

As the name BFS suggests, you are required to traverse the graph breadthwise as follows:

    1. First move horizontally and visit all the nodes of the current layer
    2. Move to the next layer


FS (G, s)                   //Where G is the graph and s is the source node
      let Q be queue.
      Q.enqueue( s ) //Inserting s in queue until all its neighbour vertices are marked.

      mark s as visited.
      while ( Q is not empty)
           //Removing that vertex from queue,whose neighbour will be visited now
           v  =  Q.dequeue( )

          //processing all the neighbours of v
          for all neighbours w of v in Graph G
               if w is not visited
                        Q.enqueue( w )             //Stores w in Q to further visit its neighbour
                        mark w as visited.

*/

func (g *Graph) BFS(start int) {
	Q := NewQueue()
	Q.Enqueue(start)
	g.Lists[start].Node.visited = true

	for !Q.IsEmpty() {
		v := Q.Dequeue().(int)
		log.Printf("DFS - visit %v", g.Lists[v].Node.V)
		//processing all the neighbours of v
		for _, adj := range g.Lists[v].Adjacent {
			if !g.Lists[adj].Node.visited {
				Q.Enqueue(adj)
				g.Lists[adj].Node.visited = true
			}
		}
	}

}

func main() {
	//testStack()
	//testDFS()
	//testQueue()
	testBFS()
}
