package piglatin

import (
	"regexp"
	"strings"
	"sync"
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
	wg            sync.WaitGroup
)

// Encrypt translates one or more english words into the PigLatin equlivent
func Encrypt(stdin string) string {
	str := reg.FindAllString(stdin, -1)
	for i, v := range str {
		if onlyLattinReg.MatchString(v) {
			wg.Add(1)
			p := &str[i]
			go asyncEncrypt(p)
		}
	}
	wg.Wait()
	return strings.Join(str, "")
}

//// private

func asyncEncrypt(w *string) {
	*w = encryptSingleWord(*w)
	wg.Done()
}

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
