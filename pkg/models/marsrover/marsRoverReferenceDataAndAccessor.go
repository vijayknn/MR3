// ////////////////////////////////////////////////////////////////////////////////////////
// Author - Vijay Namboodiri(vijay.namboodiri@hotmail.com)
// /////////////////////////////////////////////////////////////////////////////////////////
package marsrover

import (
	utils "GoProjects/MR3/pkg/utils/utilFuncs"
	"errors"
)

var (
	ErrPlateauNotSet          = errors.New("marsrover : Plateau Not Set!")
	ErrInvalidPlateauRefData  = errors.New("marsrover : Invalid Plateau Reference data!")
	ErrInvalidX               = errors.New("marsrover : Invalid X!")
	ErrInvalidY               = errors.New("marsrover : Invalid Y!")
	ErrInvalidCaridnal        = errors.New("marsrover : Invalid cardinal!")
	ErrInvalidCaridnalOrTurn  = errors.New("marsrover : Invalid cardinal or Turn!")
	ErrInvalidCaridnalRefData = errors.New("marsrover : Invalid cardinal reference data!")
	ErrInvalidCommand         = errors.New("marsrover : Invalid command! ")
	ErrInvalidCommandsRefData = errors.New("marsrover : Invalid commands reference data!")
	ErrInvalidTurnRefData     = errors.New("marsrover : Invalid Turn reference data!")
	ErrInvalidXMoveRefData    = errors.New("marsrover : Invalid X axis move reference data!")
	ErrInvalidYMoveRefData    = errors.New("marsrover : Invalid Y axis move reference data!")
	ErrRefDataNotSet          = errors.New("marsrover : Reference data not set")
)

// Go needs a struct with public (i.e memebers with first letter in caps) fields
// to unmarshall data from jason files and hence the identical structs below - second one being private to
// to this package
type MarsRoverReferenceDataDepInj struct {
	Plateau          []int             `json:"plateau"`
	Cardinals        []string          `json:"cardinals"`
	Commands         []string          `json:"commands"`
	RoverTurnHelper  map[string]string `json:"roverTurnHelper"`
	RoverXmoveHelper map[string]int    `json:"roverXmoveHelper"`
	RoverYmoveHelper map[string]int    `json:"roverYmoveHelper"`
}

type MarsRoverReferenceData struct {
	plateau          []int
	cardinals        []string
	commands         []string
	roverTurnHelper  map[string]string
	roverXmoveHelper map[string]int
	roverYmoveHelper map[string]int
	initialized      bool
}

type MarsRoverReferenceDataAccessor interface {
	Initialize(pt []int, confFile string) error

	setPlateau(pt []int) error
	setCardinals(cdnls []string) error
	setCommands(cmds []string) error
	setRoverTurnHelper(rth map[string]string) error
	setRoverXmoveHelper(rxmh map[string]int) error
	setRoverYmoveHelper(rymh map[string]int) error

	validateCardinal(cardinal string) error
	validateCommand(command string) error
	validateXUpdate(nextX int) error
	validateYUpdate(nextX int) error

	GetTurn(cardinal string, turn string) (string, error)
	GetXChange(cardinal string) (int, error)
	GetYChange(cardinal string) (int, error)
}

func (rmp *MarsRoverReferenceData) Initialize(pt []int, confFile string) error {

	rmp.initialized = false
	var err error = nil
	var dIRefdata MarsRoverReferenceDataDepInj

	err = utils.LoadJasonFiletoStruct(confFile, &dIRefdata)
	if err != nil {
		return err
	}

	err = rmp.setPlateau(pt)
	if err != nil {
		return err
	}

	//Set reference data / logic - this also avoids hard coding in the code and easy enhancebility
	//Various Set functions of 'marsRoverReferenceDataAccessor' interface 'deep copies' these reference data
	//so that each rover can operate wtih different sets of data if required. Where such different data to be provided
	//to each rover, required reference data should be altered before passing on to Set functions.
	err = rmp.setCardinals(dIRefdata.Cardinals)
	if err != nil {
		return err
	}

	err = rmp.setCommands(dIRefdata.Commands)
	if err != nil {
		return err
	}

	err = rmp.setRoverTurnHelper(dIRefdata.RoverTurnHelper)
	if err != nil {
		return err
	}

	err = rmp.setRoverXmoveHelper(dIRefdata.RoverXmoveHelper)
	if err != nil {
		return err
	}

	err = rmp.setRoverYmoveHelper(dIRefdata.RoverYmoveHelper)
	if err != nil {
		return err
	}

	rmp.initialized = true
	return nil
}

// various Set functions below.
func (rmp *MarsRoverReferenceData) setPlateau(pt []int) error {

	if pt == nil || len(pt) < 4 {
		return ErrInvalidPlateauRefData
	}

	if pt[0] == pt[2] || pt[1] == pt[3] {
		return ErrInvalidPlateauRefData
	}

	if pt[0] == pt[1] && pt[1] == pt[2] && pt[2] == pt[3] {
		return ErrInvalidPlateauRefData
	}

	rmp.plateau = make([]int, len(pt))
	for i, value := range pt {
		rmp.plateau[i] = value
	}

	return nil
}

func (rmp *MarsRoverReferenceData) setCardinals(cdnls []string) error {

	if cdnls == nil || len(cdnls) < 1 {
		return ErrInvalidCaridnalRefData
	}

	rmp.cardinals = make([]string, len(cdnls))
	for i, crdn := range cdnls {
		rmp.cardinals[i] = crdn
	}

	return nil
}

