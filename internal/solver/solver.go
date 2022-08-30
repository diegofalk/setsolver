package solver

import "setsolver/models"

func Solve(cards []models.Card) (sets [][]models.Card, setsIndex [][]int) {
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			for k := j + 1; k < len(cards); k++ {
				set := []models.Card{cards[i], cards[j], cards[k]}
				if checkSet(set) {
					sets = append(sets, set)
					setsIndex = append(setsIndex, []int{i, j, k})
				}
			}
		}
	}
	return
}

func checkSet(cards []models.Card) bool {
	colorMap := map[models.Color]int{}
	shapeMap := map[models.Shape]int{}
	fillingMap := map[models.Filling]int{}
	countMap := map[int]int{}
	for i, c := range cards {
		colorMap[c.Color] = i
		shapeMap[c.Shape] = i
		fillingMap[c.Filling] = i
		countMap[c.Count] = i
	}
	if len(colorMap) == 2 {
		return false
	}
	if len(shapeMap) == 2 {
		return false
	}
	if len(fillingMap) == 2 {
		return false
	}
	if len(countMap) == 2 {
		return false
	}
	return true
}
