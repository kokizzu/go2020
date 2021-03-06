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

type bullet struct {
	name        string
	x           float64
	imageXShift float64
	imageYShift float64
	y           float64
	vx          float64
	vy          float64
	ax          float64
	ay          float64
	xSize       float64
	ySize       float64
	collision   bool
	hullSet     bool
	cHull       []point
	xMin        float64
	yMin        float64
	xMax        float64
	yMax        float64
	isBig       bool
	image       *ebiten.Image
}

const (
	basicBulletSize = 20
)

func (b *bullet) update() {
	b.vx += b.ax
	b.vy += b.ay
	b.x += b.vx
	b.y += b.vy
	b.hullSet = false
	b.xMin = b.x - b.xSize/2
	b.xMax = b.x + b.xSize/2
	b.yMin = b.y - b.ySize/2
	b.yMax = b.y + b.ySize/2
}

func (b bullet) isOut() bool {
	return b.collision || b.xmax() < 0 || b.ymax() < 0 || b.xmin() >= screenWidth+200 || b.ymin() >= screenHeight
}

func (b bullet) draw(screen *ebiten.Image, color color.Color) {
	if b.image != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(b.xmin()+b.imageXShift, b.ymin()+b.imageYShift)
		screen.DrawImage(
			b.image,
			op,
		)
	}
	if isDebug() {
		cHull := b.convexHull()
		for i := 0; i < len(cHull); i++ {
			ii := (i + 1) % len(cHull)
			ebitenutil.DrawLine(screen, cHull[i].x, cHull[i].y, cHull[ii].x, cHull[ii].y, color)
		}
	}
}

func (b *bullet) xmin() float64 {
	return b.xMin
}

func (b *bullet) xmax() float64 {
	return b.xMax
}

func (b *bullet) ymin() float64 {
	return b.yMin
}

func (b *bullet) ymax() float64 {
	return b.yMax
}

func (b *bullet) convexHull() []point {
	if !b.hullSet {
		if b.isBig {
			b.cHull = []point{
				point{b.xmin(), b.ymin()},
				point{b.xmin(), b.ymax()},
				point{b.xmax(), (b.ymax()+b.ymin())/2 + 30},
				point{b.xmax(), (b.ymax()+b.ymin())/2 - 30},
			}
		} else {
			b.cHull = []point{
				point{b.xmin(), b.ymin()},
				point{b.xmax(), b.ymin()},
				point{b.xmax(), b.ymax()},
				point{b.xmin(), b.ymax()},
			}
		}
		b.hullSet = true
	}
	return b.cHull
}

func (b *bullet) hasCollided() {
	b.collision = true
}

type bulletSet struct {
	numBullets int
	bullets    []*bullet
}

func initBulletSet() bulletSet {
	return bulletSet{
		numBullets: 0,
		bullets:    make([]*bullet, 0),
	}
}

func (bs *bulletSet) addBullet(b bullet) {
	bb := b
	bb.xSize = basicBulletSize
	bb.ySize = basicBulletSize
	bb.xMin = bb.x - bb.xSize/2
	bb.xMax = bb.x + bb.xSize/2
	bb.yMin = bb.y - bb.ySize/2
	bb.yMax = bb.y + bb.ySize/2
	bs.numBullets++
	bs.bullets = append(bs.bullets, &bb)
}

func (bs *bulletSet) addBigBullet(b bullet) {
	bb := b
	bb.xSize = 20
	bb.ySize = 140
	bb.xMin = bb.x - bb.xSize/2
	bb.xMax = bb.x + bb.xSize/2
	bb.yMin = bb.y - bb.ySize/2
	bb.yMax = bb.y + bb.ySize/2
	bb.isBig = true
	bs.numBullets++
	bs.bullets = append(bs.bullets, &bb)
}

func (bs *bulletSet) update() {
	for pos := 0; pos < bs.numBullets; pos++ {
		bs.bullets[pos].update()
		if bs.bullets[pos].isOut() {
			bs.numBullets--
			bs.bullets[pos] = bs.bullets[bs.numBullets]
			bs.bullets = bs.bullets[:bs.numBullets]
			pos--
		}
	}
}

func (bs *bulletSet) draw(screen *ebiten.Image, color color.Color) {
	for pos := 0; pos < bs.numBullets; pos++ {
		bs.bullets[pos].draw(screen, color)
	}
}
