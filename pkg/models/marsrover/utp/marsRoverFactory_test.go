// ////////////////////////////////////////////////////////////////////////////////
// Author - Vijay Namboodiri(vijay.namboodiri@hotmail.com)
// ////////////////////////////////////////////////////////////////////////////////
package marsrover_test

//This 'Table test' test program tests two methods exposed by MarsRoverFactory struct.
//Initialize() and CreateMarsRover()
//Dependency - file 'marsRoverFactory_test_input.jason' in the same directory.
//This file contains the test data used for table test data.
//currently this file contains both negative and positive test scenarios.
//fmt.Print... commands below emits output only when this test case is run with -v flag
//This test to be invoked with command :  go test <path> marsRoverFactory_test.go (include -v optionally)

import (
	marsrover "GoProjects/MR3/pkg/models/marsrover"
	utils "GoProjects/MR3/pkg/utils/utilFuncs"
	"testing"
)

const marsRoverFactoryinputFile = "marsRoverFactory_test_input.jason"

type marsFactoryTestInput struct {
	RefData           string `jason:"RefData"`
	Plateau           []int  `json:"Plateau"`
	FactoryCreationER string `json:"FactoryCreationER"`
	RoverCreationER   string `json:"RoverCreationER"`
}

func TestRoverFactory(t *testing.T) {

	var marsFactoryTestInputs []marsFactoryTestInput = nil
	var mf marsrover.MarsroverFactory
	var err error = nil

	err = utils.LoadJasonFiletoStruct(marsRoverFactoryinputFile, &marsFactoryTestInputs)

	if err != nil {
		t.Errorf("\nTestRoverFactory() : Can't run the tests due to error : (%s)", err.Error())
		return
	}

	if marsFactoryTestInputs == nil {
		t.Errorf("\nTestRoverFactory() : Can't run the tests due to unknown error!")
		return
	}

	for i := 0; i < len(marsFactoryTestInputs); i++ {
		err = mf.Initialize(marsFactoryTestInputs[i].Plateau, marsFactoryTestInputs[i].RefData)
		utils.HandleTestResponse(t, err, marsFactoryTestInputs[i].FactoryCreationER, "TestRoverFactory() - mf.Initialize(marsFactoryInputs[i].Plateau, marsFactoryInputs[i].RefData)")

		_, err = mf.CreateMarsRover(1, 2, "N") // it is ok for this to be hard coded here as the goal of this test case is limited to MarsFactory class not MarsRover
		utils.HandleTestResponse(t, err, marsFactoryTestInputs[i].RoverCreationER, "TestRoverFactory() - mf.CreateMarsRover(1, 2, N)")
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
