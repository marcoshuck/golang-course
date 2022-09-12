package main

import "fmt"

type Engine struct {
	Trademark  string
	HorsePower float64
}

// NewEngine is an initialization function
func NewEngine(tm string, hp float64) Engine {
	return Engine{
		Trademark:  tm,
		HorsePower: hp,
	}
}

type Seat struct {
	Trademark string
}

func NewSeat(tm string) Seat {
	return Seat{
		Trademark: tm,
	}
}

// Composition over inheritance

type Car struct {
	engine Engine // compose
	seat   *Seat  // aggregate (nil)
}

func NewCar(engine Engine, seat *Seat) Car {
	return Car{
		engine: engine,
		seat:   seat,
	}
}

func main() {
	e := NewEngine("Audi", 1200.0)
	s := NewSeat("SuperComfy")
	c := NewCar(e, &s)

	fmt.Printf("Car: %+v\n", c)
	fmt.Printf("Seat: %+v\n", *c.seat)
}
