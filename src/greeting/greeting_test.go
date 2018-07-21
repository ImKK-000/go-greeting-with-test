package greeting_test

import (
	. "greeting"
	"testing"
)

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
