package text

import uuid "github.com/satori/go.uuid"

func GenerateRandomUUID() string {
	return uuid.NewV1().String()
}
