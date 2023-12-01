package randomizer

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomCode(role, name string) string {
	// Split the name into words
	words := strings.Fields(name)

	// Initialize a result string
	result := ""

	// Iterate through the words and use the first character of each word
	for _, word := range words {
		if len(word) > 0 {
			result += string(word[0])
		}
	}

	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 0 and 100
	randomInt := rand.Intn(1000)

	res := fmt.Sprintf("%s-%s-%d", role, result, randomInt)

	return strings.ToUpper(res)
}
