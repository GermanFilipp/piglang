package piglatin

import (
	"regexp"
	"strings"
	"unicode"
)

const (
	vowel              string = "aeiou"
	suffix             string = "ay"
	nonConsonantSuffix string = "way"
)

var (
	reg           = regexp.MustCompile(`[A-Za-z']+|[^A-Za-z]`)
	onlyLattinReg = regexp.MustCompile(`[A-Za-z']`)
)

// Encrypt translates one or more english words into the PigLatin equlivent
func Encrypt(stdin string) string {
	var encryptedWords []string
	for _, word := range reg.FindAllString(stdin, -1) {
		if onlyLattinReg.MatchString(word) {
			word = encryptSingleWord(word)
		}
		encryptedWords = append(encryptedWords, word)
	}
	return strings.Join(encryptedWords, "")
}

//// private

// encryptSingleWord one english words into the PigLatin equlivent
func encryptSingleWord(str string) string {
	firstByte := str[0]
	str = strings.ToLower(str)
	if !isConsonant(str[0:1]) {
		return str + nonConsonantSuffix
	}

	for i := 0; i <= len(str) && isConsonant(str[0:1]); i++ {
		var tmpChar string
		tmpChar, str = str[0:1], str[1:]
		str = str + tmpChar
	}

	if unicode.IsUpper(rune(firstByte)) {
		str = strings.Title(str)
	}
	return str + suffix
}

//isConsonant return true if letter consonant
func isConsonant(char string) bool {
	return !strings.Contains(vowel, char)
}
