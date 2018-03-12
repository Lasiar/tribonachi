package system

import "testing"

func TestTrib(t *testing.T) {
	actualResult, _ := Trib(76)

	expectedResult := "12903063846126135669"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
