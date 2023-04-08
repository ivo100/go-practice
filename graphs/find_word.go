package main

/**

https://leetcode.com/problems/word-search/

Given an m x n grid of characters board and a string word,
return true if word exists in the grid.
The word can be constructed from letters of sequentially adjacent cells,
where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once.
*/
import "fmt"

type node struct {
	Row     int
	Col     int
	Visited bool
	Val     string
	Adj     []node
}

func exist(board [][]byte, word string) bool {
	w := len(board[0])
	h := len(board)
	g := make([]node, 0)

	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			// fmt.Printf("\nr %d c %d -> %q\n", r, c, board[r][c])
			n := node{
				Row:     r,
				Col:     c,
				Visited: false,
				Val:     string(board[r][c]),
				// N, W, S, E
				Adj: nil,
			}
			// N
			adj(board, &n, r-1, c, w, h)
			// W
			adj(board, &n, r, c+1, w, h)
			// S
			adj(board, &n, r+1, c, w, h)
			adj(board, &n, r, c-1, w, h)
			// fmt.Printf("%+v\n", n)
			g = append(g, n)
		}
	}
	return find(g, word)
}

func find(g []node, word string) bool {
	L := 0
	LW := len(word)
	wi := 0
	for wi < LW-1 {
		fmt.Printf("crnt %v\n", word[wi:wi+1])
		for _, n := range g {
			if n.Visited {
				continue
			}
			if n.Val == word[wi:wi+1] {
				fmt.Printf("match %v\n", n.Val)
				n.Visited = true
				L++
				wi++
				break
			}
			// no match, check adj
			for _, adj := range n.Adj {
				if adj.Visited {
					continue
				}
				if adj.Val == word[wi:wi+1] {
					L++
					wi++
					adj.Visited = true
					fmt.Printf("match adj %v\n", adj.Val)
					break
				}
			}
		}
	}
	return L == len(word)
}

func adj(board [][]byte, n *node, r int, c int, w int, h int) {
	if r < 0 || c < 0 || r >= h || c >= w {
		return
	}
	if n.Adj == nil {
		n.Adj = make([]node, 0)
	}
	nn := node{
		Row:     r,
		Col:     c,
		Visited: false,
		Val:     string(board[r][c]),
	}
	n.Adj = append(n.Adj, nn)
}

func main() {
	line1 := []byte{'E', 'H', 'E'}
	line2 := []byte{'P', 'P', 'L'}
	board := [][]byte{line1, line2}

	// fmt.Printf("%q\n", line1)
	//fmt.Printf("%q\n", board)
	fmt.Printf("%v\n", exist(board, "HELP"))
}
