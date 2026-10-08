package main

import (
	"fmt"
	"gomalloc/allocator"
	"os"
)

func main() {
	h, err := allocator.NewHeap(256)
	if err != nil {
		panic(err)
	}

	out := os.Stdout
	h.Dump(out, "fresh heap")

	a, _ := h.Malloc(20)
	b, _ := h.Malloc(40)
	c, _ := h.Malloc(10)
	copy(h.Bytes(a), "yoooo")
	copy(h.Bytes(b), "Meeee")
	fmt.Printf("a=%d b=%d c=%d  a=%q b=%q\n", a, b, c,
		h.Bytes(a)[:5], h.Bytes(b)[:5])
	h.Dump(out, "after 3 mallocs")

	h.Free(a)
	h.Free(b)
	h.Dump(out, "after frreeing a and b")
	d, _ := h.Malloc(50)
	fmt.Printf("d=%d (reused a+b space)\n", d)
	h.Dump(out, "after malloc(50)")

	fmt.Println("double free:", h.Free(d), h.Free(d))
	_, err = h.Malloc(1000)
	fmt.Println("too big:", err)

}
