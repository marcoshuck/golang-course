package main

import "fmt"

type HouseBuilder struct {
	House House
	Error error
}

type House struct {
	Doors   int
	Windows int
}

func (b *HouseBuilder) SetDoors(n int) *HouseBuilder {
	b.House.Doors = n
	return b
}

func (b *HouseBuilder) SetWindows(n int) *HouseBuilder {
	b.House.Windows = n
	return b
}

func (b *HouseBuilder) Find(result interface{}) *HouseBuilder {
	result = b.House
	b.Error = nil
}

func (b *HouseBuilder) Build() (House, error) {
	var h House
	h = b.House
	b.House = House{}
	return h, nil
}

func main() {
	var houseBuilder HouseBuilder

	var h House

	err := houseBuilder.SetDoors(3).SetWindows(2).Find(&h).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("House:", h)

	h, err = houseBuilder.SetWindows(2).Build()

	fmt.Println("House 2:", h)
}
