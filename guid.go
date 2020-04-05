package helpers

import (
	"encoding/base64"
	"math/rand"
	"regexp"
	"time"

	"github.com/google/uuid"
)

const alphaAndDigitsBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomStringHelper(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = alphaAndDigitsBytes[rand.Intn(len(alphaAndDigitsBytes))]
	}
	return string(b)
}

// RandomString -
func RandomString(length int) string {
	// head -c 1000 /dev/urandom | base64 | tr -cd '[:alnum:]' | cut -c 1-100

	buffer := make([]byte, 1000*length)
	_, err := rand.Read(buffer)
	if err != nil {
		return randomStringHelper(length)
	}

	bufferAsB64String := base64.StdEncoding.EncodeToString(buffer)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return randomStringHelper(length)
	}

	theWholeThing := reg.ReplaceAllString(bufferAsB64String, "")
	if len(theWholeThing) > length {
		return theWholeThing[:length]
	}

	theWholeThing = theWholeThing + randomStringHelper(length)

	return theWholeThing[:length]
}

// NewGUID -
func NewGUID() string {
	guidAsString := RandomString(50)
	id, err := uuid.NewUUID()
	if err == nil {
		guidAsString = id.String()
	}

	return guidAsString
}

// NewGUIDWithLength -
func NewGUIDWithLength(length int) string {
	guidAsString := RandomString(length)
	id, err := uuid.NewUUID()
	if err == nil {
		guidAsString = id.String()
	}

	return guidAsString
}
