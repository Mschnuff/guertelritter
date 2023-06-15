package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	inp "github.com/mschnuff/guertelritter/input"
	text "github.com/mschnuff/guertelritter/text"
)

type moveableobj struct {
	img  *ebiten.Image
	xpos int
	ypos int
	rot  int
}

var player moveableobj
var mouseAngle float64

func init() {
	var err error

	player.img, _, err = ebitenutil.NewImageFromFile("trust.png")
	if err != nil {
		log.Fatal(err)
	}
	player.xpos = 320
	player.ypos = 240

	// initialise the debug-window
	text.InitDebug()

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
	middleX := float64(player.img.Bounds().Dx()) / 2
	middleY := float64(player.img.Bounds().Dy()) / 2

	// translate object to half of its width and height before and after rotating to make it spin around its center
	// we temporarily deactivated reversing the translation
	op.GeoM.Translate(-middleX, -middleY)

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
