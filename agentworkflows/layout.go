package agentworkflows

import (
	"encoding/json"
	"strconv"

	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

const (
	columnWidth = 300
	rowHeight   = 200
)

// layout gives every node its editor position. Nodes that already have one keep it. The rest go
// in column = breadth-first depth from the start node, row = order of discovery within that column,
// skipping grid cells a positioned node already occupies. Unreachable nodes go in the column after
// the deepest one, in insertion order.
func layout(g *graph) map[string]any {
	depths, visited := depthsFrom(g)
	lastColumn := 0
	for _, depth := range depths {
		if depth > lastColumn {
			lastColumn = depth
		}
	}
	lastColumn++

	order := append([]string{}, visited...)
	for _, id := range g.nodeOrder {
		if _, ok := depths[id]; !ok {
			order = append(order, id)
		}
	}

	occupied := map[string]bool{}
	for _, id := range g.nodeOrder {
		if position, ok := existingPosition(g, id); ok {
			occupied[cellKey(position)] = true
		}
	}

	rows := map[int]int{}
	positions := map[string]any{}
	for _, id := range order {
		if position, ok := existingPosition(g, id); ok {
			positions[id] = position
			continue
		}
		column, ok := depths[id]
		if !ok {
			column = lastColumn
		}
		row := rows[column]
		for occupied[gridKey(column, row)] {
			row++
		}
		rows[column] = row + 1
		occupied[gridKey(column, row)] = true
		positions[id] = map[string]any{
			"x": json.Number(strconv.Itoa(column * columnWidth)),
			"y": json.Number(strconv.Itoa(row * rowHeight)),
		}
	}
	return positions
}

// existingPosition is a node's own position. A JSON null counts as no position.
func existingPosition(g *graph, id string) (any, bool) {
	position, ok := g.nodes[id]["position"]
	if !ok || position == nil {
		return nil, false
	}
	return position, true
}

func gridKey(column, row int) string {
	return strconv.Itoa(column*columnWidth) + "," + strconv.Itoa(row*rowHeight)
}

// cellKey is the grid cell a loaded position occupies, written as the JavaScript reference keys it
// (`${x},${y}`), so a position given as 300.0 occupies the same cell as an automatic 300.
func cellKey(position any) string {
	m, _ := position.(map[string]any)
	return coordText(m, "x") + "," + coordText(m, "y")
}

func coordText(m map[string]any, key string) string {
	value, ok := m[key]
	if !ok {
		return "undefined"
	}
	switch v := value.(type) {
	case nil:
		return "null"
	case json.Number:
		return wirejson.FormatNumber(string(v))
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	}
	return "[object]"
}
