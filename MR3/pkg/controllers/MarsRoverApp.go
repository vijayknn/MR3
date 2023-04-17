// ////////////////////////////////////////////////////////////////////////////////
// Author - Vijay Namboodiri(vijay.namboodiri@hotmail.com)
// ////////////////////////////////////////////////////////////////////////////////
package marsroverapp

import (
	marsrover "GoProjects/MR3/pkg/models/marsrover"
	utils "GoProjects/MR3/pkg/utils/utilFuncs"
	view "GoProjects/MR3/pkg/view"
	"log"
	"os"
)

// Below are keys to extract config info from App config file - which is passed as command line arguments to the program.
const appconfigfileKey = "MarsRoverRefdataFile"
const logfileKey = "MarsRoverLogFile"
const genericErrorMsgKey = "GenericErrorMsg"

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Run() {

	//App can only be run along with a config file mentioned in that order - nothing else.
	if len(os.Args) != 2 {
		log.Fatal("\nTo run this app, enter : <path><program name> <path><config file>!\n")
		return
	}

	var mf marsrover.MarsroverFactory
	//In case any error with MarsRover factory initialization, system can't continue
	err := intializeMarsRoverFactory(&mf)
	if err != nil {
		return
	}

	//Now system should send as many Mars rovers as user want
	//Each Mars rovers will be intialized with X, Y and Cardinal values user provides
	//And, Mars rovers will be sent the commands user enters

	var mrs map[int]marsrover.MarsRover
	sendRoversAndCommands(mf, &mrs)

	//Now display the position of each Mars rovers.
	view.DisplayRoversData(mrs)
}

func intializeMarsRoverFactory(mf *marsrover.MarsroverFactory) error {

	//Get upper cords of plateau from user...
	//Only upper cord is required as per the requirements that only upper cords user provided.

	var err error
	x2, y2 := view.GetPlateauUpperCords()

	//get the ref data file name from app config file and pass to factory to refer it.
	//by passing ref data to factory rather than letting it help itself that, factories with difference reference data can be created as required.
	err = (*mf).Initialize([]int{0, 0, x2, y2}, utils.GetFromJasonFile(os.Args[1], appconfigfileKey))
	if err != nil {
		view.DisplayGenericError(utils.GetFromJasonFile(os.Args[1], genericErrorMsgKey))
		utils.LogError(utils.GetFromJasonFile(os.Args[1], logfileKey), err)
	}
	return err
}

func sendRoversAndCommands(mf marsrover.MarsroverFactory, mrs *map[int]marsrover.MarsRover) {

	var i int = 0
	*mrs = make(map[int]marsrover.MarsRover)
	for {
		m, err := mf.CreateMarsRover(view.GetMarsRoverInitialroverProperties())
		if err == nil {
			(*mrs)[i] = m
			err = (*mrs)[i].ProcessCommands(view.GetCommands())
			i++
		}

		if err != nil {
			view.DisplayGenericError(utils.GetFromJasonFile(os.Args[1], genericErrorMsgKey))
			utils.LogError(utils.GetFromJasonFile(os.Args[1], logfileKey), err)
		}

		if !view.ContinueRoverSend() {
			break
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
