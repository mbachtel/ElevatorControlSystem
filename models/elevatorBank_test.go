package models

import (
	"testing"
	"github.com/stretchr/testify/require"
)


func Test_AddElevatorSuccess(t *testing.T){
	elevatorBank := &ElevatorBank{}
	elevatorToAdd := NewElevator()
	err := elevatorBank.AddElevator(&elevatorToAdd)
	require.NoError(t, err)
}
