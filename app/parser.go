package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
)

const (
	Integer = ':'
	String  = '+'
	Bulk    = '$'
	ArraySymbol   = '*'
	Error   = '-'
)

type RESPData interface {
	String() string
}

func parseCommand(buf []byte) (RESPData, error) {
	//*2\r\n$4\r\nECHO\r\n$5\r\nmango\r\n
	log.Printf("parseCommand: %s", buf)
	var data RESPData

	switch buf[0] {
		case ArraySymbol:
		data, _ = CreateArray(buf)
	default:
		return nil, fmt.Errorf("Unknown type: %s", buf[0])
	}

	return data, nil
}

type BulkString struct {
	Value string
}

func (b BulkString) String() string {
	return b.Value
}

type Array struct {
	Values []RESPData
}

// *2\r\n$4\r\nECHO\r\n$5\r\nmango\r\n -> ECHO mango
// *<number-of-elements>\r\n<element-1>...<element-n>
func CreateArray(buf []byte) (Array, error) {
	
	r := bytes.IndexByte(buf, '\r') // Find first \r
	log.Printf("CreateArray: %d", r)
	if r == -1 {
		return Array{}, fmt.Errorf("Invalid RESP format")
	}

	n, err := strconv.Atoi(string(buf[1:r]))
	if err != nil {
		return Array{}, err
	}

	Values := make([]RESPData, n)

	i := 0
	c := r + 2
	for i < n {
		if c >= len(buf) {
			return Array{}, fmt.Errorf("Buffer too short")
		}

		switch buf[c] {
			case Bulk:
				c++
				r = bytes.IndexByte(buf[c:], '\r') // Find next \r
				if r == -1 {
					return Array{}, fmt.Errorf("Invalid bulk string")
				}
				length, err := strconv.Atoi(string(buf[c : c+r]))
				if err != nil {
					return Array{}, err
				}
				c += r + 2
				if c + length >= len(buf) {
					return Array{}, fmt.Errorf("Unexpected end of buffer for bulk string")
				}
				Values[i] = BulkString{Value: string(buf[c : c+length])}
				c += length + 2


		default:
			return Array{}, fmt.Errorf("Unknown type: %s", buf[c])

		}
		i++
	}

	return Array{Values: Values}, nil
}

func (a Array) String() string {
	s := ""
	for i, v := range a.Values {
		if i > 0 {
			s += " "
		}
		s += v.String()
	}
	log.Printf("Array: %s", s)
	return s
}

// func main() {
// 	buf := []byte("*2\r\n$4\r\nECHO\r\n$5\r\nmango\r\n")
// 	array, err := parseCommand(buf)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println(array.String()) 
// 	}
// }