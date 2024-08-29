package utils

import "testing"

/*
func Test_IsPrime_13(t *testing.T) {

		// arrange
		no := 13
		expectedResult := true

		// act
		actualResult := IsPrime(no)

		// assert
		if actualResult != expectedResult {
			t.Errorf("IsPrime(13) : expected = %v, actual = %v", expectedResult, actualResult)
		}
	}
*/
func Test_IsPrime(t *testing.T) {
	testData := []struct {
		name           string
		no             int
		expectedResult bool
	}{
		{name: "IsPrime_13", no: 13, expectedResult: true},
		{name: "IsPrime_14", no: 14, expectedResult: false},
		{name: "IsPrime_15", no: 15, expectedResult: false},
		{name: "IsPrime_16", no: 16, expectedResult: false},
		{name: "IsPrime_17", no: 17, expectedResult: true},
	}
	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			actualResult := IsPrime(td.no)

			// assert
			if actualResult != td.expectedResult {
				t.Errorf("%s : expected = %v, actual = %v", td.name, td.expectedResult, actualResult)
			}
		})
	}
}
