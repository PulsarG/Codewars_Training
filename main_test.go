package main

import (
	"fmt"
	//"strconv"
	//"fmt"
	"testing"
	//"strings"
	//"unicode"
)

var testNum = []int{2, 1, 3, 1, 2, 2, 2, 2}
var resPos = []int{2}

func TestOk(t *testing.T) {
	p := &PosPeaks{}
	r := searchPicks(testNum, p)

	if r.Pos[0] != resPos[0] {
		t.Errorf("TEST FAILED")
		fmt.Println(r)
		fmt.Println(resPos)
		return
	}
	if r.Pos[0] == resPos[0] {
		t.Errorf("TEST OK !!!")
		return
	}
	/* if res != testResult {
		t.Errorf("Test for OK Failed")
	} */
	/* if TowerBuilder(testNum) != testResult {
			t.Errorf("Test for OK Failed")
		}
	} */
}
