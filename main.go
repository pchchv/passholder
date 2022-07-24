package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	uppercase     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase     = "abcdefghijklmnopqrstuvwxyz"
	numbers       = "0123456789"
	symbols       = "!@#$%^&*"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // Number of letter indices fitting in 63 bits
)

var envURL string
var src = rand.NewSource(time.Now().UnixNano())

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func passwordGenerator(characters string, length int) string {
	var char string
	if length == 0 {
		length = int(src.Int63())
		for length > 50 {
			length /= 3
		}
		for length < 5 {
			length *= 3
		}
	}
	switch {
	case strings.Contains(characters, "u"):
		char += uppercase
	case strings.Contains(characters, "l"):
		char += lowercase
	case strings.Contains(characters, "n"):
		char += numbers
	case strings.Contains(characters, "s"):
		char += symbols
	}
	return generator(length, char)
}

func generator(n int, symbols string) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(symbols) {
			b[i] = symbols[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func main() {
	envURL = getEnvValue("HOST") + ":" + getEnvValue("PORT")
	server()
}
