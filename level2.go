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
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loig/go2020/assets"
)

func (g *game) setUpLevel2() {

	loadLevel2Enemies()
	loadLevel2Background()

	var l level
	l.firstPlanHeight = 271
	l.secondPlanHeight = 419
	l.thirdPlanHeight = 465
	l.fourthPlanYShift = 7
	l.fourthPlanTotalXShift = screenWidth - 488
	l.levelTotalFrames = 0
	l.spawnSequence = level2SpawnSequence

	g.level = l
	g.bulletSet = initBulletSet()
	g.player.initialPosition()
	g.enemySet = initEnemySet()
	g.bossSet = initBossSet()
	g.powerUpSet = initPowerUpSet()
}

func loadLevel2Enemies() {
	img, _, err := image.Decode(bytes.NewReader(assets.Ennemi4))
	//img, _, err := ebitenutil.NewImageFromFile("assets/Ennemi4.png")
	if err != nil {
		panic(err)
	}
	movingFiringEnemyImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Ennemi42))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi4.2.png")
	if err != nil {
		panic(err)
	}
	movingFiringEnemyImage2 = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Ennemi43))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi4.3.png")
	if err != nil {
		panic(err)
	}
	movingFiringEnemyImage3 = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Ennemi1))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi1.png")
	if err != nil {
		panic(err)
	}
	staticEnemyImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Ennemi2))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Ennemi2.png")
	if err != nil {
		panic(err)
	}
	staticExplodingEnemyImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Midboss))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Midboss.png")
	if err != nil {
		panic(err)
	}
	midBoss1Images[0] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Midboss2))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Midboss2.png")
	if err != nil {
		panic(err)
	}
	midBoss1Images[1] = ebiten.NewImageFromImage(img)
	midBoss1Images[3] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Midboss3))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Midboss3.png")
	if err != nil {
		panic(err)
	}
	midBoss1Images[2] = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Boss2))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Boss2.png")
	if err != nil {
		panic(err)
	}
	boss2Image = ebiten.NewImageFromImage(img)
}

func disposeLevel2Enemies() {
	movingFiringEnemyImage.Dispose()
	movingFiringEnemyImage2.Dispose()
	movingFiringEnemyImage3.Dispose()
	staticEnemyImage.Dispose()
	staticExplodingEnemyImage.Dispose()
	midBoss1Images[0].Dispose()
	midBoss1Images[1].Dispose()
	midBoss1Images[2].Dispose()
	midBoss1Images[3].Dispose()
	boss2Image.Dispose()
}

func loadLevel2Background() {
	img, _, err := image.Decode(bytes.NewReader(assets.Level21))
	//img, _, err := ebitenutil.NewImageFromFile("assets/Level2-1.png")
	if err != nil {
		panic(err)
	}
	levelFirstPlan = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Level22))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Level2-2.png")
	if err != nil {
		panic(err)
	}
	levelSecondPlan = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Level23))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Level2-3.png")
	if err != nil {
		panic(err)
	}
	levelThirdPlan = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Level24))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Level2-4.png")
	if err != nil {
		panic(err)
	}
	levelFourthPlan = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(assets.Level2back))
	//img, _, err = ebitenutil.NewImageFromFile("assets/Level2-back.png")
	if err != nil {
		panic(err)
	}
	levelBackground = ebiten.NewImageFromImage(img)
}

