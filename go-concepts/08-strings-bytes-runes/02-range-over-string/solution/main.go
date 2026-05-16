package main

import "fmt"

func main() {
	s := "héllo"

	fmt.Println("range (rune-aware):")
	for i, r := range s {
		fmt.Printf("  byte_index=%d rune=%q (%d)\n", i, r, r)
	}

	fmt.Println("\nbyte loop:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("  i=%d byte=0x%02x\n", i, s[i])
	}
}
