package main

import "fmt"

type bcounter int

func (b *bcounter) Write(p []byte) (int, error) {
	*b += bcounter(len(p))
	return len(p), nil
}

func main() {
	var b bcounter
	b.Write([]byte("THISISATEST"))
	fmt.Println(b)

	b = 0

	str := "TEST"

	fmt.Fprintf(&b, "THISISA%sé", str)

	fmt.Println(b)
	é
}
