package es

import (
	"testing"
)

type UserID struct {
	*UniqueID
}

func NewUserID() *UserID {
	uuid := Generate("user")
	return &UserID{uuid}
}

func FromStringToUserID(s string) *UserID {
	return &UserID{ParseUUIDString(s)}
}

func TestGeneratedIdsAreDifferent(t *testing.T) {
	s := make([]string, 10000)

	for i := range s {
		s[i] = NewUserID().ToString()
		for j := 0; j < i; j++ {
			if s[j] != s[i] {
				continue
			}
			t.Errorf("Ids at index %d is the same as index %d", i, j)
		}
	}
}

func TestItCanParseUniqueID(t *testing.T) {
	sid := "user_9e57dc585b4090ff80b3276b8e9f2e2b"

	uuid := FromStringToUserID(sid)

	if uuid.prefix != "user" && uuid.uuid != "9e57dc585b4090ff80b3276b8e9f2e2b" {
		t.Error("Failed to parse user unique id")
	}
}
