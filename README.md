# About this repository:
With this project I test the capabilities of the **ebiten** game engine while also practicing working with golang at the same time by building a game prototype. The game is themed around an astronaut flying through space. The astronaut is also supposed to encounter asteroids on which he can land and thus switch to a different mode of movement. 

## Links
* [ebiten tutorials and info](https://ebitengine.org/)

## next Steps (TODO):
- [X] make player object rotate towards the mouse Cursor
- [X] draw arrow on player object to make it obvious what it is facing atm
- [ ] refactor the main file to generalise the functionality so it can be used for other objects (besides player-object)
- [X] figure out how embed other go files (in the same directory and in other directories)
- [X] add keyboard input for moving the player-object
- [ ] move player-assoctiated code from init function to a different function dedicated to the player-object only
- [ ] figure out if 'op.GeoM.Translate' etc should be called in the draw or the update function
- [ ] change game logic so the player-object does not leave screen. instead it moves relative to the background
- [ ] create tile-based background (tiles are supposed to rearrange automatically)
- [ ] look into os dependant paths to files
- [ ] learn how to write tests in golang