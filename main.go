package main

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	fps     = 10
	rows    = 100
	columns = 100
)

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	grid := makeGrid()

	for !window.ShouldClose() {
		t := time.Now()

		draw(window, program, grid)

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

func draw(window *glfw.Window, program uint32, grid [][]*cell) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for x := range grid {
		for _, c := range grid[x] {
			c.draw()
		}
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
