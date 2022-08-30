package models

type Shape string

const (
	Round   Shape = "round"
	Diamond Shape = "diamond"
	Snake   Shape = "snake"
)

type Filling string

const (
	Outline Filling = "outline"
	Filled  Filling = "filled"
	Texture Filling = "texture"
)

type Color string

const (
	Green  Color = "green"
	Purple Color = "purple"
	Red    Color = "red"
)

type Card struct {
	Shape   Shape
	Color   Color
	Filling Filling
	Count   int
}
