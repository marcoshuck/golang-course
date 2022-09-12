package main

import (
	"fmt"
)

type PrinterInitializer interface {
	Start() error
}

type PrinterManager struct{}

func (p *PrinterManager) Start() error {
	fmt.Println("Initializing the worst nightmare!!")
	return nil
}

func NewPrinterManager() PrinterInitializer {
	return &PrinterManager{}
}

type allGoodPrinter struct {
}

func (a *allGoodPrinter) Start() error {
	fmt.Println("It's all good man! Saul goodman")
	return nil
}

func NewPrinterAllGood() PrinterInitializer {
	return &allGoodPrinter{}
}

type Application struct {
	printerInitializer PrinterInitializer
	//db                 *interface{}
}

func NewApplication(p PrinterInitializer) *Application {
	return &Application{
		printerInitializer: p,
		//db:                 db,
	}
}

func main() {
	//p := NewPrinterManager()
	saul := NewPrinterAllGood()

	app := NewApplication(saul)

	app.printerInitializer.Start()
}
