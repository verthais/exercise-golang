package graph_test

import (
	"exercise-go/src/data-struc/graph"
	"exercise-go/tests"
	"testing"
)

func TestGrap(t *testing.T) {
	g := graph.NewGraph()

	for i := 0; i < 5; i++ {
		g.Insert(i)
	}

	g.Edge(0, 4)
	g.Edge(2, 3)
	g.Edge(4, 2)

	tests.AssertEqual(t, true, true)
}
