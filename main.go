package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type projectile struct {
	xVel   float32
	yVel   float32
	xAccel float32
	yAccel float32
	body   rl.Rectangle
}

var started bool = false
var object projectile = projectile{10, -10, 0, 0.3, rl.Rectangle{10, 399, 10, 10}}

func update() {
	if started == true && object.body.Y < 400 {
		object.body.X += object.xVel
		object.xVel += object.xAccel
		object.body.Y += object.yVel
		object.yVel += object.yAccel
	}
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawLine(0, 410, 800, 410, rl.Black)
	rl.DrawText("Press any key to start", 5, 5, 15, rl.Gray)
	rl.DrawText("R - restart", 5, 26, 15, rl.Gray)

	rl.DrawRectangleRec(object.body, rl.Red)

	rl.EndDrawing()
}

func restart() {
	started = false
	object = projectile{10, -10, 0, 0.3, rl.Rectangle{10, 399, 10, 10}}
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		getInput := rl.GetKeyPressed()
		if getInput != 0 {
			started = true
		}
		if getInput == 82 {
			restart()
		}
		update()
		render()
	}
}
