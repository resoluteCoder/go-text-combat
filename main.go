package main

import (
	"fmt"
	"math"
)

type Player struct {
	Living
}

type Enemy struct {
	Living
}

type Living struct {
	Name           string
	Health         int
	PhysicalAttack int
	PhysicalArmor  int
}

type Combatant interface {
	TakeDamage(int)
}

func (living *Living) TakeDamage(damage int) {
	totalDamage := living.PhysicalArmor - damage
	living.Health -= int(math.Abs(float64(totalDamage)))
}

func (living Living) DisplayStats() {
	fmt.Printf("%s is at %d health and does %d damage.\n", living.Name, living.Health, living.PhysicalAttack)
}

func NewPlayer(name string) *Player {
	playerStats := Living{
		Name:           name,
		Health:         100,
		PhysicalAttack: 10,
		PhysicalArmor:  5,
	}
	player := Player{
		playerStats,
	}
	return &player
}

func NewEnemy(name string) *Enemy {
	enemyStats := Living{
		Name:           name,
		Health:         100,
		PhysicalAttack: 5,
		PhysicalArmor:  1,
	}
	enemy := Enemy{
		enemyStats,
	}
	return &enemy
}

func main() {
	player := NewPlayer("Chris")
	enemy := NewEnemy("Orc")

	combatants := []Combatant{player, enemy}

	for _, combatant := range combatants {
		fmt.Println(combatant)
	}
}
