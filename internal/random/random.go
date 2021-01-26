package random

import "crypto/rand"

const mapStr = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Create a pseudo-random string from a set
// containing 62 alphanumeric characters
func String62 (length int) string {
	p := make([]byte, length)
	rand.Read(p)

	var out string = ""
	for _, b := range p {
		out += string(mapStr[int(b) % len(mapStr)])
	}
	return out
}