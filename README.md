# About this repository:
With this project I test the capabilities of the **ebiten** game engine while also practicing working with golang at the same time by building a game prototype. The game is themed around an astronaut flying through space. The astronaut is also supposed to encounter asteroids on which he can land and thus switch to a different mode of movement. 

![Random Screenshot of early Development](https://github.com/Mschnuff/guertelritter/blob/main/static/images/example_screenshot.PNG?raw=true)

## Links
* [ebiten tutorials and info](https://ebitengine.org/)
* [golang boot dev](https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/10b92b5a-6687-474e-9df0-e215b7d0a46d/9aedf839-7d94-43f7-82d0-1d27e5d0b79c)
* [golang code camp video](https://www.youtube.com/watch?v=un6ZyFkqFKo&t=3080s)


## New Features
* Added Background (now also rearranges itself correctly)
* Allow multiple window resolutions
* Allow different player speed
* added camera offset


## next Steps (TODO):
- [X] make player object rotate towards the mouse Cursor
- [X] draw arrow on player object to make it obvious what it is facing atm
- [ ] refactor the main file to generalise the functionality so it can be used for other objects (besides player-object)
- [X] figure out how embed other go files (in the same directory and in other directories)
- [X] add keyboard input for moving the player-object
- [ ] move player-assoctiated code from init function to a different function dedicated to the player-object only
- [X] figure out if 'op.GeoM.Translate' etc should be called in the draw or the update function
- [X] change game logic so the player-object does not leave screen. instead it moves relative to the background
- [X] create tile-based background (tiles are supposed to rearrange automatically)
- [ ] look into os dependant paths to files
- [ ] learn how to write tests in golang
- [X] figure out how to fix movement with events (player stops if two opposite buttons are pressed at the same time)
- [ ] add comments to code that explain what's going on
- [ ] fix background (now messed up due to camera offset)
- [ ] distinguish between logical screen and graphics