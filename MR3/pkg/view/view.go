package view

import (
	"GoProjects/MR3/pkg/models/marsrover"
	"fmt"
)

// this function accepts commands for the rover - Like LMRM.. etc
func GetCommands() string {
	var roverCommands string
	fmt.Print("\nEnter commands to send to Mars Rover without space - like LMRM... : ")
	fmt.Scan(&roverCommands)
	return roverCommands
}

// Global function to help driver function (main) to get inputs from user to upper X & Y of the Plateau
func GetPlateauUpperCords() (int, int) {
	var x2, y2 int

	for {
		fmt.Print("\nEnter Plateau upper coordinates (lower coordinates are 0,0) within which all rovers operate to : ")
		n, err := fmt.Scan(&x2, &y2)
		if n != 2 || err != nil {
			fmt.Println("Wrong number of inputs or format...please try again.")
		} else {
			return x2, y2
		}
	}
}

// This function accepts initial inputs from user for mars rovers - like 1 2 N
func GetMarsRoverInitialroverProperties() (int, int, string) {
	for {
		var x int
		var y int
		var cardinal string

		fmt.Print("\nEnter Mars Rover intial values (X Y and Cardinals) : ")
		n, err := fmt.Scan(&x, &y, &cardinal)

		if n != 3 || err != nil {
			fmt.Println("Wrong number of inputs or format...please try again.")
		} else {
			return x, y, cardinal
		}
	}
}

// Displays X, Y, & Cardinal data of given map of rover data
func DisplayRoversData(mrs map[int]marsrover.MarsRover) {

	j := 1
	c := len(mrs) // to check if this could be a nil pointer
	if c == 0 {
		fmt.Println("\nNo rover data to display")
		return
	}

	fmt.Println("\nPositions of Rovers are as below:")
	for i := 0; i < c; i++ {
		if mrs[i] != nil { //this is just to avoid a null pointer exception
			fmt.Println("\n Rover data[", j, "] : ", mrs[i].GetCurrentPositionX(), mrs[i].GetCurrentPositionY(), mrs[i].GetCurrentCardinal())
			j++
		}
	}
}

// simple Y /y input from user to continue operation!
func ContinueRoverSend() bool {
	var cont string

	fmt.Print("\nTo send more rovers to Mars press Y/y or press any other key:  ")

	fmt.Scan(&cont)
	if cont == "Y" || cont == "y" {
		return true
	}
	return false
}

func DisplayGenericError(msg string) {
	fmt.Println(msg)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
