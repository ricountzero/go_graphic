package main

import (
    "image/color"
    "math"

    ebiten "github.com/hajimehoshi/ebiten/v2"
    ebitenutil "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
    width  = 800
    height = 600
    radius = 100
)

type Game struct{}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.RGBA{255, 255, 255, 255})

    for x := 0; x <= radius*2; x++ {
        for y := 0; y <= radius*2; y++ {
            dx := float64(x - radius)
            dy := float64(y - radius)
            distanceSquared := dx*dx + dy*dy

            if distanceSquared <= float64(radius*radius) {
                intensity := 1.0 - (math.Sqrt(distanceSquared) / float64(radius))
                if intensity < 0 {
                    intensity = 0
                }

                shade := uint8(255 * intensity)
                c := color.RGBA{shade, 0, 0, 255}
                screen.Set(width/2+x-radius, height/2+y-radius, c)
            }
        }
    }

    shadowY := height / 2 + radius/2
    ebitenutil.DrawCircle(screen, float64(width/2), float64(shadowY), float64(radius), color.RGBA{0, 0, 0, 128})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return width, height
}

func main() {
    game := &Game{}
    ebiten.SetWindowSize(width, height)
    ebiten.SetWindowTitle("3D Sphere with Shadows in Ebiten")

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
}

