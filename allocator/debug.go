package allocator

import (
	"fmt"
	"io"
)

func (h *Heap) Dump(w io.Writer, title string) {
	fmt.Fprintln(w, "--", title)
	for off := 0; off < len(h.mem); {
		size, used := h.header(off)
		state := "free"
		if used {
			state = "USED"
		}

		fmt.Fprintf(w, "   offset %4d  size %4d  %s\n", off, size, state)
		off += size

	}

}

func (h *Heap) validate() error {
	off := 0
	for off < len(h.mem) {
		size, _ := h.header(off)
		if size < minBlock {
			return fmt.Errorf("block at %d: size %d below minimum", off, size)

		}
		if size%align != 0 {
			return fmt.Errorf("block at %d: size %d not aligned", off, size)
		}

		if off+size > len(h.mem) {
			return fmt.Errorf("block at %d: size %d runs past heap end", off, size)
		}

		off += size
		if off != len(h.mem) {
			return fmt.Errorf("blocks end at %d, heap is %d bytes", off, len(h.mem))
		}
	}
	return nil

}
