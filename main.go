// ============================================================================
// = main.go
// = 	Description		The file from which to run any of the solutions
// = 	Date			October 08, 2021
// ============================================================================

package main

import (
	"errors"
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"time"

	prob "project-euler/problems"
	"project-euler/utils"
)

func main() {
	// program initialization (flags)
	probid := flag.Int("p", 0, "Which sequence to run. Example: -p 10")
	comptime := flag.Bool("t", true, "Enables approximate time-of-computation")
	
	flag.Parse()		// remember to parse!

	_, exists := StubStorage[*probid]

	if *probid == 0 {				// user must specify a problem to run
		utils.HandleError(errors.New("you need to specify a problem to run! "))
	} else if !exists {				// user must specify a sequence that exists
		utils.HandleError(errors.New("either this problem has not been implemented yet, or your id is invalid! "))
	}

	start := time.Now()
	temp := handler(*probid)
	fmt.Println(temp)
	duration := time.Since(start)

	// output time if requested
	if *comptime {
		utils.PrintInfo("Computed Problem" + strconv.FormatInt(int64(*probid), 10) + " in " + duration.String())
	}
}

// this handles the call to make life easier
func handler(name int, params ...interface{}) (interface{}) {
	out1, err := call(name, params...)
	utils.HandleError(err)
	return out1
}

// this is based upon:
// https://medium.com/@vicky.kurniawan/go-call-a-function-from-string-name-30b41dcb9e12
func call(name int, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(StubStorage[name])
	_ = params		// quiet warnings

	// build result interface
	var res []reflect.Value = f.Call(nil)
	result = res[0].Interface()
	return
}

// the following is a (large) mapping from strings to the corresponding function
var StubStorage = map[int]interface{}{
	1: prob.P1,
	2: prob.P2,
	3: prob.P3,
	4: prob.P4,
	5: prob.P5,
	6: prob.P6,
	7: prob.P7,
	8: prob.P8,
	9: prob.P9,
	10: prob.P10,
	11: prob.P11,
	12: prob.P12,
	13: prob.P13,
	14: prob.P14,
	15: prob.P15,
	16: prob.P16,
	17: prob.P17,
	18: prob.P18,
	19: prob.P19,
	20: prob.P20,
}