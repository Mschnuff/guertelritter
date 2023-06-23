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
type previousDirections struct {
	right bool
	down  bool
	left  bool
	up    bool
}

var mouse mouseCursor
var pD previousDirections

func GetCursorToPlayerAngle(screenMiddleX int, screenMiddleY int) float64 {

	mouse.xPos, mouse.yPos = ebiten.CursorPosition()
	mouse.angle = math.Atan2(float64(screenMiddleX)-float64(mouse.xPos), float64(screenMiddleY)-float64(mouse.yPos))

	return mouse.angle
}
func calculatePlayerX(cl bool, cr bool) int {
	pX := 0
	if cl && !cr {
		pX = -1
		pD.left = true
		pD.right = false
	} else if cr && !cl {
		pX = 1
		pD.left = false
		pD.right = true
	} else if cr && cl {
		//rand2 := rand.Intn(2)*2 - 2
		//playerX += rand2
		if pD.left && !pD.right {
			pX = 1
			pD.left = true
			pD.right = false

		} else if !pD.left && pD.right {
			pX = -1
			pD.left = false
			pD.right = true

		}
	} // !cl && !cr faellt erstmal weg
	return pX
}
func calculatePlayerY(cu bool, cd bool) int {
	pY := 0
	if cu && !cd {
		pY = -1
		pD.up = true
		pD.down = false
	} else if cd && !cu {
		pY = 1
		pD.up = false
		pD.down = true
	} else if cd && cu {
		//rand2 := rand.Intn(2)*2 - 2
		//playerX += rand2
		if pD.up && !pD.down {
			pY = 1
			pD.up = true
			pD.down = false

		} else if !pD.up && pD.down {
			pY = -1
			pD.up = false
			pD.down = true

		}
	} // !cl && !cr faellt erstmal weg
	return pY
}

func HandlePlayerMovement(playerX int, playerY int) (int, int) {
	cl := false
	cr := false
	cu := false
	cd := false
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		cr = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		cd = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		cl = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		cu = true
	}
	playerX += calculatePlayerX(cl, cr)
	playerY += calculatePlayerY(cu, cd)
	return playerX, playerY
}
