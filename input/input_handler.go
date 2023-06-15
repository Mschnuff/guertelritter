package input

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type mouseCursor struct {
	xPos  int
	yPos  int
	angle float64
}

var mouse mouseCursor

func GetCursorToPlayerAngle(playerX int, playerY int) float64 {

	mouse.xPos, mouse.yPos = ebiten.CursorPosition()
	mouse.angle = math.Atan2(float64(playerX)-float64(mouse.xPos), float64(playerY)-float64(mouse.yPos))

	return mouse.angle
}

func HandlePlayerMovement(playerX int, playerY int) (int, int) {
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		playerX += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		playerY += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		playerX -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		playerY -= 1
	}
	return playerX, playerY
}