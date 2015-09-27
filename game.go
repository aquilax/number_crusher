package main

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

type Game struct {
	level  int
	game   *tl.Game
	board  *Board
	player *Player
	status *tl.Text
}

func NewGame() *Game {
	game := &Game{
		level:  startLevel,
		game:   tl.NewGame(),
		board:  NewBoard(),
		player: NewPlayer(),
		status: tl.NewText(20, 0, "", tl.ColorWhite, tl.ColorBlack),
	}
	// TODO: This is ugly
	game.player.setGame(game)
	game.updateStatus()
	return game
}

func (g *Game) Run() {
	rand.Seed(time.Now().UTC().UnixNano())
	g.addChrome()
	g.buildLevel(1)
	g.game.Start()
}

func (g *Game) buildLevel(gameLevel int) {
	level := tl.NewBaseLevel(tl.Cell{})
	// TODO: Remove this abomination
	level.AddEntity(tl.NewRectangle(1, 1, 65, 33, tl.ColorGreen))
	for y := 2; y < 63; y = y + 8 {
		for x := 2; x < 31; x = x + 4 {
			level.AddEntity(tl.NewRectangle(y, x, 7, 3, tl.ColorBlue))
		}
	}
	g.board.populateBoard(gameLevel, answersPerLevel, level)
	level.AddEntity(g.player)
	foe := NewFoe(g)
	level.AddEntity(foe)
	g.game.Screen().SetLevel(level)
	g.updateStatus()
}

func (g *Game) addChrome() {
	g.game.Screen().AddEntity(tl.NewText(1, 0, " Number crusader! ", tl.ColorBlack, tl.ColorGreen))
	g.game.Screen().AddEntity(g.status)
}

func (g *Game) updateStatus() {
	statusText := fmt.Sprintf("Lvl %2d | ❤ %d | Score %06d", g.level, g.player.lives, g.player.score)
	g.status.SetText(statusText)
}

func (g *Game) nextLevel() {
	g.level++
	g.buildLevel(g.level)
}

func (g *Game) restartGame() {
	g.level = 1
	g.player.Init()
	g.buildLevel(g.level)
}

func (g *Game) gameOver() {
	g.game.Screen().Level().AddEntity(tl.NewText(28, 17, " GAME OVER ", tl.ColorBlack, tl.ColorRed))
}
