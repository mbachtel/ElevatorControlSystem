package models

type Elevator struct {
	CurrentFloor int
}

func NewElevator () (Elevator){
	newElevator := Elevator{CurrentFloor:0}
	return newElevator
}
