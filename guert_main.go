package main

// --------------------------- Imports --------------------------- //
import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/mschnuff/guertelritter/graphics"
	inp "github.com/mschnuff/guertelritter/input"
	text "github.com/mschnuff/guertelritter/text"
)

// --------------------------- Structs --------------------------- //
type moveableobj struct {
	img     *ebiten.Image
	xpos    int
	ypos    int
	rot     int
	middleX float64
	middleY float64
	scale   float64
}

type backgroundGraphics struct {
	img  [16]*ebiten.Image
	xpos [16]int
	ypos [16]int
	op   [16]*ebiten.DrawImageOptions
	//imgSize [2]int
	width        int
	height       int
	xScaleFactor float64
	yScaleFactor float64
}

type graphiceSettings struct {
	winWidth  int
	winHeight int
	sh        graphics.ScreenHandler
}
type mouseHandler struct {
	x       int
	y       int
	angle   float64
	offsetX float64
	offsetY float64
}

// Game implements ebiten.Game interface.
type Game struct{}

// --------------------------- Variables --------------------------- //
var gset graphiceSettings
var player moveableobj
var mouseAngle float64
var bg backgroundGraphics
var mouse mouseHandler

// --------------------------- Init - Start --------------------------- //

func init() {
	initGraphics()
	inp.SetMovementSpeed(3)

	var imgFolder string = "./static/images/"
	initPlayer(imgFolder)
	initBackground(imgFolder)
	// initialise the debug-window
	text.InitDebug()

}
func initPlayer(imgFolder string) {
	var err error
	// this path depends on the operating system. TODO: look into this problem while working from home (on windows system)
	// 25.06.2023 works fine from home

	player.img, _, err = ebitenutil.NewImageFromFile(imgFolder + "trust.png")
	if err != nil {
		log.Fatal(err)
	}
	player.xpos = gset.winWidth / 2
	player.ypos = gset.winHeight / 2
	player.scale = 0.3
	player.middleX = float64(player.img.Bounds().Dx()) / 2
	player.middleY = float64(player.img.Bounds().Dy()) / 2
}
func initGraphics() {
	// 96 * 16, 96 * 9
	gset.winWidth = 1536
	gset.winHeight = 864
	gset.sh = graphics.InitScreenHandler(gset.winWidth, gset.winHeight)
}
func initBackground(imgFolder string) {
	var err error

	bgImgPositions := [16][2]int{
		{0, 0}, // 0 xpos ypos
		{1, 0}, // 1
		{2, 0}, // 2
		{3, 0}, // 3
		{0, 1}, // 4
		{1, 1}, // 5
		{2, 1}, // 6
		{3, 1}, // 7
		{0, 2}, // 8
		{1, 2}, // 9
		{2, 2}, // 10
		{3, 2}, // 11
		{0, 3}, // 12
		{1, 3}, // 13
		{2, 3}, // 14
		{3, 3}, // 15
	}

	for c := 0; c < len(bg.xpos); c++ {
		currentImgPath := imgFolder + "background_template_" + text.IntToStr(c+1) + ".png"
		bg.img[c], _, err = ebitenutil.NewImageFromFile(currentImgPath)
		if err != nil {
			log.Fatal(err)
		}

		bg.width = bg.img[c].Bounds().Dx()
		bg.height = bg.img[c].Bounds().Dy()
		bg.xpos[c] = bgImgPositions[c][0] * bg.width
		bg.ypos[c] = bgImgPositions[c][1] * bg.height

		// ScaleFactor not in use
		bg.xScaleFactor = (float64((gset.winWidth / 4)) / float64(bg.width))
		bg.yScaleFactor = (float64((gset.winHeight / 4)) / float64(bg.height))

	}

}

// --------------------------- Init - End --------------------------- //

// --------------------------- Update - Start --------------------------- //

