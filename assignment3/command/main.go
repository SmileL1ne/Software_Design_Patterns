package main

import "fmt"

type Switch struct {
	command Command
}

func (s *Switch) press() {
	s.command.execute()
}

type Command interface {
	execute()
}

type OnCommand struct {
	lightSource LightSource
}

func (c *OnCommand) execute() {
	c.lightSource.turnOn()
}

type OffCommand struct {
	lightSource LightSource
}

func (c *OffCommand) execute() {
	c.lightSource.turnOff()
}

type LightSource interface {
	turnOn()
	turnOff()
}

type LightBulb struct {
	isOn bool
}

func (l *LightBulb) turnOn() {
	l.isOn = true
	fmt.Println("Turning on the light")
}

func (l *LightBulb) turnOff() {
	l.isOn = false
	fmt.Println("Turning off the light")
}

func main() {
	lightBulb := &LightBulb{}

	onCommand := &OnCommand{
		lightSource: lightBulb,
	}

	offCommand := &OffCommand{
		lightSource: lightBulb,
	}

	onSwitch := &Switch{
		command: onCommand,
	}

	onSwitch.press()

	offSwitch := &Switch{
		command: offCommand,
	}

	offSwitch.press()
}
