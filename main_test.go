package main

import (
	"fmt"
	//"strconv"
	//"fmt"
	"testing"
	//"strings"
	//"unicode"
	"math/big"
)

func TestMain(t *testing.T) {
	a := big.NewInt(55)
	test := fibAdd(10)

	result := a.Cmp(test)

	switch result {
	case -1:
		fmt.Println("Test for OK Failed")
	case 0:
		fmt.Println("Test OK")
	case 1:
		fmt.Println("Test for OK Failed")
	}

	/* if result != a {
		t.Errorf("Test for OK Failed")
		t.Log(result, a)
	} */
}
