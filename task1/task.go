package task1

type GraphMatrix struct {
	costMatrix [][]int
}
type emptyVal struct{}

func (g *GraphMatrix) bfsCost(startVertex int) []int {
	path := make([]int, 0, len(g.costMatrix))
	visited := make(map[int]emptyVal)
	var queue []int

	visited[startVertex] = emptyVal{}
	queue = append(queue, startVertex)

	for len(queue) != 0 {
		currentVert := queue[0]
		path = append(path, currentVert)

		for adjVert, cost := range g.costMatrix[currentVert] {
			// Предполагается, что положительные значения обозначают наличие ребра
			if cost > 0 {
				if _, ok := visited[adjVert]; !ok {
					visited[adjVert] = emptyVal{}
					queue = append(queue, adjVert)
				}
			}
		}

		queue = queue[1:]
	}

	return path
}
