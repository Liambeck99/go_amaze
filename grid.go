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
	scale float32 = 1.00

	wallColour = [3]float32{0.27, 0.28, 0.35}
)

type point struct {
	drawable uint32 ``

	isEdge bool
	isWall bool

	colour [3]float32

	x int
	y int
}

func (p *point) draw(program uint32) {
	r := p.colour[0]
	g := p.colour[1]
	b := p.colour[2]

	gl.Uniform3f(gl.GetUniformLocation(program, gl.Str("uColor"+"\x00")), r, g, b)

	gl.BindVertexArray(p.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
}

func (p *point) setColour(r, g, b int) {
	p.colour = [3]float32{float32(r), float32(g), float32(b)}
}

func (p *point) setWall() {
	p.isWall = true
	p.colour = wallColour
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
				point.colour = wallColour
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

	colour := [3]float32{1.0, 1.0, 1.0} // set to white by default

	return &point{
		drawable: makeVao(points),

		colour: colour,

		x: x,
		y: y,
	}
}
