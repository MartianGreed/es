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
