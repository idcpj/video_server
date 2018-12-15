package utils

import (
	"crypto/rand"
	"fmt"
)

func NewUuid() (string, error) {
	data := make([]byte, 16)
	_, err := rand.Read(data)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", data[0:4], data[4:6], data[6:8], data[8:10], data[10:])
	return uuid, nil

}
