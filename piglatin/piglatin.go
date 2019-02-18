package piglatin

import (
	"regexp"
	"strings"
	"unicode"
)

const suffix string = "ay"
const vowel string = "aeiou"
const nonConsonantSuffix string = "way"

var reg = regexp.MustCompile(`[A-Za-z']+|[^A-Za-z]`)
var onlyLattinReg = regexp.MustCompile(`[A-Za-z']`)

// Encrypt translates one or more english words into the PigLatin equlivent
func Encrypt(stdin string) string {
	var encryptedWords []string
	for _, word := range reg.FindAllString(stdin, -1) {
		if onlyLattinReg.MatchString(word) {
			word = EncryptSingleWord(word)
		}
		encryptedWords = append(encryptedWords, word)
	}
	return strings.Join(encryptedWords, "")
}

// EncryptSingleWord one english words into the PigLatin equlivent
func EncryptSingleWord(str string) string {
	first := rune(str[0])
	str = strings.ToLower(str)
	if !IsConsonant(str[0:1]) {
		return str + nonConsonantSuffix
	}
	strArr := strings.Split(str, "")

	for i := 0; i <= len(strArr) && IsConsonant(strArr[0]); i++ {
		var tmpChar string
		tmpChar, strArr = strArr[0], strArr[1:]
		strArr = append(strArr, tmpChar)
	}
	result := strings.Join(strArr, "") + suffix

	if unicode.IsUpper(first) {
		result = strings.Title(result)
	}
	return result
}

//IsConsonant return true if letter consonant
func IsConsonant(char string) bool {
	return !strings.Contains(vowel, char)
}
