package main

import "fmt"

type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%d\t\t\t %o\t\t\t %x\t\t\t %b\n", KB, KB, KB, KB)
	fmt.Printf("%d\t\t\t %o\t\t\t %x\t\t\t %b\n", MB, MB, MB, MB)
	fmt.Printf("%d\t\t %b\n", GB, GB)
	fmt.Printf("%d\t\t %b\n", TB, TB)
	fmt.Printf("%d\t %b\n", PB, PB)
	fmt.Printf("%d\t %b\n", EB, EB)

}
