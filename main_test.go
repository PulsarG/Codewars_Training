package main

import (
	//"fmt"
	//"strconv"
	//"fmt"
	"testing"
	//"strings"
	//"unicode"
)

// var testNum = 25
var testResult = "TEST OK"

func TestOk(t *testing.T) {
	/* res := fooOne()
	if res != testResult {
		fmt.Println(res)
		t.Errorf("Test for OK Failed")
	} else {
		fmt.Println(res)
	} */

	if fooOne() != testResult {
		t.Errorf("Test for OK Failed")
	}
}
