package main

import "fmt"

const (
	// Vert - кол-во вершин
	Vert = 7
)

var used = make(map[int]bool)
var graph = [Vert][]int{
	{1, 2},
	{0, 2},
	{0, 1, 3, 4, 5},
	{2},
	{2, 6},
	{2, 6},
	{4, 5},
}

func main() {
	fmt.Println(dfs(3, 6))
	fmt.Println(bfs(5))
}

func dfs(start int, needle int) int {
	used[start] = true
	for _, n := range graph {
		for _, v := range n {
			if _, ok := used[v]; !ok {
				if v == needle {
					return v
				}
				dfs(v, needle)
			}
		}
	}
	return -1
}

func bfs(needle int) int {
	graph := [Vert][]int{
		{1, 2},
		{0, 2},
		{0, 1, 3, 4, 5},
		{2},
		{2, 6},
		{2, 6},
		{4, 5},
	}
	start := 2
	q := make([][]int, len(graph))[len(graph):]
	q = append(q, graph[start])

	used := make(map[int]bool, len(graph))
	used = map[int]bool{start: true}

	for len(q) > 0 {
		for _, key := range q[len(q)-1] {
			if _, ok := used[key]; !ok {
				q = append(q, graph[key])
			}
			if key == needle {
				return key
			}
			used[key] = true
		}
		q = q[:len(q)-1]
	}
	return -1
}
