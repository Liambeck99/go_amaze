package main

import "math/rand"

// ? Could maybe allow colour schemes in the future i.e. catpuccin ??
var (
	colourInStack = [3]float32{0.98, 0.70, 0.52}
	colourInMaze  = [3]float32{0.65, 0.89, 0.63}
	colourHead    = [3]float32{0.95, 0.54, 0.66}
)

type cell struct {
	*point
	index     int
	isVisited bool
}

type rdfs struct {
	grid  [][]*point
	cells []*cell
	stack *stack
}

func makerdfs(grid [][]*point) *rdfs {
	// display all walls to begin with will delete them as we progress
	for x := range grid {
		for _, p := range grid[x] {
			if p.isEdge {
				continue
			}
			p.setWall()
		}
	}

	cells := getcells(grid)

	// create stack and assign a cell to the stack to start algorithm
	stack := NewStack()
	rand_idx := rand.Intn(len(cells))
	starting_cell := cells[rand_idx]
	starting_cell.isVisited = true
	stack.Push(starting_cell)

	return &rdfs{
		grid:  grid,
		cells: cells,
		stack: stack,
	}
}

func getcells(grid [][]*point) []*cell {
	cells := make([]*cell, 0)

	for x := range grid {
		for y, p := range grid[x] {
			if p.isEdge {
				continue
			}

			if x%2 == 1 && y%2 == 1 {
				cell := &cell{p, len(cells), false}

				cells = append(cells, cell)
			}
		}
	}

	return cells
}

func (r *rdfs) step() {
	//* 1. pop a cell off stack
	current_cell, err := r.stack.Pop()

	if err != nil {
		println("RDFS Complete")
		return
	}
	unvisNeighbours := r.getUnvisitedNeighbours(current_cell.index)

	//* 2. if the current cell has any unvisted neighbours push onto stack
	if len(unvisNeighbours) > 0 {
		r.stack.Push(current_cell)

		//* 3. choose one of unvisited neighbours (random)
		rand_idx := rand.Intn(len(unvisNeighbours))

		next_cell := unvisNeighbours[rand_idx]

		//* 4. remove wall between current and chosen
		r.getPointBetweenCells(current_cell.index, next_cell.index).isWall = false

		//* 5. mark chosen cell as visited and push to stack
		next_cell.isVisited = true
		next_cell.isWall = false
		r.stack.Push(next_cell)
	}

	r.updateColours()
}
func (r *rdfs) getNeighbours(index int) []*cell {
	n := make([]*cell, 0)

	cellRowCount := (rows - 1) / 2

	// not left
	if index%cellRowCount != 0 {
		n = append(n, r.cells[index-1])
	}

	// not right
	if index%cellRowCount != cellRowCount-1 {
		n = append(n, r.cells[index+1])
	}

	// not top
	if index+cellRowCount < len(r.cells) {
		n = append(n, r.cells[index+cellRowCount])
	}

	// not bottom
	if index-cellRowCount >= 0 {
		n = append(n, r.cells[index-cellRowCount])
	}

	return n
}
func (r *rdfs) getUnvisitedNeighbours(i int) []*cell {
	u := make([]*cell, 0)

	for _, c := range r.getNeighbours(i) {
		if !c.isVisited {
			u = append(u, c)
		}
	}

	return u
}

func (r *rdfs) getPointBetweenCells(a, b int) *point {
	if b < a {
		tmp := b
		b = a
		a = tmp
	}

	dt := b - a

	x := r.cells[a].point.x
	y := r.cells[a].point.y

	if dt == 1 { // on horizontal row
		return r.grid[x][y+1]
	} else {
		return r.grid[x+1][y]
	}
}

func (r *rdfs) updateColours() {
	// set all non walls to included in grid
	for x := range r.grid {
		for _, p := range r.grid[x] {
			if !p.isWall {
				p.colour = colourInMaze
			}
		}
	}

	s := *r.stack

	for i := range s {
		s[i].colour = colourInStack
		if i == len(s)-1 {
			continue
		}
		r.getPointBetweenCells(s[i].index, s[i+1].index).colour = colourInStack
	}

	head, err := r.stack.Last()
	if err == nil {
		head.colour = colourHead
	}

}
