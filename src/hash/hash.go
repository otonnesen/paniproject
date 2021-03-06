package hash

import (
	"errors"
	"math"
)

// HashFromNumber converts an integer into a 6 digit string using characters
// from [a-z0-9].
func HashFromNumber(id int) (string, error) {
	// Our hashes all have length 6, so the corresponding ID can't be larger than 36^6
	if id < 0 || id > int(math.Pow(36, 6)) {
		return "", errors.New("Hash calculation failed, ID is out of range [0, 36^6]")
	}

	// We just convert our decimal number to base 36 and pad it with 0s to
	// get our hash; this makes our hash function invertible, which is
	// pretty handy.
	chars := []byte{'0', '0', '0', '0', '0', '0'}

	q, r := id/36, id%36

	var i int

	for i = len(chars) - 1; q != 0; i-- {
		digit, _ := decToBase36(r)
		chars[i] = digit
		r = q % 36
		q = q / 36
	}

	digit, _ := decToBase36(r)
	chars[i] = digit

	return string(chars), nil
}

// decToBase36 takes an integer 0 <= q <= 35 and returns its base 36 representation.
func decToBase36(q int) (byte, error) {
	if q < 0 || q > 35 {
		return 0, errors.New("Digit cannot be lesser than 0 or greater than 35")
	}

	if q < 10 {
		return byte('0' + q), nil
	} else {
		return byte('a' - 10 + q), nil
	}
}

// base36ToDec takes a base 36 digit and returns an integer from 0 to 35.
func base36ToDec(q byte) (int, error) {
	if !('0' <= q && q <= '9') && !('a' <= q && q <= 'z') {
		return 0, errors.New("Digit must be in [a-z0-9]")
	}

	if '0' <= q && q <= '9' {
		return int(q - '0'), nil
	} else {
		return int(q - 'a' + 10), nil
	}
}

// NumberFromHash converts a 6 digit string containing characters from [a-z0-9] to
func NumberFromHash(hash string) (int, error) {
	// TODO: Custom back-halfs can have lengths other than 6.
	if len(hash) != 6 {
		return 0, errors.New("ID calculation failed, hash has the wrong length")
	}

	s := 0

	for i, c := range hash {
		value, err := base36ToDec(byte(c))
		if err != nil {
			return 0, errors.New("ID calculation failed, hash contains illegal characters")
		}
		s += value * int(math.Pow(36, float64(len(hash)-1-i)))
	}

	return s, nil
}
