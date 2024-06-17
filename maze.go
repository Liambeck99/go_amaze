package main

import (
	"math/rand"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var (
	square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
)

type cell struct {
	drawable uint32

	display bool

	x int
	y int
}

func (c *cell) draw() {
	if !c.display {
		return
	}

	gl.BindVertexArray(c.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
}

func makeGrid() [][]*cell {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	cells := make([][]*cell, rows)
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			c := newCell(x, y)

			c.display = rand.Float64() < 0.1

			cells[x] = append(cells[x], c)
		}
	}

	return cells
}

func newCell(x, y int) *cell {
	// create a copy of the square definition which means we can modify the coords and not affect others
	points := make([]float32, len(square))
	copy(points, square)

	// iterate over points, use mod to only change x and y, ignore z
	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 3 {
		// calcs based on board dimensions
		case 0:
			size = 1.0 / float32(columns)
			position = float32(x) * size
		case 1:
			size = 1.0 / float32(rows)
			position = float32(y) * size
		default:
			continue

		}

		// update points accordingly
		if points[i] < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}

	return &cell{
		drawable: makeVao(points),

		x: x,
		y: y,
	}
}
