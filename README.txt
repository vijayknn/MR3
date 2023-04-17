// /////////////////////////////////////////////////////////////////////////////////////
// Author - Vijay Namboodiri(vijay.namboodiri@hotmail.com)
// /////////////////////////////////////////////////////////////////////////////////////
About:
======*
This program provides solution to Mars Rover challenge problem.

Program implemented in Go Lang

Program tested against various test scenarios successfully.
For unit testing, Go Testing module is used - pls refer project folder MR3\pkg\models\marsrover\utp

How to run
==========*
There is an executable (MarsRover.exe) in the given Git link. This may be downloaded and run.
User can also download all the source files and run locally. Command is as below.

   go run <full path>main.go <full path>MarsRoverAppConfig.jason

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

Project folders and files:
==========================
\MR3\cmd\main\main.go
\MR3\config
 MarsRoverAppConfig.jason
 MarsRoverRefData.jason

\MR3\logs\MarsRoverAppErrorLog.txt
\MR3\pkg\
   <DIR>          controllers
   <DIR>          models
   <DIR>          utils
   <DIR>          view

\MR3\pkg\controllers\MarsRoverApp.go
\MR3\pkg\models\marsrover
	marsrover.go
	marsRoverFactory.go
	marsRoverReferenceDataAndAccessor.go
	
\MR3\pkg\models\marsrover\utp
	marsRoverFactory_test.go
	marsRoverFactory_test_input.jason
	MarsRoverRefDataT.jason
	marsRover_test.go
	marsRover_test_input.jason
	
\MR3\pkg\utils\remoteRover\remoteRover.go
\MR3\pkg\utils\utilFuncs\utilFuncs.go
\MR3\pkg\view\view.go

Design and build philosophy
============================*
Well designed, elegant neat code with least / no hard coding - follows SOLID priciples.
MVC & Factory Facotory patterns used and influenced by other design patterns like Composite, Proxy, builder, Inversion of control, facade etc
Easily understandable and maintainable.
Uses configuration data and logic for flexibility that will be useful beyond current requirements
Designed following Go OOPS way using Interfaces, Structs and methods.

Reference data used:
=====================
<Contents from - MarsRoverAppConfig.jason>
{
	"MarsRoverRefdataFile": "/GoProjects/MR3/config/MarsRoverRefData.jason",
	"MarsRoverLogFile": "/GoProjects/MR3/logs/MarsRoverAppErrorLog.txt",
	"GenericErrorMsg": "An error occured. Pease raise with support team!"
}

<Contents from MarsRoverRefData.jason>

{
	"plateau": [0, 0, 0, 0],
	"cardinals": ["N", "S", "W", "E"],
	"commands": ["L", "R", "M"],
	"roverTurnHelper": {
		"NL": "W",
		"WL": "S",
		"SL": "E",
		"EL": "N",
		"NR": "E",
		"ER": "S",
		"SR": "W",
		"WR": "N"
	},
	"roverXmoveHelper": {
		"N": 0,
		"W": -1,
		"S": 0,
		"E": 1
	},
	"roverYmoveHelper": {
		"N": 1,
		"W": 0,
		"S": -1,
		"E": 0
	}
}
========================================================================================*
