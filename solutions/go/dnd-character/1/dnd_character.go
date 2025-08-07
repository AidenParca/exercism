package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// Character represents a D&D character with six abilities and hitpoints.
type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// init is used to seed the random number generator to ensure different results on each run.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}
func Ability() int {
	rolls := make([]int, 4)
	for i := 0; i < 4; i++ {
		rolls[i] = rand.Intn(6) + 1
	}
	sort.Ints(rolls)
	sum := 0
	for i := 1; i < 4; i++ {
		sum += rolls[i]
	}
	return sum
}
func GenerateCharacter() Character {
	char := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	// Initial hitpoints are 10 + the character's constitution modifier.
	char.Hitpoints = 10 + Modifier(char.Constitution)

	return char
}
