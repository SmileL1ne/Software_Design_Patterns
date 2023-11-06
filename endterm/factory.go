package main

import "log"

type CharacterType int

const (
	Warrior CharacterType = iota
	Mage
)

type CharacterFactory struct{}

func (f *CharacterFactory) CreateCharacter(t CharacterType, name string) *Character {
	switch t {
	case Warrior:
		return &Character{name, new(SwordSlash)}
	case Mage:
		return &Character{name, new(SignMagic)}
	default:
		log.Fatalln("Invalid character type. Try again")
	}
	return new(Character)
}
