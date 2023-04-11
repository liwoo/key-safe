package passgen

import (
	"math/rand"
	"strings"
	"time"
)

type Config struct {
	MaxLength       int
	MinLength       int
	MinSpecialChars int
	MinUppercase    int
	MinNumbers      int
}

func DefaultConfig() Config {
	return Config{
		MaxLength:       15,
		MinLength:       6,
		MinSpecialChars: 1,
		MinUppercase:    1,
		MinNumbers:      1,
	}
}

func GeneratePassword(config Config) string {
	passwordLength := rand.Intn(config.MaxLength-config.MinLength) + config.MinLength

	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	specialChars := "@#$_"

	allChars := uppercaseLetters + lowercaseLetters + digits + specialChars

	var passwordBuilder strings.Builder

	passwordBuilder.WriteByte(uppercaseLetters[rand.Intn(len(uppercaseLetters))])
	passwordBuilder.WriteByte(lowercaseLetters[rand.Intn(len(lowercaseLetters))])
	passwordBuilder.WriteByte(digits[rand.Intn(len(digits))])

	//ensure minimum special characters from config
	for i := 0; i < config.MinSpecialChars; i++ {
		nextChar := specialChars[rand.Intn(len(specialChars))]
		passwordBuilder.WriteByte(nextChar)
	}

	//ensure minimum uppercase characters from config
	for i := 0; i < config.MinUppercase; i++ {
		nextChar := uppercaseLetters[rand.Intn(len(uppercaseLetters))]
		passwordBuilder.WriteByte(nextChar)
	}

	//ensure minimum numbers from config
	for i := 0; i < config.MinNumbers; i++ {
		nextChar := digits[rand.Intn(len(digits))]
		passwordBuilder.WriteByte(nextChar)
	}

	//fill in the rest of the password based on what is left to fill
	alreadyUsedUsedLength := len(passwordBuilder.String())

	for i := alreadyUsedUsedLength; i < passwordLength; i++ {
		nextChar := allChars[rand.Intn(len(allChars))]

		if isConsecutive(passwordBuilder.String(), nextChar, 5) || isRepeated(passwordBuilder.String(), nextChar, 8) {
			i--
			continue
		}

		passwordBuilder.WriteByte(nextChar)
	}

	return shuffle(passwordBuilder.String())

}

func isConsecutive(input string, nextChar byte, maxConsecutive int) bool {
	count := strings.Count(input[len(input)-maxConsecutive+1:], string(nextChar))
	return count >= maxConsecutive-1
}

func isRepeated(input string, nextChar byte, maxRepeated int) bool {
	count := strings.Count(input, string(nextChar))
	return count >= maxRepeated
}

func shuffle(input string) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	runes := []rune(input)
	n := len(runes)
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
