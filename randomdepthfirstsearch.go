package main

import "math/rand"

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
			p.isWall = true
		}
	}

	cells := getcells(grid)

	// unset all cells to begin algorithm
	for _, c := range cells {
		c.point.isWall = false
	}

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

func (maze *rdfs) step() {
	//* 1. pop a cell off stack
	current_cell, err := maze.stack.Pop()
	if err != nil {
		println("RDFS Complete")
		return
	}
	unvisNeighbours := maze.getUnvisitedNeighbours(current_cell.index)

	//* 2. if the current cell has any unvisted neighbours push onto stack
	if len(unvisNeighbours) > 0 {
		maze.stack.Push(current_cell)

		//* 3. choose one of unvisited neighbours (random)
		rand_idx := rand.Intn(len(unvisNeighbours))

		next_cell := unvisNeighbours[rand_idx]

		//* 4. remove wall between current and chosen
		maze.removeWallBetween(current_cell.index, next_cell.index)

		//* 5. mark chosen cell as visited and push to stack
		next_cell.isVisited = true
		maze.stack.Push(next_cell)
	}
}

func (maze *rdfs) getUnvisitedNeighbours(index int) []*cell {
	unvisitedNeighbours := make([]*cell, 0)

	cellRowCount := (rows - 1) / 2

	// not left
	if index%cellRowCount != 0 {
		if !maze.cells[index-1].isVisited {
			unvisitedNeighbours = append(unvisitedNeighbours, maze.cells[index-1])
		}
	}

	// not right
	if index%cellRowCount != cellRowCount-1 {
		if !maze.cells[index+1].isVisited {
			unvisitedNeighbours = append(unvisitedNeighbours, maze.cells[index+1])
		}
	}

	// not top
	if index+cellRowCount < len(maze.cells) {
		if !maze.cells[index+cellRowCount].isVisited {
			unvisitedNeighbours = append(unvisitedNeighbours, maze.cells[index+cellRowCount])
		}
	}

	// not bottom
	if index-cellRowCount >= 0 {
		if !maze.cells[index-cellRowCount].isVisited {
			unvisitedNeighbours = append(unvisitedNeighbours, maze.cells[index-cellRowCount])
		}
	}
	return unvisitedNeighbours
}
func (maze *rdfs) removeWallBetween(current int, next int) {
	l := current
	r := next

	if next < current {
		l = next
		r = current
	}

	dt := r - l

	x := maze.cells[l].point.x
	y := maze.cells[l].point.y

	if dt == 1 { // on horizontal row
		maze.grid[y+1][x].isWall = false
	} else {
		maze.grid[y][x+1].isWall = false
	}
}
