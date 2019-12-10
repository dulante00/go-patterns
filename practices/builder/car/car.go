package car

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type carBuilder struct {
	wheels Wheels
	color  Color
	speed  Speed
}

func NewBuilder() *carBuilder {
	return &carBuilder{}
}

func (cb *carBuilder) Paint(c Color) Builder {
	cb.color = c
	return cb
}

func (cb *carBuilder) Color(color Color) Builder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Wheels(wheels Wheels) Builder {
	cb.wheels = wheels
	return cb
}

func (cb *carBuilder) TopSpeed(speed Speed) Builder {
	cb.speed = speed
	return cb
}

func (cb carBuilder) Stop() error {
	fmt.Printf("The %s %s car stopped.\n", cb.wheels, cb.color)
	return nil
}

func (cb *carBuilder) Drive() error {
	fmt.Printf("The %s %s car as %vm/s running.\n", cb.wheels, cb.color, cb.speed)
	return nil
}

func (cb *carBuilder) Build() Interface {
	return &carBuilder{
		wheels: cb.wheels,
		color:  cb.color,
		speed:  cb.speed,
	}
}