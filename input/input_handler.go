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
