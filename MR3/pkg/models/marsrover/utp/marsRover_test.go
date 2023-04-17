// ////////////////////////////////////////////////////////////////////////////////
// Author - Vijay Namboodiri(vijay.namboodiri@hotmail.com)
// ////////////////////////////////////////////////////////////////////////////////
package marsrover_test

import (
	marsrover "GoProjects/MR3/pkg/models/marsrover"
	utils "GoProjects/MR3/pkg/utils/utilFuncs"
	"fmt"
	"strconv"
	"testing"
)

const marsRoverTestInputFile = "marsRover_test_input.jason"

type marsRoverTestInput struct {
	RefData           string   `jason:"RefData"`
	Plateau           []int    `json:"Plateau"`
	MarsRoverInput    []string `json:"MarsRoverInput"`
	MarsRoverCommands string   `json:"MarsRoverCommands"`
	MarsRoverOutput   string   `json:"MarsRoverOutput"`
	FactoryErrorExp   string   `json:"FactoryErrorExp"`
	RoverErrorExp     string   `json:"RoverErrorExp"`
}

func TestMarsRover(t *testing.T) {
	var marsRoverTestInputs []marsRoverTestInput = nil
	var mf marsrover.MarsroverFactory
	var mr marsrover.MarsRover

	var err error = nil

	err = utils.LoadJasonFiletoStruct(marsRoverTestInputFile, &marsRoverTestInputs)

	if err != nil {
		t.Errorf("\nTestMarsRover() : Can't run the tests due to error : (%s)", err.Error())
		return
	}

	if marsRoverTestInputs == nil {
		t.Errorf("\nTestMarsRover() : Can't run the tests due to unknown error!")
		return
	}

	for i := 0; i < len(marsRoverTestInputs); i++ {

		//Create Marsrover thru Marsrover factory
		err = mf.Initialize(marsRoverTestInputs[i].Plateau, marsRoverTestInputs[i].RefData)
		utils.HandleTestResponse(t, err, marsRoverTestInputs[i].FactoryErrorExp, "TestMarsRover() - mf.Initialize(marsFactoryInputs[i].Plateau, marsFactoryInputs[i].RefData)")

		//below two lines of code supports x,y & cardinal stored in input file as just as on one string(otherwise will need a nested struct and jason attributes)
		x, _ := strconv.Atoi(marsRoverTestInputs[i].MarsRoverInput[0])
		y, _ := strconv.Atoi(marsRoverTestInputs[i].MarsRoverInput[1])

		mr, err = mf.CreateMarsRover(x, y, marsRoverTestInputs[i].MarsRoverInput[2])
		utils.HandleTestResponse(t, err, marsRoverTestInputs[i].RoverErrorExp, "TestMarsRover() - mf.CreateMarsRover()")

		//Now send commands to Mars rover - only if tests expected to continue having Marsrover been created.
		if mr == nil { //this test scenario ends at failed creation of marsRover.
			continue
		}

		mr.ProcessCommands(marsRoverTestInputs[i].MarsRoverCommands)
		utils.HandleTestResponse(t, err, marsRoverTestInputs[i].RoverErrorExp, "TestMarsRover() - mr.ProcessCommands()")

		ar := strconv.Itoa(mr.GetCurrentPositionX()) + strconv.Itoa(mr.GetCurrentPositionY()) + mr.GetCurrentCardinal()
		if ar == marsRoverTestInputs[i].MarsRoverOutput {
			fmt.Println("TestMarsRover() - Get current X,Y, Cardinal TEST SUCCESS! (Actual result == Expected result)")
		} else {
			t.Errorf("TestMarsRover() - Get current X,Y, Cardinal TEST FAIL! (Actual result != Expected result)")
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