var level2SpawnSequence []spawn = []spawn{
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 13},
		},
		frameDelay: 250,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 24,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(10*screenHeight) / 14},
		},
		frameDelay: 38,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 27,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 7},
		},
		frameDelay: 18,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(8*screenHeight) / 14},
		},
		frameDelay: 17,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 11,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(screenHeight) / 2},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 19,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(6*screenHeight) / 14},
		},
		frameDelay: 55,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(11*screenHeight) / 15},
		},
		frameDelay: 44,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(2*screenHeight) / 15},
		},
		frameDelay: 68,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 6},
		},
		frameDelay: 67,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(10*screenHeight) / 11},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 10},
		},
		frameDelay: 53,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(7*screenHeight) / 13},
		},
		frameDelay: 55,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 51,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 10},
		},
		frameDelay: 53,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 37,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(8*screenHeight) / 9},
		},
		frameDelay: 26,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(3*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(6*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 5},
		},
		frameDelay: 47,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(8*screenHeight) / 13},
		},
		frameDelay: 63,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(11*screenHeight) / 13},
		},
		frameDelay: 48,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(6*screenHeight) / 8},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 14},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 28,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 34,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 4},
		},
		frameDelay: 61,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(17*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(14*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(11*screenHeight) / 13},
		},
		frameDelay: 31,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 63,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 6},
		},
		frameDelay: 22,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 12},
		},
		frameDelay: 37,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 10},
		},
		frameDelay: 34,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(8*screenHeight) / 11},
		},
		frameDelay: 62,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(12*screenHeight) / 13},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 13},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 6},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 8},
		},
		frameDelay: 32,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(8*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(12*screenHeight) / 20},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(8*screenHeight) / 9},
		},
		frameDelay: 68,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 14},
		},
		frameDelay: 52,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 5},
		},
		frameDelay: 16,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 7},
		},
		frameDelay: 54,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(8*screenHeight) / 9},
		},
		frameDelay: 61,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 12},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(4*screenHeight) / 7},
		},
		frameDelay: 40,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 9},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 16},
		},
		frameDelay: 58,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 7},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(2*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(18*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 11},
		},
		frameDelay: 21,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 13},
		},
		frameDelay: 69,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 9},
		},
		frameDelay: 38,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(4*screenHeight) / 10},
		},
		frameDelay: 68,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 6},
		},
		frameDelay: 47,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 63,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 4},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(6*screenHeight) / 9},
		},
		frameDelay: 51,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(7*screenHeight) / 12},
		},
		frameDelay: 64,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(7*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(13*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 31,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 8},
		},
		frameDelay: 69,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 15},
		},
		frameDelay: 14,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(12*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(11*screenHeight) / 16},
		},
		frameDelay: 69,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(14*screenHeight) / 16},
		},
		frameDelay: 45,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 13},
		},
		frameDelay: 24,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(5*screenHeight) / 7},
		},
		frameDelay: 48,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 6},
		},
		frameDelay: 64,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(4*screenHeight) / 11},
		},
		frameDelay: 16,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(12*screenHeight) / 14},
		},
		frameDelay: 52,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 4},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(screenHeight) / 4},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(screenHeight) / 2},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 6},
		},
		frameDelay: 59,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 13},
		},
		frameDelay: 61,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 10},
		},
		frameDelay: 49,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 10},
		},
		frameDelay: 38,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 6},
		},
		frameDelay: 65,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(7*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(13*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 36,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 11},
		},
		frameDelay: 49,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 5},
		},
		frameDelay: 57,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: midBoss1, y: float64(screenHeight) / 5},
			enemySpawn{enemyType: midBoss1, y: float64(4*screenHeight) / 5},
		},
		frameDelay: 450,
	},
	// Part 2
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 13},
		},
		frameDelay: 250,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 24,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(10*screenHeight) / 14},
		},
		frameDelay: 38,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 27,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 7},
		},
		frameDelay: 18,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(8*screenHeight) / 14},
		},
		frameDelay: 17,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 11,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 19,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(6*screenHeight) / 14},
		},
		frameDelay: 55,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(11*screenHeight) / 15},
		},
		frameDelay: 44,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(2*screenHeight) / 15},
		},
		frameDelay: 68,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 6},
		},
		frameDelay: 67,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(10*screenHeight) / 11},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 10},
		},
		frameDelay: 53,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(7*screenHeight) / 13},
		},
		frameDelay: 55,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 51,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 10},
		},
		frameDelay: 53,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 37,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(8*screenHeight) / 9},
		},
		frameDelay: 26,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 5},
		},
		frameDelay: 47,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(8*screenHeight) / 13},
		},
		frameDelay: 63,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(11*screenHeight) / 13},
		},
		frameDelay: 48,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(6*screenHeight) / 8},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 14},
		},
		frameDelay: 35,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 25,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 28,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 34,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 4},
		},
		frameDelay: 61,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(11*screenHeight) / 13},
		},
		frameDelay: 31,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 63,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 6},
		},
		frameDelay: 22,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 12},
		},
		frameDelay: 37,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 10},
		},
		frameDelay: 34,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(8*screenHeight) / 11},
		},
		frameDelay: 62,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(12*screenHeight) / 13},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 13},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 6},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 8},
		},
		frameDelay: 32,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(8*screenHeight) / 9},
		},
		frameDelay: 68,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 14},
		},
		frameDelay: 52,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 5},
		},
		frameDelay: 16,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 7},
		},
		frameDelay: 54,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(8*screenHeight) / 9},
		},
		frameDelay: 61,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 12},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(4*screenHeight) / 7},
		},
		frameDelay: 40,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 9},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 16},
		},
		frameDelay: 58,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 7},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 11},
		},
		frameDelay: 21,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 13},
		},
		frameDelay: 69,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 9},
		},
		frameDelay: 38,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(4*screenHeight) / 10},
		},
		frameDelay: 68,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(3*screenHeight) / 6},
		},
		frameDelay: 47,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 63,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 4},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 3},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(6*screenHeight) / 9},
		},
		frameDelay: 51,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(7*screenHeight) / 12},
		},
		frameDelay: 64,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 31,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 8},
		},
		frameDelay: 69,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 15},
		},
		frameDelay: 14,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(11*screenHeight) / 16},
		},
		frameDelay: 69,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 50,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(14*screenHeight) / 16},
		},
		frameDelay: 45,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 13},
		},
		frameDelay: 24,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(5*screenHeight) / 7},
		},
		frameDelay: 48,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 6},
		},
		frameDelay: 64,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(4*screenHeight) / 11},
		},
		frameDelay: 16,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(12*screenHeight) / 14},
		},
		frameDelay: 52,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 4},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 6},
		},
		frameDelay: 59,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 13},
		},
		frameDelay: 61,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 10},
		},
		frameDelay: 49,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 10},
		},
		frameDelay: 38,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 6},
		},
		frameDelay: 65,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 36,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 11},
		},
		frameDelay: 49,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 5},
		},
		frameDelay: 57,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 6},
		},
		frameDelay: 41,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 15},
		},
		frameDelay: 46,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(5*screenHeight) / 9},
		},
		frameDelay: 46,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 59,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(7*screenHeight) / 9},
		},
		frameDelay: 59,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 6},
		},
		frameDelay: 32,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 23,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 55,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 11},
		},
		frameDelay: 46,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 12},
		},
		frameDelay: 27,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(4*screenHeight) / 5},
		},
		frameDelay: 43,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 5},
		},
		frameDelay: 15,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 6},
		},
		frameDelay: 13,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 5},
		},
		frameDelay: 23,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 4},
		},
		frameDelay: 18,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(7*screenHeight) / 13},
		},
		frameDelay: 36,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 13},
		},
		frameDelay: 67,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 6},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 4},
		},
		frameDelay: 14,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 5},
		},
		frameDelay: 26,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(10*screenHeight) / 14},
		},
		frameDelay: 64,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(5*screenHeight) / 11},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(8*screenHeight) / 16},
		},
		frameDelay: 14,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(6*screenHeight) / 10},
		},
		frameDelay: 45,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(15*screenHeight) / 16},
		},
		frameDelay: 55,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(8*screenHeight) / 13},
		},
		frameDelay: 41,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(4*screenHeight) / 11},
		},
		frameDelay: 45,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 9},
		},
		frameDelay: 67,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(7*screenHeight) / 10},
		},
		frameDelay: 29,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(6*screenHeight) / 7},
		},
		frameDelay: 16,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(4*screenHeight) / 11},
		},
		frameDelay: 64,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 10},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(3*screenHeight) / 7},
		},
		frameDelay: 13,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(8*screenHeight) / 16},
		},
		frameDelay: 55,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: movingFiringEnemy, y: float64(4*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(10*screenHeight) / 20},
			enemySpawn{enemyType: movingFiringEnemy, y: float64(16*screenHeight) / 20},
		},
		frameDelay: 10,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 10},
		},
		frameDelay: 64,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(2*screenHeight) / 11},
		},
		frameDelay: 57,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(1*screenHeight) / 2},
		},
		frameDelay: 17,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticEnemy, y: float64(9*screenHeight) / 11},
		},
		frameDelay: 30,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(7*screenHeight) / 8},
		},
		frameDelay: 43,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: staticExplodingEnemy, y: float64(2*screenHeight) / 3},
		},
		frameDelay: 20,
	},
	spawn{
		enemies: []enemySpawn{
			enemySpawn{enemyType: boss2, y: float64(screenHeight) / 2},
		},
		frameDelay: 450,
	},
}
