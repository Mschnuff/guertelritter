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
type movementSettings struct {
	speed int
}

var mS movementSettings
var mouse mouseCursor
var pD previousDirections

func SetMovementSpeed(speed int) {
	mS.speed = speed
}

func GetCursorToPlayerAngle(screenMiddleX int, screenMiddleY int) float64 {

	mouse.xPos, mouse.yPos = ebiten.CursorPosition()
	mouse.angle = math.Atan2(float64(screenMiddleX)-float64(mouse.xPos), float64(screenMiddleY)-float64(mouse.yPos))

	return mouse.angle
}
func calculatePlayerX(cl bool, cr bool) int {
	pX := 0
	if cl && !cr {
		pX = -mS.speed
		pD.left = true
		pD.right = false
	} else if cr && !cl {
		pX = mS.speed
		pD.left = false
		pD.right = true
	} else if cr && cl {
		//rand2 := rand.Intn(2)*2 - 2
		//playerX += rand2
		if pD.left && !pD.right {
			pX = mS.speed
			pD.left = true
			pD.right = false

		} else if !pD.left && pD.right {
			pX = -mS.speed
			pD.left = false
			pD.right = true

		}
	} // !cl && !cr faellt erstmal weg
	return pX
}
func calculatePlayerY(cu bool, cd bool) int {
	pY := 0
	if cu && !cd {
		pY = -mS.speed
		pD.up = true
		pD.down = false
	} else if cd && !cu {
		pY = mS.speed
		pD.up = false
		pD.down = true
	} else if cd && cu {
		//rand2 := rand.Intn(2)*2 - 2
		//playerX += rand2
		if pD.up && !pD.down {
			pY = mS.speed
			pD.up = true
			pD.down = false

		} else if !pD.up && pD.down {
			pY = -mS.speed
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
