// Package secret implements a function which converts a decimal number into a secret handshake.
package secret

const testVersion = 1

// Handshake converts a decimal number into a secret handshake.
func Handshake(number uint) []string {
	handshake := []string{}
	if number&1 > 0 {
		handshake = append(handshake, "wink")
	}
	if number&2 > 0 {
		handshake = append(handshake, "double blink")
	}
	if number&4 > 0 {
		handshake = append(handshake, "close your eyes")
	}
	if number&8 > 0 {
		handshake = append(handshake, "jump")
	}
	if number&16 > 0 {
		handshake = reverseHandshake(handshake)
	}

	return handshake
}

func reverseHandshake(handshake []string) []string {
	for i, j := 0, len(handshake)-1; i < j; i, j = i+1, j-1 {
		handshake[i], handshake[j] = handshake[j], handshake[i]
	}

	return handshake
}
