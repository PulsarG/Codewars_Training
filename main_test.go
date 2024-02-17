package main

import (
	//"fmt"
	//"strconv"
	//"fmt"
	"testing"
	//"strings"
	//"unicode"
)

var testNum = 3
var testResult = [3]string{"  *  ", " *** ", "*****"}

func TestOk(t *testing.T) {

	res := TowerBuilder(testNum)

	for i, _ := range res {
		if res[i] != testResult[i] {
			t.Errorf("TEST FAILED")
			return
		}
	}
	/* if res != testResult {
		t.Errorf("Test for OK Failed")
	} */
	/* if TowerBuilder(testNum) != testResult {
			t.Errorf("Test for OK Failed")
		}
	} */
}
