package main

import "fmt"

// Printers ------------------------------------------------------------------------------------------------------------

type Printer interface {
	Print(value interface{})
}

type InkJet struct{}

func (p *InkJet) Print(value interface{}) {
	fmt.Println("[InkJet]", value)
}

func NewInkJet() Printer {
	return &InkJet{}
}

type Laser struct{}

func (p *Laser) Print(value interface{}) {
	fmt.Println("[Laser]", value)
}

func NewLaser() Printer {
	return &Laser{}
}

// Scanners ------------------------------------------------------------------------------------------------------------

type Scanner interface {
	Scan() string
}

type BasicScanner struct{}

func (s *BasicScanner) Scan() string {
	return fmt.Sprintf("[BasicScanner] Reading from scanner...")
}

func NewBasicScanner() Scanner {
	return &BasicScanner{}
}

// All In One ----------------------------------------------------------------------------------------------------------

type PrintScanner interface {
	Printer
	Scanner
}

type AllInOnePrinter struct {
	Printer
	Scanner
}

func NewAllInOne(p Printer, s Scanner) PrintScanner {
	return &AllInOnePrinter{
		Printer: p,
		Scanner: s,
	}
}

// Main program  -------------------------------------------------------------------------------------------------------

func main() {
	inkJet := NewInkJet()
	laser := NewLaser()

	inkJet.Print("Test")
	laser.Print("Test")

	scanner := NewBasicScanner()
	fmt.Println(scanner.Scan())

	aio := NewAllInOne(inkJet, scanner)

	aio.Print(aio.Scan())
}
