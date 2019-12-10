package main

import (
	"github.com/sharelinux/go-patterns/practices/builder/car"
)

func main() {

	assembly := car.NewBuilder().Paint(car.RedColor)

	familyCar := assembly.Wheels(car.SteelWheels).TopSpeed(50 * car.MPH).Build()
	_ = familyCar.Drive()

	sportsCar := assembly.Wheels(car.SportsWheels).TopSpeed(150 * car.MPH).Build()
	_ = sportsCar.Drive()

}
