// /////////////////////////////////////////////////////////////////////////////////////
// Author - Vijay Namboodiri(vijay.namboodiri@hotmail.com)
// /////////////////////////////////////////////////////////////////////////////////////
About:
======*
This program provides solution to Mars Rover challenge problem.

Program implemented in Go Lang

Program tested against various test scenarios successfully.

How to run
==========*
There is an executable (MarsRover.exe) in the given Git link. This may be downloaded and run.
User can also download all the source files and run locally. Command is as below.

   "go run marsRover.go marsRoverController.go marsRoverReferenceDataAndAccessor.go"

Program runs in an infinite loop till user ends.

User Inputs are as below (as specified in the requirements). All are command line inputs
-----------------------------------------------------------------------------------------
Step 1 - Accepts upper coordinates of the plateau
Step 2 - Accepts initial position of the rover (i.e X, Y and cardinal)
Step 3 - Accepts commands (like LMLMMRM etc)
Step 4 - Takes Y/y input to continue with more rovers. 
  If Y/y is pressed, continues from step 2 adding another rover; otherwise exits the program printing output.

User can stop the program anytime by pressing Ctrl + C

All letters are accepted and processed ignoring the case.


Output - to command line
==========================*
  Prints the rover position upon executing the commands of each rover.
  Applicable error message are displayed as applicable
  Commands execution  


Design and build philosophy
============================*
Well designed, elegant neat code with least / no hard coding.
Easily understandable and maintainable.
Uses configuration data and logic for flexibility that will be useful beyond current requirements
Designed following Go OOPS way using Interfaces, Structs and methods.
Influenced by design patterns/ principles like Composite, Proxy, builder, Inversion of control, facade etc

Interfaces, Structs, functions
==============================*'
'marsRoverReferenceData' struct & type 'marsRoverReferenceDataAccessor' interface works together with associated functions to deliver
most features of Mars rover. They hold reference data that can be injected from outside. There are Setters, Getters and validation functions.

File - marsRoverReferenceDataAndAccessor.go

'marsRoverPosition' struct & 'MarsRover' Interface delivers functionality of the system and maintains the state.
File - MarsRover.go

File MarsRoverController.go - main() function and helper functions of main for testing.

Reference data used:
=====================
var plateau = []int{0, 0, 0, 0}  (Note: 3rd and 4th will be updated based on user input.)
var cardinals = []string{"N", "S", "W", "E"}
var roverTurnHelper = map[string]string{"NL": "W", "WL": "S", "SL": "E", "EL": "N", "NR": "E", "ER": "S", "SR": "W", "WR": "N"}
var roverXmoveHelper = map[string]int{"N": 0, "W": -1, "S": 0, "E": 1}
var roverYmoveHelper = map[string]int{"N": 1, "W": 0, "S": -1, "E": 0}

const (
	LEFT  string = "L"
	RIGHT string = "R"
	MOVE  string = "M"
)
========================================================================================*
