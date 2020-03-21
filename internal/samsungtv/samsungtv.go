package samsungtv

import (
	"time"
)

const timeout = time.Duration(1 * time.Second)

// CheckConnection checks if a connection to the TV is possible
func CheckConnection() bool {
	return false
}

// SendKey sends a key to the TV
func SendKey(key string) error {
	return nil
}
