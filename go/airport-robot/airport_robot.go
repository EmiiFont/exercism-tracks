package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greeter(visitorName string) string
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greeter(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

type Portuguese struct{}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}

func (i Portuguese) Greeter(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func SayHello(visitorName string, concrete Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", concrete.LanguageName(), concrete.Greeter(visitorName))
}
