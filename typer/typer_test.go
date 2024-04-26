package typer

import (
	"fmt"
	"log"
	"testing"
)

func TestPrint(t *testing.T) {
	typer := Typer{}

	if err := typer.Print("Bonjour!"); err != nil {
		if err == ErrFloatRangeInvalid {
			log.Println("FloatRange checked")
		} else {
			t.Fatal(err)
		}
	}

	typer.FloatRange = 200
	if err := typer.Print("Bonjour!"); err != nil {
		if err == ErrPrinterIsNull {
			log.Println("Printer checked")
		} else {
			t.Fatal(err)
		}
	}

	typer.Printer = fmt.Print
	if err := typer.Print("Bonjour!"); err != nil {
		t.Fatal(err)
	}

	if err := typer.Println("Au revoir~"); err != nil {
		t.Fatal(err)
	}
}
