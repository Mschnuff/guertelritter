package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	inp "github.com/mschnuff/guertelritter/input"
	text "github.com/mschnuff/guertelritter/text"
)

type moveableobj struct {
	img     *ebiten.Image
	xpos    int
	ypos    int
	rot     int
	middleX float64
	middleY float64
}

type backgroundGraphics struct {
	img  [16]*ebiten.Image
	xpos [16]int
	ypos [16]int
}

var player moveableobj
var mouseAngle float64
var bg backgroundGraphics

func init() {
	// init player TODO: move this code to a different function
	var err error

	// this path depends on the operating system. TODO: look into this problem while working from home (on windows system)
	var imgFolder string = "./static/images/"
	player.img, _, err = ebitenutil.NewImageFromFile(imgFolder + "trust.png")
	if err != nil {
		log.Fatal(err)
	}
	player.xpos = 320
	player.ypos = 240
	player.middleX = float64(player.img.Bounds().Dx()) / 2
	player.middleY = float64(player.img.Bounds().Dy()) / 2

	initBackground(imgFolder)
	// initialise the debug-window
	text.InitDebug()

}
func initBackground(imgFolder string) {
	var err error
	imgSize := []int{512, 512}
	bgImgPositions := [16][2]int{
		{imgSize[0] * 0, imgSize[1] * 0}, // 0 xpos ypos
		{imgSize[0] * 1, imgSize[1] * 0}, // 1
		{imgSize[0] * 2, imgSize[1] * 0}, // 2
		{imgSize[0] * 3, imgSize[1] * 0}, // 3
		{imgSize[0] * 0, imgSize[1] * 1}, // 4
		{imgSize[0] * 1, imgSize[1] * 1}, // 5
		{imgSize[0] * 2, imgSize[1] * 1}, // 6
		{imgSize[0] * 3, imgSize[1] * 1}, // 7
		{imgSize[0] * 0, imgSize[1] * 2}, // 8
		{imgSize[0] * 1, imgSize[1] * 2}, // 9
		{imgSize[0] * 2, imgSize[1] * 2}, // 10
		{imgSize[0] * 3, imgSize[1] * 2}, // 11
		{imgSize[0] * 0, imgSize[1] * 3}, // 12
		{imgSize[0] * 0, imgSize[1] * 3}, // 13
		{imgSize[0] * 0, imgSize[1] * 3}, // 14
		{imgSize[0] * 0, imgSize[1] * 3}, // 15
	}

	for c := 0; c < len(bg.xpos); c++ {
		bg.img[c], _, err = ebitenutil.NewImageFromFile(imgFolder + "background_template_" + text.IntToStr(c+1) + ".png")
		if err != nil {
			log.Fatal(err)
		}
		bg.xpos[c] = bgImgPositions[c][0]
		bg.ypos[c] = bgImgPositions[c][1]
		fmt.Print("background image num: " + text.IntToStr(c+1) + ", x: " + text.IntToStr(bg.xpos[c]) + ", y: " + text.IntToStr(bg.ypos[c]) + "\n")
	}

}

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// we have to clear debug on every update
	text.ClearDebug()
	player.xpos, player.ypos = inp.HandlePlayerMovement(player.xpos, player.ypos)
	text.AddToDebug("player x-postion: " + text.IntToStr(player.xpos) + ", player y-position: " + text.IntToStr(player.ypos))

	// addtodebug can be used to conveniently add debug messages
	mouseAngle = inp.GetCursorToPlayerAngle(player.xpos, player.ypos)
	text.AddToDebug("mouse angle (atan2): " + text.FloatToStr(mouseAngle))
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.CompositeMode = 0
	//centerX := screen.Bounds().Dx()
	//centerY := screen.Bounds().Dy()

	// translate object to half of its width and height before and after rotating to make it spin around its center
	// we temporarily deactivated reversing the translation
	op.GeoM.Translate(-player.middleX, -player.middleY)

	op.GeoM.Rotate(-mouseAngle)
	//op.GeoM.Translate(middleX, middleY)

	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(float64(player.xpos), float64(player.ypos))

	// steps:
	// #1 translate to middle of object
	// #2 rotate
	// #3 scale
	// #4 translate to actual position

	// draw player object
	screen.DrawImage(player.img, op)

	// draw debug-background and debug-message
	screen.DrawImage(text.DebugBackground(), nil)
	ebitenutil.DebugPrintAt(screen, text.DebugMsg(), text.DebugXpos(), text.DebugYpos())

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}

	// Specify the window size as you like. Here, a doubled size is specified.
	var winWidth, winHeight int = 640, 480
	ebiten.SetWindowSize(game.Layout(winWidth, winHeight))
	ebiten.SetWindowTitle("GuertelRitter")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
