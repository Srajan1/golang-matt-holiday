package hello

import "testing"

func TestSayHello(t *testing.T) {
	subTests := []struct {
		testParameters []string
		expectedResult string
	}{
		{
			expectedResult: "Hello Anonymous!",
		},
		{
			expectedResult: "Hello Srajan!",
			testParameters: []string{"Srajan"},
		},
		{
			expectedResult: "Hello Srajan, Rutuja!",
			testParameters: []string{"Srajan", "Rutuja"},
		},
	}

	for _, subTest := range subTests {
		s := Say(subTest.testParameters)
		if s != subTest.expectedResult {
			t.Errorf("Expected [%s], got [%s], against input (%v)", subTest.expectedResult, s, subTest.testParameters)
		}
	}
}
