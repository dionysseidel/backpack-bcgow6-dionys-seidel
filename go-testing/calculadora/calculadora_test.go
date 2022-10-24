package calculadora

import "testing"

func TestRestar(t *testing.T) {
	errorMessage := "Restar function returned %v, but the expected result is %v"

	number1 := 6
	number2 := 4
	expectedResult := 2

	result := Restar(number1, number2)

	if result != expectedResult {
		t.Errorf(errorMessage, result, expectedResult)
	}

	number1 = 6
	number2 = 6
	expectedResult = 0

	result = Restar(number1, number2)

	if result != expectedResult {
		t.Errorf(errorMessage, result, expectedResult)
	}

	number1 = 6
	number2 = 10
	expectedResult = -4

	result = Restar(number1, number2)

	if result != expectedResult {
		t.Errorf(errorMessage, result, expectedResult)
	}

	number1 = 6
	number2 = -10
	expectedResult = 16

	result = Restar(number1, number2)

	if result != expectedResult {
		t.Errorf(errorMessage, result, expectedResult)
	}

	number1 = -6
	number2 = -10
	expectedResult = 4

	result = Restar(number1, number2)

	if result != expectedResult {
		t.Errorf(errorMessage, result, expectedResult)
	}

	number1 = -16
	number2 = -10
	expectedResult = -6

	result = Restar(number1, number2)

	if result != expectedResult {
		t.Errorf(errorMessage, result, expectedResult)
	}

	number1 = 1_600_000_000_000_000_000
	number2 = 1_500_000_000_000_000_000
	expectedResult = 100_000_000_000_000_000

	result = Restar(number1, number2)

	if result != expectedResult {
		t.Errorf(errorMessage, result, expectedResult)
	}

	number1 = 1_600_000_000_000_000_000
	number2 = -1_500_000_000_000_000_000
	expectedResult = 3_100_000_000_000_000_000

	result = Restar(number1, number2)

	if result != expectedResult {
		t.Errorf(errorMessage, result, expectedResult)
	}

	number1 = 1_600_000_000_000_000_000
	number2 = -1_500_000_000_000_000_000
	expectedResult = 3_100_000_000_000_000_000

	result = Restar(number1, number2)

	if result != expectedResult {
		t.Errorf(errorMessage, result, expectedResult)
	}

	number1 = -0
	number2 = -0
	expectedResult = 0

	result = Restar(number1, number2)

	if result != expectedResult {
		t.Errorf(errorMessage, result, expectedResult)
	}
}
