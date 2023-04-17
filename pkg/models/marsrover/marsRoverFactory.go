package marsrover

import (
	"errors"
)

var (
	ErrFactoryNotInitialized = errors.New("marsrover : Mars Rover Factory is not initialised!")
)

type MarsroverFactory struct {
	intialized bool
	mrrda      MarsRoverReferenceDataAccessor //interface
	mrd        MarsRoverReferenceData         //struct
}

func (mrf *MarsroverFactory) Initialize(p []int, confFile string) error {
	//load the config file and populate the memeber variables
	//this function can be used for initialize and re-initialize and hence not checking if it is alreday been intialized.

	mrf.intialized = false
	mrf.mrrda = &mrf.mrd

	err := mrf.mrd.Initialize(p, confFile)

	if err != nil {
		return err
	}

	mrf.intialized = true
	return nil
}

func (mrf *MarsroverFactory) CreateMarsRover(x int, y int, c string) (MarsRover, error) {

	//Shouldn't proceed if factory is not initialized.
	if mrf.intialized == false {
		return nil, ErrFactoryNotInitialized
	}

	var mr MarsRover
	var mri MarsRoverImpl
	mr = &mri
	mr.SetMarsRoverReferenceDataAccessor(mrf.mrrda)
	err := mr.Initialize(x, y, c)

	//if initialization failed, the object shouldn't be used /referred and set nil
	if err != nil {
		mr = nil
	}

	return mr, err
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
