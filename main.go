package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Sprite struct {
	Name string

	SkinColor color.RGBA
	EyeColor  color.RGBA

	HasBeard   bool
	BeardColor color.RGBA

	HairColor  color.RGBA
	HairLength int

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

	if Sprite.HasBeard {
		ColorMap(hair, Sprite.BeardColor, Sprite.Image.Set)
	} else {
		ColorMap(hair, Sprite.SkinColor, Sprite.Image.Set)
	}
}

func (Sprite Sprite) drawHair() {
	hair := make(map[int]([]int))

	hair[0] = []int{3, 4, 5, 6, 7, 8}
	hair[1] = []int{2, 3, 4, 5, 6, 7, 8, 9}
	hair[2] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	hair[3] = []int{0, 1, 2, 3, 4, 10, 11}
	hair[4] = []int{0, 1, 2, 11}
	hair[5] = []int{0, 11}

	if Sprite.HairLength >= 1 {
		hair[6] = []int{11}
		hair[7] = []int{11}
		hair[8] = []int{0, 11}
	}

	if Sprite.HairLength >= 2 {
		hair[9] = []int{0, 11}
		hair[10] = []int{0, 11}
		hair[11] = []int{0, 11}
		hair[12] = []int{0, 11}
	}

	if Sprite.HairLength >= 3 {
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

	Lipp := make(map[int]([]int))
	Lipp[11] = []int{5, 6, 7}

	lippColor := color.RGBA{129, 57, 53, 255}

	ColorMap(Lipp, lippColor, Sprite.Image.Set)
}

func ColorMap(points map[int]([]int), color color.RGBA, fun func(int, int, color.Color)) {
	for y, row := range points {
		for _, x := range row {
			fun(x, y, color)
		}
	}
}

func main() {
	Skins := []color.RGBA{
		color.RGBA{141, 85, 36, 255},
		color.RGBA{198, 134, 66, 255},
		color.RGBA{224, 172, 105, 255},
		color.RGBA{241, 194, 125, 255},
		color.RGBA{255, 219, 172, 255},
	}

	Hairs := []color.RGBA{
		color.RGBA{9, 8, 6, 255},
		color.RGBA{113, 99, 90, 255},
		color.RGBA{183, 166, 158, 255},
		color.RGBA{214, 196, 194, 255},
		color.RGBA{202, 191, 177, 255},
		color.RGBA{220, 208, 186, 255},
		color.RGBA{255, 245, 225, 255},
		color.RGBA{222, 188, 153, 255},
		color.RGBA{184, 151, 120, 255},
		color.RGBA{165, 107, 70, 255},
		color.RGBA{181, 82, 57, 255},
		color.RGBA{141, 74, 67, 255},
		color.RGBA{145, 85, 61, 255},
		color.RGBA{83, 61, 50, 255},
		color.RGBA{59, 48, 36, 255},
		color.RGBA{85, 72, 56, 255},
		color.RGBA{78, 67, 63, 255},
		color.RGBA{106, 78, 66, 255},
		color.RGBA{167, 133, 106, 255},
		color.RGBA{151, 121, 97, 255},
	}

	Eyes := []color.RGBA{
		color.RGBA{94, 72, 30, 255},
		color.RGBA{84, 42, 14, 255},
		color.RGBA{99, 57, 15, 255},
		color.RGBA{96, 49, 1, 255},
		color.RGBA{69, 24, 0, 255},
		color.RGBA{75, 114, 72, 255},
		color.RGBA{51, 122, 44, 255},
		color.RGBA{25, 163, 55, 255},
		color.RGBA{155, 240, 157, 255},
		color.RGBA{238, 232, 170, 255},
		color.RGBA{203, 206, 134, 255},
		color.RGBA{168, 180, 97, 255},
		color.RGBA{134, 153, 61, 255},
		color.RGBA{99, 127, 25, 255},
		color.RGBA{74, 108, 110, 255},
		color.RGBA{71, 98, 105, 255},
		color.RGBA{65, 97, 86, 255},
		color.RGBA{67, 101, 128, 255},
		color.RGBA{44, 76, 99, 255},
	}

	// BeardHaving := []bool{true, false, false}

	// Count := 0

	// for _, Gender := range Genders {
	// }

	MySprite := Sprite{
		Name:       "bob",
		SkinColor:  Skins[2],
		HairColor:  Hairs[2],
		BeardColor: Hairs[0],
		EyeColor:   Eyes[4],
		HasBeard:   true,
		HairLength: 2,
	}

	MySprite.Generate()
}
