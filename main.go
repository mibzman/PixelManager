package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Sprite struct {
	SkinColor  color.RGBA
	BeardColor color.RGBA
	EyeColor   color.RGBA
	HairColor  color.RGBA
	Image      *image.RGBA
}

func (Sprite Sprite) Draw() *image.RGBA {
	var width, height int = 12, 18
	Sprite.Image = image.NewRGBA(image.Rect(0, 0, width, height))

	Sprite.drawHead()
	// Sprite.drawBeard()
	// Sprite.drawEye()
	// Sprite.drawHair()

	return Sprite.Image
}

type Pixel struct {
	X     int
	Y     int
	Color color.RGBA
}

func (Sprite Sprite) drawHead() {
	// Pixels := []Pixel

	skin := make(map[int]([]int))

	skin[3] = []int{5, 6, 7, 8, 9}
	skin[4] = []int{3, 4, 5, 6, 7, 8, 9}
	skin[5] = []int{1, 2, 5, 6, 7}
	skin[6] = []int{1, 2, 5, 6, 7}
	skin[7] = []int{2, 3, 4, 5, 6, 7, 8, 9}
	skin[8] = []int{2, 3, 4, 5, 8, 9}
	skin[9] = []int{2, 3, 4, 5, 6, 7, 8, 9}
	skin[10] = []int{2, 3, 9}
	skin[11] = []int{2, 3, 9}
	skin[12] = []int{2, 3, 9}
	skin[13] = []int{2, 3, 4, 8}
	skin[14] = []int{2, 3, 4}
	skin[15] = []int{2, 3}
	skin[16] = []int{2, 3}
	skin[17] = []int{2, 3}

	ColorMap(skin, Sprite.SkinColor, Sprite.Image.Set)

	Head := make(map[int]([]int))

	Head[4] = []int{10}
	Head[5] = []int{10}
	Head[6] = []int{0, 10}
	Head[7] = []int{0, 1, 10}
	Head[8] = []int{1, 6, 7, 10}
	Head[9] = []int{1, 10}
	Head[10] = []int{1, 10}
	Head[11] = []int{1, 5, 6, 7, 10}
	Head[12] = []int{1, 10}
	Head[13] = []int{1, 9}
	Head[14] = []int{1, 8}
	Head[15] = []int{1, 4, 5, 6, 7}
	Head[16] = []int{1, 4}
	Head[17] = []int{1, 4}

	black := color.RGBA{
		0,
		0,
		0,
		255,
	}

	ColorMap(Head, black, Sprite.Image.Set)
}

func ColorMap(points map[int]([]int), color color.RGBA, fun func(int, int, color.Color)) {
	for y, row := range points {
		for _, x := range row {
			fun(x, y, color)
		}
	}
}

func main() {
	// var width, height int = 12, 18

	// outputImage := image.NewRGBA(image.Rect(0, 0, width, height))
	// for x := 0; x < width; x++ {
	// 	for y := 0; y < height; y++ {
	// 		c := color.RGBA{
	// 			255,
	// 			255,
	// 			255,
	// 			255,
	// 		}
	// 		outputImage.Set(x, y, c)
	// 	}
	// }

	// c := color.RGBA{
	// 	0,
	// 	0,
	// 	0,
	// 	255,
	// }

	// outputImage.Set(0, 0, c)

	Skin := color.RGBA{
		255,
		200,
		110,
		255,
	}

	MySprite := Sprite{SkinColor: Skin}

	outputImage := MySprite.Draw()

	f, err := os.OpenFile("rgb.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	png.Encode(f, outputImage)
}
