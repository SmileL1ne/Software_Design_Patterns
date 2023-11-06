package main

import "fmt"

type CharacterDecorator struct {
	character  *Character
	addedState string
}

func (c *CharacterDecorator) PerformAttack() {
	c.character.PerformAttack()
	fmt.Println(c.addedState)
}

func NewCharacterDecorator(character *Character, addedState string) *CharacterDecorator {
	return &CharacterDecorator{
		character:  character,
		addedState: addedState,
	}
}
