// Global utility fuction to check given string is a string array(not case sensitive) - there is none part of standard package!
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

// This function is to search a string in slice array of strings.
func StringSearch(searchedArray []string, searchString string) bool {
	searchString = strings.ToUpper(searchString)
	for _, value := range searchedArray {
		if strings.ToUpper(value) == searchString {
			return true
		}
	}
	return false
}

// this logs error to the error log file.
func LogError(logFile string, err error) {
	f, e := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	//if, the logfile can't be opened, just send the error to console.
	if e != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(err)
}

// This function returns value of the given key from a given jason file. if not found, returns 0 in string format.
func GetFromJasonFile(jsfile string, key string) string {
	var err error
	jsonFile, err := os.Open(jsfile)
	if err != nil {
		return ""
	}
	defer jsonFile.Close()

	var mapdata map[string]interface{}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &mapdata)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s", mapdata[key])
}

func LoadJasonFiletoStruct(inpf string, stv interface{}) error {

	jsonFile, err := os.Open(inpf)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return json.Unmarshal(byteValue, &stv)
}

// This generic function provides details (success / fail and why) of the test exeuction whee -v option is specified along with test.
// Note that this is only meant to use where functionality / test case is checking for error only - i.e not specific return values.
func HandleTestResponse(t *testing.T, err error, expr string, msg string) {

	var actr string
	if err == nil {
		actr = ""
	} else {
		actr = "error" //this is approach is by design - to handle n number of possible error messages.
	}

	if actr == expr {
		fmt.Println(msg, " TEST SUCCESS! (Actual result == Expected result)")
	} else {
		t.Errorf("%s%s", msg, " TEST FAIL! (Actual result != Expected result)")
	}

	fmt.Println()
}

////////////////////////////////////////////////////////////////////////////////////////
