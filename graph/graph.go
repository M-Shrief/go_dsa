package graph

import (
	"fmt"

	"github.com/M-Shrief/go-dsa-practice/queue"
)

type Graph[T comparable] struct {
	adjList    map[T][]T
	isDirected bool
}

func NewGraph[T comparable](isDirected bool) *Graph[T] {
	mp := map[T][]T{}
	return &Graph[T]{adjList: mp, isDirected: isDirected}
}

func (g *Graph[T]) HasVertex(n T) bool {
	_, exist := g.adjList[n]
	return exist
}

func (g *Graph[T]) AddVertex(n T) error {
	exist := g.HasVertex(n)
	if exist {
		return fmt.Errorf("already exists")
	}
	g.adjList[n] = []T{}
	return nil
}

func (g *Graph[T]) AddEdge(n1, n2 T) {
	if !g.HasVertex(n1) {
		g.AddVertex(n1)
	}
	if !g.HasVertex(n2) {
		g.AddVertex(n2)
	}

	g.adjList[n1] = append(g.adjList[n1], n2)
	if g.isDirected {
		g.adjList[n2] = append(g.adjList[n2], n1)
	}
}

func (g *Graph[T]) BFT(n T) ([]T, error) {
	if !g.HasVertex(n) {
		return nil, fmt.Errorf("element doesn't exists")
	}

	list := []T{}
	visited := map[T]bool{}
	queue := queue.NewQueue[T]()

	queue.Enqueue(n)
	visited[n] = true

	for queue.GetSize() != 0 {
		current, _ := queue.Dequeue()
		list = append(list, current)

		for _, key := range g.adjList[current] {
			isVisited := visited[key]
			if !isVisited {
				visited[key] = true
				queue.Enqueue(key)
			}
		}
	}

	return list, nil
}

func (g *Graph[T]) DFT(n T) ([]T, error) {
	if !g.HasVertex(n) {
		return nil, fmt.Errorf("element doesn't exists")
	}

	list := []T{}
	visited := map[T]bool{}

	res, err := g.dft(n, &list, visited)
	return *res, err
}

func (g *Graph[T]) dft(n T, list *[]T, visited map[T]bool) (*[]T, error) {

	visited[n] = true
	*list = append(*list, n)

	for _, val := range g.adjList[n] {
		if !visited[val] {
			g.dft(val, list, visited)
		}
	}

	return list, nil
}
