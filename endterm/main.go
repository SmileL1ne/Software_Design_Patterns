package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Game struct {
	factory   CharacterFactory
	notifier  EventNotifier
	character *Character
}

func NewGame() *Game {
	return &Game{
		factory:   CharacterFactory{},
		notifier:  EventNotifier{},
		character: nil,
	}
}

func (g *Game) CreateCharacter() {
	fmt.Println("Choose your character (type the number):\n1. Warrior\n2. Mage\n3. Archer")
	reader := bufio.NewReader(os.Stdin)
	charType, _ := reader.ReadString('\n')

	var characterType CharacterType
	var chType string
	switch strings.TrimSpace(charType) {
	case "1":
		characterType = Warrior
		chType = "Warrior"
	case "2":
		characterType = Mage
		chType = "Mage"
	case "3":
		characterType = Archer
		chType = "Archer"
	default:
		log.Fatalln("Invalid choice. Try again.")
	}

	// Ask for attack method
	fmt.Println("\nChoose your attack method (type the number):\n1. Sword Slash\n2. Sign Magic\n3. Mighty Bow")
	attackChoice, _ := reader.ReadString('\n')

	var attackMethod AttackMethod
	switch strings.TrimSpace(attackChoice) {
	case "1":
		attackMethod = new(SwordSlash)
	case "2":
		attackMethod = new(SignMagic)
	case "3":
		attackMethod = new(MightyBow)
	default:
		log.Fatalln("Invalid choice. Try again.")
	}

	fmt.Println("\nEnter your character's name:")
	input, _ := reader.ReadString('\n')
	name := strings.TrimSpace(input)

	char := g.factory.CreateCharacter(characterType, name, attackMethod)

	// Ask for additional state
	fmt.Println("\nChoose an additional ability (type the number):\n1. Enchanted Armor\n2. Fire Sword\n3. No additional ability")
	ability, _ := reader.ReadString('\n')

	var addedState string
	switch strings.TrimSpace(ability) {
	case "1":
		addedState = "with Enchanted Armor"
	case "2":
		addedState = "wielding a Fire Sword"
	case "3":
		addedState = ""
	default:
		log.Fatalln("Invalid choice. Try again.")
	}

	if addedState != "" {
		char = NewCharacterDecorator(char, addedState).character
	}
	g.notifier.RegisterObserver(char)

	g.character = char

	fmt.Printf("\nCharacter %s the %s has been created!\n", g.character.Name, chType)
	time.Sleep(2 * time.Second)
	if addedState != "" {
		fmt.Printf("\nAnd character has additionaly ability - %s\n", addedState)
	} else {
		fmt.Printf("\nCharacter has no additional ability or items\n")
	}
	time.Sleep(2 * time.Second)
}

func (g *Game) RunEventLoop() {
	events := []string{
		"A dragon has been spotted in the skies!",
		"A mysterious stranger offers a powerful sword.",
		"An earthquake reveals a hidden cavern.",
	}

	for _, eventDescription := range events {
		event := Event{Desciption: eventDescription}
		fmt.Println("\nEVENT:")
		g.notifier.Notify(event)
		fmt.Println()
		time.Sleep(2 * time.Second)

		fmt.Println("\nType 'attack' to attack, 'wait' to do nothing.")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		fmt.Println()

		if strings.TrimSpace(input) == "attack" {
			g.character.PerformAttack()
		} else {
			fmt.Println(g.character.Name, "waits for the next opportunity.")
		}
		time.Sleep(2 * time.Second)
	}
	fmt.Println("\nGame is finished! Thank you!")
}

func main() {
	game := NewGame()
	game.CreateCharacter()
	game.RunEventLoop()
}
