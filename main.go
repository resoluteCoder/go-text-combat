package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

// type Combatant interface {
// 	takeDamage(int)
// 	displayStatus()
// }

type Living struct {
	entityId       int
	name           string
	health         int
	physicalAttack int
	physicalArmor  int
}

func (living *Living) takeDamage(damage int) {
	totalDamage := living.physicalArmor - damage
	living.health -= int(math.Abs(float64(totalDamage)))
}

func (living *Living) displayStatus() {
	fmt.Printf("%s is at %d health.\n", living.name, living.health)
}

func NewPlayer(name string) *Player {
	playerStats := Living{
		entityId:       0,
		name:           name,
		health:         50,
		physicalAttack: 25,
		physicalArmor:  10,
	}
	player := Player{
		playerStats,
	}
	return &player
}

func NewEnemy(name string) *Enemy {
	enemyStats := Living{
		entityId:       rand.Int() + 1,
		name:           name,
		health:         25,
		physicalAttack: 15,
		physicalArmor:  5,
	}
	enemy := Enemy{
		enemyStats,
	}
	return &enemy
}

type CombatManager struct {
	enemies          []*Enemy
	player           *Player
	isCombatFinished bool
	isPlayerTurn     bool
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
		fmt.Printf("%d) %s - %d HP\n", index+1, enemy.name, enemy.health)
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
	enemy.takeDamage(combatManager.player.physicalAttack)
	fmt.Printf("%s attacks %s for %d damage\n", combatManager.player.name, enemy.name, enemy.physicalAttack)

	if enemy.health <= 0 {
		fmt.Printf("%s has been defeated!\n", enemy.name)
		combatManager.enemies = combatManager.removeDefeatedEnemy(enemy)
	} else {
		enemy.displayStatus()
	}

	combatManager.isPlayerTurn = false
}

func (combatManager *CombatManager) handleEnemyAttack() {
	for _, enemy := range combatManager.enemies {
		time.Sleep(1 * time.Second)
		combatManager.player.takeDamage(enemy.physicalAttack)

		fmt.Printf("%s attacks %s for %d damage\n", enemy.name, combatManager.player.name, enemy.physicalAttack)
		combatManager.player.displayStatus()
	}
	combatManager.isPlayerTurn = true
}

func (combatManager *CombatManager) removeDefeatedEnemy(enemy *Enemy) []*Enemy {
	remainingEnemies := []*Enemy{}
	for _, e := range combatManager.enemies {
		if enemy.entityId != e.entityId {
			remainingEnemies = append(remainingEnemies, e)
		}
	}
	return remainingEnemies
}

func main() {
	player := NewPlayer("Chris")
	giant := NewEnemy("Giant")
	skeleton := NewEnemy("Skeleton")

	enemies := []*Enemy{giant, skeleton}

	combatManager := CombatManager{
		enemies:      enemies,
		player:       player,
		isPlayerTurn: true,
	}

	for len(combatManager.enemies) != 0 || combatManager.player.health <= 0 {
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
	fmt.Println("Combat is finished!")
}
