package text

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type debugWindow struct {
	out             string
	xpos            int
	ypos            int
	xsize           int
	ysize           int
	background      *ebiten.Image
	backgroundColor color.RGBA
}

var debug debugWindow

func InitDebug() {

	debug.xsize = 220
	debug.ysize = 30
	debug.backgroundColor = color.RGBA{0, 0, 0, 100}
	debug.background = ebiten.NewImage(debug.xsize, debug.ysize)
	debug.background.Fill(debug.backgroundColor)
}
func ClearDebug() {
	debug.out = ""
}
func AddToDebug(addition string) {
	debug.out = debug.out + "\n" + addition
}
func DebugBackground() *ebiten.Image {
	return debug.background
}
func DebugXpos() int {
	return debug.xpos
}
func DebugYpos() int {
	return debug.ypos
}
func DebugMsg() string {
	return debug.out
}
