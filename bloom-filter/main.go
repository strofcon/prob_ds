package main

import (
	"fmt"
)

func stringToInt(s string) uint64 {
	x := uint64(0)
	for _, r := range s {
		x += uint64(r)
	}
	return x
}

func hashCountOddFromRight(x uint64, n uint64) uint64 {
	y := uint64(0)

	for i := 0; i < 16; i++ {
		oddBit := x & 1
		y |= (oddBit << uint64(i))
		x >>= 2
	}

	return y % n
}

func hashCountEvenFromRight(x uint64, n uint64) uint64 {
	y := uint64(0)

	for i := 0; i < 16; i++ {
		x >>= 1
		evenBit := x & 1
		y |= (evenBit << uint64(i))
		x >>= 1
	}

	return y % n
}

func insertString(s string, e *uint64, n uint64) []uint64 {
	i := stringToInt(s)
	h1 := hashCountOddFromRight(i, n)
	h2 := hashCountEvenFromRight(i, n)

	fmt.Printf("incoming: %064b\n", *e)
	*e |= (uint64(1) << h1)
	fmt.Printf("addh1:    %064b\n", *e)

	*e |= (uint64(1) << h2)
	fmt.Printf("addh2:    %064b\n\n", *e)

	return []uint64{h1, h2}
}

func seenBefore(s string, e *uint64, n uint64) bool {
	i := stringToInt(s)
	h1 := hashCountOddFromRight(i, n)
	if *e&(uint64(1)<<h1) == 0 {
		return false
	}

	h2 := hashCountEvenFromRight(i, n)
	if *e&(uint64(1)<<h2) == 0 {
		return false
	}

	return true
}

func main() {
	n := uint64(64)
	e := uint64(0)
	a := insertString("hello world", &e, n)
	b := insertString("floobety", &e, n)
	c := insertString("Chowny is the mediocre-est ever", &e, n)

	fmt.Printf("a: %+v\n", a)
	fmt.Printf("b: %+v\n", b)
	fmt.Printf("c: %+v\n", c)

	aSeen := seenBefore("hello world", &e, n)
	bSeen := seenBefore("floobety", &e, n)
	cSeen := seenBefore("Chowny is the greatest ever", &e, n)
	dSeen := seenBefore("Chowny is the mideocre-est ever", &e, n)

	fmt.Printf("aSeen: %+v\n", aSeen)
	fmt.Printf("bSeen: %+v\n", bSeen)
	fmt.Printf("cSeen: %+v\n", cSeen)
	fmt.Printf("dSeen: %+v\n", dSeen)
}
