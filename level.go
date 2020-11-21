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
)

type level struct {
	firstPlan      *ebiten.Image
	secondPlan     *ebiten.Image
	thirdPlan      *ebiten.Image
	fourthPlan     *ebiten.Image
	background     *ebiten.Image
	spawnSequence  []spawn
	currentSpawn   int
	currentFrame   int
	lastSpawnFrame int
	bossBattle     bool
}

type spawn struct {
	enemies    []enemySpawn
	frameDelay int
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

func (l *level) update(es *enemySet, bs *bossSet) {
	if !l.bossBattle {
		l.currentFrame++
		if l.currentSpawn < len(l.spawnSequence) {
			if l.lastSpawnFrame+l.spawnSequence[l.currentSpawn].frameDelay == l.currentFrame {
				for _, e := range l.spawnSequence[l.currentSpawn].enemies {
					if e.enemyType >= midBoss1 {
						bs.addBoss(e.enemyType, screenWidth-1, e.y)
						l.bossBattle = true
						for ePos := 0; ePos < es.numEnemies; ePos++ {
							es.enemies[ePos].vx += firstPlanPxPerFrame
						}
					} else {
						es.addEnemy(e.enemyType, screenWidth-1, e.y)
					}
				}
				l.currentSpawn++
				l.lastSpawnFrame = l.currentFrame
			}
		}
	} else {
		if bs.numBosses == 0 {
			l.bossBattle = false
			for ePos := 0; ePos < es.numEnemies; ePos++ {
				es.enemies[ePos].vx -= firstPlanPxPerFrame
			}
		}
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
