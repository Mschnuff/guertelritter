package graphics

import (
	"fmt"

	"github.com/mschnuff/guertelritter/text"
)

type calcScreen struct {
}

type ScreenHandler struct {
	width      int
	height     int
	xScreenMin int
	xScreenMax int
	yScreenMin int
	yScreenMax int
	MiddleX    float64
	MiddleY    float64
	Fuck       int
}

func InitScreenHandler(inwidth int, inheight int) ScreenHandler {

	sh := ScreenHandler{
		width:      inwidth,
		height:     inheight,
		xScreenMin: 0,
		xScreenMax: inwidth,
		yScreenMin: 0,
		yScreenMax: inheight,
		MiddleX:    float64(inwidth / 2),
		MiddleY:    float64(inheight / 2),
	}
	return sh
}
func isOnScreen(sh ScreenHandler, xmin, ymin, width, height int) bool {
	xmax := xmin + width
	ymax := ymin + height
	onScreen := false
	leftToScreen := xmax < sh.xScreenMin
	rightToScreen := xmin > sh.xScreenMax
	belowScreen := ymin > sh.yScreenMax
	aboveScreen := ymax < sh.yScreenMin
	onScreen = !aboveScreen && !leftToScreen && !rightToScreen && !belowScreen
	return onScreen
}

func CalcOffset(sh ScreenHandler, mouseX int, mouseY int) (float64, float64) {

	cameraOffsetX := (sh.MiddleX - float64(mouseX)) / 2
	cameraOffsetY := (sh.MiddleY - float64(mouseY)) / 2

	return cameraOffsetX, cameraOffsetY
}
func UpdateScreenPos(sh *ScreenHandler, playerX, playerY int, ofx, ofy float64) {
	sh.xScreenMin = playerX - sh.width/2 - int(ofx)
	sh.yScreenMin = playerY - sh.height/2 - int(ofy)
	sh.xScreenMax = sh.xScreenMin + sh.width
	sh.yScreenMax = sh.yScreenMin + sh.height
	sh.Fuck = sh.xScreenMin
	text.AddToDebug(fmt.Sprintf("sh pivot x = %v, sh pivot y = %v", sh.xScreenMin, sh.yScreenMin))
}
func ScreenPivot(sh ScreenHandler) (int, int) {

	return sh.xScreenMin, sh.yScreenMin
}
func ScreenBounds(sh ScreenHandler) (int, int) {
	return sh.xScreenMax, sh.yScreenMax
}
