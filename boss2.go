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
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	boss2Points                  = 15000
	boss2PV                      = 10 //750
	boss2DeathAnimationFrames    = 180
	boss2BulletXShift            = -250
	boss2BulletMaxY              = 900
	boss2BulletMinY              = 200
	boss2LastPhaseFramePerBullet = 50
	boss2Phase1NumBullets        = 14
	boss2Phase1NumBulletsMore    = 10
	boss2Phase1FramePerBullet    = 150
	boss2Phase1BulletSpeed       = 3
)

func makeBoss2(y float64) boss {
	x := float64(screenWidth - 1 + 711/2)
	hitB1 := bossHitBox{
		x:     x,
		xrel:  120,
		y:     y,
		yrel:  -20,
		xSize: 100,
		ySize: 100,
		hullShape: []point{
			point{x: -20, y: 20},
			point{x: -20, y: -20},
			point{x: 20, y: -20},
			point{x: 20, y: 20},
		},
	}
	hitB1.updateBox()
	hurtB1 := bossHitBox{
		x:     x,
		xrel:  0,
		y:     y,
		yrel:  -210,
		xSize: 700,
		ySize: 400,
		hullShape: []point{
			point{x: -350, y: -110},
			point{x: -260, y: -200},
			point{x: 350, y: 100},
			point{x: 350, y: 200},
			point{x: -220, y: 200},
		},
	}
	hurtB1.updateBox()
	hurtB2 := bossHitBox{
		x:     x,
		xrel:  0,
		y:     y,
		yrel:  202,
		xSize: 700,
		ySize: 416,
		hullShape: []point{
			point{x: -350, y: 110},
			point{x: -200, y: 208},
			point{x: 350, y: -120},
			point{x: 350, y: -208},
			point{x: -220, y: -208},
		},
	}
	hurtB2.updateBox()
	return boss{
		x:              x,
		xSize:          711,
		y:              y + 30,
		ySize:          825,
		pv:             boss2PV,
		bossType:       boss2,
		points:         boss2Points,
		hitBoxes:       []bossHitBox{hitB1},
		hurtBoxes:      []bossHitBox{hurtB1, hurtB2},
		numDeathFrames: boss2DeathAnimationFrames,
	}
}

func (b *boss) boss2Update(bs *bulletSet, p *player) bool {
	if b.pv <= 0 {
		b.pv = 1
		b.phase = 100
		b.hitBoxes = []bossHitBox{}
		b.hurtBoxes = []bossHitBox{}
	}
	switch b.phase {
	case 0:
		if b.x < 6*screenWidth/7 {
			b.phase = 1
			b.hitBoxes[0].hitable = true
		}
		b.x -= 5
	case 1:
		b.frame++
		if b.frame >= boss2Phase1FramePerBullet {
			b.frame = 0
			noBulletNum := rand.Intn(boss2Phase1NumBullets-2) + 1
			for bNum := 0; bNum < boss2Phase1NumBullets; bNum++ {
				if bNum != noBulletNum {
					y := float64(boss2BulletMinY) + float64(bNum*(boss2BulletMaxY-boss2BulletMinY))/float64(boss2Phase1NumBullets-1)
					bs.addBullet(bullet{
						x:     b.x + boss2BulletXShift,
						y:     y,
						vx:    -boss2Phase1BulletSpeed,
						vy:    0,
						image: enemyBasicBullet,
					})
				}
			}
			for bNum := 0; bNum < boss2Phase1NumBulletsMore; bNum++ {
				bs.addBullet(bullet{
					x:     b.x + boss2BulletXShift,
					y:     boss2BulletMaxY,
					vx:    -boss2Phase1BulletSpeed,
					vy:    +0.125 * float64(bNum),
					image: enemyBasicBullet,
				})
				bs.addBullet(bullet{
					x:     b.x + boss2BulletXShift,
					y:     boss2BulletMinY,
					vx:    -boss2Phase1BulletSpeed,
					vy:    -0.125 * float64(bNum),
					image: enemyBasicBullet,
				})
			}
		}
	case 100:
		b.frame++
		hb := bossHitBox{
			x:     b.x,
			xrel:  120,
			y:     b.y,
			yrel:  -20,
			xSize: 100,
			ySize: 100,
			hullShape: []point{
				point{x: -20, y: 20},
				point{x: -20, y: -20},
				point{x: 20, y: -20},
				point{x: 20, y: 20},
			},
		}
		hb.updateBox()
		if collideNoHarm(p, &hb) {
			b.pv = 0
		}
		if b.frame > boss2LastPhaseFramePerBullet {
			y := float64(rand.Intn(boss2BulletMaxY-boss2BulletMinY+1) + boss2BulletMinY)
			vy := float64(rand.Intn(5) - 2)
			bs.addBullet(bullet{
				x:     b.x + boss2BulletXShift,
				y:     y,
				vx:    -2,
				vy:    vy,
				image: enemyBasicBullet,
			})
			b.frame = 0
		}
	}
	return false
}

func (b *boss) boss2Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x-b.xSize/2, b.y-b.ySize/2)
	if b.pv <= 0 {
		alpha := float64(b.deathAnimationFrame) / float64(b.numDeathFrames)
		op.ColorM.Translate(0, 0, 0, -alpha)
	}
	screen.DrawImage(
		boss2Image,
		op,
	)
}