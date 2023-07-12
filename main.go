package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Player struct {
	Living
	// Coins           int
	// CurrentExperience int
	// MaxExperience     int
	// Level             int
	// Items             []string
}

type Enemy struct {
	Living
	// ExperienceGiven int
	// ItemsDropped    []string
}

type Combatant interface {
	TakeDamage(int)
	DisplayStatus()
}

type Living struct {
	EntityId       int
	Name           string
	Health         int
	PhysicalAttack int
	PhysicalArmor  int
}

func (living *Living) TakeDamage(damage int) {
	totalDamage := living.PhysicalArmor - damage
	living.Health -= int(math.Abs(float64(totalDamage)))
}

func (living *Living) DisplayStatus() {
	fmt.Printf("%s is at %d health.\n", living.Name, living.Health)
}

func NewPlayer(name string) *Player {
	playerStats := Living{
		EntityId:       0,
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
		EntityId:       1,
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
	giant := NewEnemy("Giant")
	skeleton := NewEnemy("Skeleton")

	combatants := []Combatant{player, giant, skeleton}

	for _, combatant := range combatants {
		switch combatant.(type) {
		case *Player:
			combatant.DisplayStatus()
			fmt.Println("Your turn")
			fmt.Println("1) Attack")
			fmt.Print("> ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}

			switch strings.TrimSuffix(input, "\n") {
			case "1":
				fmt.Println("Which target?")
				for index, combatant := range combatants {
					switch combatant.(type) {
					case *Enemy:
						fmt.Printf("%d) %s\n", index, combatant.(*Enemy).Name)
					}
				}
				fmt.Print("> ")
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println(err)
				}
				fmt.Print(input)
			}
		case *Enemy:
			fmt.Printf("%s took their turn\n", combatant.(*Enemy).Name)
		}
	}

}
