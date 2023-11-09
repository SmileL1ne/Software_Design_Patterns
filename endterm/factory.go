package main

type CharacterType int

const (
	Warrior CharacterType = iota
	Mage
	Archer
)

type CharacterFactory struct{}

func (f *CharacterFactory) CreateCharacter(t CharacterType, name string, attack AttackMethod) *Character {
	return &Character{name, attack}
}
