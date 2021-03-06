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
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type powerUp struct {
	x         float64
	y         float64
	vx        float64
	vy        float64
	xSize     float64
	ySize     float64
	collision bool
}

func (p *powerUp) update() {
	p.x += p.vx
	p.y += p.vy
}

func (p powerUp) isOut() bool {
	return p.collision || p.xmax() < 0 || p.ymax() < 0 || p.xmin() >= screenWidth || p.ymin() >= screenHeight
}

func (p powerUp) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.xmin(), p.ymin())
	screen.DrawImage(bonusImage, op)
	if isDebug() {
		cHull := p.convexHull()
		for i := 0; i < len(cHull); i++ {
			ii := (i + 1) % len(cHull)
			ebitenutil.DrawLine(screen, cHull[i].x, cHull[i].y, cHull[ii].x, cHull[ii].y, color.RGBA{255, 0, 0, 255})
		}
	}
}

func (p *powerUp) xmin() float64 {
	return p.x - p.xSize/2
}

func (p *powerUp) xmax() float64 {
	return p.x + p.xSize/2
}

func (p *powerUp) ymin() float64 {
	return p.y - p.ySize/2
}

func (p *powerUp) ymax() float64 {
	return p.y + p.ySize/2
}

func (p *powerUp) convexHull() []point {
	return []point{
		point{p.xmin(), p.ymin()},
		point{p.xmax(), p.ymin()},
		point{p.xmax(), p.ymax()},
		point{p.xmin(), p.ymax()},
	}
}

func (p *powerUp) hasCollided() {
	p.collision = true
}

type powerUpSet struct {
	numPowerUps int
	powerUps    []*powerUp
}

func initPowerUpSet() powerUpSet {
	return powerUpSet{
		numPowerUps: 0,
		powerUps:    make([]*powerUp, 0),
	}
}

func (ps *powerUpSet) addPowerUp(p powerUp) {
	p.xSize = 44
	p.ySize = 44
	ps.numPowerUps++
	ps.powerUps = append(ps.powerUps, &p)
}

func (ps *powerUpSet) update() {
	for pos := 0; pos < ps.numPowerUps; pos++ {
		ps.powerUps[pos].update()
		if ps.powerUps[pos].isOut() {
			ps.numPowerUps--
			ps.powerUps[pos] = ps.powerUps[ps.numPowerUps]
			ps.powerUps = ps.powerUps[:ps.numPowerUps]
			pos--
		}
	}
}

func (ps *powerUpSet) draw(screen *ebiten.Image) {
	for pos := 0; pos < ps.numPowerUps; pos++ {
		ps.powerUps[pos].draw(screen)
	}
}
