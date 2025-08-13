package cipher

import "strings"

type shift struct {
	distance int
}
type vigenere struct {
	key string
}
func NewCaesar() Cipher {
	return NewShift(3)
}
func NewShift(distance int) Cipher {
	if distance == 0 || distance <= -26 || distance >= 26 {
		return nil
	}
	return shift{distance: distance}
}

func (c shift) Encode(input string) string {
	var result strings.Builder
	for _, r := range strings.ToLower(input) {
		if r >= 'a' && r <= 'z' {
			shifted := r + rune(c.distance)
			if shifted > 'z' {
				shifted -= 26
			} else if shifted < 'a' {
				shifted += 26
			}
			result.WriteRune(shifted)
		}
	}
	return result.String()
}
func (c shift) Decode(input string) string {
	inverseShift := shift{distance: -c.distance}
	return inverseShift.Encode(input)
}
func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}
	allAs := true
	for _, r := range key {
		if r != 'a' {
			allAs = false
		}
		if r < 'a' || r > 'z' {
			return nil 
		}
	}
	if allAs {
		return nil 
	}
	return vigenere{key: key}
}
func (v vigenere) Encode(input string) string {
	var result strings.Builder
	keyIndex := 0
	for _, r := range strings.ToLower(input) {
		if r >= 'a' && r <= 'z' {
			shiftDist := rune(v.key[keyIndex%len(v.key)] - 'a')
			shifted := r + shiftDist
			if shifted > 'z' {
				shifted -= 26
			}
			result.WriteRune(shifted)
			keyIndex++
		}
	}
	return result.String()
}
func (v vigenere) Decode(input string) string {
	var result strings.Builder
	keyIndex := 0
	for _, r := range strings.ToLower(input) {
		if r >= 'a' && r <= 'z' {
			shiftDist := rune(v.key[keyIndex%len(v.key)] - 'a')
			shifted := r - shiftDist
			if shifted < 'a' {
				shifted += 26
			}
			result.WriteRune(shifted)
			keyIndex++
		}
	}
	return result.String()
}