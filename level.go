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
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type level struct {
	firstPlan     *ebiten.Image
	secondPlan    *ebiten.Image
	thirdPlan     *ebiten.Image
	fourthPlan    *ebiten.Image
	background    *ebiten.Image
	spawnSequence []spawn
	currentSpawn  int
	currentFrame  int
}

type spawn struct {
	enemies []enemySpawn
	frame   int
}

type enemySpawn struct {
	enemyType int
	y         float64
}

const (
	firstPlanPxPerFrame  = 4
	secondPlanPxPerFrame = 2
	thirdPlanPxPerFrame  = 1
	fourthPlanPxPerFrame = 0.25
	planImageWidth       = 3824
)

func initLevel() level {
	var l level
	img, _, err := ebitenutil.NewImageFromFile("assets/Montagnes-1.png")
	if err != nil {
		panic(err)
	}
	l.firstPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Montagnes-2.png")
	if err != nil {
		panic(err)
	}
	l.secondPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Montagnes-3.png")
	if err != nil {
		panic(err)
	}
	l.thirdPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Lune.png")
	if err != nil {
		panic(err)
	}
	l.fourthPlan = img

	img, _, err = ebitenutil.NewImageFromFile("assets/Etoiles.png")
	if err != nil {
		panic(err)
	}
	l.background = img

	l.spawnSequence = []spawn{
		spawn{
			enemies: []enemySpawn{
				enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 4},
				enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 4},
			},
			frame: 100,
		},
		spawn{
			enemies: []enemySpawn{
				enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 8},
				enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 8},
				enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 8},
				enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
			},
			frame: 200,
		},
		spawn{
			enemies: []enemySpawn{
				enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 8},
				enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 8},
				enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 8},
				enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
			},
			frame: 300,
		},
		spawn{
			enemies: []enemySpawn{
				enemySpawn{enemyType: staticEnemy, y: float64(screenHeight) / 8},
				enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 8},
				enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 8},
				enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
			},
			frame: 400,
		},
		spawn{
			enemies: []enemySpawn{
				enemySpawn{enemyType: staticFiringEnemy, y: float64(screenHeight) / 6},
				enemySpawn{enemyType: staticFiringEnemy, y: float64(3*screenHeight) / 6},
				enemySpawn{enemyType: staticFiringEnemy, y: float64(5*screenHeight) / 6},
			},
			frame: 500,
		},
		spawn{
			enemies: []enemySpawn{
				enemySpawn{enemyType: staticFiringDownEnemy, y: float64(5)},
				enemySpawn{enemyType: staticFiringUpEnemy, y: float64(screenHeight - 5)},
			},
			frame: 800,
		},
	}
	return l
}

func (l *level) update(es *enemySet) {
	l.currentFrame++
	if l.currentSpawn < len(l.spawnSequence) {
		if l.spawnSequence[l.currentSpawn].frame == l.currentFrame {
			for _, e := range l.spawnSequence[l.currentSpawn].enemies {
				es.addEnemy(e.enemyType, screenWidth-1, e.y)
			}
			l.currentSpawn++
		}
	} else {
		l.currentFrame = 0
		l.currentSpawn = 0
	}
}

func (l level) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(
		l.background,
		op,
	)

	fourthPlanStart := (int(fourthPlanPxPerFrame * float64(l.currentFrame))) % planImageWidth
	drawPlan(screen, l.fourthPlan, fourthPlanStart, op)

	thirdPlanStart := (thirdPlanPxPerFrame * l.currentFrame) % planImageWidth
	drawPlan(screen, l.thirdPlan, thirdPlanStart, op)

	secondPlanStart := (secondPlanPxPerFrame * l.currentFrame) % planImageWidth
	drawPlan(screen, l.secondPlan, secondPlanStart, op)

	firstPlanStart := (firstPlanPxPerFrame * l.currentFrame) % planImageWidth
	drawPlan(screen, l.firstPlan, firstPlanStart, op)
}

func drawPlan(screen, plan *ebiten.Image, start int, op *ebiten.DrawImageOptions) {
	screen.DrawImage(
		plan.SubImage(image.Rect(start, 0, start+screenWidth, screenHeight)).(*ebiten.Image),
		op,
	)
	if start+screenWidth > planImageWidth {
		missingPx := start + screenWidth - planImageWidth
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(screenWidth-float64(missingPx), 0)
		screen.DrawImage(
			plan.SubImage(image.Rect(0, 0, missingPx, screenHeight)).(*ebiten.Image),
			op2,
		)
	}
}
