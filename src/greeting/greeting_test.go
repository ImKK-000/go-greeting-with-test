package greeting_test

import (
	. "greeting"
	"testing"
)

func TestHelloEmptyString(t *testing.T) {
	//Arrange
	expectedResult := "hello thailand"
	g := Greeting{Prefix: "hello"}

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
	g := Greeting{Prefix: "hello"}

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
	g := Greeting{Prefix: "hello"}

	//Action
	actualResult := g.Message("world")

	//Assert
	if expectedResult != actualResult {
		t.Errorf("Expect %s but it got %s", expectedResult, actualResult)
	}
}
