package piglatin

import "testing"

func TestPiglatinEncrypt(t *testing.T) {
	for _, test := range encryptTestCases {
		actual := Encrypt(test.input)
		if actual != test.expected {
			t.Errorf("Encrypt test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkPiglatinEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range encryptTestCases {
			Encrypt(test.input)
		}
	}
}

func TestEncryptSingleWord(t *testing.T) {
	for _, test := range encryptSingleWordTestCases {
		actual := encryptSingleWord(test.input)
		if actual != test.expected {
			t.Errorf("encryptSingleWord test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkEncryptSingleWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range encryptSingleWordTestCases {
			encryptSingleWord(test.input)
		}
	}
}

func TestIsConsonant(t *testing.T) {
	for _, test := range alpha {
		actual := isConsonant(test.input)
		if actual != test.expected {
			t.Errorf("encryptSingleWord test [%s], expected [%t], actual [%t]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkIsConsonant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range alpha {
			isConsonant(test.input)
		}
	}
}
