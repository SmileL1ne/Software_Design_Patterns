package main

import "fmt"

type AttackMethod interface {
	Attack()
}

type SwordSlash struct{}

func (s *SwordSlash) Attack() {
	fmt.Println("Performing sword slash! *Whooooosh*")
}

type SignMagic struct{}

func (f *SignMagic) Attack() {
	fmt.Println("Performing sign magic, Aard! *Pheeeeeew")
}

type Character struct {
	Name   string
	attack AttackMethod
}

func (c *Character) SetAttackMethod(method AttackMethod) {
	c.attack = method
}

func (c *Character) PerformAttack() {
	fmt.Printf("%s ", c.Name)
	c.attack.Attack()
}
