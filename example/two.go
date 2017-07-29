package example

/*
This file has some examples of imported type usage
*/

import (
	"bytes"
	realsync "sync"
)

// GimmeBuffer gives you a bufferino
func GimmeBuffer() *bytes.Buffer {
	return new(bytes.Buffer)
}

// CheckSync locks and unlocks a lock
func CheckSync(s realsync.Locker) {
	s.Lock()
	s.Unlock()
}
