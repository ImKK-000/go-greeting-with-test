package greeting_test

import (
	. "greeting"
	"testing"
)

var g = Greeting{Prefix: "hello", DefaultEmptyString: "thailand"}

func TestHelloEmptyString(t *testing.T) {
	//Arrange
	expectedResult := "hello thailand"

	//Action
	actualResult := g.Message("")

	//Assert
	if expectedResult != actualResult {
		t.Errorf("Expect %s but it got %s", expectedResult, actualResult)
	}
}

func TestHelloBangkok(t *testing.T) {
	//Arrange
	expectedResult := "hello bangkok"

	//Action
	actualResult := g.Message("bangkok")

	//Assert
	if expectedResult != actualResult {
		t.Errorf("Expect %s but it got %s", expectedResult, actualResult)
	}
}

func TestHelloWorld(t *testing.T) {
	//Arrange
	expectedResult := "hello world"

	//Action
	actualResult := g.Message("world")

	//Assert
	if expectedResult != actualResult {
		t.Errorf("Expect %s but it got %s", expectedResult, actualResult)
	}
}
