package main

import (
	"bytes"

	"github.com/vmihailenco/msgpack/v5"
)

// Message -
type MessagePacket struct {
	// SPENT SO LONG FOR THIS, SOMEHOW IT DOESNT WORK IF THE STRUCT KEY IS SAME AS THE MSG PACKK 😠
	Payload string `msgpack:"payload"`
}

// MarshallMsg -
func MarshallMsg(msg *MessagePacket) ([]byte, error) {
	var buf bytes.Buffer
	e := msgpack.NewEncoder(&buf)
	if err := e.Encode(msg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalMsg -
func UnmarshalMsg(b []byte, msg *MessagePacket) error {
	d := msgpack.NewDecoder(bytes.NewReader(b))
	return d.Decode(msg)
}
