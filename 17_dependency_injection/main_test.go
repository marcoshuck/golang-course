package main

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestApplication_ConnectDatabase(t *testing.T) {
	p := NewPrinterManagerTest()

	db := gorm.Open(owadpjapoatjatw)

	app := NewApplication(p, db)

	app.printerInitializer.Start()
}

type mockedPrinterManager struct {
}

func (m *mockedPrinterManager) Start() error {
	fmt.Println("This has been tested!")
	return nil
}

func NewPrinterManagerTest() PrinterInitializer {
	return &mockedPrinterManager{}
}
