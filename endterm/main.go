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
	fmt.Println("Choose your character (type the number):\n1. Warrior\n2. Mage")
	reader := bufio.NewReader(os.Stdin)
	charType, _ := reader.ReadString('\n')

	var characterType CharacterType
	switch strings.TrimSpace(charType) {
	case "1":
		characterType = Warrior
		charType = "Warrior"
	case "2":
		characterType = Mage
		charType = "Mage"
	default:
		log.Fatalln("Invalid choice. Try again.")
	}

	fmt.Println("\nEnter your character's name:")
	input, _ := reader.ReadString('\n')
	name := strings.TrimSpace(input)

	g.character = g.factory.CreateCharacter(characterType, name)
	g.notifier.RegisterObserver(g.character)
	fmt.Printf("\nCharacter %s the %s has been created!\n", g.character.Name, charType)
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
