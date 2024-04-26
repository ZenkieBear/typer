package typer

import (
	"fmt"
	"math/rand"
	"time"
)

// As same as fmt.Print
type Printer = func(a ...any) (n int, err error)

type Typer struct {
	// Minimum time to wait before each character input (millisecond).
	Base int
	// The interval of extra waiting time (millisecond).
	// It makes the output appear more vivid.
	FloatRange int
	// The time to wait when encountering spaces (millisecond).
	// Gap feels like thinking time before output the next word.
	Gap int

	// The function to print characters
	Printer Printer
}

func New(printer Printer, base, floatRange, gap int) *Typer {
	return &Typer{
		Printer:    printer,
		Base:       base,
		FloatRange: floatRange,
		Gap:        gap,
	}
}

func (t Typer) Print(message string) {
	for _, c := range message {
		delay := rand.Intn(t.FloatRange) + t.Base

		time.Sleep(time.Duration(delay) * time.Millisecond)

		t.Printer(string(c))

		if string(c) == " " {
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func (t Typer) Println(message string) {
	t.Print(fmt.Sprintf("%s\n", message))
}

var DefaultTyper Typer = Typer{
	FloatRange: 200,
	Base:       10,
	Gap:        200,
	Printer:    fmt.Print,
}

func Print(message string) {
	DefaultTyper.Print(message)
}

func Println(message string) {
	DefaultTyper.Println(message)
}
