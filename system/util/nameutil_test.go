package util

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {

	s := "myPath/yourPath1"
	if !IsPathForm(s) {
		t.Error(s + " KO found not path form")
	}

	s = "myPath.yourPath2"
	if IsPathForm(s) {
		t.Error(s + " KO found path form")
	}

	s = "myPath3"
	if !IsPathForm(s) {
		t.Error(s + " KO found not path form")
	}

	s = "./myPath4"
	if !IsPathForm(s) {
		t.Error(s + " KO found not path form")
	}

	/*
	 * Dotted
	 */
	s = "myPath/yourPath1"
	if IsDottedForm(s) {
		t.Error(s + " KO found dotted form")
	}

	s = "myPath.yourPath2"
	if !IsDottedForm(s) {
		t.Error(s + " KO found not dotted form")
	}

	s = "myPath3"
	if !IsDottedForm(s) {
		t.Error(s + " KO found not dotted form")
	}

	s = "./myPath4"
	if IsDottedForm(s) {
		t.Error(s + " KO found dotted form")
	}

}

func TestCapitalCase(t *testing.T) {

	s := "f3_ArrayNameStruct"
	fmt.Println(ToCapitalCase(s))

	s = "f3_array-name_struct"
	fmt.Println(ToCapitalCase(s))
}

func TestFormatIdentifier(t *testing.T) {
	/*
		s2 := FirstToLower("TestMio")
		fmt.Println(s2)
	*/
	s := "books.[].title"
	s1 := FormatIdentifier(s, "", lowerCase, suppress, indexIjk)
	fmt.Println(s1)

	printIt("The value is %s %s\n", "ciao", "pippo")
	printIt("No value here \n")

}

func printIt(f string, args ...interface{}) {

	if args == nil {
		fmt.Printf(f)
	} else {
		fmt.Printf(f, args...)
	}

}
