package piglatin

import "testing"

func TestPiglatinEncrypt(t *testing.T) {
	for _, test := range stringTestCases {
		actual := Encrypt(test.input)
		if actual != test.expected {
			t.Errorf("Encrypt test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkPiglatinEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Encrypt(test.input)
		}
	}
}
