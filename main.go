package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const screenWidth int32 = 800
	const screenHeight int32 = 450

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 3d camera free")
	defer rl.CloseWindow()

	// Define the camera to look into our 3d world
	var camera rl.Camera3D
	camera.Position = rl.Vector3{10.0, 10.0, 10.0} // Camera position
	camera.Target = rl.Vector3{0.0, 0.0, 0.0}      // Camera looking at point
	camera.Up = rl.Vector3{0.0, 1.0, 0.0}          // Camera up vector (rotation towards target)
	camera.Fovy = 45.0                             // Camera field-of-view Y
	camera.Projection = rl.CameraPerspective       // Camera projection type

	cubePosition := rl.Vector3{0.0, 0.0, 0.0}

	rl.DisableCursor() // Limit cursor to relative movement inside the window

	rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !rl.WindowShouldClose() {
		// Update
		//----------------------------------------------------------------------------------
		rl.UpdateCamera(&camera, rl.CameraFree)

		if rl.IsKeyPressed('Z') {
			camera.Target = rl.Vector3{0.0, 0.0, 0.0}
		}
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
		rl.DrawCubeWires(cubePosition, 2.0, 2.0, 2.0, rl.Maroon)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 320, 93, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 320, 93, rl.Blue)

		rl.DrawText("Free camera default controls:", 20, 20, 10, rl.Black)
		rl.DrawText("- Mouse Wheel to Zoom in-out", 40, 40, 10, rl.DarkGray)
		rl.DrawText("- Mouse Wheel Pressed to Pan", 40, 60, 10, rl.DarkGray)
		rl.DrawText("- Z to zoom to (0, 0, 0)", 40, 80, 10, rl.DarkGray)

		rl.EndDrawing()
		//----------------------------------------------------------------------------------
	}
}
