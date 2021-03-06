/* A game for Game Off 2020
// Copyright (C) 2020 Loïg Jezequel
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>
*/

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		}
	}

	g.updateMusic()

	switch g.state {
	case gameWelcome:
		err := g.welcomeUpdate()
		if err != nil {
			return err
		}
	case gameHelp:
		g.helpUpdate()
	case gameJoypadSetup:
		g.joypadSetupUpdate()
	case gameInfo:
		g.infoUpdate()
	case gameIntro:
		g.introUpdate()
	case gameInLevel1, gameInLevel2:
		if g.isEnterJustPressed() {
			if g.state == gameInLevel1 {
				g.state = gameInLevel1Paused
			} else {
				g.state = gameInLevel2Paused
			}
		}
		if g.stateFrame < framesBeforeLevel {
			g.stateFrame++
		}
		g.bulletSet.update()
		g.enemySetUpdate()
		g.bossSetUpdate()
		g.powerUpSet.update()
		//if g.stateFrame >= framesBeforeLevel {
		g.playerUpdate()
		//}
		g.explosionSetUpdate()
		g.levelUpdate()
		g.checkCollisions()
	case gameInLevel1Paused, gameInLevel2Paused:
		if g.isEnterJustPressed() {
			if g.state == gameInLevel1Paused {
				g.state = gameInLevel1
			} else {
				g.state = gameInLevel2
			}
		}
	case gameTransition:
		g.transitionUpdate()
	case gameFinished:
		err := g.finishedUpdate()
		if err != nil {
			return err
		}
	case gameOver:
		g.gameOverUpdate()
	}
	return nil
}
