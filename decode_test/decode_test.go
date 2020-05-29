package main

import (
	"io"
	"testing"
)

func TestDecode(t *testing.T) {
	bytes := []byte{0, 41, 101, 54, 97, 54, 53, 54, 102, 56, 45, 48, 55, 51, 49, 45, 49, 49, 101, 57, 45, 98, 101, 51, 50, 45, 54, 97, 102, 99, 98, 49, 56, 50, 55, 55, 99, 100, 58, 57, 51, 52, 50}
	t.Log(string(bytes))
	v, _, vCount, err := readLengthEncodedString(bytes[0:])
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
	t.Log(vCount)
	v1, _, v1Count, err := readLengthEncodedString(bytes[0+vCount:])
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(v1))
	t.Log(v1Count)
	if string(v1) != "e6a656f8-0731-11e9-be32-6afcb18277cd:9342" {
		t.Error("error gtid")
	}
}

// returns the number read, whether the value is NULL and the number of bytes read
func readLengthEncodedInteger(b []byte) (uint64, bool, int) {
	// See issue #349
	if len(b) == 0 {
		return 0, true, 1
	}

	switch b[0] {
	// 251: NULL
	case 0xfb:
		return 0, true, 1

	// 252: value of following 2
	case 0xfc:
		return uint64(b[1]) | uint64(b[2])<<8, false, 3

	// 253: value of following 3
	case 0xfd:
		return uint64(b[1]) | uint64(b[2])<<8 | uint64(b[3])<<16, false, 4

	// 254: value of following 8
	case 0xfe:
		return uint64(b[1]) | uint64(b[2])<<8 | uint64(b[3])<<16 |
				uint64(b[4])<<24 | uint64(b[5])<<32 | uint64(b[6])<<40 |
				uint64(b[7])<<48 | uint64(b[8])<<56,
			false, 9
	}

	// 0-250: value of first byte
	return uint64(b[0]), false, 1
}

// returns the string read as a bytes slice, wheter the value is NULL,
// the number of bytes read and an error, in case the string is longer than
// the input slice
func readLengthEncodedString(b []byte) ([]byte, bool, int, error) {
	// Get length
	num, isNull, n := readLengthEncodedInteger(b)
	if num < 1 {
		return b[n:n], isNull, n, nil
	}

	n += int(num)

	// Check data length
	if len(b) >= n {
		return b[n-int(num) : n : n], false, n, nil
	}
	return nil, false, n, io.EOF
}
