package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Sprite struct {
	Name   string
	Gender int

	SkinColor  color.RGBA
	BeardColor color.RGBA
	EyeColor   color.RGBA
	HairColor  color.RGBA

	Image *image.RGBA
}

func (Sprite Sprite) Generate() error {
	Image := Sprite.Draw()
	return Print(Sprite.Name, Image)
}

func (Sprite Sprite) Draw() *image.RGBA {
	var width, height int = 12, 18
	Sprite.Image = image.NewRGBA(image.Rect(0, 0, width, height))

	Sprite.drawHead()
	Sprite.drawBeard()
	Sprite.drawHair()
	Sprite.drawEyes()

	return Sprite.Image
}

func Print(Name string, Image *image.RGBA) error {
	f, err := os.OpenFile(Name+".png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, Image)
	return err
}

func (Sprite Sprite) drawEyes() {
	pupils := make(map[int]([]int))
	pupils[6] = []int{3, 8}

	black := color.RGBA{0, 0, 0, 255}

	ColorMap(pupils, black, Sprite.Image.Set)

	iris1 := make(map[int]([]int))
	iris1[5] = []int{3, 4, 8, 9}

	ColorMap(iris1, Sprite.EyeColor, Sprite.Image.Set)

	iris2 := make(map[int]([]int))
	iris2[6] = []int{4, 9}

	iris2Color := color.RGBA{
		Sprite.EyeColor.R + 15,
		Sprite.EyeColor.G + 15,
		Sprite.EyeColor.B + 15,
		255,
	}

	ColorMap(iris2, iris2Color, Sprite.Image.Set)
}

func (Sprite Sprite) drawBeard() {
	hair := make(map[int]([]int))

	hair[10] = []int{4, 5, 6, 7, 8}
	hair[11] = []int{4, 8}
	hair[12] = []int{4, 5, 6, 7, 8}
	hair[13] = []int{5, 6, 7}
	hair[14] = []int{5, 6, 7}

	ColorMap(hair, Sprite.BeardColor, Sprite.Image.Set)
}

func (Sprite Sprite) drawHair() {
	hair := make(map[int]([]int))

	hair[0] = []int{3, 4, 5, 6, 7, 8}
	hair[1] = []int{2, 3, 4, 5, 6, 7, 8, 9}
	hair[2] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	hair[3] = []int{0, 1, 2, 3, 4, 10, 11}
	hair[4] = []int{0, 1, 2, 11}
	hair[5] = []int{0, 11}

	if Sprite.Gender == 1 {
		hair[6] = []int{11}
		hair[7] = []int{11}
		hair[8] = []int{0, 11}
		hair[9] = []int{0, 11}
		hair[10] = []int{0, 11}
		hair[11] = []int{0, 11}
		hair[12] = []int{0, 11}
		hair[13] = []int{0, 10, 11}
		hair[14] = []int{0, 9, 10, 11}
		hair[15] = []int{0, 8, 9, 10, 11}
		hair[16] = []int{0, 5, 6, 7, 8, 9, 10, 11}
	}

	ColorMap(hair, Sprite.HairColor, Sprite.Image.Set)
}

func (Sprite Sprite) drawHead() {
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

	Skin := color.RGBA{255, 200, 110, 255}
	Hair := color.RGBA{144, 101, 60, 255}
	Eye := color.RGBA{226, 153, 38, 255}

	MySprite := Sprite{
		Name:       "bob",
		SkinColor:  Skin,
		HairColor:  Hair,
		BeardColor: Hair,
		EyeColor:   Eye,
		Gender:     1,
	}

	MySprite.Generate()
}
