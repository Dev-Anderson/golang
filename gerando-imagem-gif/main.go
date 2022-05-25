package main 

import (
	"github.com/llgcode/draw2d/draw2dimg"
    "github.com/llgcode/draw2d/draw2dkit"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math"
	"math/rand"
	"os"
)

var w, h float64 = 500, 250
var palette color.Palette = color.Palette{}
var zCycle float64 = 8
var zMin, zMax float64 = 1, 15

type Point struct {
	X, Y float64
}

type Circle struct {
	X, Y, Z, R float64
}

func (c *Circle) Draw(gc * draw2dimg.GraphicContext, ratio float64) {
	z := c.Z - ratio*zCycle

	for z < zMax {
		if z >= zMin {
			x, y, r := c.X/z, c.Y/z, c.R/z
			gc.SetFillColor(color.White)
			gc.Fill() 
			draw2dkit.Circle(gc, w/2+x, h/2+y, r)
			gc.Close() 
		}
		z += zCycle
	}
}

func drawFrame(circles []Circle, ratio float64) *image.Paletted {
	img := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	gc := draw2dimg.NewGraphicContext(img)

	gc.SetFillColor(color.Gray{0x11})
	draw2dkit.Rectangle(gc, 0, 0, w, h)
	gc.Fill() 
	gc.Close() 

	for _, circle := range circles {
		circle.Draw(gc, ratio)
	}

	pm := image.NewPaletted(img.Bounds(), palette)
	draw.FloydSteinberg.Draw(pm, img.Bounds(), img, image.ZP)
	return pm 
}

func main() { 
	circles := []Circle{}
    for len(circles) < 4000 {
        x, y := rand.Float64()*8-4, rand.Float64()*8-4
        if math.Abs(x) < 0.5 && math.Abs(y) < 0.5 {
            continue
        }
        z := rand.Float64() * zCycle
        circles = append(circles, Circle{x * w, y * h, z, 5})
    }

	palette = color.Palette{}
	for i := 0; i < 16; i++ {
		palette  = append(palette, color.Gray{uint8(i) * 0x11})
	}

	var images []*image.Paletted
	var delays []int
	count := 30 
	for i := 0; i < count; i++ {
		pm := drawFrame(circles, float64(i)/float64(count))
		images = append(images, pm)
		delays = append(delays, 4)
	}

	f, _ := os.OpenFile("space.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close() 
	gif.EncodeAll(f, &gif.GIF {
		Image: images, 
		Delay: delays,
	})
}