package main

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	fps = 6
	// TODO: currently row must eq columns, but potentially worth changing in future
	// TODO: row/col to take user inputs
	rows    = 21
	columns = 21
)

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	grid := makeGrid()
	rdfs := makerdfs(grid)

	for !window.ShouldClose() {
		t := time.Now()

		draw(window, program, rdfs.grid)

		rdfs.step()

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

func draw(window *glfw.Window, program uint32, grid [][]*point) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for x := range grid {
		for _, p := range grid[x] {
			p.draw()
		}
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
