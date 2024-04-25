package typer

import (
	"fmt"
	"math/rand"
	"time"
)

type Printer = func(a ...any) (n int, err error)

type Typer struct {
	floatRange int
	base       int
	printer    Printer
}

var DefaultTyper Typer = Typer{
	floatRange: 200,
	base:       10,
	printer:    fmt.Print,
}

func (t Typer) print(message string) {
	for _, c := range message {
		delay := rand.Intn(t.floatRange) + t.base

		time.Sleep(time.Duration(delay) * time.Millisecond)

		t.printer(string(c))

		if string(c) == " " {
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func (t Typer) println(message string) {
	t.print(fmt.Sprintf("%s\n", message))
}

func Print(message string) {
	DefaultTyper.print(message)
}

func Println(message string) {
	DefaultTyper.println(message)
}
