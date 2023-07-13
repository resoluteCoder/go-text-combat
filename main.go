package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
	takeDamage(int)
	displayStatus()
}

type Living struct {
	entityId       int
	name           string
	health         int
	physicalAttack int
	physicalArmor  int
}

func (living *Living) TakeDamage(damage int) {
	totalDamage := living.physicalArmor - damage
	living.health -= int(math.Abs(float64(totalDamage)))
}

func (living *Living) DisplayStatus() {
	fmt.Printf("%s is at %d health.\n", living.name, living.health)
}

func NewPlayer(name string) *Player {
	playerStats := Living{
		entityId:       0,
		name:           name,
		health:         100,
		physicalAttack: 10,
		physicalArmor:  5,
	}
	player := Player{
		playerStats,
	}
	return &player
}

func NewEnemy(name string) *Enemy {
	enemyStats := Living{
		entityId:       1,
		name:           name,
		health:         100,
		physicalAttack: 5,
		physicalArmor:  1,
	}
	enemy := Enemy{
		enemyStats,
	}
	return &enemy
}

type CombatManager struct {
	enemies          []*Enemy
	player           Player
	isCombatFinished bool
	isPlayerTurn     bool
	currentIndex     int
}

func (combatManager CombatManager) printPlayerMenu() {
	playerMenu := []string{"Your turn.", "1) Attack"}
	for _, text := range playerMenu {
		fmt.Println(text)
	}
}

func (combatManager CombatManager) printAttackMenu() {
	fmt.Println("Which target?")
	for index, enemy := range combatManager.enemies {
		fmt.Printf("%d) %s\n", index+1, enemy.name)
	}
}

func (combatManager CombatManager) getInputFromPlayer() (string, error) {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return strings.TrimSuffix(input, "\n"), nil
}

func (combatManager *CombatManager) handlePlayerAttack(targetIndex string) {
	index, _ := strconv.ParseInt(targetIndex, 10, 16)
	enemy := combatManager.enemies[index-1]
	enemy.TakeDamage(combatManager.player.physicalAttack)
	enemy.DisplayStatus()
	combatManager.isPlayerTurn = false
}

func (combatManager *CombatManager) handleEnemyAttack() {
	for _, enemy := range combatManager.enemies {
		fmt.Println(enemy.name)
	}
	combatManager.isPlayerTurn = true
}

func main() {
	player := NewPlayer("Chris")
	giant := NewEnemy("Giant")
	skeleton := NewEnemy("Skeleton")

	enemies := []*Enemy{giant, skeleton}

	combatManager := CombatManager{
		enemies:      enemies,
		player:       *player,
		isPlayerTurn: true,
		currentIndex: 0,
	}

	for len(combatManager.enemies) != 0 {
		if combatManager.isPlayerTurn {
			combatManager.printPlayerMenu()
			input, _ := combatManager.getInputFromPlayer()
			if input == "1" {
				combatManager.printAttackMenu()
				input, _ := combatManager.getInputFromPlayer()
				combatManager.handlePlayerAttack(input)
			}
		} else {
			combatManager.handleEnemyAttack()
		}
	}
}
