package es

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
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
		return nil
	}

	return &UniqueID{
		uuid:   uuid,
		prefix: p,
	}
}

// ParseUUIDString takes a string as an input argument and returns a unique id
func ParseUUIDString(s string) *UniqueID {
	p := strings.Split(s, "_")

	if len(p) != 2 {
		log.Fatalf("%s unique ID was not generated properly. Accepted format is 'prefix_generateduuid'", s)
		return nil
	}

	return &UniqueID{
		uuid:   p[1],
		prefix: p[0],
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
