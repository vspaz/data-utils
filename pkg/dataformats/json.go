package dataformats

import (
	"encoding/json"
)

func FromJson(deserializable []byte, deserialized any) error {
	return json.Unmarshal(deserializable, deserialized)
}

func ToJson(serializable any) []byte {
	encodedMessage, _ := json.Marshal(serializable)
	return encodedMessage
}
