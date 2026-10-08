package allocator

type Ptr int

type Heap struct {
	mem []byte
}

func NewHeap(size int) (*Heap, error) {
	size &^= align - 1 // multiple of 8
	if size < minBlock {
		return nil, ErrHeapTooSmall
	}
	h := &Heap{mem: make([]byte, size)}
	h.setHeader(0, size, false)
	return h, nil

}

func (h *Heap) Malloc(n int) (Ptr, error) {
	if n <= 0 {
		return 0, ErrInvalidSize
	}
	need := blockSize(n)

	for off := 0; off < len(h.mem); {
		size, used := h.header(off)
		if used {
			off += size
			continue
		}

		for off+size < len(h.mem) {
			nextSize, nextUsed := h.header(off + size)
			if nextUsed {
				break
			}
			size += nextSize
		}

		if size >= need {
			if size-need >= minBlock {
				h.setHeader(off+need, size-need, false)
				size = need
			}

			h.setHeader(off, size, true)
			return Ptr(off + headerSize), nil
		}
		h.setHeader(off, size, false)
		off += size

	}
	return 0, ErrOutOfMemory
}

func (h *Heap) Free(p Ptr) error {
	target := int(p) - headerSize //returns the targetOffset
	for off := 0; off < len(h.mem); {
		size, used := h.header(off)
		if off == target {
			if !used {
				return ErrDoubleFree
			}
			h.setHeader(off, size, false)
			return nil
		}

		off += size

	}
	return ErrInvalidFree
}

func (h *Heap) Bytes(p Ptr) []byte {
	start := int(p) - headerSize
	size, _ := h.header(start)
	end := start + size
	return h.mem[p:end:end] //capping the capacity so it doesnt reach next Header
}
