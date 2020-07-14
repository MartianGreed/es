package es

// Serializable converts any struct to something commonly understandable
type Serializable interface {
	Serialize() map[string]interface{}
}

// Deserializable converts any map[string]interface{} to the corresponding struct
type Deserializable interface {
	Deserialize(d map[string]interface{}) Event
}

// Event interface is a convenience wrapper with the needed interfaces
type Event interface {
	Serializable
	Deserializable
}
