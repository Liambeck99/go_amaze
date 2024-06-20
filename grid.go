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
	scale float32 = 0.8
)

type point struct {
	drawable uint32 ``

	isEdge bool
	isWall bool

	x int
	y int
}

func (p *point) draw() {
	if !p.isEdge && !p.isWall {
		return
	}

	gl.BindVertexArray(p.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
}

func makeGrid() [][]*point {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	points := make([][]*point, rows)
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			point := newpoint(x, y)

			if x == 0 || x == rows-1 || y == 0 || y == columns-1 {
				point.isEdge = true
				point.isWall = true
			}

			points[x] = append(points[x], point)
		}
	}

	return points
}

func newpoint(x, y int) *point {
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
			size = scale / float32(columns)
			position = float32(x) * size
		case 1:
			size = scale / float32(rows)
			position = float32(y) * size
		default:
			continue

		}

		// update points accordingly
		if points[i] < 0 {
			points[i] = (position * 2) - scale
		} else {
			points[i] = ((position + size) * 2) - scale
		}
	}

	return &point{
		drawable: makeVao(points),

		x: x,
		y: y,
	}
}
