package graphics

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
func isOnScreen(xpos float64, ypos float64, width int, height int) bool {
	return false
}

func CalcOffset(sh ScreenHandler, mouseX int, mouseY int) (float64, float64) {

	cameraOffsetX := (sh.MiddleX - float64(mouseX)) / 2
	cameraOffsetY := (sh.MiddleY - float64(mouseY)) / 2

	return cameraOffsetX, cameraOffsetY
}
