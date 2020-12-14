package main

import (
	"fmt"
	"gorm.io/gorm"
)

type PrinterInitializer interface {
	Start() error
}

type PrinterManager struct {
}

func (p *PrinterManager) Start() error {
	fmt.Println("Initializing the worst nightmare!!")
	return nil
}

func NewPrinterManager() PrinterInitializer {
	return &PrinterManager{}
}

type Application struct {
	printerInitializer PrinterInitializer
	db                 *interface{}
}

func NewApplication(p PrinterInitializer, db *gorm.DB) *Application {
	return &Application{
		printerInitializer: p,
		db:                 db,
	}
}

func main() {
	p := NewPrinterManager()
	app := NewApplication(p)

	app.printerInitializer.Start()
}
