//go:build exclude

package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	klength = 32
	vlength = 1000
)

func main() {
	mbs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	size := mbs * 1024 * 1024

	// All of this code is terrible. DO NOT DO EXACTLY THIS. But it's fine
	// for an example.
	b := bytes.Buffer{}
	b.WriteString(`//go:build prod
package main

var data = map[string]string{
	`)

	for b.Len() < size {
		b.WriteByte('"')
		for i := 0; i < klength; i++ {
			b.WriteByte(charset[rand.Intn(len(charset))])
		}
		b.WriteString(`": `)

		b.WriteByte('"')
		for i := 0; i < vlength; i++ {
			b.WriteByte(charset[rand.Intn(len(charset))])
		}
		b.WriteString("\",\n")
	}

	b.WriteString(`}`)

	fmt.Println(b.String())
}
