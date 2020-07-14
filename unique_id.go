package es

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

// UniqueID is a unique string ID generator
type UniqueID struct {
	uuid   string
	prefix string
}

// Generate is a simple constructor to generate a unique id
func Generate(p string) *UniqueID {
	uuid, err := generateUUID()

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	return &UniqueID{
		uuid:   uuid,
		prefix: p,
	}
}

// ToString converts the generated unique id to a string
func (u *UniqueID) ToString() string {
	return fmt.Sprintf("%s_%s", u.prefix, u.uuid)
}

func generateUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	uuid[8] = 0x80
	uuid[5] = 0x40

	return hex.EncodeToString(uuid), nil
}
