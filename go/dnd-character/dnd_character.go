package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	res := math.Floor(float64(score-10) / float64(2))
	return int(res)
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rand := rand.Intn(6-1) + 1

	nums := []int{}
	for i := 0; i < 4; i++ {
		nums = append(nums, rand)
	}

	min := nums[0]
	idx := 0
	for i, val := range nums {
		if min > val {
			min = val
			idx = i
		}
	}

	res := 0
	for i, val := range nums {
		if i == idx {
			continue
		}
		res += val
	}

	return res
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	dd := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}

	dd.Hitpoints = 10 + Modifier(dd.Constitution)
	return dd
}
