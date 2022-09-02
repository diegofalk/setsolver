package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/color"
	"setsolver/internal/alligner"
	"setsolver/internal/cardfinder"
	"setsolver/internal/preprocess"
	"setsolver/internal/solver"
	"setsolver/models"
)

func main() {
	fmt.Println("Hola solver")

	window := gocv.NewWindow("Hello")
	img := gocv.IMRead("setimages/5.jpg", gocv.IMReadColor)
	imgopen := gocv.NewMat()

	preprocess.Run(img, &imgopen)
	cardContours := cardfinder.Run(imgopen)

	imgcontours := img.Clone()

	ptsv := gocv.NewPointsVector()
	for _, c := range cardContours {
		ptsv.Append(c.Contour)
		tr := alligner.Run(img, c)

		window.IMShow(tr)
		window.WaitKey(0)
	}

	gocv.DrawContours(&imgcontours, ptsv, -1, color.RGBA{0, 255, 0, 255}, 3)

	window.IMShow(img)
	window.WaitKey(0)
	window.IMShow(imgopen)
	window.WaitKey(0)
	window.IMShow(imgcontours)
	window.WaitKey(0)
	window.Close()
	return

	cards := []models.Card{
		{
			Shape:   models.Round,
			Color:   models.Purple,
			Filling: models.Texture,
			Count:   2,
		},
		{
			Shape:   models.Round,
			Color:   models.Purple,
			Filling: models.Texture,
			Count:   1,
		},
		{
			Shape:   models.Snake,
			Color:   models.Purple,
			Filling: models.Outline,
			Count:   2,
		},
		{
			Shape:   models.Snake,
			Color:   models.Purple,
			Filling: models.Outline,
			Count:   3,
		},
		{
			Shape:   models.Round,
			Color:   models.Green,
			Filling: models.Outline,
			Count:   1,
		},
		{
			Shape:   models.Round,
			Color:   models.Red,
			Filling: models.Outline,
			Count:   1,
		},
		{
			Shape:   models.Diamond,
			Color:   models.Green,
			Filling: models.Filled,
			Count:   1,
		},
		{
			Shape:   models.Diamond,
			Color:   models.Green,
			Filling: models.Texture,
			Count:   1,
		},
		{
			Shape:   models.Diamond,
			Color:   models.Purple,
			Filling: models.Outline,
			Count:   3,
		},
		{
			Shape:   models.Diamond,
			Color:   models.Purple,
			Filling: models.Texture,
			Count:   2,
		},
		{
			Shape:   models.Round,
			Color:   models.Green,
			Filling: models.Outline,
			Count:   3,
		},
		{
			Shape:   models.Round,
			Color:   models.Red,
			Filling: models.Texture,
			Count:   1,
		},
	}
	sets, index := solver.Solve(cards)
	fmt.Printf("Sets: %+v\n", sets)
	fmt.Printf("Indexes: %+v\n", index)
}
