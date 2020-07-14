package es

import (
	"fmt"
	"testing"
)

type userRegistered struct {
	email    string
	password string
}

func (u userRegistered) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"email":    u.email,
		"password": u.password,
	}
}

func (u userRegistered) Deserialize(d map[string]interface{}) Event {
	return userRegistered{
		email:    fmt.Sprint(d["email"]),
		password: fmt.Sprint(d["password"]),
	}
}

type TestAggregateRoot struct {
	*AggregateRoot
	id *UserID
}

func (ar *TestAggregateRoot) AggregateRootID() string {
	return ar.id.ToString()
}

func NewUserAggregateRoot() *TestAggregateRoot {
	return &TestAggregateRoot{
		AggregateRoot: NewAggregateRoot(),
		id:            NewUserID(),
	}
}

func TestItCanApplyEvent(t *testing.T) {
	ure := userRegistered{
		email:    "valentin.dosimont@gmail.com",
		password: "valentin",
	}
	ar := NewUserAggregateRoot()
	ar.Apply(ure)

	if 1 != len(ar.RaiseEvents()) {
		t.Errorf("After applying an event RaiseEvents should have len == 1, it has len == %d", len(ar.RaiseEvents()))
	}
}
