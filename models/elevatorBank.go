package models

import "github.com/go-errors/errors"

const(
	MAX_ELEVATOR_COUNT = 16
)

type ElevatorBank struct {
	numberOfFloors int
	elevators []Elevator
}

 func (EB ElevatorBank) AddElevator(elevatorToAdd *Elevator)(error){
	 if len(EB.elevators) + 1 <= MAX_ELEVATOR_COUNT{
	 	EB.elevators = append(EB.elevators,NewElevator())
	 	return nil
	 } else{
	 	return errors.New("Can not add elevator to elevator bank as it exceeds the bank limit of 16")
	 }
}

