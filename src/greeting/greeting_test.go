package greeting_test

import (
	. "greeting"
	"testing"
)

func TestHelloEmptyString(t *testing.T) {
	//Arrange
	expectedResult := "hello thailand"

	//Action
	actualResult := Message("")

	//Assert
	if expectedResult != actualResult {
		t.Errorf("Expect %s but it got %s", expectedResult, actualResult)
	}
}

func TestHelloBangkok(t *testing.T) {
	//Arrange
	expectedResult := "hello bangkok"

	//Action
	actualResult := Message("bangkok")

	//Assert
	if expectedResult != actualResult {
		t.Errorf("Expect %s but it got %s", expectedResult, actualResult)
	}
}

func TestHelloWorld(t *testing.T) {
	//Arrange
	expectedResult := "hello world"

	//Action
	actualResult := Message("world")

	//Assert
	if expectedResult != actualResult {
		t.Errorf("Expect %s but it got %s", expectedResult, actualResult)
	}
}
