package main

import (
	"fmt"
	"setsolver/internal/solver"
	"setsolver/models"
)

func main() {
	fmt.Println("Hola solver")
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
