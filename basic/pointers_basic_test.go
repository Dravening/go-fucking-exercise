package basic

import (
	"fmt"
	"testing"
)

func TestBasicPointer(t *testing.T) {
	answer := 42
	fmt.Println(&answer) // & is address operator

	address := &answer
	fmt.Println(*address)                     // * is dereferencing, which providers the value that a memory address refers to.
	fmt.Printf("address is a %T \n", address) // print the pointer type

	var address2 *int  // declare a pointer
	address2 = address // address2 can store some pinter type
	fmt.Println(*address2)

}

func TestPointer(t *testing.T) {
	var test *string = new(string)
	*test = "123"
	fmt.Println(test)
}