func (rmp *MarsRoverReferenceData) setCommands(cmds []string) error {

	if cmds == nil || len(cmds) < 1 {
		return ErrInvalidCommandsRefData
	}

	rmp.commands = make([]string, len(cmds))
	for i, cmd := range cmds {
		rmp.commands[i] = cmd
	}

	return nil
}

func (rmp *MarsRoverReferenceData) setRoverTurnHelper(rth map[string]string) error {

	if len(rth) < 1 {
		return ErrInvalidTurnRefData
	}

	rmp.roverTurnHelper = make(map[string]string)
	for k, v := range rth {
		rmp.roverTurnHelper[k] = v
	}

	return nil
}

func (rmp *MarsRoverReferenceData) setRoverXmoveHelper(rxmh map[string]int) error {

	if len(rxmh) < 1 {
		return ErrInvalidXMoveRefData
	}

	rmp.roverXmoveHelper = make(map[string]int)
	for k, v := range rxmh {
		rmp.roverXmoveHelper[k] = v
	}

	return nil
}

func (rmp *MarsRoverReferenceData) setRoverYmoveHelper(rymh map[string]int) error {

	if len(rymh) < 1 {
		return ErrInvalidYMoveRefData
	}

	rmp.roverYmoveHelper = make(map[string]int)
	for k, v := range rymh {
		rmp.roverYmoveHelper[k] = v
	}

	return nil
}

// This function checks if the X axis movement about to happen will be within the plateau
func (rmp *MarsRoverReferenceData) validateXUpdate(nextX int) error {
	var min int
	var max int

	if !rmp.initialized {
		return ErrRefDataNotSet
	}

	if len(rmp.plateau) < 4 {
		return ErrPlateauNotSet
	}

	//below logic is to handle plateau with negetive or positive X1,Y1, X2, Y2
	if rmp.plateau[2] < rmp.plateau[0] {
		min = rmp.plateau[2]
		max = rmp.plateau[0]
	} else {
		min = rmp.plateau[0]
		max = rmp.plateau[2]
	}

	if nextX >= min && nextX <= max {
		return nil
	} else {
		return ErrInvalidX
	}
}

// This method checks if the Y axis movement about to happen will be within the plateau
func (rmp *MarsRoverReferenceData) validateYUpdate(nextY int) error {
	var min int
	var max int

	if !rmp.initialized {
		return ErrRefDataNotSet
	}

	if len(rmp.plateau) < 4 {
		return ErrPlateauNotSet
	}

	//below logic is to handle plateau with negetive or positive X1,Y1, X2, Y2
	if rmp.plateau[3] < rmp.plateau[1] {
		min = rmp.plateau[3]
		max = rmp.plateau[1]
	} else {
		min = rmp.plateau[1]
		max = rmp.plateau[3]
	}

	if nextY >= min && nextY <= max {
		return nil
	} else {
		return ErrInvalidY
	}
}

// This checks if the rover turn to E/W/S/N is valid - by referring the given list.
// This helps to set fewer than 4 sides as the maximum allowed for a rover.
func (rmp *MarsRoverReferenceData) validateCardinal(cardinal string) error {

	if !rmp.initialized {
		return ErrRefDataNotSet
	}

	if rmp.cardinals == nil {
		return ErrInvalidCaridnalRefData
	}

	if utils.StringSearch(rmp.cardinals, cardinal) {
		return nil
	}

	return ErrInvalidCaridnal
}

// This checks if the rover turn to E/W/S/N is valid - by referring the given list.
// This helps to set fewer than 4 sides as the maximum allowed for a rover.
func (rmp *MarsRoverReferenceData) validateCommand(command string) error {

	if !rmp.initialized {
		return ErrRefDataNotSet
	}

	if rmp.commands == nil {
		return ErrInvalidCommandsRefData
	}

	if utils.StringSearch(rmp.commands, command) {
		return nil
	}

	return ErrInvalidCommand
}

// this is to check the name of the turning cardinal - i.e from North to L will take to West
func (rmp *MarsRoverReferenceData) GetTurn(cardinal string, turn string) (string, error) {

	err := rmp.validateCardinal(cardinal)

	if err != nil {
		return "", err
	}

	val, ok := rmp.roverTurnHelper[cardinal+turn]

	if !ok {
		return "", ErrInvalidCaridnalOrTurn
	}

	return val, nil
}

// This returns the advancement required on X axis when M command to move is issued. it could be 1, -1 or 0
func (rmp *MarsRoverReferenceData) GetXChange(cardinal string) (int, error) {

	err := rmp.validateCardinal(cardinal)

	if err != nil {
		return 0, err
	}

	val, ok := rmp.roverXmoveHelper[cardinal]

	if !ok {
		return 0, ErrInvalidCaridnal
	}

	return val, nil
}

// This returns the advancement required on X axis when M command to move is issued. it could be 1, -1 or 0
func (rmp *MarsRoverReferenceData) GetYChange(cardinal string) (int, error) {

	err := rmp.validateCardinal(cardinal)

	if err != nil {
		return 0, err
	}

	val, ok := rmp.roverYmoveHelper[cardinal]

	if !ok {
		return 0, ErrInvalidCaridnal
	}

	return val, nil
}

//////////////////////////////////////////////////////////////////////////////////////////
