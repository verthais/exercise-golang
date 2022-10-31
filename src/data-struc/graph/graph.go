package graph

type Vertex struct {
	key      int
	vertices []*Vertex
}

func (v *Vertex) adjacence(key int) bool {
	for _, o := range v.vertices {
		if o.key == key {
			return true
		}
	}

	return false
}

type Graph struct {
	vertices []*Vertex
}

func (g *Graph) Insert(k int) {
	if !g.contains(k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

func (g *Graph) Edge(from, to int) {
	if !g.contains(from) {
		g.Insert(from)
	}

	if !g.contains(to) {
		g.Insert(to)
	}

	v_from, v_to := g.getVertex(from), g.getVertex(to)

	if !v_from.adjacence(to) {
		v_from.vertices = append(v_from.vertices, v_to)
	}
}

func (g *Graph) getVertex(key int) *Vertex {
	for _, v := range g.vertices {
		if v.key == key {
			return v
		}
	}

	return nil
}

func (g *Graph) contains(key int) bool {
	for _, v := range g.vertices {
		if v.key == key {
			return true
		}
	}

	return false
}

func NewGraph() Graph {
	return Graph{}
}
