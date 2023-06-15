package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type moveableobj struct {
	img  *ebiten.Image
	xpos int
	ypos int
	rot  int
}

type debugWindow struct {
	out             string
	xpos            int
	ypos            int
	xsize           int
	ysize           int
	background      *ebiten.Image
	backgroundColor color.RGBA
}
type mouseCursor struct {
	xPos  int
	yPos  int
	angle float64
}

// var img *ebiten.Image
var player moveableobj
var debug debugWindow
var mouse mouseCursor

func init() {
	var err error

	player.img, _, err = ebitenutil.NewImageFromFile("trust.png")
	if err != nil {
		log.Fatal(err)
	}
	player.xpos = 100
	player.ypos = 100

	debug.xsize = 220
	debug.ysize = 30
	debug.backgroundColor = color.RGBA{0, 0, 0, 100}
	debug.background = ebiten.NewImage(debug.xsize, debug.ysize)
	debug.background.Fill(debug.backgroundColor)

}

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	player.xpos += 0
	player.ypos += 0
	player.rot += 1

	mouse.xPos, mouse.yPos = ebiten.CursorPosition()
	mouse.angle = math.Atan2(float64(player.xpos)-float64(mouse.xPos), float64(player.ypos)-float64(mouse.yPos))

	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.CompositeMode = 0
	centerX := screen.Bounds().Dx()
	centerY := screen.Bounds().Dy()
	middleX := float64(player.img.Bounds().Dx()) / 2
	middleY := float64(player.img.Bounds().Dy()) / 2

	// translate object to half of its width and height before and after rotating to make it spin around its center
	op.GeoM.Translate(-middleX, -middleY)
	op.GeoM.Rotate(float64(player.rot) / 360)
	op.GeoM.Translate(middleX, middleY)

	op.GeoM.Translate(float64(centerX), float64(centerY))
	op.GeoM.Scale(0.3, 0.3)

	// zuerst in Ursprung verschieben?
	//op.GeoM.Translate(float64(player.xpos), float64(player.ypos))

	screen.DrawImage(player.img, op)
	// Write your game's rendering.
	myangle := fmt.Sprintf("%f", mouse.angle)
	var playerPosition string = "player x-postion: " + strconv.Itoa(player.xpos) + ", player y-position: " + strconv.Itoa(player.ypos) + "\n"
	var cursorPosition string = "cursor x-postion: " + strconv.Itoa(mouse.xPos) + ", cursor y-position: " + strconv.Itoa(mouse.yPos) + "\n"
	var cursorAngle string = "mouse angle (atan2): " + myangle + "\n"
	debug.out = playerPosition + cursorPosition + cursorAngle

	screen.DrawImage(debug.background, nil)
	ebitenutil.DebugPrintAt(screen, debug.out, debug.xpos, debug.ypos)

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
	//fmt.Printf("player x-postion: %d, player y-position: %d\n", player.xpos, player.ypos)
}