func updateBackground() {
	//TODO: move this in separate file aswell
	//offX, offY := graphics.CalcOffset(gset.sh, mouse.x, mouse.y)

	xBoundScreenMax := gset.winWidth/2 + player.xpos - int(mouse.offsetX)
	xBoundScreenMin := player.xpos - gset.winWidth/2 - int(mouse.offsetX)
	yBoundScreenMax := gset.winHeight/2 + player.ypos - int(mouse.offsetY)
	yBoundScreenMin := player.ypos - gset.winHeight/2 - int(mouse.offsetY)

	/*minX := "min Screen xpos: " + text.IntToStr(xBoundScreenMin)
	maxX := "max Screen xpos: " + text.IntToStr(xBoundScreenMax)
	minY := "min Screen ypos: " + text.IntToStr(yBoundScreenMin)
	maxY := "max Screen ypos: " + text.IntToStr(yBoundScreenMax)
	text.AddToDebug(minX + ", " + maxX + ", " + minY + ", " + maxY)*/

	for i := 0; i < len(bg.xpos); i++ {
		/*onScreen := false

		xMax := bg.xpos[i] < xBoundScreenMax
		xMin := bg.xpos[i] >= xBoundScreenMin-bg.width
		yMax := bg.ypos[i] < yBoundScreenMax
		yMin := bg.ypos[i] >= yBoundScreenMin-bg.height
		onScreen = xMax && xMin && yMax && yMin
		*/
		if bg.xpos[i]+bg.width < xBoundScreenMin {
			bg.xpos[i] = bg.xpos[i] + bg.width*4
		}
		if bg.xpos[i] > xBoundScreenMax {
			bg.xpos[i] = bg.xpos[i] - bg.width*4
		}
		if bg.ypos[i]+bg.height < yBoundScreenMin {
			bg.ypos[i] = bg.ypos[i] + bg.height*4
		}
		if bg.ypos[i] > yBoundScreenMax {
			bg.ypos[i] = bg.ypos[i] - bg.height*4
		}
		/*text.AddToDebug(
		"xpos background image [" + text.IntToStr(i+1) + "]: " +
			text.IntToStr(bg.xpos[i]) + ", " +
			text.IntToStr(bg.ypos[i]) + " -> image is on screen: " +
			text.BooltoStr(onScreen))
		*/
	}

}
func updateMouse() {
	mouse.x, mouse.y = ebiten.CursorPosition()
	mouse.angle = inp.GetCursorToPlayerAngle(gset.winWidth/2, gset.winHeight/2)
	mouse.offsetX, mouse.offsetY = graphics.CalcOffset(gset.sh, mouse.x, mouse.y)
	//text.AddToDebug("mouse angle (atan2): " + text.FloatToStr(mouse.angle))
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// we have to clear debug on every update
	text.ClearDebug()
	player.xpos, player.ypos = inp.HandlePlayerMovement(player.xpos, player.ypos)
	updateMouse()
	updateBackground()
	// addtodebug can be used to conveniently add debug messages
	text.AddToDebug("player x-postion: " + text.IntToStr(player.xpos) + ", player y-position: " + text.IntToStr(player.ypos))

	// Write your game's logical update.
	return nil
}

// --------------------------- Update - End --------------------------- //

// --------------------------- Draw - Start --------------------------- //

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	// translate object to half of its width and height before and after rotating to make it spin around its center
	// we temporarily deactivated reversing the translation

	// steps:
	// #1 translate to middle of object
	// #2 rotate
	// #3 scale
	// #4 translate to actual position

	text.AddToDebug(fmt.Sprintf("screen: (%v, %v) offset: (%v, %v)", gset.sh.MiddleX, gset.sh.MiddleY, mouse.offsetX, mouse.offsetY))

	for i := 0; i < len(bg.img); i++ {

		bgop := &ebiten.DrawImageOptions{}

		pivotX := player.xpos - gset.winWidth/2
		pivotY := player.ypos - gset.winHeight/2
		calcPosX := float64(bg.xpos[i] - pivotX)
		calcPosY := float64(bg.ypos[i] - pivotY)

		bgop.GeoM.Translate(mouse.offsetX, mouse.offsetY)
		bgop.GeoM.Translate(calcPosX, calcPosY)
		// draw background first
		screen.DrawImage(bg.img[i], bgop)
	}

	// player object
	op := &ebiten.DrawImageOptions{}
	op.CompositeMode = 0
	op.GeoM.Translate(-player.middleX, -player.middleY)
	op.GeoM.Rotate(-mouse.angle)
	op.GeoM.Translate(mouse.offsetX*(1/player.scale), mouse.offsetY*(1/player.scale))
	op.GeoM.Scale(player.scale, player.scale)
	op.GeoM.Translate(gset.sh.MiddleX, gset.sh.MiddleY)
	// draw player object
	screen.DrawImage(player.img, op)

	// draw debug-background and debug-message
	screen.DrawImage(text.DebugBackground(), nil)
	ebitenutil.DebugPrintAt(screen, text.DebugMsg(), text.DebugXpos(), text.DebugYpos())

}

// --------------------------- Draw - End --------------------------- //

// --------------------------- Main - Start --------------------------- //

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}

	// window size  is specified in the init()-Function. feel free to change it.

	ebiten.SetWindowSize(game.Layout(gset.winWidth, gset.winHeight))
	ebiten.SetWindowTitle("GuertelRitter")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
