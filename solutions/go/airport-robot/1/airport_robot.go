package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
}
type German struct{}
func (g German) LanguageName() string {
	return "German"
}
type Italian struct{}
func (i Italian) LanguageName() string {
	return "Italian"
}
type Portuguese struct{}
func (p Portuguese) LanguageName() string {
	return "Portuguese"
}
func SayHello(name string, g Greeter) string {
	switch g.(type) {
	case German:
		return fmt.Sprintf("I can speak German: Hallo %s!", name)
	case Italian:
		return fmt.Sprintf("I can speak Italian: Ciao %s!", name)
	case Portuguese:
		return fmt.Sprintf("I can speak Portuguese: Olá %s!", name)
	default:
		return "I don't know that language."
	}
}
