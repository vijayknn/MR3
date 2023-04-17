// / ////////////////////////////////////////////////////////////////////////////////////////
// Author - Vijay Namboodiri(vijay.namboodiri@hotmail.com)
// ///////////////////////////////////////////////////////////////////////////////////////
package marsrover

import (
	utils "GoProjects/MR3/pkg/utils/remoteRover"
	"errors"
	"strings"
)

var (
	ErrInvalidRoverManeuverReferenceData = errors.New("marsrover : Invalid Rover Maneuver Reference data!")
)

// Note: ideally i would like  this const definition to be part of the struct; but Go won't allow!.
const (
	LEFT  string = "L"
	RIGHT string = "R"
	MOVE  string = "M"
)

type MarsRoverImpl struct {
	currX        int
	currY        int
	currCardinal string
	mrrda        MarsRoverReferenceDataAccessor
}

type MarsRover interface {
	utils.RemoteRover
	SetMarsRoverReferenceDataAccessor(mrrda MarsRoverReferenceDataAccessor) error
	ProcessCommands(commands string) error
}

func (mrp *MarsRoverImpl) SetMarsRoverReferenceDataAccessor(mrrda MarsRoverReferenceDataAccessor) error {

	if mrrda == nil {
		return ErrInvalidRoverManeuverReferenceData
	} else {
		mrp.mrrda = mrrda
		return nil
	}
}

// This function sets the initial position of the mars rover - like 1 2 E / 3  4 S
func (mrp *MarsRoverImpl) Initialize(x int, y int, c string) error {
	if mrp.mrrda == nil {
		return ErrInvalidRoverManeuverReferenceData
	}

	var err error

	err = mrp.mrrda.validateXUpdate(x)
	if err != nil {
		return err
	}

	err = mrp.mrrda.validateYUpdate(y)
	if err != nil {
		return err
	}

	err = mrp.mrrda.validateCardinal(c)
	if err != nil {
		return err
	}

	mrp.currX = x
	mrp.currY = y
	mrp.currCardinal = strings.ToUpper(c)

	return nil
}

// This will read the X and Y move reference table and update both X and Y axis.
// It is expected that the reference table will store 0 for X / Y for X or Y movements.
// This design approach avoids lots of coding and hard coding.
func (mrp *MarsRoverImpl) Move() error {

	if mrp.mrrda == nil {
		return ErrInvalidRoverManeuverReferenceData
	}

	var nextX, xChange int
	var nextY, yChange int
	var err error

	xChange, err = mrp.mrrda.GetXChange(mrp.currCardinal) //get the value to be added / deducted from current X

	if err != nil {
		return err
	}

	nextX = mrp.currX + xChange
	err = mrp.mrrda.validateXUpdate(nextX) //Check if the move is within the defined plateau

	if err != nil {
		return err
	}

	yChange, err = mrp.mrrda.GetYChange(mrp.currCardinal) //get the value to be added / deducted from current X
	if err != nil {
		return err
	}

	nextY = mrp.currY + yChange
	err = mrp.mrrda.validateYUpdate(nextY) //Check if the move is within the defined plateau
	if err != nil {
		return err
	}

	mrp.currX = nextX
	mrp.currY = nextY

	return nil
}

// This function turns the rover to the given side (i.e Left / Right etc) - called from process commands function of this interface
func (mrp *MarsRoverImpl) Turn(side string) error {

	if mrp.mrrda == nil {
		return ErrInvalidRoverManeuverReferenceData
	}

	val, err := mrp.mrrda.GetTurn(mrp.currCardinal, side)

	if err != nil {
		return err
	}

	mrp.currCardinal = val
	return nil
}

// This function process the series of commands passed on to the Mars over
func (mrp *MarsRoverImpl) ProcessCommands(commands string) error {

	var err error
	var command string

	//it is expected that untill it encounters first invalid command or untill all the commands are processed this continues.
	//and X, Y & Cardinal values are retained the point it reached.
	for i := 0; i < len(commands); i++ {
		//check if this command is allowed to be processed by this marsrover.
		command = string(commands[i] & '_')
		err = mrp.mrrda.validateCommand(command)
		if err != nil {
			return err
		}

		switch string(command) {
		case LEFT:
			err = mrp.Turn(LEFT)
			if err != nil {
				return err
			}
		case RIGHT:
			err = mrp.Turn(RIGHT)
			if err != nil {
				return err
			}
		case MOVE:
			err = mrp.Move()
			if err != nil {
				return err
			}
		default: // this is required just in case config ref data contains commands that are not supported by this
			return ErrInvalidCommand
		}
	}
	return nil
}

// Getter functions - Will return default values (0 / "") if called before intializing.
// This approach in this case is to avoid 'not so important' error checking by callers.
// This return the current X position of Mars rover
func (mrp *MarsRoverImpl) GetCurrentPositionX() int {
	return mrp.currX
}

// This return the current Y position of Mars rover
func (mrp *MarsRoverImpl) GetCurrentPositionY() int {
	return mrp.currY
}

// This return the current cardinal position of Mars rover
func (mrp *MarsRoverImpl) GetCurrentCardinal() string {
	return mrp.currCardinal
}

///////////////////////////////////////////////////////////////////////////////////////////////////
